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
var PrivateKey string
var Publickey string

type DATABASE string

const (
	MONGO DATABASE= "MONGO"
	MYSQL DATABASE= "MYSQL"
	IN_MEMORY DATABASE= "IN_MEMORY"
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
	PrivateKey =os.Getenv("PRIVATE_KEY")
	Publickey=os.Getenv("PUBLIC_KEY")

	if Database==string(MONGO){
		DatabaseConnectionString = "mongodb://" + DbUsername + ":" + DbPassword + "@" + DbServer + ":" + DbPort
	}
	if Database==string(MONGO){
		DatabaseConnectionString = DbUsername+":"+DbPassword+"@tcp("+DbServer+":"+DbPort+")/"+DatabaseName
	}



}

func InitMongoEnvironmentVariables(){

}
func InitMysqlEnvironmentVariables(){

}