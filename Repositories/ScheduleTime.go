package db

import (
	"fmt"

	models "github.com/cs161079/godbLib/Models"
	logger "github.com/cs161079/godbLib/Utils/goLogger"
	"gorm.io/gorm"
)

func SelectScheduleTime(lineCode int64, sdcCode int32) ([]models.ScheduleTime, error) {
	//var selectedPtr *oasaSyncModel.Busline
	var selectedVal []models.ScheduleTime
	r := DB.Table("SCHEDULETIME").Where("line_code = ? and sdc_code = ?", lineCode, sdcCode).Find(&selectedVal)
	if r != nil {
		if r.Error != nil {
			return nil, r.Error
		}
		if r.RowsAffected == 0 {
			logger.WARN(fmt.Sprintf("SCHEDULE NOT FOUND. [line_code: %d].", lineCode))
			return nil, nil
		}
	}
	return selectedVal, nil
}

func SelectScheduleTimeByKey(lineCode int64, sdcCode int32, tTime string, typ int) ([]models.ScheduleTime, error) {
	//var selectedPtr *oasaSyncModel.Busline
	var selectedVal []models.ScheduleTime
	r := DB.Table("SCHEDULETIME").Where("line_code = ? and sdc_code = ? and start_time = ? and type = ?", lineCode, sdcCode, tTime, typ).Find(&selectedVal)
	if r != nil {
		if r.Error != nil {
			return nil, r.Error
		}
		if r.RowsAffected == 0 {
			logger.WARN(fmt.Sprintf("SCHEDULE NOT FOUND. [line_code: %d, %d, %s, %d].", lineCode, sdcCode, tTime, typ))
			return nil, nil
		}
	}
	return selectedVal, nil
}

func SaveScheduleTime(input models.ScheduleTime) error {
	var r *gorm.DB = nil
	sdcTime, err := SelectScheduleTimeByKey(input.Line_Code, input.Sdc_Code, input.Start_Time, int(input.Type))
	if err != nil {
		return err
	}
	if sdcTime != nil {
		return nil
	}
	r = DB.Table("SCHEDULETIME").Create(&input)
	if r.Error != nil {
		return r.Error
	}
	return nil
}

func DeleteScheduleTime(trans *gorm.DB) error {
	if err := trans.Table("SCHEDULETIME").Where("1=1").Delete(&models.ScheduleTime{}).Error; err != nil {
		trans.Rollback()
		return err
	}
	return nil
}
