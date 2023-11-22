package utils

import (
	"time"
)

type DelayedResult[T any] struct {
	Value T
	HasReturn bool
}

func delay[T any](f func() T, timeout int, hasReturn bool) DelayedResult[T] {
	resultChan := make(chan T, 1)

	go func() {
		time.Sleep(time.Duration(timeout) * time.Second)
		result := f()
		if hasReturn {
			resultChan <- result
		}
	}()

	if hasReturn {
		return DelayedResult[T]{Value: <-resultChan, HasReturn: true}
	}

	return DelayedResult[T]{HasReturn: false}
}

func DelayF[T any](f func() T, timeout int) DelayedResult[T] {
	return delay(f, timeout, true)
}

func DelayV(f func(), timeout int) DelayedResult[struct{}] {
	return delay(func() struct{} {
		f()
		return struct{}{}
	}, timeout, false)
}