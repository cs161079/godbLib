package db

import (
	"fmt"

	models "github.com/cs161079/godbLib/Models"
	logger "github.com/cs161079/godbLib/Utils/goLogger"
	"gorm.io/gorm"
)

func SequenceGetNextVal(seqName string) (*int64, error) {
	var nextVal int64
	var sequnece models.Sequence
	r0 := DB.Table("SEQUENCE").Where("SEQ_GEN=?", seqName).Find(&sequnece)
	if r0 != nil && r0.RowsAffected == 0 {
		sequnece = models.Sequence{
			SEQ_GEN: seqName,
		}
		nextVal = 1
	} else {
		nextVal = sequnece.SEQ_COUNT + 1
	}
	sequnece.SEQ_COUNT = nextVal
	r1 := DB.Table("SEQUENCE").Save(&sequnece)
	if r1 != nil {
		if r1.Error != nil {
			return nil, r1.Error
		}
	}
	return &nextVal, nil
}

func SequenceList01() ([]models.Sequence, error) {
	var selectedData []models.Sequence
	r := DB.Table("SEQUENCE").Find(&selectedData)
	if r != nil {
		if r.Error != nil {
			return nil, r.Error
		}
		if r.RowsAffected == 0 {
			logger.WARN(fmt.Sprintf("DOES NOT EXIST DATA IN SEQUENCE TABLE"))
			return nil, nil
		}
	}
	return selectedData, nil
}

func UpdateSequence(trans *gorm.DB, seq models.Sequence) error {
	if err := trans.Table("SEQUENCE").Where("SEQ_GEN=?", seq.SEQ_GEN).Update("SEQ_COUNT", seq.SEQ_COUNT).Error; err != nil {
		trans.Rollback()
		return err
	}
	logger.WARN(fmt.Sprintf("ROWS AFFECTED %d", trans.RowsAffected))
	return nil
}

func ZeroAllSequence(trans *gorm.DB) error {
	selectedData, err := SequenceList01()
	if err != nil {
		return err
	}
	for _, seq := range selectedData {
		seq.SEQ_COUNT = 0
		UpdateSequence(trans, seq)
	}
	return nil
}
