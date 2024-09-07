package models

// ********* Struct for Route Details OASA **************
type Route01Oasa struct {
	Routed_x     float32 `oasa:"routed_x"`
	Routed_y     float32 `oasa:"routed_y"`
	Routed_order int16   `oasa:"routed_order"`
}

// ********* Struct for Route Details Entities **************
type Route01 struct {
	Route_code   int32   `json:"route_code" gorm:"PrimaryKey"`
	Routed_x     float32 `json:"routed_x"`
	Routed_y     float32 `json:"routed_y"`
	Routed_order int16   `json:"routed_order" gorm:"PrimaryKey"`
}

//**********************************************************
