package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func fazerRequest(ctx context.Context, url string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	// time.Sleep(3 * time.Second) // Simula um atraso na requisição
	// como o tempo limite do contexto é de 2 segundos, a requisição será cancelada antes de ser concluída
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	url := "https://jsonplaceholder.typicode.com/posts/1"
	data, err := fazerRequest(ctx, url)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("Request cancelada devido ao tempo limite")
			return
		}

		if ctx.Err() == context.Canceled {
			fmt.Println("Request cancelada manualmente")
			return
		}

		fmt.Println("Erro ao fazer request:", err)
	}
	fmt.Println("Dados recebidos:", data)
}
