package main

import (
	"fmt"
	"sync"
	"time"
)

// Cache guarda valores por chave. Como um map não permite leitura e escrita
// concorrentes com segurança, os acessos feitos por Get e Set usam o mesmo mutex.
type Cache struct {
	data map[string]any
	// RWMutex permite vários leitores ao mesmo tempo ou um único escritor.
	// Leitura: RLock() -> RUnlock(). Escrita: Lock() -> Unlock().
	mu sync.RWMutex
}

// NewCache inicializa o map para que ele possa receber valores.
// O RWMutex já está pronto para uso com seu valor zero.
func NewCache() *Cache {
	return &Cache{
		data: make(map[string]any),
	}
}

func (c *Cache) Set(key string, value any) {
	logf("Set: tentando escrever")
	// Lock espera os leitores e qualquer outro escritor liberarem o mutex.
	// Enquanto houver um escritor aguardando, novas chamadas a RLock bloqueiam
	// até que esse escritor consiga escrever e libere o mutex.
	c.mu.Lock()
	logf("Set: bloqueado ")
	// Simula uma escrita demorada, mantendo o bloqueio durante esse período.
	time.Sleep(300 * time.Millisecond)

	c.data[key] = value

	// Libera o acesso exclusivo, permitindo que outras operações prossigam.
	c.mu.Unlock()
	logf("Set: desbloqueado ")
}

func (c *Cache) Get(key string) (any, bool) {
	logf("Get: tentando ler")
	// RLock compartilha o bloqueio com outros leitores, mas espera se houver
	// um escritor com o bloqueio exclusivo ou aguardando para adquiri-lo.
	c.mu.RLock()
	logf("Get: bloqueado ")

	// Simula uma leitura demorada para tornar a sobreposição dos leitores visível.
	time.Sleep(100 * time.Millisecond)
	// ok indica se a chave existe; result contém o valor encontrado.
	result, ok := c.data[key]

	// Libera apenas o bloqueio deste leitor. Um escritor precisa esperar
	// até que todos os leitores que já adquiriram o bloqueio o liberem.
	c.mu.RUnlock()
	logf("Get: desbloqueado ")
	return result, ok
}

// Marca o início da demonstração para comparar os instantes dos eventos.
var start = time.Now()

// Registra o tempo decorrido em milissegundos junto da mensagem de cada operação.
func logf(msg string) {
	fmt.Printf("[4%d ms] - %s \n", time.Since(start).Milliseconds(), msg)
}

func main() {
	var (
		cache = NewCache()
		wg    sync.WaitGroup
	)

	// O WaitGroup conta as quatro goroutines e mantém main ativa até todas terminarem.
	wg.Add(4)
	// Aqui o acesso direto é seguro porque nenhuma goroutine foi iniciada ainda.
	// Depois que elas começam, os acessos ao map passam por Get e Set.
	cache.data["foo"] = "bar"

	// O primeiro leitor tenta começar imediatamente e mantém RLock por cerca de 100 ms.
	go func() {
		// Done reduz o contador ao sair da goroutine, mesmo se houver um retorno antecipado.
		defer wg.Done()
		logf("Goroutine 1: tentando ler\n")
		_, _ = cache.Get("foo")
		logf("Goroutine 1: leitura concluída\n")
	}()

	// O segundo leitor tenta começar após 10 ms. Na sequência esperada,
	// ele lê junto com o primeiro, pois ambos usam o bloqueio compartilhado.
	go func() {
		defer wg.Done()
		time.Sleep(10 * time.Millisecond)
		logf("Goroutine 2: tentando ler\n")
		_, _ = cache.Get("foo")
		logf("Goroutine 2: leitura concluída\n")
	}()

	// O escritor tenta começar após 30 ms e espera os dois leitores terminarem.
	// Ao adquirir Lock, altera "foo" de "bar" para "baz" com acesso exclusivo.
	go func() {
		defer wg.Done()
		time.Sleep(30 * time.Millisecond)
		logf("Goroutine 2: tentando escrever\n")
		cache.Set("foo", "baz")
		logf("Goroutine 2: escrita concluída\n")
	}()

	// O último leitor tenta começar após 300 ms, quando a escrita deve estar em curso.
	// Seu RLock espera o Unlock do escritor; depois, Get encontra o valor "baz".
	go func() {
		defer wg.Done()
		time.Sleep(300 * time.Millisecond)
		logf("Goroutine 4: tentando ler\n")
		_, _ = cache.Get("foo")
		logf("Goroutine 4: leitura concluída\n")
	}()

	// Os Sleep ajudam a demonstrar a ordem esperada, mas não garantem o agendamento.
	// A segurança do map depende do RWMutex; a conclusão das goroutines, do WaitGroup.
	wg.Wait()

	logf("Final \n")
}
