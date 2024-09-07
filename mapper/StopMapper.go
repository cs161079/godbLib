package mapper

import models "github.com/cs161079/godbLib/Models"

type StopMapper interface {
	StopOasaToStop(source models.StopOasa) models.Stop
	StopMapper(source any) models.StopOasa
}

type stopMapper struct {
}

func (m stopMapper) StopMapper(source any) models.StopOasa {
	var busStopOb models.StopOasa
	vMap, ok := source.(map[string]interface{})
	if !ok {
		panic("Προέκυψε σφάλμα στην ανάλυση του αντικειμένου.")
	}
	internalMapper(vMap, &busStopOb)
	return busStopOb
}

func (m stopMapper) StopOasaToStop(source models.StopOasa) models.Stop {
	var busStop models.Stop
	structMapper02(source, &busStop)
	return busStop
}
