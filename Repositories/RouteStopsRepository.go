package db

import (
	"fmt"

	models "github.com/cs161079/godbLib/Models"
	logger "github.com/cs161079/godbLib/Utils/goLogger"
	"gorm.io/gorm"
)

func DeleteStopByRoute(routeCode int64) error {
	var routeStops []models.RouteStops
	r := DB.Table("ROUTESTOPS").Where("route_code=?", routeCode).Delete(&routeStops)
	if r.Error != nil {
		// logger.ERROR(r.Error.Error())
		return r.Error
	}
	logger.INFO(fmt.Sprintf("DELETED ROWS %d", r.RowsAffected))
	return nil
}

func SaveRouteStops(input models.RouteStops) error {
	r := DB.Table("ROUTESTOPS").Create(&input)
	if r.Error != nil {
		return r.Error
	}
	logger.INFO(fmt.Sprintf("STOP [%d] SAVED SUCCESFULL IN ROUTE [%d].", input.Stop_code, input.Route_code))
	return nil
}

func DeleteRouteStops(trans *gorm.DB) error {
	if err := trans.Table("ROUTESTOPS").Where("1=1").Delete(&models.RouteStops{}).Error; err != nil {
		trans.Rollback()
		return err
	}
	return nil
}
