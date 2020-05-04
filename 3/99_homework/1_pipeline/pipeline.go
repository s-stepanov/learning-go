package pipeline

import (
    "sync"
)

type job func(in, out chan interface{})

func jobExecutor(worker job, in chan interface{}, out chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	worker(in, out)
	close(out)
}

func Pipe(jobs ...job) {
	var channels []chan interface{}
	wg := &sync.WaitGroup{}

	for i := 0; i < len(jobs); i++ {
		channels = append(channels, make(chan interface{}))
	}
	
	for i, worker := range(jobs) {
		wg.Add(1)
		index := i
		if i == 0 {
			go jobExecutor(worker, nil, channels[index], wg)
		} else {
			go jobExecutor(worker, channels[index - 1], channels[index], wg)
		}
	}

	wg.Wait()
}