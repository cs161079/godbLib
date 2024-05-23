package models

import "time"

type rawTime []byte

func (t rawTime) Time() time.Time {
	results, err := time.Parse("15:04:05", string(t))
	if err != nil {
		panic(err.Error())
	}
	return results
}

type ScheduleOasa struct {
	Sdc_Descr     string `json:"scheduleDescr" oasa:"sdc_descr"`
	Sdc_Descr_Eng string `json:"scheduleDescrEng" oasa:"sdc_descr_eng"`
	Sdc_Code      int32  `json:"scheduleCode" oasa:"sdc_code"`
}

type ScheduleMasterDto struct {
	Sdc_Descr     string        `json:"sdc_descr" oasa:"sdc_descr"`
	Sdc_Descr_Eng string        `json:"sdc_descr_eng" oasa:"sdc_descr_eng"`
	Sdc_Code      int32         `json:"sdc_code" oasa:"sdc_code"`
	ShcedeLine    ScheduleTimes `json:"scheduleLine"`
}

type ScheduleTimes struct {
	Go   []ScheduleTimeDto `oasa:"go"`
	Come []ScheduleTimeDto `oasa:"come"`
}
type ScheduleTimeOasa struct{}

type ScheduleTimeDto struct {
	//Line_Code  int64     `json:"line_code" oasa:"line_code"`
	Start_Time string `json:"start_time" oasa:"sde_start1" type:"time"`
	Type       int8   `json:"type"`
}

type ScheduleTime struct {
	Sdc_Code   int32  `json:"sdc_code" gorm:"primaryKey"`
	Line_Code  int64  `json:"line_code" gorm:"primaryKey"`
	Start_Time string `json:"start_time" gorm:"primaryKeys"`
	Type       int8   `json:"type" gorm:"primaryKey"`
}

// type ScheduleLineDto struct {
// 	Line_Code  int64     `json:"line_code"`
// 	Sdc_Code   int32     `json:"sdc_code"`
// 	Start_Time time.Time `json:"start_time"`
// 	Type       int8      `json:"type"`
// }

// type ScheduleLine struct {
// 	Line_Code  int64   `json:"line_code"`
// 	Sdc_Code   int32   `json:"sdc_code"`
// 	Start_Time rawTime `json:"start_time"`
// 	Type       int8    `json:"type"`
// }

type ScheduleMaster struct {
	Id            int64  `json:"id"`
	Sdc_Descr     string `json:"sdc_descr"`
	Sdc_Descr_Eng string `json:"sdc_descr_eng"`
	Sdc_Code      int32  `json:"sdc_code"`
	Line_Code     int64  `json:"line_code"`
}
