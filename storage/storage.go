package storage

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

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

// SaveLink save a new link in the db
func SaveLink(link *Link) error {
	err := db.Update(func(txn *badger.Txn) error {
		payload, _ := json.Marshal(link)
		duration := time.Millisecond * time.Duration(link.TTL)
		e := badger.NewEntry([]byte(fmt.Sprintf("link_%s", link.Slug)), []byte(string(payload))).WithTTL(time.Hour).WithTTL(duration)
		err := txn.SetEntry(e)
		return err
	})
	return err
}

func FindLink(slug string) *Link {
	var link *Link = nil
	db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(fmt.Sprintf("link_%s", slug)))
		if err != nil {
			return err
		}
		item.Value(func(v []byte) error {
			err = json.Unmarshal(v, &link)
			return nil
		})
		return err
	})
	return link
}

func DeleteLink(slug string) error {
	err := db.Update(func(txn *badger.Txn) error {
		err := txn.Delete([]byte(fmt.Sprintf("link_%s", slug)))
		return err
	})
	return err
}

func FindAllLinks() []*Link {
	links := []*Link{}
	return links
}
