package mapper

import (
	models "github.com/cs161079/godbLib/Models"
)

type LineMapper interface {
	GeneralLine(any) models.LineOasa
	OasaToLineDto(models.LineOasa) models.LineDto
	LineDtoToLine(models.LineDto) models.Line
}

type lineMapper struct {
}

func (m lineMapper) GeneralLine(source any) models.LineOasa {
	var busLineOb models.LineOasa
	vMap, ok := source.(map[string]interface{})
	if !ok {
		panic("Προέκυψε σφάλμα στην ανάλυση του αντικειμένου.")
	}
	internalMapper(vMap, &busLineOb)

	return busLineOb
}

func (m lineMapper) OasaToLineDto(source models.LineOasa) models.LineDto {
	var target models.LineDto
	structMapper02(source, &target)
	return target
}

func (m lineMapper) LineDtoToLine(source models.LineDto) models.Line {
	var target models.Line
	structMapper02(source, &target)
	return target
}
