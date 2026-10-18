package json

import "encoding/json"

type ConversorJSON struct{}

func (c ConversorJSON) Conversor(data string) (map[string]any, error) {
	var result map[string]any
	err := json.Unmarshal([]byte(data), &result)

	return result, err
}

func (c ConversorJSON) Formato() string {
	return "JSON"
}
