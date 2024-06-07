package mapper

import models "github.com/cs161079/godbLib/Models"

func ScheduleOasaToScheduleDto(source models.ScheduleOasa) models.ScheduleMasterDto {
	var target models.ScheduleMasterDto
	structMapper02(source, &target)
	return target
}

func ScheduleTimesMapper(source any) models.ScheduleTimes {
	var result models.ScheduleTimes
	vMap, ok := source.(map[string]interface{})
	if !ok {
		panic("Προέκυψε σφάλμα στην ανάλυση του αντικειμένου.")
	}
	for _, rec := range vMap["go"].([]interface{}) {
		var current models.ScheduleTimeDto
		internalMapper(rec.(map[string]interface{}), &current)
		result.Go = append(result.Go, current)
	}
	for _, rec := range vMap["come"].([]interface{}) {
		var current models.ScheduleTimeDto
		internalMapper(rec.(map[string]interface{}), &current)
		result.Come = append(result.Come, current)
	}
	return result
}

func ScheduleMasterDtoToScheduleMaster(source models.ScheduleMasterDto) models.ScheduleMaster {
	var target models.ScheduleMaster
	structMapper02(source, &target)
	return target
}

func ScheduleTimeDtoToScheduleTime(source models.ScheduleTimeDto) models.ScheduleTime {
	var target models.ScheduleTime
	structMapper02(source, &target)
	return target
}
