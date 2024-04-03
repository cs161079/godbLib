package BusRouteStops

import (
	"fmt"

	models "github.com/cs161079/godbLib/Models"
	db "github.com/cs161079/godbLib/Repositories"
	logger "github.com/cs161079/godbLib/Utils/goLogger"
)

func DeleteStopByRoute(routeCode int64) error {
	var routeStops []models.BusRouteStops
	r := db.DB.Table("BUSROUTESTOPS").Where("route_code=?", routeCode).Delete(&routeStops)
	if r.Error != nil {
		// logger.ERROR(r.Error.Error())
		return r.Error
	}
	logger.INFO(fmt.Sprintf("DELETED ROWS %d", r.RowsAffected))
	return nil
}

func SaveRouteStops(input models.BusRouteStops) error {
	r := db.DB.Table("BUSROUTESTOPS").Create(&input)
	if r.Error != nil {
		return r.Error
	}
	logger.INFO(fmt.Sprintf("STOP [%d] SAVED SUCCESFULL IN ROUTE [%d].", input.Stop_code, input.Route_code))
	return nil
}
