package mapper

import models "github.com/cs161079/godbLib/Models"

type Route01Mapper interface {
	RouteDetailDtoMapper(source any) models.Route01Oasa
	RouteDetailOasaToRouteDetailDto(source models.Route01Oasa) models.Route01
}

type route01Mapper struct {
}

func (m route01Mapper) RouteDetailDtoMapper(source any) models.Route01Oasa {
	var routeDetailDto models.Route01Oasa
	vMap, ok := source.(map[string]interface{})
	if !ok {
		panic("Προέκυψε σφάλμα στην ανάλυση του αντικειμένου.")
	}
	internalMapper(vMap, &routeDetailDto)
	return routeDetailDto
}

func (m route01Mapper) RouteDetailOasaToRouteDetailDto(source models.Route01Oasa) models.Route01 {
	var target models.Route01
	structMapper02(source, &target)
	return target
}
