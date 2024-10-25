package db

import (
	"marcadors/internal/entities"
)

type Operations interface {
	GetBookmarks() ([]entities.Bookmark, error)
	AddBookmark(bookmark entities.Bookmark) error
}

type Db struct {
	bookmarks []entities.Bookmark
}

func BuildDbClient() (Db, error) {
	bookmark1, err := entities.NewBookmark("uri", "1", "https://diccionari.icarns.xyz", "DCM")
	if err != nil {
		return Db{}, err
	}
	return Db{bookmarks: []entities.Bookmark{bookmark1}}, nil
}

func (db *Db) GetBookmarks() ([]entities.Bookmark, error) {
	return db.bookmarks, nil
}

func (db *Db) AddBookmark(bookmark entities.Bookmark) error {
	db.bookmarks = append(db.bookmarks, bookmark)
	return nil
}
