package backend

import (
	"golb/server/model"
	"time"
)

type Store interface {
	GetAll() []model.BlogEntry
}

type MockStore struct {
	Entries []model.BlogEntry
}

func (mb MockStore) GetAll() []model.BlogEntry {
	return mb.Entries
}

func NewMockStore() MockStore {
	loc, _ := time.LoadLocation("Europe/Budapest")

	cdate1 := time.Unix(1405544146, 0).In(loc)
	udate1 := cdate1
	pdate1 := cdate1

	return MockStore{
		Entries: []model.BlogEntry{
			{
				ID:          "one",
				Author:      "Joe",
				Tags:        []string{"first tag", "second tag"},
				Categories:  []string{"first category", "second category"},
				Title:       "first title",
				Content:     "first content",
				CreateDate:  cdate1,
				UpdateDate:  udate1,
				PublishDate: pdate1,
			},
			{
				ID:          "two",
				Author:      "Joe",
				Tags:        []string{"first tag", "third tag"},
				Categories:  []string{"first category", "third category"},
				Title:       "second title",
				Content:     "second content",
				CreateDate:  cdate1.AddDate(0, 0, 1),
				UpdateDate:  udate1.AddDate(0, 0, 1),
				PublishDate: pdate1.AddDate(0, 0, 1),
			},
		},
	}
}
