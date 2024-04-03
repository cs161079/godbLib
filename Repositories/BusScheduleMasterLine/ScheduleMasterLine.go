package BusScheduleMasterLine

import (
	"fmt"

	models "github.com/cs161079/godbLib/Models"
	db "github.com/cs161079/godbLib/Repositories"
	logger "github.com/cs161079/godbLib/Utils/goLogger"

	"gorm.io/gorm"
)

func SelectBySdcCodeLineCode(iLine int32, iSdc int32) (*models.BusScheduleMasterLine, error) {
	var selectedVal models.BusScheduleMasterLine
	r := db.DB.Table("BUSSCHEDULEMASTERLINE").Where("sdc_code = ? AND line_code = ?", iSdc, iLine).Find(&selectedVal)
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

func Save(input models.BusScheduleMasterLine) error {
	selectedBusLine, err := SelectBySdcCodeLineCode(input.Line_code, int32(input.Sdc_code))
	if err != nil {
		return err
	}
	isNew := selectedBusLine == nil
	var r *gorm.DB = nil
	if isNew {
		newId, err := db.SequenceGetNextVal(models.BUSSCHEDULEMASTERLINE)
		if err != nil {
			return err
		}
		input.Id = *newId
		r = db.DB.Table("BUSSCHEDULEMASTERLINE").Create(&input)

	} else {
		input.Id = selectedBusLine.Id
		//input.Line_descr = input.Line_descr + " Update"
		r = db.DB.Table("BUSSCHEDULEMASTERLINE").Save(&input)
	}
	if r.Error != nil {
		return r.Error
	}
	return nil
}
