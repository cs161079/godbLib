package main

import (
	"fmt"

	logger "github.com/cs161079/godbLib/Utils/goLogger"
	"github.com/cs161079/godbLib/db"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type LineTest struct {
	gorm.Model
	Id             int64  `json:"id" gorm:"primaryKey"`
	Ml_Code        int16  `json:"ml_code"`
	Sdc_Code       int16  `json:"sdc_code"`
	Line_Code      int32  `json:"line_code" gorm:"index:LINE_CODE,unique"`
	Line_Id        string `json:"line_id"`
	Line_Descr     string `json:"line_descr"`
	Line_Descr_Eng string `json:"line_descr_eng"`
	Mld_master     int8   `json:"is_master"`
}

func initialize() {
	logger.InitLogger("test")
	if err := godotenv.Load(".env"); err != nil {
		logger.ERROR("No .env file found")
	}
	err := db.IntializeDb()

	if err != nil {
		panic(err)
	}
	logger.INFO(fmt.Sprintf("database connection established!!"))
}

func testLineSave(input LineTest) error {
	var r *gorm.DB = nil
	r = db.DB.Table("line").Save(&input)
	if r.Error != nil {
		return r.Error
	}
	return nil
}

func selectLineTest(lineCode int32) (*LineTest, error) {
	var r *gorm.DB = nil
	var result LineTest
	r = db.DB.Table("line").Where("line_code = ?", lineCode).Find(&result)
	if r.Error != nil {
		return nil, r.Error
	}
	return &result, nil
}

//func main() {
//	initialize()
//	err := testLineSave(LineTest{
//		Line_Code: 1,
//		Ml_Code:   2, Sdc_Code: 3, Line_Id: "131", Line_Descr: "ΔΑΦΝΗ - ΑΓ. ΔΗΜΗΤΡΙΟΣ", Line_Descr_Eng: "DAFNI - AG. DIMITRIOS", Mld_master: 1,
//	})
//	if err != nil {
//		panic(err.Error())
//	}
//	line, err := selectLineTest(1)
//	if err != nil {
//		panic(err.Error())
//	}
//	fmt.Printf("This is the record %+v", line)
//}
