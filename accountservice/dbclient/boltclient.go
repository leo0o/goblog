package dbclient

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/leo0o/goblog/accountservice/model"
	"log"
	"strconv"
)

type IBoltClient interface {
	OpenBoltDb()
	QueryAccount(string) (model.Account, error)
	Seed()
}

type BoltClient struct {
	boltDb *bolt.DB
}

func (bc *BoltClient) OpenBoltDb() {
	var err error
	bc.boltDb, err = bolt.Open("account.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (bc *BoltClient) QueryAccount(id string) (model.Account, error) {
	var account model.Account
	err := bc.boltDb.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("AccountBucket"))
		value := b.Get([]byte(id))
		err := json.Unmarshal(value, &account)
		return err
	})

	if err != nil {
		return model.Account{}, err
	}
	return account, nil
}

func (bc *BoltClient) Seed() {
	bc.initBucket()
	bc.seedAccounts()
}

func (bc *BoltClient) initBucket() {
	bc.boltDb.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("AccountBucket"))
		if err != nil {
			return fmt.Errorf("create bucket failed: %s", err)
		}
		return nil
	})
}

func (bc *BoltClient) seedAccounts() {
	total := 100
	for i := 0; i < total; i++ {
		key := strconv.Itoa(10000 + i)

		account := model.Account{
			Id:   key,
			Name: "person" + key,
		}
		value, _ := json.Marshal(account)
		bc.boltDb.Update(func(tx *bolt.Tx) error {
			return tx.Bucket([]byte("AccountBucket")).Put([]byte(key), value)
		})
	}
	log.Printf("seed %d fake accounts...", total)
}
