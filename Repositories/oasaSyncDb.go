package db

import (
	"fmt"
	"os"
	"strconv"
	"time"

	logger "godbLib/Utils/goLogger"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DataSource struct {
	Address      *string
	Port         *int32
	User         string
	Password     string
	DatabaseName string
}

const dataSourceFormat = "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"

type DatasourceLink interface {
	DatasourceUrl() (string, error)
}

func (ds DataSource) DatasourceUrl() string {
	return fmt.Sprintf(dataSourceFormat, ds.User, ds.Password,
		*ds.Address, *ds.Port, ds.DatabaseName)

}

func getDataSource() (*DataSource, error) {
	// err := godotenv.Load()
	// if err != nil {
	// 	return nil, err
	// }
	var ip = os.Getenv("database.ip")
	port, err := strconv.ParseInt(os.Getenv("database.port"), 10, 32)
	if err != nil {
		return nil, err
	}
	var port32 = int32(port)
	return &DataSource{
		Address:      &ip,
		Port:         &port32,
		User:         "user1",
		Password:     "user1password",
		DatabaseName: "oasaDb",
	}, nil
}

// This is core for DB

var DB *gorm.DB

func IntializeDb() error {
	dataSource, err := getDataSource()
	if err != nil {
		return err
	}
	var defaultIp = "127.0.0.1"
	var defaultPort int32 = 3306
	if dataSource.Address == nil {
		dataSource.Address = &defaultIp
	}
	if dataSource.Port == nil {
		dataSource.Port = &defaultPort
	}

	dialector := mysql.New(mysql.Config{
		DSN:                       dataSource.DatasourceUrl(), // data source name
		DefaultStringSize:         256,                        // default size for string fields
		DisableDatetimePrecision:  true,                       // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,                       // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                       // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                      // auto configure based on currently MySQL version
	})

	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	// 	logger.Config{
	// 		SlowThreshold: time.Second, // Slow SQL threshold
	// 		LogLevel:      logger.Info, // Log level
	// 		// IgnoreRecordNotFoundError: true,         // Ignore ErrRecordNotFound error for logger
	// 		ParameterizedQueries: true, // Don't include params in the SQL log
	// 		Colorful:             true, // Disable color
	// 	},
	// )

	database, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.GetGormLogger(),
	})

	if err != nil {
		// fmt.Println("An Error occured on creation of database connection")
		return err
	}

	sqlDb, err := database.DB()
	if err != nil {
		// fmt.Println("An Error Occured... ", err.Error())
		return err
	}
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetConnMaxLifetime(time.Minute)
	sqlDb.SetMaxOpenConns(100)

	DB = database
	return nil
}
