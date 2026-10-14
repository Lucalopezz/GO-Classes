package main

import (
	"fmt"
	"time"
)

type Task[T any] struct {
	Name     string
	Interval time.Duration
	Action   func(T)
}

func (t Task[T]) Start(param T, stopChan <-chan struct{}, finishChan chan<- struct{}) {
	ticker := time.NewTicker(t.Interval)
	defer ticker.Stop()
	defer close(finishChan) // Fechando o canal de finalização quando a função Start terminar

	// Sem usar o for range ticker.C, podemos usar um loop infinito com select para aguardar os ticks do ticker e também para permitir a interrupção da execução da tarefa.
	for {
		select {
		case <-stopChan:
			return
		case <-ticker.C:
			t.Action(param)
		}
	}
}

func NewTask[T any](name string, interval time.Duration, action func(T)) *Task[T] {
	return &Task[T]{
		Name:     name,
		Interval: interval,
		Action:   action,
	}
}

func main() {
	firstTask := NewTask("First Task", 2*time.Second, func(param string) {
		fmt.Println("Executing first task with param:", param)
	})

	firstStopChan := make(chan struct{})
	finalizeChan := make(chan struct{})
	go firstTask.Start("Hello", firstStopChan, finalizeChan)

	time.Sleep(10 * time.Second)

	close(firstStopChan) // Fechando o canal de parada para interromper a execução da primeira tarefa
	<-finalizeChan       // Aguardando a finalização da primeira tarefa
}
