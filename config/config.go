package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)



var RunMode string
var ServerPort string
var DbServer string
var DbPort string
var DbUsername string
var DbPassword string
var DatabaseConnectionString string
var DatabaseName string
var Database string
var ServiceVersion string
var Serializer string

type SERIALIZER string
const (
	JSON SERIALIZER= "JSON"
	XML SERIALIZER= "XML"
	MSGPACK SERIALIZER= "MSGPACK"
)

type DATABASE string
const (
	MONGO DATABASE= "MONGO"
	MYSQL DATABASE= "MYSQL"
	IN_MEMORY DATABASE= "IN_MEMORY"
)

type SERVICE_VERSION string
const (
	V1 SERVICE_VERSION= "V1"
)
func InitEnvironmentVariables() {
	RunMode = os.Getenv("RUN_MODE")
	if RunMode == "" {
		RunMode = DEVELOP
	}

	log.Println("RUN MODE:", RunMode)

	if RunMode != PRODUCTION {
		//Load .env file
		err := godotenv.Load()
		if err != nil {
			log.Println("ERROR:", err.Error())
			return
		}
	}

	ServerPort = os.Getenv("SERVER_PORT")
	DbServer = os.Getenv("MONGO_SERVER")
	DbPort = os.Getenv("MONGO_PORT")
	DbUsername = os.Getenv("MONGO_USERNAME")
	DbPassword = os.Getenv("MONGO_PASSWORD")
	DatabaseName = os.Getenv("DATABASE_NAME")
	Database=os.Getenv("DATABASE")
	ServiceVersion=os.Getenv("SERVICE_VERSION")
	Serializer=os.Getenv("SERIALIZER")
	if Database==string(MONGO){
		DatabaseConnectionString = "mongodb://" + DbUsername + ":" + DbPassword + "@" + DbServer + ":" + DbPort
	}
	if Database==string(MYSQL){
		DatabaseConnectionString = DbUsername+":"+DbPassword+"@tcp("+DbServer+":"+DbPort+")/"+DatabaseName
	}



}

func InitMongoEnvironmentVariables(){

}
func InitMysqlEnvironmentVariables(){

}