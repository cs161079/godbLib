package Busstop

import (
	"fmt"

	models "github.com/cs161079/godbLib/Models"
	db "github.com/cs161079/godbLib/Repositories"

	"gorm.io/gorm"
)

func SelectByStopCode(stopCode int64) (*models.BusStop, error) {
	var selectedVal models.BusStop
	r := db.DB.Table("BUSSTOP").Where("stop_code = ?", stopCode).Find(&selectedVal)
	if r != nil {
		if r.Error != nil {
			fmt.Println(r.Error.Error())
			return nil, r.Error
		}
		if r.RowsAffected == 0 {
			return nil, nil
		}
	}
	return &selectedVal, nil
}

func Save(busStop models.BusStop) error {
	selectedBusStop, err := SelectByStopCode(busStop.Stop_code)
	if err != nil {
		return err
	}
	isNew := selectedBusStop == nil
	var r *gorm.DB = nil
	if isNew {
		newId, err := db.SequenceGetNextVal(models.BUSSTOP_SEQ)
		if err != nil {
			return err
		}
		busStop.Id = *newId
		r = db.DB.Table("BUSSTOP").Create(&busStop)

	} else {
		busStop.Id = selectedBusStop.Id
		r = db.DB.Table("BUSSTOP").Save(&busStop)
	}
	if r.Error != nil {
		return r.Error
	}
	return nil
}

func StopList01(routeCode int32) (*[]models.StopDto, error) {
	var result []models.StopDto
	r := db.DB.Table("BUSSTOP").
		Select("BUSSTOP.*, "+
			"BUSROUTESTOPS.senu").
		Joins("LEFT JOIN BUSROUTESTOPS ON BUSSTOP.stop_code=BUSROUTESTOPS.stop_code").
		Where("BUSROUTESTOPS.route_code=?", routeCode).Order("senu").Find(&result)
	if r.Error != nil {
		return nil, r.Error
	}
	return &result, nil
}
