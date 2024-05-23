package mapper

import models "github.com/cs161079/godbLib/Models"

func StopMapper(source map[string]interface{}) models.StopOasa {
	var busStopOb models.StopOasa
	internalMapper(source, &busStopOb)
	return busStopOb
}

func StopOasaToStop(source models.StopOasa) models.Stop {
	var busStop models.Stop
	structMapper02(source, &busStop)
	return busStop
}
