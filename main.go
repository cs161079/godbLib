package main

import (
	"fmt"
	"os"
	"strings"

	models "github.com/cs161079/godbLib/Models"
	logger "github.com/cs161079/godbLib/Utils/goLogger"
	oasaSyncWeb "github.com/cs161079/godbLib/Web"
	"github.com/cs161079/godbLib/db"
	mapper "github.com/cs161079/godbLib/mapper"
	"github.com/cs161079/godbLib/repository"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func initEnviroment() {
	// loads values from .env into the system
	if err := godotenv.Load(".env"); err != nil {
		logger.ERROR("No .env file found")
	}
}

func initProgram() {
	logger.InitLogger("godbLib")
	initEnviroment()
	err := db.IntializeDb()

	if err != nil {
		panic(err)
	}
	logger.INFO(fmt.Sprintf("database connection established!!"))
}

func deleteAllData(connection *gorm.DB) error {
	//trans := db.DB.Begin()
	//if trans.Error != nil {
	//	return trans.Error
	//}

	if err := db.DeleteRouteStops(connection); err != nil {
		//trans.Rollback()
		return err
	}
	if err := db.DeleteStop(connection); err != nil {
		//trans.Rollback()
		return err
	}
	if err := db.DeleteRoute(connection); err != nil {
		//trans.Rollback()
		return err
	}
	if err := db.DeleteScheduleTime(connection); err != nil {
		//trans.Rollback()
		return err
	}
	if err := db.DeleteScheduleMaster(connection); err != nil {
		//trans.Rollback()
		return err
	}
	//if err := db.DeleteLines(connection); err != nil {
	//	//trans.Rollback()
	//	return err
	//}
	//
	//if err := db.ZeroAllSequence(trans); err != nil {
	//	return err
	//}

	//if err := trans.Commit().Error; err != nil {
	//	return err
	//}
	logger.INFO("ERASE ALL DATA FROM DATABASE FINISHED")
	return nil

}

func saveAllData(connection *gorm.DB, allData *[]models.LineDto) error {
	//logger.INFO("INSERTING DATA IN DATABASE START")
	for _, line := range *allData {
		//err := db.SaveLine(connection, mapper.LineDtoToLine(line))
		//if err != nil {
		//	return err
		//}
		for _, route := range line.Routes {
			err := db.SaveRoute(mapper.RouteDtoToRoute(route))
			if err != nil {
				return err
			}

			for i, stop := range route.Stops {
				err := db.SaveStop(stop)
				if err != nil {
					return err
				}
				err = db.SaveRouteStops(models.RouteStops{
					Route_code: route.Route_Code,
					Stop_code:  stop.Stop_code,
					Senu:       int16(i + 1),
				})
				if err != nil {
					return err
				}
			}
		}

		for _, schedule := range line.Schedules {
			finalSchedule := mapper.ScheduleMasterDtoToScheduleMaster(schedule)
			finalSchedule.Line_Code = int64(line.Line_Code)
			err := db.SaveScheduleMaster(finalSchedule)
			if err != nil {
				return err
			}
			for _, scheduleTime := range schedule.ShcedeLine.Go {
				finalScheduleTime := mapper.ScheduleTimeDtoToScheduleTime(scheduleTime)
				finalScheduleTime.Sdc_Code = schedule.Sdc_Code
				finalScheduleTime.Line_Code = int64(line.Line_Code)
				err := db.SaveScheduleTime(finalScheduleTime)
				if err != nil {
					return err
				}
			}
			for _, scheduleTime := range schedule.ShcedeLine.Come {
				finalScheduleTime := mapper.ScheduleTimeDtoToScheduleTime(scheduleTime)
				finalScheduleTime.Sdc_Code = schedule.Sdc_Code
				finalScheduleTime.Line_Code = int64(line.Line_Code)
				err := db.SaveScheduleTime(finalScheduleTime)
				if err != nil {
					return err
				}
			}

		}
	}
	//logger.INFO("INSERTING DATA IN DATABASE FINISHED")
	return nil
}

//func SyncData(allData *[]models.LineDto) error {
//	//logger.INFO("SYNC DATA FROM OASA START")
//	// ****** Http Request to get All Bus Lines ******
//	lines, err := GetBusLines()
//	if err != nil {
//		return err
//	}
//
//	// ***********************************************
//	//var currentLineId = lines[0].Line_Id
//	// ******** For loop in array of lines ***********
//	for i := range lines {
//		// **** Http Request to get all Routes for every Line *******
//		lines[i].Routes, err = GetBusRoutes(lines[i].Line_Code)
//		if err != nil {
//			return err
//		}
//		// **********************************************************
//		// ************* For Loop in array of Routes ****************
//		for j := range lines[i].Routes {
//			lines[i].Routes[j].Line_Code = lines[i].Line_Code
//			// *** Http Request to Get stop for every route in array ****
//			lines[i].Routes[j].Stops, err = GetBusStops(lines[i].Routes[j].Route_Code)
//			if err != nil {
//				return err
//			}
//			// ***********************************************************
//			// ******* Http Request to get details for every route *******
//			lines[i].Routes[j].RouteDetails, err = GetRouteDetails(lines[i].Routes[j].Route_Code)
//			if err != nil {
//				return err
//			}
//			// **********************************************************
//		}
//		// ********* Http Request get Schedule Master Line **************
//		lines[i].Schedules, err = GetScheduleMasterLine(int64(lines[i].Line_Code))
//		if err != nil {
//			return err
//		}
//		// **************************************************************
//		for k := range lines[i].Schedules {
//			result, err := GetScheduleTime(int32(lines[i].Ml_Code), lines[i].Line_Code, lines[i].Schedules[k].Sdc_Code)
//			if err != nil {
//				return err
//			}
//			lines[i].Schedules[k].ShcedeLine = *result
//		}
//		dailyTimes, err := GetDailySchedule(lines[i].Line_Code)
//		if err != nil {
//			return err
//		}
//		lines[i].Schedules = append(lines[i].Schedules, *dailyTimes)
//		//printProgressBar(i, len(tempLines), "Progress", "Complete", 25, "=")
//	}
//	// **************************************************
//	*allData = lines
//	///logger.INFO("SYNC DATA FROM OASA FINISHED")
//	return nil
//}

func GetDailySchedule(lineCode int32) (*models.ScheduleMasterDto, error) {
	response := oasaSyncWeb.OasaRequestApi("getDailySchedule", map[string]interface{}{"line_code": int64(lineCode)})
	if response.Error != nil {
		return nil, response.Error
	}
	var dailySchedule models.ScheduleMasterDto = models.ScheduleMasterDto{
		Sdc_Code:      0,
		Sdc_Descr:     "ΚΑΘΗΜΕΡΙΝΟ ΠΡΟΓΡΑΜΜΑ",
		Sdc_Descr_Eng: "DAILY SCHEDULE",
	}
	dailySchedule.ShcedeLine = mapper.ScheduleTimesMapper(response.Data.(map[string]interface{}))
	return &dailySchedule, nil
}

func GetScheduleMasterLine(lineCode int64) ([]models.ScheduleMasterDto, error) {
	var result []models.ScheduleMasterDto
	response := oasaSyncWeb.OasaRequestApi("getScheduleDaysMasterline", map[string]interface{}{"p1": lineCode})
	if response.Error != nil {
		return nil, response.Error
	}

	if response.Data != nil {
		for _, record := range response.Data.([]interface{}) {
			result = append(result, mapper.ScheduleOasaToScheduleDto(mapper.ScheduleMasterLine(record.(map[string]interface{}))))
		}
	}

	return result, nil
}

func GetScheduleTime(masteLineCode int32, lineCode int32, sdcCode int32) (*models.ScheduleTimes, error) {
	var result models.ScheduleTimes
	response := oasaSyncWeb.OasaRequestApi("getSchedLines", map[string]interface{}{"p1": int64(masteLineCode), "p2": int64(sdcCode), "p3": int64(lineCode)})
	if response.Error != nil {
		return nil, response.Error
	}
	result = mapper.ScheduleTimesMapper(response.Data.(map[string]interface{}))
	return &result, nil
}

func GetLineData() ([]models.LineDto, error) {
	var result []models.LineDto
	response := oasaSyncWeb.OasaRequestApi("webGetLinesWithMLInfo", nil)
	if response.Error != nil {
		return nil, response.Error
	}

	if response.Data != nil {
		for _, record := range response.Data.([]interface{}) {
			result = append(result, mapper.LineOasaToLine(mapper.LineMapper(record.(map[string]interface{}))))
		}
	}

	return result, nil
}

func GetBusRoutes(linedId int32) ([]models.RouteDto, error) {
	var result []models.RouteDto
	response := oasaSyncWeb.OasaRequestApi("webGetRoutes", map[string]interface{}{"p1": int64(linedId)})
	if response.Error != nil {
		return nil, response.Error
	}

	if response.Data != nil {
		for _, record := range response.Data.([]interface{}) {
			// fmt.Println("at Index Route ", index)
			result = append(result, mapper.RouteOasaToRouteDto(mapper.RouteMapper(record.(map[string]interface{}))))
		}
	} // fmt.Printf("Routes results %d \n", len(result))

	return result, nil
}

func GetBusStops(routeCode int32) ([]models.Stop, error) {
	var result []models.Stop
	response := oasaSyncWeb.OasaRequestApi("webGetStops", map[string]interface{}{"p1": int64(routeCode)})
	if response.Error != nil {
		return nil, response.Error
	}

	if response.Data != nil {
		for _, record := range response.Data.([]interface{}) {
			// fmt.Println("at Index Route ", index)
			result = append(result, mapper.StopOasaToStop(mapper.StopMapper(record.(map[string]interface{}))))
		}
	}
	// fmt.Printf("Routes results %d \n", len(result))

	return result, nil
}

func GetRouteDetails(routeCode int32) ([]models.RouteDetail, error) {
	var result []models.RouteDetail
	response := oasaSyncWeb.OasaRequestApi("webRouteDetails", map[string]interface{}{"p1": int64(routeCode)})
	if response.Error != nil {
		return nil, response.Error
	}

	if response.Data != nil {
		for _, record := range response.Data.([]interface{}) {
			// fmt.Println("at Index Route ", index)
			result = append(result, mapper.RouteDetailOasaToRouteDetailDto(mapper.RouteDetailDtoMapper(record.(map[string]interface{}))))
		}
	}
	// fmt.Printf("Routes results %d \n", len(result))

	return result, nil
}

//	func GetDataVersions() error {
//		var result interface{}
//		response := oasaSyncWeb.OasaRequestApi("getUVersions", nil)
//		if response.Error != nil {
//			return response.Error
//		}
//	}

func testForConnection(conn *gorm.DB) error {
	return nil
}

// ****************************** TEST ***************************************************
// ***** Αυτό είνα ένα Τεστ για αν αποθηκευσώ τα δεδομένα από τα API συγχρονισμού ********
// ***** σε αρχεία και να αξιολογήσω τα δεδομένα και να αποφασίσω πως θα κάνω     ********
// ***** τον συγχρονισμό                                                          ********
// ***************************************************************************************
func testSyncData(action string) error {
	resp, err := oasaSyncWeb.MakeRequest(action)
	if err != nil {
		return err
	}
	orgResp := *resp
	trimmedResp := orgResp[1 : len(orgResp)-2]
	syncDataLines := strings.Split(trimmedResp, "),(")
	f, err := os.Create(fmt.Sprintf("tmp/%s.txt", action))
	defer f.Close()
	if err != nil {
		return err
	}
	for _, line := range syncDataLines {
		if _, err = f.Write([]byte(fmt.Sprintf("%s\n", line))); err != nil {
			return err
		}
	}
	return nil
}

func getUIVersions() (*models.UVersions, error) {
	var result models.UVersions
	response := oasaSyncWeb.OasaRequestApi("getUVersions", nil)
	if response.Error != nil {
		return nil, response.Error
	}
	result = mapper.UVersionsOasaToUVersions(mapper.UVersionsMapper(response.Data.(map[string]interface{})))
	return &result, nil
}

func arrayToMap(dataArray []models.UVersions) map[string]int64 {
	// Build a config map:
	confMap := map[string]int64{}
	for _, v := range dataArray {
		confMap[v.Uv_descr] = v.Uv_lastupdatelong
	}
	return confMap
}
func checkUVerisios(dbVersions []models.UVersions, oasaVersions []models.UVersions) error {
	var oasaVersionsMp = arrayToMap(oasaVersions)
	var dbVersionsMp = arrayToMap(dbVersions)
	var routeDetailsUpdate bool = false
	if oasaVersionsMp[models.UVERSIONS_LINE] > dbVersionsMp[models.UVERSIONS_LINE] {
		// Download Data for Lines, Update Database
		_, err := GetLineData()
		if err != nil {
			return err
		}
		// Update Uversion Table with new revision
	}
	if oasaVersionsMp[models.UVERSIONS_ROUTE] > dbVersionsMp[models.UVERSIONS_ROUTE] {
		//Dowload all Data for Routes , update Databaseo
		// Update Uversion Table with new revision
		// if Routes updated must update and RouteDetails
		routeDetailsUpdate = true
		// if Stop Updated must update RouteStops Probably the revision of ROUTESTOPS will have been updated as well
	}
	if oasaVersionsMp[models.UVERSIONS_STOP] > dbVersionsMp[models.UVERSIONS_STOP] {
		//Dowload all Data for Stops , update Database
		// Update Uversion Table with new revision
		// if Stop Updated must update RouteStops Probably the revision of ROUTESTOPS will have been updated as well
	}
	if oasaVersionsMp[models.UVERSIONS_ROUTESTOPS] > dbVersionsMp[models.UVERSIONS_ROUTESTOPS] {
		//Dowload all Data for RouteStops , update Database
		// Update Uversion Table with new revision
	}
	if routeDetailsUpdate {
		// Update Route details from Oasa Server
	}
	return nil

}

func main() {
	defer func() {
		if r := recover(); r != nil {
			logger.ERROR(fmt.Sprintf("Panic occurred: %+v\n", r))
		}
	}()

	connection, err := db.CreateConnection()
	if err != nil {
		panic(err.Error())
	}
	repository.NewLineRepository(connection)
	// repo := &db.Repository{Db: db.DB}
	// service := &db.LineService{Repo: repo}
	// err := service.PostLineArray([]models.Line{
	// 	models.Line{
	// 		Line_Code:      1,
	// 		Line_Descr:     "Αυτή είναι δοκιμαστική γραμμή",
	// 		Line_Descr_Eng: "This is a test Line",
	// 		Line_Id:        "131",
	// 		Sdc_Code:       1,
	// 		Mld_master:     1,
	// 		Ml_Code:        1,
	// 	},
	// 	models.Line{
	// 		Line_Code:      1,
	// 		Line_Descr:     "Αυτή είναι δοκιμαστική γραμμή 2",
	// 		Line_Descr_Eng: "This is a test Line 2",
	// 		Line_Id:        "131_2",
	// 		Sdc_Code:       1,
	// 		Mld_master:     1,
	// 		Ml_Code:        1,
	// 	},
	// })
	// if err != nil {
	// 	logger.ERROR(fmt.Sprintf("Transaction failed %v", err))
	// } else {
	// 	logger.INFO("Transaction Succeed")
	// }
	//if err := testSyncData("getLines"); err != nil {
	//	logger.ERROR(fmt.Sprintf("Synchronize data failed. %v", err))
	//	return
	//}
	//logger.INFO("Sychronization succed.")
	//return
	//repo := &repository{db: db.DB}
	//lineService := &LineService{repo: repo}
	//ctx := context.Background()
	//if err := lineService.LinePost(ctx, &models.Line{
	//	Line_Code:      1,
	//	Line_Descr:     "Αυτή είναι δοκιμαστική γραμμή",
	//	Line_Descr_Eng: "This is a test Line",
	//	Line_Id:        "131",
	//	Sdc_Code:       1,
	//	Mld_master:     1,
	//	Ml_Code:        1,
	//}); err != nil {
	//	logger.ERROR(fmt.Sprintf("Transaction failed %v", err))
	//} else {
	//	logger.INFO("Transaction Succeed")
	//}
	//return
	//var err error
	//// ********** Make Request for test ******************************
	//
	// resp, err := oasaSyncWeb.MakeRequest("getSched_entriesW")
	// if err != nil {
	// 	logger.ERROR(err.Error())
	// 	return
	// }
	// logger.INFO(*resp)
	// return
	//
	//// **************** Delete all data from database *****************
	//err = deleteAllData()
	//if err != nil {
	//	panic(err)
	//}
	//
	//// ****************************************************************
	//// **************** Sync Data from OASA Server ********************
	//logger.INFO("Synchronization start...")
	//var startSync = time.Now()
	//var lineData []models.LineDto
	//err = SyncData(&lineData)
	//if err != nil {
	//	panic(err.Error())
	//	//logger.ERROR(err.Error())
	//}
	//var syncDuration = time.Since(startSync)
	//logger.INFO(fmt.Sprintf("Synchronization Data from OASA Server take %.2f minutes to finish.", syncDuration.Minutes()))
	////*****************************************************************
	//
	//// **************** Save all Line data in Database *****************
	//logger.INFO("Save data in database start...")
	//var startSave = time.Now()
	//err = saveAllData(&lineData)
	//if err != nil {
	//	panic(err.Error())
	//}
	//var saveDuration = time.Since(startSave)
	//logger.INFO(fmt.Sprintf("Save data in database take %.2f minutes to finih.", saveDuration.Minutes()))
	////*****************************************************************

}
