package mapper

import models "github.com/cs161079/godbLib/Models"

type RouteMapper interface {
	GeneralRoute(source any) models.RouteOasa
	OasaToRouteDto(source models.RouteOasa) models.RouteDto
	DtoToRoute(source models.RouteDto) models.Route
}

type routeMapper struct {
}

func (m routeMapper) GeneralRoute(source any) models.RouteOasa {
	var busRouteOb models.RouteOasa
	vMap, ok := source.(map[string]interface{})
	if !ok {
		panic("Προέκυψε σφάλμα στην ανάλυση του αντικειμένου.")
	}
	internalMapper(vMap, &busRouteOb)

	return busRouteOb
}

func (m routeMapper) OasaToRouteDto(source models.RouteOasa) models.RouteDto {
	var target models.RouteDto
	structMapper02(source, &target)
	return target
}

func (m routeMapper) DtoToRoute(source models.RouteDto) models.Route {
	var target models.Route
	structMapper02(source, &target)
	return target
}
