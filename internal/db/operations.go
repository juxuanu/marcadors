package db

import "marcadors/internal/entities"

type BookmarksOperations interface {
	GetBookmarks() ([]entities.Bookmark, error)
	AddBookmark(bookmark entities.Bookmark) error
}
