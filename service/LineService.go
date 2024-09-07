package service

import (
	"context"
	"time"

	models "github.com/cs161079/godbLib/Models"
	"github.com/cs161079/godbLib/mapper"
	"github.com/cs161079/godbLib/repository"
	"gorm.io/gorm"
)

type LineService interface {
	SelectByLineCode(lineCode int32) (*models.Line, error)
	PostLine(line *models.Line) (*models.Line, error)
	PostLineArray(context.Context, []models.Line) ([]models.Line, error)
	WithTrx(*gorm.DB)
}

type lineService struct {
	Repo   repository.LineRepository
	Mapper mapper.LineMapper
}

func (s lineService) SelectByLineCode(line_code int32) (*models.Line, error) {
	return s.Repo.SelectByLineCode(line_code)

}

func (s lineService) PostLine(line *models.Line) (*models.Line, error) {
	time.Sleep(10 * time.Second)
	var selectedLine *models.Line = nil
	var err error = nil
	//selectedBusLine, err := s.Repo.SelectByLineCode(line.Line_Code)
	if err != nil {
		return nil, err
	}
	isNew := selectedLine == nil
	if isNew {
		return s.Repo.InsertLine(line)
	} else {
		line.Id = selectedLine.Id
		return s.Repo.UpdateLine(line)
	}
}

func (s lineService) WithTrx(trxHandle *gorm.DB) lineService {
	s.Repo = s.Repo.LineWithTx(trxHandle)
	return s
}

func (s lineService) PostLineArray(ctx context.Context, lines []models.Line) error {
	// var result []models.Line = make([]models.Line, 0)
	var trx = ctx.Value("db_tx").(*gorm.DB).Begin()
	for _, line := range lines {
		_, err := s.WithTrx(trx).PostLine(&line)
		if err != nil {
			trx.Rollback()
			return err
		}
		// result = append(result, *ln)
	}
	if err := trx.Commit().Error; err != nil {
		return err
	}
	return nil
}
