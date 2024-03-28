package models

//type Line struct {
//	ID       int64   `json:"LineCode" gorm:"primaryKey"`
//	CODE     string  `json:"LineID"`
//	DESCR    string  `json:"LineDescr"`
//	DESCRENG string  `json:"LineDescrEng"`
//	NUM1     int     `json:"Num1"`
//	NUM2     float32 `json:"Num2"`
//	NUM3     int     `json:"Num3"`
//	NUM4     float32 `json:"Num4"`
//	NUM5     int     `json:"Num5"`
//	NUM6     float32 `json:"Num6"`
//	NUM7     int     `json:"Num7"`
//	NUM8     float32 `json:"Num8"`
//	NUM9     int     `json:"Num9"`
//	NUM10    float32 `json:"Num10"`
//	NUM11    int     `json:"Num11"`
//	NUM12    float32 `json:"Num12"`
//	NUM13    int     `json:"Num13"`
//	NUM14    float32 `json:"Num14"`
//}
//
//type LineList []Line

type Busline struct {
	Id             int64  `gorm:"primaryKey"`
	Ml_code        int16  `json:"ml_code"`
	Sdc_code       int16  `json:"sdc_code"`
	Line_code      int32  `json:"line_code" validate:"required" gorm:"index:LINE_CODE,unique"`
	Line_id        string `json:"line_id"`
	Line_descr     string `json:"line_descr" validate:"required,min=2,max=5"`
	Line_descr_eng string `json:"line_descr_eng" validate:"required,min=2,max=5"`
}
type LineDto struct {
	ID             int64
	Ml_Code        int64
	Sdc_Code       int64
	Line_Code      int64
	Line_Id        string
	Line_Descr     string `validate:"required,min=2,max=5`
	Line_Descr_Eng string
}

// type LineDto01 struct {
// 	ml_code        string
// 	sdc_code       string
// 	line_code      string
// 	line_id        string
// 	line_descr     string
// 	line_descr_eng string
// }
