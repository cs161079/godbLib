package mapper

import (
	models "github.com/cs161079/godbLib/Models"
)

func LineMapper(source map[string]interface{}) models.LineOasa {
	var busLineOb models.LineOasa
	internalMapper(source, &busLineOb)

	return busLineOb
}

func LineOasaToLine(source models.LineOasa) models.LineDto {
	var target models.LineDto
	structMapper02(source, &target)
	return target
}

func LineDtoToLine(source models.LineDto) models.Line {
	var target models.Line
	structMapper02(source, &target)
	return target
}
