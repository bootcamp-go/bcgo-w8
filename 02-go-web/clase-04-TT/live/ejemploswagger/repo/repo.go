package repo

import "ejemploswagger/model"

// base de datos del repositorio
var albums = []model.Album{
	{ID: "1", Title: "cualquiera", Artist: "cualquiera", Year: 2002},
}

func GetAlbums() []model.Album {
	return albums
}

func Create(album model.Album) []model.Album{
	albums = append(albums, album)
	return albums
}