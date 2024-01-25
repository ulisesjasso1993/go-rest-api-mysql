package queries

import (
	"database/sql"
	"fmt"

	"connector"
)

type Album struct {
	Id     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func startConnection() *sql.DB {
	return connector.GetConnection()
}

func GetAlbums() ([]Album, error) {
	var albums []Album
	conn := startConnection()

	rows, err := conn.Query("Select * from album")
	if err != nil {
		return nil, fmt.Errorf("error retrieving data %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var album Album
		if err := rows.Scan(&album.Id, &album.Title, &album.Artist, &album.Price); err != nil {
			return nil, fmt.Errorf("error reading data %v", err)
		}

		albums = append(albums, album)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error on data %v", err)
	}

	return albums, nil
}
