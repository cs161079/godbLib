package db

import (
	"fmt"

	models "github.com/cs161079/godbLib/Models"
	logger "github.com/cs161079/godbLib/Utils/goLogger"

	"gorm.io/gorm"
)

func SelectBySdcCodeLineCode(iLine int64, iSdc int32) (*models.ScheduleMaster, error) {
	var selectedVal models.ScheduleMaster
	r := DB.Table("SCHEDULEMASTER").Where("sdc_code = ? AND line_code = ?", iSdc, iLine).Find(&selectedVal)
	if r.Error != nil {
		fmt.Println(r.Error.Error())
		return nil, r.Error
	}
	if r.RowsAffected == 0 {
		logger.INFO(fmt.Sprintf("BUS SCHEDULE MASTER LINE NOT FOUND [line_code: %d, sdc_code: %d].", iLine, iSdc))
		return nil, nil
	}
	return &selectedVal, nil
}

func SaveScheduleMaster(input models.ScheduleMaster) error {
	selectedBusLine, err := SelectBySdcCodeLineCode(input.Line_Code, int32(input.Sdc_Code))
	if err != nil {
		return err
	}
	isNew := selectedBusLine == nil
	var r *gorm.DB = nil
	if !isNew {
		input.Id = selectedBusLine.Id
		//input.Line_descr = input.Line_descr + " Update"
	}
	r = DB.Table("SCHEDULEMASTER").Save(&input)
	if r.Error != nil {
		return r.Error
	}
	return nil
}

func DeleteScheduleMaster(trans *gorm.DB) error {
	if err := trans.Table("SCHEDULEMASTER").Where("1=1").Delete(&models.ScheduleMaster{}).Error; err != nil {
		trans.Rollback()
		return err
	}
	return nil
}
