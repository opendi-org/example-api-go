/**
 * Database factory. Abstracts the configuration and construction of a couple types of databases,
 * so the API can use any GORM-supported database type.
 *
 * Currently implemented: SQLite, MySQL
 *
 * Want to add support for another database type? Check this project's README for detailed instructions.
 *
 */

package db

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/gorm"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
)

/**
 * Simple struct for holding environment variables related to database construction/configuration
 */
type DBConfig struct {
	DBType     string
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
}

/**
 * Get a populated DBConfig struct by reading the OS environment variables directly.
 * If environment variables are not set, they will use default values.
 */
func GetConfig() (*DBConfig, error) {
	portString := getEnvironmentValue("DB_PORT", "3306")
	port, err := strconv.Atoi(portString)
	if err != nil {
		return nil, fmt.Errorf("invalid DB_PORT: %v", err)
	}

	return &DBConfig{
		DBType:     getEnvironmentValue("DATABASE_TYPE", "sqlite"),
		DBHost:     getEnvironmentValue("DB_HOST", "http://db"),
		DBPort:     port,
		DBUser:     getEnvironmentValue("DB_USER", "myuser"),
		DBPassword: getEnvironmentValue("DB_PASSWORD", "defaultpassword"),
		DBName:     getEnvironmentValue("DB_NAME", "modelsdb"),
	}, nil
}

/**
 * Read an individual environment variable key. If the variable is not set,
 * use the provided default value instead.
 */
func getEnvironmentValue(key string, defaultValue string) string {
	if retrievedValue := os.Getenv(key); retrievedValue != "" {
		return retrievedValue
	}
	return defaultValue
}

/**
 * Construct a database. Specific database/driver is determined by environment variables,
 * retrieved with GetConfig().
 */
func GetDatabase() (*gorm.DB, error) {
	dbConfig, err := GetConfig()
	if err != nil {
		return nil, fmt.Errorf("error retrieving database configuration: %v", err)
	}

	switch dbConfig.DBType {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dbConfig.DBUser,
			dbConfig.DBPassword,
			dbConfig.DBHost,
			dbConfig.DBPort,
			dbConfig.DBName,
		)
		return gorm.Open(mysql.Open(dsn), &gorm.Config{})
	default:
		//This would also be case "sqlite":
		fullDBName := dbConfig.DBName + ".db"
		return gorm.Open(sqlite.Open("db-data/"+fullDBName), &gorm.Config{})
	}

}
