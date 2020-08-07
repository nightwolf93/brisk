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

// FindSecretByID find a secret by the id
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
		e := badger.NewEntry([]byte(fmt.Sprintf("link_%s", link.Slug)), []byte(string(payload))).WithTTL(duration)
		err := txn.SetEntry(e)
		return err
	})
	return err
}

// SaveVisitorEntry save a new visitor entry
func SaveVisitorEntry(link *Link, visitor *VisitorEntry) error {
	err := db.Update(func(txn *badger.Txn) error {
		payload, _ := json.Marshal(visitor)
		duration := time.Millisecond * time.Duration(link.TTL)
		e := badger.NewEntry([]byte(fmt.Sprintf("visitor/link_%s/%s", link.Slug, visitor.Hash)), []byte(string(payload))).WithTTL(duration)
		err := txn.SetEntry(e)
		return err
	})
	return err
}

// FindLink find a link by the slug
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

// DeleteLink delete a link
func DeleteLink(slug string) error {
	err := db.Update(func(txn *badger.Txn) error {
		err := txn.Delete([]byte(fmt.Sprintf("link_%s", slug)))
		return err
	})
	return err
}

// FindAllLinks find all links stored
func FindAllLinks() ([]*Link, error) {
	links := []*Link{}
	err := db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		prefix := []byte("link_")
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			err := item.Value(func(v []byte) error {
				var link *Link = nil
				json.Unmarshal(v, &link)
				links = append(links, link)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	return links, err
}

// FindVisitorsForLink find all visitors entry for the link
func FindVisitorsForLink(link *Link) ([]*VisitorEntry, error) {
	visitors := []*VisitorEntry{}
	err := db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		prefix := []byte(fmt.Sprintf("visitor/link_%s", link.Slug))
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			err := item.Value(func(v []byte) error {
				var visitor *VisitorEntry = nil
				json.Unmarshal(v, &visitor)
				visitors = append(visitors, visitor)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	return visitors, err
}
