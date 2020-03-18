package db

import (
	"github.com/jinzhu/gorm"
)

type PostgresStore interface {
	DB() *gorm.DB
	Begin() (PostgresStore, error)
	Commit() error
	Rollback() error
	Close()
}

type postgresStore struct {
	db *gorm.DB
}

func New() (*postgresStore, error) {
	db, err := gorm.Open("postgres", "postgres://username:password@127.0.0.1:5432/yourDB")
	if err != nil {
		return nil, err
	}

	db.DB().SetMaxIdleConns(0)

	store := NewStore(db)
	return store, nil
}

func NewStore(db *gorm.DB) *postgresStore {
	store := &postgresStore{db: db}
	return store
}

func (s *postgresStore) DB() *gorm.DB {
	return s.db
}

func (s *postgresStore) Begin() (PostgresStore, error) {
	tx := s.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	store := NewStore(tx)
	return store, nil
}

func (s *postgresStore) Commit() error {
	c := s.db.Commit()
	if c.Error != nil {
		return c.Error
	}
	return nil
}

func (s *postgresStore) Rollback() error {
	c := s.db.Rollback()
	if c.Error != nil {
		return c.Error
	}
	return nil
}

func (s *postgresStore) Close() {
	s.db.Close()
}
