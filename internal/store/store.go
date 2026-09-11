package store

import (
	"errors"

	bolt "go.etcd.io/bbolt"
)

type Store struct {
	db *bolt.DB
}

const URLS = "urls"

var (
	ErrNotFound      = errors.New("url not found")
	ErrAlreadyExists = errors.New("short code already exists")
)

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

func (s *Store) GetRedirectURL(shortCode string) (string, error) {

	var url string
	err := s.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(URLS))
		if bucket == nil {
			return errors.New("bucket not found")
		}
		result := bucket.Get([]byte(shortCode))
		if result == nil {
			return ErrNotFound
		}
		url = string(result)
		return nil
	})
	return url, err
}

func (s *Store) PutRedirectURL(shortCode, url string) error {

	return s.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(URLS))
		if bucket == nil {
			return errors.New("bucket not found")
		}
		result := bucket.Get([]byte(shortCode))
		if result != nil {
			return ErrAlreadyExists
		}
		return bucket.Put([]byte(shortCode), []byte(url))
	})
}
