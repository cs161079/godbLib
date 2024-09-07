package repository

import (
	models "github.com/cs161079/godbLib/Models"
	"github.com/cs161079/godbLib/db"
	"gorm.io/gorm"
)

type RouteRepository interface {
	SelectByLineCode(int32) (*models.Route, error)
	SelectRouteByLineCode(int32) (*[]models.Route, error)
	InsertRoute(input models.Route) (*models.Route, error)
	UpdateRoute(input models.Route) (*models.Route, error)
	RouteList01() ([]models.Route, error)
	DeleteRoute() error
}

type routeRepository struct {
	DB *gorm.DB
}

func (r routeRepository) SelectByRouteCode(routeCode int32) (*models.Route, error) {
	var selectedVal models.Route
	dbRes := r.DB.Table(db.ROUTETABLE).Where("route_code = ?", routeCode).Find(&selectedVal)
	if dbRes != nil {
		if dbRes.Error != nil {
			// fmt.Println(r.Error.Error())
			return nil, dbRes.Error
		}
		if dbRes.RowsAffected == 0 {
			//logger.WARN(fmt.Sprintf("BUS ROUTE NOT FOUND [ROUTE_CODE: %d]", routeCode))
			return nil, nil
		}
	}
	return &selectedVal, nil
}

func (r routeRepository) SelectRouteByLineCode(lineCode int32) (*[]models.Route, error) {
	var selectedVal []models.Route
	res := r.DB.Table(db.ROUTETABLE).Where("line_code = ?", lineCode).Find(&selectedVal)
	if res.Error != nil {
		return nil, res.Error
	}
	return &selectedVal, nil
}

func (r routeRepository) InsertRoute(input models.Route) (*models.Route, error) {
	res := r.DB.Table(db.ROUTETABLE).Create(&input)
	if res.Error != nil {
		return nil, res.Error
	}
	return &input, nil
}

func (r routeRepository) UpdateRoute(input models.Route) (*models.Route, error) {
	res := r.DB.Table(db.ROUTETABLE).Create(&input)
	if res.Error != nil {
		return nil, res.Error
	}
	return &input, nil
}

// NOTE: Αυτά πρέπει να μεταφερθούν στο Repository του RouteDetail
// func (r routeRepository) SelecetDetailRouteCode(routeCode int32) (*models.RouteDetail, error) {
// 	var selectedVal models.RouteDetail
// 	r := DB.Table(ROUTEDETAILTABLE).Where("route_code = ?", routeCode).Find(&selectedVal)
// 	if r != nil {
// 		if r.Error != nil {
// 			// fmt.Println(r.Error.Error())
// 			return nil, r.Error
// 		}
// 		if r.RowsAffected == 0 {
// 			//logger.WARN(fmt.Sprintf("BUS ROUTE NOT FOUND [ROUTE_CODE: %d]", routeCode))
// 			return nil, nil
// 		}
// 	}
// 	return &selectedVal, nil
// }

// func SaveRouteDetails(input models.RouteDetail) error {
// 	r := DB.Table(ROUTEDETAILTABLE).Save(&input)
// 	if r.Error != nil {
// 		return r.Error
// 	}
// 	return nil
// }

func (r routeRepository) RouteList01() ([]models.Route, error) {
	var result []models.Route
	res := r.DB.Table(db.ROUTETABLE).Order("route_code").Find(&result)
	if res.Error != nil {
		return nil, res.Error
	}
	return result, nil
}

func (r routeRepository) DeleteRoute() error {
	if err := r.DB.Table(db.ROUTETABLE).Where("1=1").Delete(&models.Route{}).Error; err != nil {
		//trans.Rollback()
		return err
	}
	return nil
}
