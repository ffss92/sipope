package sipope

import (
	"context"
	"testing"
	"time"
)

func TestListRisps(t *testing.T) {
	client := newTestClient()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.ListRisps(ctx)
	if err != nil {
		t.Errorf("unexpected error listing RISPs: %v", err)
	}
}

func TestGetRisp(t *testing.T) {
	client := newTestClient()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.GetRisp(ctx, 1)
	if err != nil {
		t.Errorf("unexpected error getting RISP: %v", err)
	}
}
