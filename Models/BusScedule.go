package models

// This is a model Describe Master Bus Line
type BusScheduleMasterLineDto struct {
	Sdc_descr     string `json: sdc_descr`
	Sdc_descr_eng string `json: sdc_descr_eng`
	Sdc_code      int32  `json: sdc_code`
}

// type BusScheduleMasterLine struct {
// 	Sdc_descr     string `json: sdc_descr`
// 	Sdc_descr_eng string `json: sdc_descr_eng`
// 	Sdc_code      int32  `json: sdc_code`
// 	Line_code     int32  `json: line_code`
// }

type BusScheduleMasterLine struct {
	Id            int64  `json: id`
	Sdc_descr     string `json: sdc_descr`
	Sdc_descr_eng string `json: sdc_descr_eng`
	Sdc_code      int16  `json: sdc_code`
	Line_code     int32  `json: line_code`
}

type BusSheduleLine struct {
	Id        int32 `json: id`
	Sdc_code  int16 `json:sdc_code`
	Ml_code   int16 `json:ml_code`
	Line_code int32 `json: line_code`
}
