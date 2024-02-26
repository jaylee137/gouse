package function

import "sync"

func Parallelize(functions ...func()) {
	var waitGroup sync.WaitGroup

	waitGroup.Add(len(functions))

	ch := make(chan struct{}, len(functions))
	done := make(chan struct{}, len(functions))

	for _, f := range functions {
		ch <- struct{}{}
		go func(copyFunc func()) {
			defer func() {
				<-ch
				waitGroup.Done()
			}()
			copyFunc()
		}(f)
	}

	go func() {
		waitGroup.Wait()
		close(done)
	}()

	<-done
}
