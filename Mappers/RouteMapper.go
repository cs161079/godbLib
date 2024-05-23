package mapper

import models "github.com/cs161079/godbLib/Models"

func RouteMapper(source map[string]interface{}) models.RouteOasa {
	var busRouteOb models.RouteOasa
	internalMapper(source, &busRouteOb)

	return busRouteOb
}

func RouteOasaToRouteDto(source models.RouteOasa) models.RouteDto {
	var target models.RouteDto
	structMapper02(source, &target)
	return target
}

func RouteDtoToRoute(source models.RouteDto) models.Route {
	var target models.Route
	structMapper02(source, &target)
	return target
}
