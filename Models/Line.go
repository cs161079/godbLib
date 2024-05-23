package models

/*
	***************************************************
	This struct is to get data from OASA Application
	***************************************************
*/
type LineOasa struct {
	Ml_Code        int16  `json:"masterCode" oasa:"ml_code"`
	Sdc_Code       int16  `json:"scheduleCode" oasa:"sdc_code"`
	Line_Code      int32  `json:"lineCode" oasa:"line_code" gorm:"index:LINE_CODE,unique"`
	Line_Id        string `json:"lineId" oasa:"line_id"`
	Line_Descr     string `json:"lineDescr" oasa:"line_descr"`
	Line_Descr_Eng string `json:"lineDescrEng" oasa:"line_descr_eng"`
}

/*
	******************************************
	Struct for Bus Lines Entities for database
	******************************************
*/
type Line struct {
	Id             int64  `json:"id" gorm:"primaryKey"`
	Ml_Code        int16  `json:"ml_code"`
	Sdc_Code       int16  `json:"sdc_code"`
	Line_Code      int32  `json:"line_code" gorm:"index:LINE_CODE,unique"`
	Line_Id        string `json:"line_id"`
	Line_Descr     string `json:"line_descr"`
	Line_Descr_Eng string `json:"line_descr_eng"`
}

/*
	*************************************************
	       This struct is for different reasons
	*************************************************
*/
type LineDto struct {
	Id             int64               `json:"id"`
	Ml_Code        int64               `json:"ml_code"`
	Sdc_Code       int64               `json:"sdc_code"`
	Line_Code      int32               `json:"line_code"`
	Line_Id        string              `json:"line_id"`
	Line_Descr     string              `json:"line_descr"`
	Line_Descr_Eng string              `json:"line_descr_eng"`
	Routes         []RouteDto          `json:"routes"`
	Schedules      []ScheduleMasterDto `json:"scheduleDay"`
}
