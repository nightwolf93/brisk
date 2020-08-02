package auth

import (
	"github.com/nightwolf93/brisk/storage"
	"os"
	"testing"
)

func TestVerifyPair(t *testing.T) {
	os.Setenv("MASTER_CLIENT_ID", "master_test")
	os.Setenv("MASTER_CLIENT_SECRET", "secret_test")
	InitMasterPair()
	pair := &storage.ClientPairCredentials{
		ClientID:     os.Getenv("MASTER_CLIENT_ID"),
		ClientSecret: os.Getenv("MASTER_CLIENT_SECRET"),
	}
	if !VerifyPair(pair) {
		t.Errorf("verify pair failed")
	}
}
