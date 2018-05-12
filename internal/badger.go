package internal

import "github.com/dgraph-io/badger"

type KeyValueStore interface {
	Set(key string, value []byte) error
	Get(key string) ([]byte, error)
}

type store struct {
	settings Settings
	*badger.DB
}

func NewKeyValueStore(s Settings) KeyValueStore {
	opts := badger.DefaultOptions
	opts.Dir = "/tmp/badger"
	opts.ValueDir = "/tmp/badger"
	db, err := badger.Open(opts)

	if err != nil {
		panic(err)
	}

	return &store{settings: s, DB: db}
}

func (s *store) Set(key string, value []byte) error {
	return s.DB.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), value)
	})
}

func (s *store) Get(key string) ([]byte, error) {
	var value []byte
	err := s.DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))

		if err != nil {
			return err
		}

		value, err = item.Value()

		if err != nil {
			return err
		}

		return nil
	})

	return value, err
}

func (s *store) Close() {
	s.DB.Close()
}