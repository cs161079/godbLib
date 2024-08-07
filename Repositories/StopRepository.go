package db

import (
	"fmt"

	models "github.com/cs161079/godbLib/Models"

	"gorm.io/gorm"
)

func SelectByStopCode(stopCode int64) (*models.Stop, error) {
	var selectedVal models.Stop
	r := DB.Table(STOPTABLE).Where("stop_code = ?", stopCode).Find(&selectedVal)
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
	if !isNew {
		busStop.Id = selectedBusStop.Id
	}
	r = DB.Table(STOPTABLE).Save(&busStop)
	if r.Error != nil {
		return r.Error
	}
	return nil
}

func StopList01(routeCode int32) (*[]models.Stop, error) {
	var result []models.Stop
	r := DB.Table(STOPTABLE).
		Select("stop.*, "+
			"routestops.senu").
		Joins("LEFT JOIN routestops ON stop.stop_code=routestops.stop_code").
		Where("routestops.route_code=?", routeCode).Order("senu").Find(&result)
	if r.Error != nil {
		return nil, r.Error
	}
	return &result, nil
}

func DeleteStop(trans *gorm.DB) error {
	if err := trans.Table(STOPTABLE).Where("1=1").Delete(&models.Stop{}).Error; err != nil {
		trans.Rollback()
		return err
	}
	return nil
}

func SelectClosestStops(point models.Point, from float32, to float32) ([]models.StopDto, error) {
	var resultList []models.StopDto
	var subQuery = DB.Table("stop s").Select("stop_code, stop_descr, stop_street," +
		fmt.Sprintf("round(haversine_distance(%f, %f, s.stop_lat, s.stop_lng), 2)", point.Lat, point.Long) +
		" AS distance")

	if err := DB.Table("(?) as b", subQuery).Select(" b. stop_code, b.stop_descr, b.stop_street, b.distance").
		Where(
			fmt.Sprintf(
				"distance > %f AND distance <= %f", from, to)).
		Order("distance").
		Find(&resultList).Error; err != nil {
		return nil, err
	}
	return resultList, nil

}
