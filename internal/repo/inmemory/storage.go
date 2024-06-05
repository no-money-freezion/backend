package inmemory

import (
	"sync"
)

type Repository struct {
	storage map[string]string
	lock    sync.RWMutex // if we wanna faster inmemory cache we can use sync.Map or sync.RWMutex `s pool
}

func NewRepository() *Repository {
	return &Repository{
		storage: make(map[string]string),
	}
}
