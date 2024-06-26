package db

import (
	"fmt"

	models "github.com/cs161079/godbLib/Models"
	logger "github.com/cs161079/godbLib/Utils/goLogger"

	"gorm.io/gorm"
)

func SelectByRouteCode(routeCode int32) (*models.Route, error) {
	var selectedVal models.Route
	r := DB.Table("ROUTE").Where("route_code = ?", routeCode).Find(&selectedVal)
	if r != nil {
		if r.Error != nil {
			// fmt.Println(r.Error.Error())
			return nil, r.Error
		}
		if r.RowsAffected == 0 {
			logger.WARN(fmt.Sprintf("BUS ROUTE NOT FOUND [ROUTE_CODE: %d]", routeCode))
			return nil, nil
		}
	}
	return &selectedVal, nil
}

func SelectRouteByLineCode(lineCode int32) (*[]models.Route, error) {
	var selectedVal []models.Route
	r := DB.Table("ROUTE").Where("line_code = ?", lineCode).Find(&selectedVal)
	if r.Error != nil {
		return nil, r.Error
	}
	return &selectedVal, nil
}

func SaveRoute(input models.Route) error {
	selectedBusLine, err := SelectByRouteCode(int32(input.Route_Code))
	if err != nil {
		return err
	}
	isNew := selectedBusLine == nil
	var r *gorm.DB = nil
	if isNew {
		newId, err := SequenceGetNextVal(models.BUSROUTE_SEQ)
		if err != nil {
			return err
		}
		input.Id = *newId
		r = DB.Table("ROUTE").Create(&input)

	} else {
		input.Id = selectedBusLine.Id
		//input.Line_descr = input.Line_descr + " Update"
		r = DB.Table("ROUTE").Save(&input)
	}
	if r.Error != nil {
		return r.Error
	}
	return nil

}

func RouteList01() ([]models.Route, error) {
	var result []models.Route
	r := DB.Table("ROUTE").Order("route_code").Find(&result)
	if r != nil {
		if r.Error != nil {
			return nil, r.Error
		}
	}
	return result, nil
}

func DeleteRoute(trans *gorm.DB) error {
	if err := trans.Table("ROUTE").Where("1=1").Delete(&models.Route{}).Error; err != nil {
		trans.Rollback()
		return err
	}
	return nil
}
