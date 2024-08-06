package db

import (
	models "github.com/cs161079/godbLib/Models"
	"gorm.io/gorm"
)

//const erroMessageTemplate = "Field validation for [%s] failed on the [%s] tag"

type OpswValidateError struct {
	Key     string
	Message string
}

func SelectByLineCode(lineCode int64) (*models.Line, error) {
	//var selectedPtr *oasaSyncModel.Busline
	var selectedVal models.Line
	r := DB.Table(LINETABLE).Where("line_code = ?", lineCode).Find(&selectedVal)
	if r != nil {
		if r.Error != nil {
			return nil, r.Error
		}
		if r.RowsAffected == 0 {
			//logger.WARN(fmt.Sprintf("BUS LINE NOT FOUND. [line_code: %d].", lineCode))
			return nil, nil
		}
	}
	return &selectedVal, nil
}

func SaveLine(input models.Line) error {
	selectedBusLine, err := SelectByLineCode(int64(input.Line_Code))
	if err != nil {
		return err
	}
	isNew := selectedBusLine == nil
	var r *gorm.DB = nil
	if !isNew {
		input.Id = selectedBusLine.Id
	}
	r = DB.Table(LINETABLE).Save(&input)
	if r.Error != nil {
		return r.Error
	}
	return nil

}

func LineList01() ([]models.Line, error) {
	var result []models.Line
	r := DB.Table(LINETABLE).Order("line_id, line_code").Find(&result)
	if r != nil {
		if r.Error != nil {
			return nil, r.Error
		}
	}
	return result, nil
}

func LineList01Distinct() ([]models.Line, error) {
	allBusLines, err := LineList01()
	if err != nil {
		return nil, err
	}
	var result []models.Line
	var currentLine = allBusLines[0]
	result = append(result, currentLine)
	for _, s := range allBusLines {
		if currentLine.Line_Id != s.Line_Id {
			result = append(result, s)
			currentLine = s
		}
	}
	return result, nil
}

func LineListBymlcode(mlcode int16) ([]models.Line, error) {
	var result []models.Line
	r := DB.Table(LINETABLE).Where("ml_code = ?", mlcode).Order("line_id, line_code").Find(&result)
	if r != nil {
		if r.Error != nil {
			return nil, r.Error
		}
	}
	return result, nil

}

func DeleteLines(trans *gorm.DB) error {
	if err := trans.Table(LINETABLE).Where("1=1").Delete(&models.Line{}).Error; err != nil {
		trans.Rollback()
		return err
	}
	return nil
}
