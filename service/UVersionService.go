package service

import (
	"github.com/cs161079/godbLib/mapper"
	"github.com/cs161079/godbLib/repository"
)

type UVersionService interface {
}

type uVersionsService struct {
	Repo   repository.UVersionRepository
	Rest   RestService
	Mapper mapper.UversionMapper
}

func (s uVersionsService) GetUversionWeb() {
	response := s.Rest.OasaRequestApi00("getUversions", nil)
	if response.Error != nil {
		// Εδώ προκείπτει Error από το Request
		//Κάτι πρέπει να κάνουμε.
	}
	// var arrVersion []models.UVersions = make([]models.UVersions, 0)
	// for index, int := range response.Data.([]interface{}) {
	// 	uVersion := s.Mapper.OasaToUVersions(s.Mapper.GeneralUVersions(response.Data))

	// }

}
