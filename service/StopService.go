package service

import "github.com/cs161079/godbLib/repository"

type StopService interface {
}

type stopService struct {
	Repo repository.StopRepository
}
