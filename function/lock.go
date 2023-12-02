package function

import (
	"reflect"
	"sync"
)

type MutexWrapper struct {
	mutex sync.Mutex
}

func (mw *MutexWrapper) lock() {
	mw.mutex.Lock()
}

func (mw *MutexWrapper) unLock() {
	mw.mutex.Unlock()
}

func LockFunc(callback interface{}) interface{} {
	m := MutexWrapper{}
	m.lock()
	defer m.unLock()

	callbackType := reflect.TypeOf(callback)
	if callbackType.Kind() != reflect.Func {
		panic("callback must be a function")
	}

	return reflect.MakeFunc(callbackType, func(params []reflect.Value) []reflect.Value {
		// Convert params to interface{} slice
		var resultInterfaces []interface{}
		for _, result := range reflect.ValueOf(callback).Call(params) {
			resultInterfaces = append(resultInterfaces, result.Interface())
		}

		// Return the results as reflect.Value
		var resultReflectValues []reflect.Value
		for _, result := range resultInterfaces {
			resultReflectValues = append(resultReflectValues, reflect.ValueOf(result))
		}

		return resultReflectValues
	}).Interface()
}

// RWLock

// type RWMutexWrapper struct {
// 	rwMutex sync.RWMutex
// }

// func (mw *RWMutexWrapper) rLock() {
// 	mw.rwMutex.RLock()
// }

// func (mw *RWMutexWrapper) rUnLock() {
// 	mw.rwMutex.RUnlock()
// }
