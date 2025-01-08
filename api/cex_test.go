package api_test

import (
	"testing"

	"cryptogo.com/crypto/api"
)

func TestApiCall(t *testing.T) {
	_, err := api.GetRate("")

	if err == nil {
		t.Error("Error")
	}
}