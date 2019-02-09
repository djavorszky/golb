package backend

import "golb/server/model"

type Store interface {
	GetAll() []model.BlogEntry
}

type MockStore struct {
	Entries []model.BlogEntry
}

func (mb MockStore) GetAll() []model.BlogEntry {
	return mb.Entries
}
