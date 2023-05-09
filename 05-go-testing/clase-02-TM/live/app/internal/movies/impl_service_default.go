package movies

import "app/internal/logger"

// constructor
func NewServiceDefault(repo Repository, logger logger.Logger) *ServiceDefault {
	return &ServiceDefault{repo, logger}
}

// ServiceDefault implements the Service interface.
type ServiceDefault struct {
	repo   Repository
	logger logger.Logger
}

func (s *ServiceDefault) GetMovie(id int) (mv Movie, err error) {
	defer func() {
		if err != nil {
			s.logger.Log("Error: " + err.Error())
		}
	}()

	// get movie from repository
	mv, err = s.repo.GetMovie(id)
	if err != nil {
		if err == ErrRepoNotFound {
			err = ErrServiceNotFound
		} else {
			err = ErrServiceInternal
		}
		return
	}

	// log
	return
}