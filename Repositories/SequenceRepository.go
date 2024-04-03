package db

import (
	models "github.com/cs161079/godbLib/Models"
)

func SequenceGetNextVal(seq_name string) (*int64, error) {
	var nextVal int64
	var sequnece models.Sequence
	r0 := DB.Table("SEQUENCES").Where("SEQ_GEN=?", seq_name).Find(&sequnece)
	if r0 != nil && r0.RowsAffected == 0 {
		sequnece = models.Sequence{
			SEQ_GEN: seq_name,
		}
		nextVal = 1
	} else {
		nextVal = sequnece.SEQ_COUNT + 1
	}
	sequnece.SEQ_COUNT = nextVal
	r1 := DB.Table("SEQUENCES").Save(&sequnece)
	if r1 != nil {
		if r1.Error != nil {
			return nil, r1.Error
		}
	}
	return &nextVal, nil
}
