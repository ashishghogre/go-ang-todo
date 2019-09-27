package main

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

type IDatabaseClient interface {
	setup()
	initialize()
	upsertItem(id string, item todoItem)
	getItem(id string) todoItem
	teardown()
}

type DatabaseClient struct {
	dbClient *bolt.DB
}

func (dc *DatabaseClient) setup() {
	var err error
	dc.dbClient, err = bolt.Open("todo.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (dc *DatabaseClient) initialize() {
	dc.dbClient.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("TodoItemsBucket"))
		if err != nil {
			return fmt.Errorf("Could not create bucket %s", err)
		}
		return nil
	})
}

func (dc *DatabaseClient) upsertItem(id string, item todoItem) {
	dc.dbClient.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("TodoItemsBucket"))
		value, err := json.Marshal(item)
		if err != nil {
			return fmt.Errorf("Could not parse item %s", item)
		}
		err = bucket.Put([]byte(id), value)
		if err != nil {
			return fmt.Errorf("Could not read bucket %s", err)
		}
		return nil
	})
}

func (dc *DatabaseClient) getItem(id string) todoItem{
	var item todoItem
	dc.dbClient.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("TodoItemsBucket"))
		json.Unmarshal(bucket.Get([]byte(id)), &item)
		return nil
	})
	return item
}

func (dc *DatabaseClient) teardown() {
	dc.dbClient.Close()
}