package models

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/sergi/go-diff/diffmatchpatch"
)

// DiffDB is the implementation of the storage interface for
// Diff.
type DiffDB struct {
	Db *bolt.DB
}

// NewDiffDB creates a new storage type.
func NewDiffDB(db *bolt.DB) *DiffDB {
	return &DiffDB{Db: db}
}

// DB returns the underlying database.
func (m *DiffDB) DB() interface{} {
	return m.Db
}

// DiffStorage represents the storage interface.
type DiffStorage interface {
	DB() interface{}
	Get(id int) ([]diffmatchpatch.Diff, error)
	Add(diffs []diffmatchpatch.Diff) error
}

// TableName overrides the table name settings in Bolt to force a specific table name
// in the database.
func (m *DiffDB) TableName() string {
	return "diffs"
}

// CRUD Functions.

// Get returns a single Diff as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *DiffDB) Get(id string) ([]diffmatchpatch.Diff, error) {
	res := []diffmatchpatch.Diff{}
	m.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Diff"))
		v := b.Get([]byte(id))
		res = DiffFromGob(v)
		return nil
	})
	return res, nil
}

// Add creates a new record.
func (m *DiffDB) Add(diff []diffmatchpatch.Diff) error {
	err := m.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Diff"))

		// Generate ID.
		id, _ := b.NextSequence()
		//model.ID = strconv.Itoa(int(id))

		buf := DiffToGob(diff)
		// Persist bytes to users bucket.
		return b.Put([]byte(strconv.Itoa(int(id))), buf)
	})

	return err
}

// DiffToGob binary encoder
func DiffToGob(m []diffmatchpatch.Diff) []byte {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	err := e.Encode(m)
	if err != nil {
		fmt.Println(`failed gob Encode`, err)
	}
	return b.Bytes()
	//return base64.StdEncoding.EncodeToString(b.Bytes())
}

// DiffFromGob binary decoder
func DiffFromGob(str []byte) []diffmatchpatch.Diff {
	m := []diffmatchpatch.Diff{}
	b := bytes.Buffer{}
	b.Write(str)
	d := gob.NewDecoder(&b)
	err := d.Decode(&m)
	if err != nil {
		fmt.Println(`failed gob Decode`, err)
	}
	return m
}
