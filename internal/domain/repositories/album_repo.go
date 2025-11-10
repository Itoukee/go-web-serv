package repositories

import "../../../GO-WEB-SERV/internal/domain/models"

type AlbumRepository interface {
	Create(album *models.Album) error
	FindById(album_id *models.Album) (*models.Album, error)
	ListAll() []*models.Album
}
