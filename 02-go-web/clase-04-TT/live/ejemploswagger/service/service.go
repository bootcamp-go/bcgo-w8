package service

import (
	"ejemploswagger/repo"
	"ejemploswagger/model"
)

func GetAlbums() []model.Album {
	return repo.GetAlbums()
}

func Create(album model.Album) []model.Album{
	return repo.Create(album)
}