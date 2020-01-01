package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mattvella07/watchlist-server/server/conn"
)

type watchlist struct {
	ID      int    `json:"id"`
	OwnerID int    `json:"owner_id"`
	Name    string `json:"name"`
	Items   []item `json:"items"`
}

type item struct {
	ID          int    `json:"id"`
	ItemType    string `json:"item_type"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ReleaseDate string `json:"release_date"`
	Rating      string `json:"rating"`
	Genre       string `json:"genre"`
	Watched     bool   `json:"watched"`
}

// GetWatchlist gets the current user's watchlists
func GetWatchlist(rw http.ResponseWriter, req *http.Request) {
	query := `SELECT id, owner_id, name
		FROM watchlist
		WHERE owner_id = $1`

	watchlistRows, err := conn.DB.Query(query, 1)
	if err != nil {
		log.Printf("DB error: %s\n", err)
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Unable to communicate with database"))
		return
	}
	defer watchlistRows.Close()

	watchlists := []watchlist{}
	for watchlistRows.Next() {
		w := watchlist{}
		err = watchlistRows.Scan(&w.ID, &w.OwnerID, &w.Name)
		if err != nil {
			log.Printf("DB error: %s\n", err)
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("Error reading from database"))
			return
		}

		query = `SELECT id, title, description, release_date, rating, genre, watched,
		(SELECT name FROM item_types it WHERE it.id = i.item_type) as item_type
		FROM items i
		WHERE watchlist_id = $1`

		itemRows, err := conn.DB.Query(query, w.ID)
		if err != nil {
			log.Printf("DB error: %s\n", err)
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("Unable to communicate with database"))
			return
		}
		defer itemRows.Close()

		for itemRows.Next() {
			i := item{}
			itemRows.Scan(&i.ID, &i.Title, &i.Description, &i.ReleaseDate, &i.Rating, &i.Genre, &i.Watched, &i.ItemType)
			if err != nil {
				log.Printf("DB error: %s\n", err)
				rw.WriteHeader(http.StatusBadRequest)
				rw.Write([]byte("Error reading from database"))
				return
			}

			w.Items = append(w.Items, i)
		}

		watchlists = append(watchlists, w)
	}

	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(watchlists)
}
