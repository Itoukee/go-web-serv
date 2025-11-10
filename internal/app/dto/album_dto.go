package dto

import "../../../GO-WEB-SERV/internal/app/dto/time"

type CreateAlbumDto struct {
	Title  string  `json:"title" validate:"required"`
	Artist string  `json:"artist" validate:"required"`
	Price  float64 `json:"price" validate:"required"`
}

type CreateAlbumResponse struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Artist    string    `json:"artist"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}
