package main

import (
	"fmt"
	"sync"
	"time"
)

type Lock struct {
	locked bool
	mu     sync.Mutex
	cond   *sync.Cond
}

func NewLock() *Lock {
	l := &Lock{}
	l.cond = sync.NewCond(&l.mu)

	return l
}

var start = time.Now()

func logf(msg string) {
	fmt.Printf("[4%d ms] - %s \n", time.Since(start).Milliseconds(), msg)
}

func worker(id int, lock *Lock, wg *sync.WaitGroup) {
	defer wg.Done()
	defer lock.mu.Unlock()

	lock.mu.Lock()

	for !lock.locked {
		logf(fmt.Sprintf("Trabalhador %d está aguardando o bloqueio", id))
		lock.cond.Wait()
	}

	logf(fmt.Sprintf("Trabalhador %d adquiriu o bloqueio", id))
	time.Sleep(1 * time.Second)
}

func main() {
	var (
		workers = 5
		lock    = NewLock()
		wg      sync.WaitGroup
	)

	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go worker(i, lock, &wg)
	}

	time.Sleep(2 * time.Second)
	lock.mu.Lock()
	lock.locked = true
	logf("A goroutine principal está notificando todos os trabalhadores")
	lock.cond.Broadcast()
	lock.mu.Unlock()

	wg.Wait()
}
