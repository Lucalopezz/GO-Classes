package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	// Criar um servidor HTTP com 2 rotas
	// Uma rota simples que retorna um ok
	// outra que retorna um ok após 5 secs
	//
	// Graceful shutdown: quando o servidor receber um sinal de interrupção (Ctrl+C),
	// ele deve esperar 10 segundos para finalizar as requisições pendentes antes de encerrar o servidor.

	e := echo.New()

	e.GET("/slow", func(c echo.Context) error {
		time.Sleep(5 * time.Second)
		return c.String(200, "ok 5 secs")
	})

	e.GET("/fast", func(c echo.Context) error {
		return c.String(200, "ok fast")
	})

	// Iniciar o servidor HTTP
	go func() {
		// só logar o erro se for diferente de http.ErrServerClosed, que é o erro retornado quando o servidor é encerrado
		// apenas logar se der erro ao iniciar o servidor
		if err := e.Start(":8080"); err != http.ErrServerClosed {
			log.Fatalf("Erro iniciando o servidor %s", err.Error())
		}
	}()
	fmt.Println("Servidor iniciado na porta 8080")

	quit := make(chan os.Signal, 1)
	// Registrar o canal para receber sinais de interrupção (Ctrl+C)
	// então quando o usuário pressionar Ctrl+C, o canal quit receberá um sinal de interrupção
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit // Bloquear até receber um sinal de interrupção
	fmt.Println("Servidor encerrando...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatalf("Erro encerrando o servidor %s", err.Error())
	} // Encerrar o servidor com um contexto de timeout de 10 segundos
}
