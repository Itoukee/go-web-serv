package services

import (
	"example-api/myself/internal/domain/models"
	"example-api/myself/internal/infra/persistence"
)

type AlbumService struct {
	repPgsql persistence.AlbumRepositoryPgsql
}

func NewAlbumService(repPgsql *persistence.AlbumRepositoryPgsql) *AlbumService {
	return &AlbumService{repPgsql: *repPgsql}
}

func (s *AlbumService) CreateAlbum(title string, artist string, price float64) (*models.Album, error) {
	album := &models.Album{
		Title:  title,
		Artist: artist,
		Price:  price,
	}

	err := s.repPgsql.Create(album)

	if err != nil {
		return nil, err
	}

	return album, nil

}
