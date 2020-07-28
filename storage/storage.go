package storage

import (
	"fmt"
	"log"

	badger "github.com/dgraph-io/badger/v2"
)

var db *badger.DB = nil

// Open the database file
func Open() {
	opts := badger.DefaultOptions("./localdb")
	opts.Logger = nil
	conn, err := badger.Open(opts)
	if err != nil {
		log.Fatal("can't open the database")
	}
	db = conn
}

// SavePair save a new pair in the db
func SavePair(pair *ClientPairCredentials) error {
	err := db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(fmt.Sprintf("pair_%s", pair.ClientID)), []byte(pair.ClientSecret))
		return err
	})
	return err
}

func FindSecretByID(id string) string {
	var secret string = ""
	db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(fmt.Sprintf("pair_%s", id)))
		if err != nil {
			return err
		}
		secret = item.String()
		return err
	})
	return secret
}
