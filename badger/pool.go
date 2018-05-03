package badger

import (
	"log"
	"runtime"

	"github.com/dgraph-io/badger"
)

var (
	// Conn db Conn
	Conn chan int
	// Pool db pool
	Pool *pool
)

// pool pool
type pool struct {
	*badger.DB
}

func init() {
	path := "./data"
	opts := badger.DefaultOptions
	opts.Dir = path
	opts.ValueDir = path
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}

	cap := runtime.NumCPU()
	Conn = make(chan int, cap)
	for i := 0; i < cap; i++ {
		Conn <- 1
	}
	Pool = &pool{db}
}
