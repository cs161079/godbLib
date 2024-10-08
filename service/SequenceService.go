package service

import (
	"github.com/cs161079/godbLib/repository"
)

type SequenceService interface {
}

type sequenceService struct {
	Repo repository.SequenceRepository
}

func (s sequenceService) ZeroAllSequence() error {
	selectedData, err := s.Repo.SequenceList01()
	if err != nil {
		return err
	}
	for _, seq := range selectedData {
		seq.SEQ_COUNT = 0
		s.Repo.UpdateSequence(seq)
	}
	return nil
}
