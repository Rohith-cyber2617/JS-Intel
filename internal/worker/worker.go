package worker

import "sync"

type Pool struct {
	Workers int
}

func New(workers int) *Pool {
	return &Pool{
		Workers: workers,
	}
}

func (p *Pool) Run(tasks []func()) {

	var wg sync.WaitGroup

	sem := make(chan struct{}, p.Workers)

	for _, task := range tasks {

		wg.Add(1)

		sem <- struct{}{}

		go func(t func()) {
			defer wg.Done()
			defer func() { <-sem }()

			t()
		}(task)
	}

	wg.Wait()
}
