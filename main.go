package main

import (
	"fmt"
	"time"

	mapper "github.com/cs161079/godbLib/Mappers"
	models "github.com/cs161079/godbLib/Models"
	db "github.com/cs161079/godbLib/Repositories"
	logger "github.com/cs161079/godbLib/Utils/goLogger"
	oasaSyncWeb "github.com/cs161079/godbLib/Web"
	"github.com/joho/godotenv"
)

func initEnviroment() {
	// loads values from .env into the system
	if err := godotenv.Load("enviroment.env"); err != nil {
		logger.ERROR("No .env file found")
	}
}

func initProgram() {
	logger.InitLogger("godbLib")
	initEnviroment()
	err := db.IntializeDb()

	if err != nil {
		logger.ERROR(err.Error())
	}
}

func main() {
	initProgram()
	start := time.Now()
	forTest()
	// var allData []models.LineDto
	// if err := deleteAllData(); err != nil {
	// 	logger.ERROR(err.Error())
	// }

	// if err := SyncData(&allData); err != nil {
	// 	logger.ERROR(err.Error())
	// }

	// if err := saveAllData(&allData); err != nil {
	// 	logger.ERROR(err.Error())
	// }
	defer func() {
		duration := time.Since(start)
		logger.INFO(fmt.Sprintf("Synchronization make %.2f to finish.", duration.Minutes()))
		if r := recover(); r != nil {
			logger.ERROR(fmt.Sprint("Panic occurred:", r))
		}
	}()

}

func forTest() {
	sdcTIme, err := db.SelectScheduleTime(1152, 54)
	if err != nil {
		logger.ERROR(err.Error())
	}
	logger.INFO(fmt.Sprintf("%+v", sdcTIme))
	err = deleteAllData()
	if err != nil {
		logger.ERROR(err.Error())
	}
	// ****** Http Request to get All Bus Lines ******
	tt, err := time.ParseInLocation("2006-01-02 15:04:05", "1900-01-01 05:10:00", time.Now().Location())
	if err != nil {
		logger.ERROR(err.Error())
	}
	err = db.SaveScheduleTime(models.ScheduleTime{
		Sdc_Code:   54,
		Line_Code:  1152,
		Start_Time: tt,
		Type:       0,
	})
	if err != nil {
		logger.ERROR(err.Error())
	}
}

func deleteAllData() error {
	trans := db.DB.Begin()
	if trans.Error != nil {
		return trans.Error
	}

	if err := db.DeleteRouteStops(trans); err != nil {
		return err
	}
	if err := db.DeleteStop(trans); err != nil {
		return err
	}
	if err := db.DeleteRoute(trans); err != nil {
		return err
	}
	if err := db.DeleteScheduleTime(trans); err != nil {
		return err
	}
	if err := db.DeleteScheduleMaster(trans); err != nil {
		return err
	}
	if err := db.DeleteLines(trans); err != nil {
		return err
	}

	if err := db.ZeroAllSequence(trans); err != nil {
		return err
	}

	if err := trans.Commit().Error; err != nil {
		return err
	}
	logger.INFO("ERASE ALL DATA FROM DATABASE FINISHED")
	return nil

}

func saveAllData(allData *[]models.LineDto) error {
	logger.INFO("INSERTING DATA IN DATABASE START")
	for _, line := range *allData {
		err := db.SaveLine(mapper.LineDtoToLine(line))
		if err != nil {
			return err
		}
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
	logger.INFO("INSERTING DATA IN DATABASE FINISHED")
	return nil
}

func SyncData(allData *[]models.LineDto) error {
	logger.INFO("SYNC DATA FROM OASA START")
	// ****** Http Request to get All Bus Lines ******
	lines, err := GetBusLines()
	if err != nil {
		return err
	}
	// ***********************************************
	// ******** For loop in array of lines ***********
	for i := range lines[0:3] {
		// **** Http Request to get all Routes for every Line *******
		lines[i].Routes, err = GetBusRoutes(lines[i].Line_Code)
		if err != nil {
			return err
		}
		// **********************************************************
		// ************* For Loop in array of Routes ****************
		for j := range lines[i].Routes {
			// *** Http Request to Get stop for every route in array ****
			lines[i].Routes[j].Stops, err = GetBusStops(lines[i].Routes[j].Route_Code)
			if err != nil {
				return err
			}
			// ***********************************************************
			// ******* Http Request to get details for every route *******
			lines[i].Routes[j].RouteDetails, err = GetRouteDetails(lines[i].Routes[j].Route_Code)
			if err != nil {
				return err
			}
			// **********************************************************
		}
		// ********* Http Request get Schedule Master Line **************
		lines[i].Schedules, err = GetScheduleMasterLine(int64(lines[i].Line_Code))
		if err != nil {
			return err
		}
		// **************************************************************
		for k := range lines[i].Schedules {
			result, err := GetScheduleTime(int32(lines[i].Ml_Code), lines[i].Line_Code, lines[i].Schedules[k].Sdc_Code)
			if err != nil {
				return err
			}
			lines[i].Schedules[k].ShcedeLine = *result
		}
		dailyTimes, err := GetDailySchedule(lines[i].Line_Code)
		if err != nil {
			return err
		}
		lines[i].Schedules = append(lines[i].Schedules, *dailyTimes)
	}
	// **************************************************
	*allData = lines[0:3]
	logger.INFO("SYNC DATA FROM OASA FINISHED")
	return nil
}

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

// func GetScheduleLines(lineCode int64, sdcCode int32)
func GetScheduleLine(mlCode int64, sdcCode int64, lineCode int64) (*models.ScheduleTimes, error) {
	var result models.ScheduleTimes
	response := oasaSyncWeb.OasaRequestApi("getSchedLines", map[string]interface{}{"p1": mlCode, "p2": sdcCode, "p3": lineCode})
	if response.Error != nil {
		return nil, response.Error
	}
	result = mapper.ScheduleTimesMapper(response.Data.(map[string]interface{}))
	return &result, nil

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

func GetBusLines() ([]models.LineDto, error) {
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
			result = append(result, mapper.RouteDetailDtoMapper(record.(map[string]interface{})))
		}
	}
	// fmt.Printf("Routes results %d \n", len(result))

	return result, nil
}
