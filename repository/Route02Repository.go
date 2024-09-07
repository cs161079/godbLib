package repository

import (
	"fmt"

	models "github.com/cs161079/godbLib/Models"
	logger "github.com/cs161079/godbLib/Utils/goLogger"
	"github.com/cs161079/godbLib/db"
	"gorm.io/gorm"
)

type Route02Repository interface {
	DeleteStopByRoute(routeCode int64) error
	InsertRoute02(input models.Route02) error
	UpdateRoute02(input models.Route02) error
	DeleteRoute02() error
}

type route02Repository struct {
	DB *gorm.DB
}

func (r route02Repository) DeleteStopByRoute(routeCode int64) error {
	var routeStops []models.Route02
	res := r.DB.Table(db.ROUTESTOPSTABLE).Where("route_code=?", routeCode).Delete(&routeStops)
	if res.Error != nil {
		// logger.ERROR(r.Error.Error())
		return res.Error
	}
	logger.INFO(fmt.Sprintf("DELETED ROWS %d", res.RowsAffected))
	return nil
}

func (r route02Repository) InsertRoute02(input models.Route02) error {
	res := r.DB.Table(db.ROUTESTOPSTABLE).Create(&input)
	if res.Error != nil {
		return res.Error
	}
	//logger.INFO(fmt.Sprintf("STOP [%d] SAVED SUCCESFULL IN ROUTE [%d].", input.Stop_code, input.Route_code))
	return nil
}

func (r route02Repository) UpdateRoute02(input models.Route02) error {
	res := r.DB.Table(db.ROUTESTOPSTABLE).Create(&input)
	if res.Error != nil {
		return res.Error
	}
	//logger.INFO(fmt.Sprintf("STOP [%d] SAVED SUCCESFULL IN ROUTE [%d].", input.Stop_code, input.Route_code))
	return nil
}

func (r route02Repository) DeleteRoute02() error {
	if err := r.DB.Table(db.ROUTESTOPSTABLE).Where("1=1").Delete(&models.Route02{}).Error; err != nil {
		//trans.Rollback()
		return err
	}
	return nil
}
