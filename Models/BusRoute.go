package models

type BusRoute struct {
	Id              int64   `json:"Id" gorm:"PrimaryKey"`
	Route_code      int32   `json:"RouteCode" gorm:"index:ROUTE_CODE_UN,unique"`
	Line_code       int32   `json:"LineCode" gorm:"index:LINE_CODE_INDX"`
	Route_descr     string  `json:"RouteDescr"`
	Route_descr_eng string  `json:"RouteDescrEng"`
	Route_type      int8    `json:"RouteType"`
	Route_distance  float32 `json:"RouteDistance"`
}
