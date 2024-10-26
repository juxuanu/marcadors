package db

import (
	"marcadors/internal/entities"
)

type Repository struct {
	bookmarks []entities.Bookmark
}

func Build() (Repository, error) {
	bookmark1, err := entities.NewBookmark("uri", "1", "https://diccionari.icarns.xyz", "DCM")
	bookmark2, err := entities.NewBookmark("uri", "1", "https://diccionari.icarns.xyz", "DCM")
	if err != nil {
		return Repository{}, err
	}
	return Repository{bookmarks: []entities.Bookmark{bookmark1, bookmark2}}, nil
}

func (repo *Repository) GetBookmarks() ([]entities.Bookmark, error) {
	return repo.bookmarks, nil
}

func (repo *Repository) AddBookmark(bookmark entities.Bookmark) error {
	repo.bookmarks = append(repo.bookmarks, bookmark)
	return nil
}
