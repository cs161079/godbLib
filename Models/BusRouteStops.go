package models

type BusRouteStops struct {
	Route_code int32 `json:"route_code" gorm:"primaryKey"`
	Stop_code  int64 `json:"stop_code" gorm:"primaryKey"`
	Senu       int16 `json:"senu" gorm:"primaryKey"`
}
