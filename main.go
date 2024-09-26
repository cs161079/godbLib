package main

import (
	"fmt"
	"os"
	"time"

	models "github.com/cs161079/godbLib/Models"
	logger "github.com/cs161079/godbLib/Utils/goLogger"
	"github.com/cs161079/godbLib/db"
	mapper "github.com/cs161079/godbLib/mapper"
	"github.com/cs161079/godbLib/repository"
	"github.com/cs161079/godbLib/service"
	"github.com/joho/godotenv"
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
	_, err := db.CreateConnection()

	if err != nil {
		panic(err)
	}
	logger.INFO(fmt.Sprintf("database connection established!!"))
}

// ****************************** TEST ***************************************************
// ***** Αυτό είνα ένα Τεστ για αν αποθηκευσώ τα δεδομένα από τα API συγχρονισμού ********
// ***** σε αρχεία και να αξιολογήσω τα δεδομένα και να αποφασίσω πως θα κάνω     ********
// ***** τον συγχρονισμό                                                          ********
// ***************************************************************************************
func testSyncData(action string) error {
	restSrv := service.NewRestService()
	response := restSrv.OasaRequestApi02(action)
	if response.Error != nil {
		return response.Error
	}
	orgResp := (*response).Data.([]string)
	//trimmedResp := orgResp[1 : len(orgResp)-2]
	//syncDataLines := strings.Split(trimmedResp, "),(")
	f, err := os.Create(fmt.Sprintf("tmp/%s%s.txt", action, time.Now().Format("2006-01-02_15_04_05")))
	defer f.Close()
	if err != nil {
		return err
	}
	for _, line := range orgResp {
		if _, err = f.Write([]byte(fmt.Sprintf("%s\n", line))); err != nil {
			return err
		}
	}
	return nil
}

func getUIVersions() (*models.UVersions, error) {
	restSrv := service.NewRestService()
	var result models.UVersions
	response := restSrv.OasaRequestApi00("getUVersions", nil)
	if response.Error != nil {
		return nil, response.Error
	}
	result = mapper.OasaToUVersions(mapper.GeneralUVersions(response.Data.(map[string]interface{})))
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
	lineRepository := repository.NewLineRepository(connection)
	service.NewLineService(lineRepository)

	err = testSyncData("getSched_series")
	if err != nil {
		panic(err.Error())
	}
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
