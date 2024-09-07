package service

import (
	"github.com/cs161079/godbLib/mapper"
	"github.com/cs161079/godbLib/repository"
)

type RouteService interface {
}

type routeService struct {
	Repo      repository.RouteRepository
	Mapper    mapper.RouteMapper
	R01Mapper mapper.Route01Mapper
}
