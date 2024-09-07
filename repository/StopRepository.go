package repository

import (
	"fmt"

	models "github.com/cs161079/godbLib/Models"
	"github.com/cs161079/godbLib/db"

	"gorm.io/gorm"
)

type StopRepository interface{}

type stopRepository struct {
	DB *gorm.DB
}

func (r stopRepository) SelectByStopCode(stopCode int64) (*models.Stop, error) {
	var selectedVal models.Stop
	res := r.DB.Table(db.STOPTABLE).Where("stop_code = ?", stopCode).Find(&selectedVal)
	if res.Error != nil {
		fmt.Println(res.Error.Error())
		return nil, res.Error
	}
	return &selectedVal, nil
}

func (r stopRepository) InsertStop(busStop models.Stop) (*models.Stop, error) {
	res := r.DB.Table(db.STOPTABLE).Create(&busStop)
	if res.Error != nil {
		return nil, res.Error
	}
	return &busStop, nil
}

func (r stopRepository) UpdateStop(busStop models.Stop) (*models.Stop, error) {
	res := r.DB.Table(db.STOPTABLE).Save(&busStop)
	if res.Error != nil {
		return nil, res.Error
	}
	return &busStop, nil
}

func (r stopRepository) StopList01(routeCode int32) (*[]models.Stop, error) {
	var result []models.Stop
	res := r.DB.Table(db.STOPTABLE).
		Select("stop.*, "+
			"routestops.senu").
		Joins("LEFT JOIN routestops ON stop.stop_code=routestops.stop_code").
		Where("routestops.route_code=?", routeCode).Order("senu").Find(&result)
	if res.Error != nil {
		return nil, res.Error
	}
	return &result, nil
}

func (r stopRepository) DeleteStop() error {
	if err := r.DB.Table(db.STOPTABLE).Where("1=1").Delete(&models.Stop{}).Error; err != nil {
		return err
	}
	return nil
}

func (r stopRepository) SelectClosestStops(point models.Point, from float32, to float32) ([]models.StopDto, error) {
	var resultList []models.StopDto
	var subQuery = r.DB.Table("stop s").Select("stop_code, stop_descr, stop_street," +
		fmt.Sprintf("round(haversine_distance(%f, %f, s.stop_lat, s.stop_lng), 2)", point.Lat, point.Long) +
		" AS distance")

	if err := r.DB.Table("(?) as b", subQuery).Select(" b. stop_code, b.stop_descr, b.stop_street, b.distance").
		Where(
			fmt.Sprintf(
				"distance > %f AND distance <= %f", from, to)).
		Order("distance").
		Find(&resultList).Error; err != nil {
		return nil, err
	}
	return resultList, nil

}
