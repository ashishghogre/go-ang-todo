package main

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

type IDatabaseClient interface {
	setup(dbName string)
	initialize()
	upsertItem(id string, item todoItem)
	getItem(id string) todoItem
	getItems() []todoItem
	updateMaxId(id string)
	getMaxId() string
	teardown()
}

type DatabaseClient struct {
	dbClient *bolt.DB
}

func (dc *DatabaseClient) setup(dbName string) {
	var err error
	dc.dbClient, err = bolt.Open(dbName, 0600, nil)
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
		_, err = tx.CreateBucketIfNotExists([]byte("MaxIdBucket"))
		if err != nil {
			return fmt.Errorf("Could not create bucket %s", err)
		}
		return nil
	})
}

func (dc *DatabaseClient) updateMaxId(id string) {
	dc.dbClient.Update(func(tx *bolt.Tx) error{
		bucket := tx.Bucket([]byte("MaxIdBucket"))
		err := bucket.Put([]byte("0"),[]byte(id))
		if err != nil {
			return fmt.Errorf("Could not read bucket %s", err)
		}
		return nil
	})
}

func (dc *DatabaseClient) getMaxId() string {
	var id string
	dc.dbClient.View(func(tx *bolt.Tx) error{
		bucket := tx.Bucket([]byte("MaxIdBucket"))
		id = string(bucket.Get([]byte("0")))
		return nil
	})
	return id
}

func (dc *DatabaseClient) upsertItem(id string, item todoItem) {
	dc.dbClient.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("TodoItemsBucket"))
		value, err := json.Marshal(item)
		if err != nil {
			return fmt.Errorf("Could not parse item %s", err)
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

func (dc *DatabaseClient) getItems() []todoItem{
	var items []todoItem
	dc.dbClient.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("TodoItemsBucket"))
		bucket.ForEach(func(key []byte,value []byte) error{
			var item todoItem
			json.Unmarshal(value, &item)
			items = append(items,item)
			return nil
		})
		return nil
	})
	return items
}

func (dc *DatabaseClient) teardown() {
	dc.dbClient.Close()
}