package main

import (
	"example-api/myself/internal/app/services"
	"example-api/myself/internal/infra/http_handler"
	"example-api/myself/internal/infra/persistence"
	"log"
	"net/http"
)

func main() {
	db, err := persistence.NewPostgresDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()

	albumRepo := persistence.NewAlbumRepoPgsql(db)
	albumService := services.NewAlbumService(albumRepo)
	albumHandler := http_handler.CreateAlbumHandler(albumService)

	mux := http.NewServeMux()
	mux.HandleFunc("/albums", albumHandler.CreateAlbum)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
