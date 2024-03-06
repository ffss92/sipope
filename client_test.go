package sipope

import "os"

// Cria um client de teste.
func newTestClient() *Client {
	apiKey := os.Getenv("SIPOPE_API_KEY")
	return NewClient(apiKey, WithHomologacaoURL())
}
