package service

import "github.com/cs161079/godbLib/repository"

type ScheduleService interface {
}

type scheduleService struct {
	Repo repository.ScheduleRepository
}
