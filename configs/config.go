package configs

import (
	"errors"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
)

type ProgrammingConfig struct {
	Server    int
	DBPort    int
	DBHost    string
	DBUser    string
	DBPass    string
	DBName    string
	Secret    string
	RefSecret string
}

func InitConfig() *ProgrammingConfig {
	var res = new(ProgrammingConfig)
	res, errorRes := loadConfig()

	logrus.Error(errorRes)
	if res == nil {
		logrus.Error("CONFIG : Cannot start program, Failed to load config")
		return nil
	}

	return res
}

func loadConfig() (*ProgrammingConfig, error) {
	var errorLoad error
	var res = new(ProgrammingConfig)
	var permit = true

	if val, found := os.LookupEnv("SERVER"); found {
		port, err := strconv.Atoi(val)
		if err != nil {
			logrus.Error("CONFIG : Invalid Port Value : ", err.Error())
			permit = false
		} else {
			res.Server = port
		}
	} else {
		permit = false
		errorLoad = errors.New("SERVER PORT UNDEFINED")
	}

	if val, found := os.LookupEnv("DBPORT"); found {
		port, err := strconv.Atoi(val)
		if err != nil {
			logrus.Error("CONFIG : Invalid DB Port Value : ", err.Error())
			permit = false
		} else {
			res.DBPort = port
		}
	} else {
		permit = false
		errorLoad = errors.New("DBPORT UNDEFINED")
	}

	if val, found := os.LookupEnv("DBHOST"); found {
		res.DBHost = val
	} else {
		permit = false
		errorLoad = errors.New("DBHOST UNDEFINED")
	}

	if val, found := os.LookupEnv("DBUSER"); found {
		res.DBUser = val
	} else {
		permit = false
		errorLoad = errors.New("DBUSER UNDEFINED")
	}

	if val, found := os.LookupEnv("DBPASS"); found {
		res.DBPass = val
	} else {
		permit = false
		errorLoad = errors.New("DBPASS UNDEFINED")
	}

	if val, found := os.LookupEnv("DBNAME"); found {
		res.DBName = val
	} else {
		permit = false
		errorLoad = errors.New("DBNAME UNDEFINED")
	}

	if val, found := os.LookupEnv("SECRET"); found {
		res.Secret = val
	} else {
		permit = false
		errorLoad = errors.New("SECRET UNDEFINED")
	}

	if val, found := os.LookupEnv("REFSECRET"); found {
		res.RefSecret = val
	} else {
		permit = false
		errorLoad = errors.New("REFSECRET UNDEFINED")
	}

	if !permit {
		return nil, errorLoad
	}

	return res, nil
}
