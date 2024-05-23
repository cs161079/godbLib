package models

type RouteOasa struct {
	Route_Code      int32   `json:"routeCode" oasa:"RouteCode"`
	Line_Code       int32   `json:"lineCode" oasa:"LineCode"`
	Route_Descr     string  `json:"routeDescr" oasa:"RouteDescr"`
	Route_Descr_Eng string  `json:"routeDescrEng" oasa:"RouteDescrEng"`
	Route_Type      int8    `json:"routeType" oasa:"RouteType"`
	Route_Distance  float32 `json:"routeDistance" oasa:"RouteDistance"`
}

// ********* Struct for Bus Route Entities ****************
type Route struct {
	Id              int64   `json:"Id" gorm:"PrimaryKey"`
	Route_Code      int32   `json:"route_code" gorm:"index:ROUTE_CODE_UN,unique" oasa:"RouteCode"`
	Line_Code       int32   `json:"line_code" gorm:"index:LINE_CODE_INDX" oasa:"LineCode"`
	Route_Descr     string  `json:"route_descr" oasa:"RouteDescr"`
	Route_Descr_eng string  `json:"route_descr_eng" oasa:"RouteDescrEng"`
	Route_Type      int8    `json:"route_type" oasa:"RouteType"`
	Route_Distance  float32 `json:"route_distance" oasa:"RouteDistance"`
}

type RouteDto struct {
	Id              int64         `json:"Id"`
	Route_Code      int32         `json:"route_code"`
	Line_Code       int32         `json:"line_code"`
	Route_Descr     string        `json:"route_descr"`
	Route_Descr_Eng string        `json:"route_descr_eng"`
	Route_Type      int8          `json:"route_type"`
	Route_Distance  float32       `json:"route_distance"`
	Stops           []Stop        `json:"stops"`
	RouteDetails    []RouteDetail `json:"routeDetails"`
}

//**********************************************************

// ********* Struct for Route Details Entities **************
type RouteDetail struct {
	Routed_x     float32 `json:"routeLati" oasa:"routed_x"`
	Routed_y     float32 `json:"routeLong" oasa:"routed_y"`
	Routed_order int16   `json:"routeOrder" oasa:"routed_order"`
}

//**********************************************************
