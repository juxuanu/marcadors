package entities

import (
	"fmt"
)

type BookmarkType = string

const (
	folderType BookmarkType = "folder"
	uriType    BookmarkType = "uri"
)

type Bookmark struct {
	ID       string       `json:"id"`
	URI      string       `json:"uri"`
	Name     string       `json:"name"`
	Type     BookmarkType `json:"type"`
	ParentId *string      `json:"parentId,omitempty"`
}

func NewBookmark(bType BookmarkType, id, uri, name string, parentId string) (Bookmark, error) {
	if !(bType == folderType || bType == uriType) || len(id) == 0 || len(name) == 0 || len(uri) == 0 {
		return Bookmark{}, fmt.Errorf("new bookmark format error")
	}

	return Bookmark{
		ID:       id,
		URI:      uri,
		Name:     name,
		Type:     bType,
		ParentId: &parentId,
	}, nil
}
