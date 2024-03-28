package Busroute

import (
	"errors"
	"fmt"
	models "godbLib/Models"
	db "godbLib/Repositories"

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
			//oasaLogger.Logger.Infof("Bus Route Not Found [route_code: %d].\n", routeCode)
			return nil, errors.New(fmt.Sprintf("Bus Route Not Found [route_code: %d].\n", routeCode))
		}
	}
	return &selectedVal, nil
}

func SelectRouteByLineCode(line_code int32) []models.BusRoute {
	var selectedVal []models.BusRoute
	r := db.DB.Table("BUSROUTE").Where("line_code = ?", line_code).Find(&selectedVal)
	if r != nil {
		if r.Error != nil {
			fmt.Println(r.Error.Error())
			return nil
		}
		//if r.RowsAffected == 0 {
		//	fmt.Println("Bus Routes Not Found [li: %d].")
		//	return nil
		//}
	}
	return selectedVal
}

func Save(input models.BusRoute) error {
	selectedBusLine, err := SelectByRouteCode(int32(input.Route_code))
	if err != nil {
		return err
	}
	isNew := selectedBusLine == nil
	var r *gorm.DB = nil
	if isNew {
		input.Id = db.SequenceGetNextVal(models.BUSROUTE_SEQ)
		//input.Line_descr = input.Line_descr + " New"
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

func BusRouteList01() []models.BusRoute {
	var result []models.BusRoute
	r := db.DB.Table("BUSROUTE").Order("route_code").Find(&result)
	if r != nil {
		if r.Error != nil {
			fmt.Println(r.Error.Error())
			return nil
		}
		//if r.RowsAffected == 0 {
		//	fmt.Println("Record does not exist!!!")
		//	return nil
		//}
	}
	return result
}
