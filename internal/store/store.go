package store

import bolt "go.etcd.io/bbolt"

type Store struct {
	db *bolt.DB
}

const URLS = "urls"

func New(path string) (*Store, error) {

	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(URLS))
		return err
	})
	if err != nil {
		db.Close()
		return nil, err
	}

	return &Store{db}, nil
}
