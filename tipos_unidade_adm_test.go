package sipope

import (
	"context"
	"testing"
	"time"
)

func TestGetTipoUnidadeAdministrativa(t *testing.T) {
	client := newTestClient()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.GetTipoUnidadeAdministrativa(ctx, 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
