package function

import (
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

type OneInOneOut func(i interface{}) interface{}
func LockOneInOneOut(callback OneInOneOut) OneInOneOut {
	m := MutexWrapper{}
	m.lock()
	defer m.unLock()
	return callback
}

type OneInTwoOut func(i interface{}) (interface{}, interface{})
func LockOneInTwoOut(callback OneInTwoOut) OneInTwoOut {
	m := MutexWrapper{}
	m.lock()
	defer m.unLock()
	return callback
}

type TwoInOneOut func(i1, i2 interface{}) interface{}
func LockTwoInOneOut(callback TwoInOneOut) TwoInOneOut {
	m := MutexWrapper{}
	m.lock()
	defer m.unLock()
	return callback
}

type TwoInTwoOut func(i1, i2 interface{}) (interface{}, interface{})
func LockTwoInTwoOut(callback TwoInTwoOut) TwoInTwoOut {
	m := MutexWrapper{}
	m.lock()
	defer m.unLock()
	return callback
}

// func LockFuncOneInOneOut[i, o any](f func(i) o) func(i) o {
// 	m := MutexWrapper{}
// 	return func(iVal i) o {
// 		m.lock()
// 		defer m.unLock()
// 		return f(iVal)
// 	}
// }

// func LockFuncOneInTwoOut[i, o1, o2 any](f func(i) (o1, o2)) func(i) (o1, o2) {
// 	m := MutexWrapper{}
// 	return func(iVal i) (o1, o2) {
// 		m.lock()
// 		defer m.unLock()
// 		return f(iVal)
// 	}
// }

// func LockFuncTwoInTwoOut[i1, i2, o1, o2 any](f func(i1, i2) (o1, o2)) func(i1, i2) (o1, o2) {
// 	m := MutexWrapper{}

// 	return func(i1Val i1, i2Val i2) (o1, o2) {
// 		m.lock()
// 		defer m.unLock()
// 		return f(i1Val, i2Val)
// 	}
// }

// // RWLock

// type RWMutexWrapper struct {
// 	rwMutex sync.RWMutex
// }

// func (mw *RWMutexWrapper) rLock() {
// 	mw.rwMutex.RLock()
// }

// func (mw *RWMutexWrapper) rUnLock() {
// 	mw.rwMutex.RUnlock()
// }