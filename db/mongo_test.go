package db

import (
	"context"
	"testing"
)

func TestConnectToMongo(t *testing.T) {
	t.Run("invalid uri", func(t *testing.T) {
		ctx := context.Background()
		_, err := ConnectToMongo(ctx, "mongodb://invalid:27017")
		if err == nil {
			t.Fatal("Expected error for invalid URI, got nil")
		}
	})
}
