package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // using postgres sql
	"github.com/spf13/viper"
)

func SetupModels() *gorm.DB {

	// Enable VIPER to read Environment Variables
	viper.AddConfigPath(".")
    viper.SetConfigName("database")
    viper.SetConfigType("env")
    viper.AutomaticEnv()

    err := viper.ReadInConfig()
    if err != nil {
        panic(err)
    }
	viper.AutomaticEnv()

	// viper package read .env
	viper_user := viper.Get("POSTGRES_USER")
	viper_password := viper.Get("POSTGRES_PASSWORD")
	viper_db := viper.Get("POSTGRES_DB")
	viper_host := viper.Get("POSTGRES_HOST")
	viper_port := viper.Get("POSTGRES_PORT")

	// https://gobyexample.com/string-formatting
	prosgret_conname := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", viper_host, viper_port, viper_user, viper_db, viper_password)

	fmt.Println("conname is\t\t", prosgret_conname)

	db, err := gorm.Open("postgres", prosgret_conname)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Book{})
	db.AutoMigrate(&User{})

	return db
}
