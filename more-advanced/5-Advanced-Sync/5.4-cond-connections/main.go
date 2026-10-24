package main

import (
	"fmt"
	"sync"
	"time"
)

type PoolConn struct {
	conn []int
	mu   sync.Mutex
	cond *sync.Cond
}

func NewPoolConn(size int) *PoolConn {
	pool := &PoolConn{
		conn: make([]int, size),
	}
	for i := 0; i < size; i++ {
		pool.conn[i] = i
	}
	pool.cond = sync.NewCond(&pool.mu)
	return pool
}

func (p *PoolConn) ReleaseConn(conn int) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.conn = append(p.conn, conn)
	fmt.Printf("Conexão liberada: %d\n", conn)

	p.cond.Signal()
}

func (p *PoolConn) GetConn() int {
	p.mu.Lock()
	defer p.mu.Unlock()

	for len(p.conn) == 0 {
		fmt.Println("Aguardando conexão...")
		p.cond.Wait()
	}

	conn := p.conn[0]
	p.conn = p.conn[1:]

	fmt.Printf("Conexão obtida: %d\n", conn)
	return conn
}

func main() {
	pool := NewPoolConn(2)

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			conn := pool.GetConn()
			fmt.Printf("Goroutine %d usando conexão: %d\n", id, conn)
			time.Sleep(3 * time.Second)
			pool.ReleaseConn(conn)
		}(i)
	}

	wg.Wait()
	fmt.Println("Todas as goroutines terminaram.")
}
