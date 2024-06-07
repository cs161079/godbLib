package mapper

import models "github.com/cs161079/godbLib/Models"

func RouteMapper(source any) models.RouteOasa {
	var busRouteOb models.RouteOasa
	vMap, ok := source.(map[string]interface{})
	if !ok {
		panic("Προέκυψε σφάλμα στην ανάλυση του αντικειμένου.")
	}
	internalMapper(vMap, &busRouteOb)

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
