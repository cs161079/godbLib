package BusRouteStops

import (
	models "godbLib/Models"
	db "godbLib/Repositories"
)

func DeleteStopByRoute(routeCode int64) error {
	var routeStops []models.BusRouteStops
	r := db.DB.Table("BUSROUTESTOPS").Where("route_code=?", routeCode).Delete(&routeStops)
	if r.Error != nil {
		// logger.ERROR(r.Error.Error())
		return r.Error
	}
	//oasaLogger.INFO(fmt.Sprintf("Deleted Rows are %d", r.RowsAffected))
	return nil
}

func SaveRouteStops(input models.BusRouteStops) error {
	r := db.DB.Table("BUSROUTESTOPS").Create(&input)
	if r.Error != nil {
		// logger.ERROR(r.Error.Error())
		return r.Error
	}
	//oasaLogger.INFO(fmt.Sprintf("Stop [%d] saved Succefully in Route [%d]", input.Stop_code, input.Route_code))
	return nil
}
