package configs

import (
	"database/sql"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/viper"
	"log"
)
import _ "github.com/go-sql-driver/mysql"

type Config struct {
	DbHost     string `mapstructure:"DB_HOST"`
	DbPort     string `mapstructure:"DB_PORT"`
	DbUser     string `mapstructure:"DB_USER"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbName     string `mapstructure:"DB_NAME"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}

	fmt.Printf("Loaded config: %+v\n", config)
	return
}

func DB() (*sql.DB, error) {
	config, err := LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	mysqlString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DbUser, config.DbPassword, config.DbHost, config.DbPort, config.DbName)
	fmt.Println("MySQL connection string:", mysqlString)

	conn, err := sql.Open("mysql", mysqlString)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	return conn, nil
}
