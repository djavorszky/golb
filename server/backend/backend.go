package backend

import "golb/server/model"

type Store interface {
	GetAll() []model.BlogEntry
}

type MockStore struct{}

func (mb MockStore) GetAll() []model.BlogEntry {
	return []model.BlogEntry{
		{"hello title", "hello content"},
		{"title two", "content two"},
	}
}
