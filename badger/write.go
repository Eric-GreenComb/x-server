package badger

import (
	"time"
)

// Write write
type Write struct{}

// NewWrite new write
func NewWrite() *Write {
	return &Write{}
}

// List list
func (w *Write) List(args map[string][]byte) error {
	<-Conn
	defer func() {
		Conn <- 1
	}()
	txn := Pool.DB.NewTransaction(true)
	defer txn.Discard()
	for k, v := range args {
		err := txn.Set([]byte(k), v)
		if err != nil {
			return err
		}
	}
	return txn.Commit(nil)
}

// Set Set item
func (w *Write) Set(k string, v []byte) error {
	<-Conn
	defer func() {
		Conn <- 1
	}()
	txn := Pool.DB.NewTransaction(true)
	defer txn.Discard()
	err := txn.Set([]byte(k), v)
	if err != nil {
		return err
	}
	return txn.Commit(nil)
}

// SetWithTTL SetWithTTL
func (w *Write) SetWithTTL(k string, v []byte, dur time.Duration) error {
	<-Conn
	defer func() {
		Conn <- 1
	}()
	txn := Pool.DB.NewTransaction(true)
	defer txn.Discard()
	err := txn.SetWithTTL([]byte(k), v, dur)
	if err != nil {
		return err
	}
	return txn.Commit(nil)
}
