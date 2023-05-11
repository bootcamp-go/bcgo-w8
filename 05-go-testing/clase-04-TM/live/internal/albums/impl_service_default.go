package albums

import (
	"errors"
	"fmt"
)

func NewServiceDefault(storage Storage) *ServiceDefault {
	return &ServiceDefault{storage: storage}
}

type ServiceDefault struct {
	storage Storage
	// other external api's
	// ...
}

func (sv *ServiceDefault) GetAlbums() (al []*Album, err error) {
	al, err = sv.storage.GetAlbums()
	if err != nil {
		err = ErrServiceInternal
		return
	}
	return
}

func (sv *ServiceDefault) Create(a *Album) (err error) {
	// validate album
	if len(a.Title) < 5 {
		err = fmt.Errorf("%w: title", ErrServiceInvalidAlbum)
		return
	}
	if a.Year < 1900 {
		err = fmt.Errorf("%w: year", ErrServiceInvalidAlbum)
		return
	}

	// create album
	err = sv.storage.Create(a)
	if err != nil {
		switch {
		case errors.Is(err, ErrStorageAttributeNotUnique):
			err = ErrServiceAttributeNotUnique
		default:
			err = ErrServiceInternal
		}
		return
	}

	return
}