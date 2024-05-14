package BuslineRepository

import (
	"fmt"

	models "github.com/cs161079/godbLib/Models"
	db "github.com/cs161079/godbLib/Repositories"
	logger "github.com/cs161079/godbLib/Utils/goLogger"

	"gorm.io/gorm"
)

//const erroMessageTemplate = "Field validation for [%s] failed on the [%s] tag"

type OpswValidateError struct {
	Key     string
	Message string
}

func SelectByLineCode(lineCode int64) (*models.Busline, error) {
	//var selectedPtr *oasaSyncModel.Busline
	var selectedVal models.Busline
	r := db.DB.Table("BUSLINE").Where("line_code = ?", lineCode).Find(&selectedVal)
	if r != nil {
		if r.Error != nil {
			return nil, r.Error
		}
		if r.RowsAffected == 0 {
			logger.INFO(fmt.Sprintf("Bus line not found. [line_code: %d].", lineCode))
			return nil, nil
		}
	}
	return &selectedVal, nil
}

func SaveBusLine(input models.Busline) error {
	selectedBusLine, err := SelectByLineCode(int64(input.Line_code))
	if err != nil {
		return err
	}
	isNew := selectedBusLine == nil
	var r *gorm.DB = nil
	if isNew {
		newId, err := db.SequenceGetNextVal(models.BUSLINE_SEQ)
		if err != nil {
			return err
		}
		input.Id = *newId
		r = db.DB.Table("BUSLINE").Create(&input)

	} else {
		input.Id = selectedBusLine.Id
		//input.Line_descr = input.Line_descr + " Update"
		r = db.DB.Table("BUSLINE").Save(&input)
	}
	if r.Error != nil {
		return r.Error
	}
	return nil

}

func BuslineList01() ([]models.Busline, error) {
	var result []models.Busline
	r := db.DB.Table("BUSLINE").Order("line_id, line_code").Find(&result)
	if r != nil {
		if r.Error != nil {
			return nil, r.Error
		}
	}
	return result, nil
}

func BuslineList01Distinct() ([]models.Busline, error) {
	allBusLines, err := BuslineList01()
	if err != nil {
		return nil, err
	}
	var result []models.Busline
	var currentLine = allBusLines[0]
	result = append(result, currentLine)
	for _, s := range allBusLines {
		if currentLine.Line_id != s.Line_id {
			result = append(result, s)
			currentLine = s
		}
	}
	return result, nil
}

func BuslineListBymlcode(mlcode int16) ([]models.Busline, error) {
	var result []models.Busline
	r := db.DB.Table("BUSLINE").Where("ml_code = ?", mlcode).Order("line_id, line_code").Find(&result)
	if r != nil {
		if r.Error != nil {
			return nil, r.Error
		}
	}
	return result, nil

}
