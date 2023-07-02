package configs

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/viper"
	"log"
)

import _ "github.com/go-sql-driver/mysql"

type Config struct {
	DbHost        string `mapstructure:"DB_HOST"`
	DbPort        string `mapstructure:"DB_PORT"`
	Migration_Url string `mapstructure:"MIGRATION_URL"`
	DbUser        string `mapstructure:"DB_USER"`
	DbPassword    string `mapstructure:"DB_PASSWORD"`
	DbName        string `mapstructure:"DB_NAME"`
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
func runDBmigration(migrationUrl string, db string) {
	migration, err := migrate.New(migrationUrl, db)
	if err != nil {
		log.Fatal("failed to migrate db")
	}
	if err = migration.Up(); err != nil && migrate.ErrNoChange != nil {
		log.Fatal("failed to migrate db")
	}
	log.Println("successfully migrated db")
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
	runDBmigration(config.Migration_Url, config.DbPort)
	return conn, nil
}
