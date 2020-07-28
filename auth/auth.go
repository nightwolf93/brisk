package auth

import (
	"os"

	"github.com/nightwolf93/brisk/storage"
)

var masterPair *storage.ClientPairCredentials = nil

// InitMasterPair init the master credentials
func InitMasterPair() {
	masterPair = &storage.ClientPairCredentials{
		ClientID:     os.Getenv("MASTER_CLIENT_ID"),
		ClientSecret: os.Getenv("MASTER_CLIENT_SECRET"),
	}
}

// VerifyPair verify a credential pair
func VerifyPair(pair *storage.ClientPairCredentials) bool {
	if IsMasterCredentials(pair) {
		return true
	}
	secret := storage.FindSecretByID(pair.ClientID)
	if secret == "" {
		return false
	}
	return secret == pair.ClientSecret
}

// IsMasterCredentials check if the pair is a master pair
func IsMasterCredentials(pair *storage.ClientPairCredentials) bool {
	return pair.ClientID == os.Getenv("MASTER_CLIENT_ID") && pair.ClientSecret == os.Getenv("MASTER_CLIENT_SECRET")
}
