package db

import (
	"fmt"

	models "github.com/cs161079/godbLib/Models"

	"gorm.io/gorm"
)

func SelectByStopCode(stopCode int64) (*models.Stop, error) {
	var selectedVal models.Stop
	r := DB.Table("STOP").Where("stop_code = ?", stopCode).Find(&selectedVal)
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

func SaveStop(busStop models.Stop) error {
	selectedBusStop, err := SelectByStopCode(busStop.Stop_code)
	if err != nil {
		return err
	}
	isNew := selectedBusStop == nil
	var r *gorm.DB = nil
	if isNew {
		newId, err := SequenceGetNextVal(models.BUSSTOP_SEQ)
		if err != nil {
			return err
		}
		busStop.Id = *newId
		r = DB.Table("STOP").Create(&busStop)

	} else {
		busStop.Id = selectedBusStop.Id
		r = DB.Table("STOP").Save(&busStop)
	}
	if r.Error != nil {
		return r.Error
	}
	return nil
}

func StopList01(routeCode int32) (*[]models.Stop, error) {
	var result []models.Stop
	r := DB.Table("STOP").
		Select("STOP.*, "+
			"ROUTESTOPS.senu").
		Joins("LEFT JOIN ROUTESTOPS ON STOP.stop_code=ROUTESTOPS.stop_code").
		Where("ROUTESTOPS.route_code=?", routeCode).Order("senu").Find(&result)
	if r.Error != nil {
		return nil, r.Error
	}
	return &result, nil
}

func DeleteStop(trans *gorm.DB) error {
	if err := trans.Table("STOP").Where("1=1").Delete(&models.Stop{}).Error; err != nil {
		trans.Rollback()
		return err
	}
	return nil
}
