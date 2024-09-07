package mapper

import models "github.com/cs161079/godbLib/Models"

func UVersionsMapper(source any) models.UVersionsOasa {
	var oasaOb models.UVersionsOasa
	vMap, ok := source.(map[string]interface{})
	if !ok {
		panic("Προέκυψε σφάλμα στην ανάλυση του αντικειμένου.")
	}
	internalMapper(vMap, &oasaOb)

	return oasaOb
}

func UVersionsOasaToUVersions(source models.UVersionsOasa) models.UVersions {
	var target models.UVersions
	structMapper02(source, &target)
	return target
}
