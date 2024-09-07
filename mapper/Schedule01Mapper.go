package mapper

import models "github.com/cs161079/godbLib/Models"

type Schedule01Mapper interface {
	DtoToSchedule01(source models.Schedule01Dto) models.Schedule01
	Schedule01Mapper(source any) models.Schedule
}

type schedule01Mapper struct {
}

func (m schedule01Mapper) Schedule01Mapper(source any) models.Schedule {
	var result models.Schedule
	vMap, ok := source.(map[string]interface{})
	if !ok {
		panic("Προέκυψε σφάλμα στην ανάλυση του αντικειμένου.")
	}
	for _, rec := range vMap["go"].([]interface{}) {
		var current models.Schedule01
		internalMapper(rec.(map[string]interface{}), &current)
		result.Go = append(result.Go, current)
	}
	for _, rec := range vMap["come"].([]interface{}) {
		var current models.Schedule01
		internalMapper(rec.(map[string]interface{}), &current)
		result.Come = append(result.Come, current)
	}
	return result
}

func (m schedule01Mapper) DtoToSchedule01(source models.Schedule01Dto) models.Schedule01 {
	var target models.Schedule01
	structMapper02(source, &target)
	return target
}
