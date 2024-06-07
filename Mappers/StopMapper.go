package mapper

import models "github.com/cs161079/godbLib/Models"

func StopMapper(source any) models.StopOasa {
	var busStopOb models.StopOasa
	vMap, ok := source.(map[string]interface{})
	if !ok {
		panic("Προέκυψε σφάλμα στην ανάλυση του αντικειμένου.")
	}
	internalMapper(vMap, &busStopOb)
	return busStopOb
}

func StopOasaToStop(source models.StopOasa) models.Stop {
	var busStop models.Stop
	structMapper02(source, &busStop)
	return busStop
}
