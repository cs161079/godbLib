package repository

import (
	models "github.com/cs161079/godbLib/Models"
	logger "github.com/cs161079/godbLib/Utils/goLogger"
	"github.com/cs161079/godbLib/db"
	"gorm.io/gorm"
)

func NewLineRepository(iConnection *db.Connection) LineRepository {
	return lineRepository{
		DB: (*iConnection).GetConnection().DB,
	}
}

type lineRepository struct {
	DB *gorm.DB
}

type LineRepository interface {
	SelectByLineCode(lineCode int32) (*models.Line, error)
	InsertLine(line *models.Line) (*models.Line, error)
	UpdateLine(line *models.Line) (*models.Line, error)
	LineList01() ([]models.Line, error)
	DeleteLines() error
	LineWithTx(*gorm.DB) lineRepository
}

// withTx creates a new repository instance with the given transaction
func (r lineRepository) LineWithTx(tx *gorm.DB) lineRepository {
	if tx == nil {
		logger.WARN("Database Tranction not exist.")
		return r
	}
	r.DB = tx
	return r
}

func (r lineRepository) SelectByLineCode(lineCode int32) (*models.Line, error) {
	var selectedVal models.Line
	res := r.DB.Table(db.LINETABLE).Where("line_code = ?", lineCode).Find(&selectedVal)
	if res.Error != nil {
		return nil, res.Error
	}
	return &selectedVal, nil
}

func (r lineRepository) InsertLine(line *models.Line) (*models.Line, error) {
	trxRes := r.DB.Table(db.LINETABLE).Create(line)
	if trxRes.Error != nil {
		return nil, trxRes.Error
	}
	return line, nil
}

func (r lineRepository) UpdateLine(line *models.Line) (*models.Line, error) {
	trxRes := r.DB.Table(db.LINETABLE).Save(line)
	if trxRes.Error != nil {
		return nil, trxRes.Error
	}
	return line, nil
}

func (r lineRepository) LineList01() ([]models.Line, error) {
	var result []models.Line
	res := r.DB.Table(db.LINETABLE).Order("line_id, line_code").Find(&result)
	if res != nil {
		if res.Error != nil {
			return nil, res.Error
		}
	}
	return result, nil
}

func (r lineRepository) DeleteLines() error {
	if err := r.DB.Table(db.LINETABLE).Where("1=1").Delete(&models.Line{}).Error; err != nil {
		return err
	}
	return nil
}
