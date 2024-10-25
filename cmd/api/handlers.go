package main

import (
	"encoding/json"
	"fmt"
	"marcadors/internal/db"
	"marcadors/internal/entities"
	"net/http"
	"os"
	"time"
)

type pathAndHandler struct {
	path     string
	function http.HandlerFunc
}

var ApiHandlers = []pathAndHandler{
	{
		path: "GET /bookmarks",
		function: func(w http.ResponseWriter, r *http.Request) {
			bookmarks, err := db.GetBookmarks()
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "[500] %s\n", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			err = json.NewEncoder(w).Encode(bookmarks)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "[500] %s\n", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		},
	},
	{
		path: "POST /bookmarks",
		function: func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			b, err := entities.NewBookmark(entities.Bookmark{
				ID:   time.Now().String(),
				URI:  "https://example.com",
				Name: "Some Example",
				Type: "uri",
			})
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "[400] %s\n", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			err = db.AddBookmark(*b)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "[507] %s\n", err)
				w.WriteHeader(http.StatusInsufficientStorage)
				return
			}
			w.WriteHeader(http.StatusCreated)
		},
	},
}
