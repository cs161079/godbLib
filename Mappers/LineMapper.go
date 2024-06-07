package mapper

import (
	models "github.com/cs161079/godbLib/Models"
)

func LineMapper(source any) models.LineOasa {
	var busLineOb models.LineOasa
	vMap, ok := source.(map[string]interface{})
	if !ok {
		panic("Προέκυψε σφάλμα στην ανάλυση του αντικειμένου.")
	}
	internalMapper(vMap, &busLineOb)

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
