package persistence

import (
	"database/sql"
	"example-api/myself/internal/domain/models"
)

type AlbumRepositoryPgsql struct {
	db *sql.DB
}

func initAlbums(db *sql.DB) error {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS albums (
            id SERIAL PRIMARY KEY,
            title VARCHAR(255) NOT NULL,
            artist VARCHAR(255) NOT NULL,
            price DECIMAL(10, 2) NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );
    `)
	return err
}

func NewAlbumRepoPgsql(db *sql.DB) *AlbumRepositoryPgsql {
	initAlbums(db)
	return &AlbumRepositoryPgsql{db: db}
}

func (r *AlbumRepositoryPgsql) Create(album *models.Album) error {
	_, err := r.db.Exec("INSERT INTO albums (title,artist,price) VALUES ($1, $2, $3)",
		album.Title, album.Artist, album.Price)
	return err
}
