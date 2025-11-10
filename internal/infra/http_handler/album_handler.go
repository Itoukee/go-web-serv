package http_handler

import (
	"encoding/json"
	"example-api/myself/internal/app/dto"
	"example-api/myself/internal/app/services"
	"net/http"
)

type AlbumHandler struct {
	albumService *services.AlbumService
}

func CreateAlbumHandler(albumService *services.AlbumService) *AlbumHandler {
	return &AlbumHandler{albumService: albumService}
}

func (h *AlbumHandler) CreateAlbum(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateAlbumDto
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	album, err := h.albumService.CreateAlbum(req.Title, req.Artist, req.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(album)

}
