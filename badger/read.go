package badger

// Read read
type Read struct{}

// NewRead new read
func NewRead() *Read {
	return &Read{}
}

// List list
func (r *Read) List(args ...string) map[string][]byte {
	<-Conn
	defer func() {
		Conn <- 1
	}()
	txn := Pool.DB.NewTransaction(false)
	defer txn.Discard()
	data := map[string][]byte{}
	for _, k := range args {
		item, err := txn.Get([]byte(k))
		if err != nil {
			continue
		}
		v, err := item.Value()
		if err != nil {
			continue
		}
		data[k] = v
	}
	return data
}

// Get Get
func (r *Read) Get(k string) ([]byte, error) {
	<-Conn
	defer func() {
		Conn <- 1
	}()
	txn := Pool.DB.NewTransaction(false)
	defer txn.Discard()
	item, err := txn.Get([]byte(k))
	if err != nil {
		return nil, err
	}
	return item.Value()
}
