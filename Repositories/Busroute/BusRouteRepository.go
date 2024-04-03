package Busroute

import (
	"fmt"

	models "github.com/cs161079/godbLib/Models"
	db "github.com/cs161079/godbLib/Repositories"

	"gorm.io/gorm"
)

func SelectByRouteCode(routeCode int32) (*models.BusRoute, error) {
	var selectedVal models.BusRoute
	r := db.DB.Table("BUSROUTE").Where("route_code = ?", routeCode).Find(&selectedVal)
	if r != nil {
		if r.Error != nil {
			// fmt.Println(r.Error.Error())
			return nil, r.Error
		}
		if r.RowsAffected == 0 {
			return nil, fmt.Errorf("BUS ROUTE NOT FOUND [ROUTE_CODE: %d].", routeCode)
		}
	}
	return &selectedVal, nil
}

func SelectRouteByLineCode(line_code int32) (*[]models.BusRoute, error) {
	var selectedVal []models.BusRoute
	r := db.DB.Table("BUSROUTE").Where("line_code = ?", line_code).Find(&selectedVal)
	if r.Error != nil {
		return nil, r.Error
	}
	return &selectedVal, nil
}

func Save(input models.BusRoute) error {
	selectedBusLine, err := SelectByRouteCode(int32(input.Route_code))
	if err != nil {
		return err
	}
	isNew := selectedBusLine == nil
	var r *gorm.DB = nil
	if isNew {
		newId, err := db.SequenceGetNextVal(models.BUSROUTE_SEQ)
		if err != nil {
			return err
		}
		input.Id = *newId
		r = db.DB.Table("BUSROUTE").Create(&input)

	} else {
		input.Id = selectedBusLine.Id
		//input.Line_descr = input.Line_descr + " Update"
		r = db.DB.Table("BUSROUTE").Save(&input)
	}
	if r.Error != nil {
		return r.Error
	}
	return nil

}

func BusRouteList01() ([]models.BusRoute, error) {
	var result []models.BusRoute
	r := db.DB.Table("BUSROUTE").Order("route_code").Find(&result)
	if r != nil {
		if r.Error != nil {
			return nil, r.Error
		}
	}
	return result, nil
}
