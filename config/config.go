package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configurations *Config

type DBConfig struct {
	DB_Host       string
	DB_Port       int
	DB_Name       string
	DB_User       string
	DB_Password   string
	EnableSSLMode bool
}

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
	JWTSecret   string
	DB          DBConfig
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load the env varibles", err)
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("VERSION IS Required")
		os.Exit(-1)
	}

	dbhost := os.Getenv("DB_HOST")
	if dbhost == "" {
		fmt.Println("Host IS Required")
		os.Exit(-1)
	}

	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		fmt.Println("Name is Required")
		os.Exit(-1)
	}

	dbuser := os.Getenv("DB_USER")
	if dbuser == "" {
		fmt.Println("User is Required")
		os.Exit(-1)
	}
	dbpassword := os.Getenv("DB_PASSWORD")
	if dbpassword == "" {
		fmt.Println("Password IS Required")
		os.Exit(-1)
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		fmt.Println("DB Port is required")
		os.Exit(1)
	}

	dbport, err := strconv.Atoi(dbPort)
	if err != nil {
		fmt.Println("DB Port must be a number")
		os.Exit(1)
	}

	enableSSLMode := os.Getenv("DB_ENABLE_SSL_MODE")
	enbleSSLMode, err := strconv.ParseBool(enableSSLMode)
	if err != nil {
		fmt.Println("Invalid Enable ssl Mode", err)
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("Service Name is Reequired")
		os.Exit(-1)
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		fmt.Println("HTTP port is required")
		os.Exit(-1)
	}

	port, err := strconv.Atoi(httpPort)
	if err != nil {
		fmt.Println("http Port Must Be Number")
		os.Exit(-1)
	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		fmt.Println("Jwt Secret key isd required")
		os.Exit(1)
	}

	dbConfig := &DBConfig{
		DB_Host:       dbhost,
		DB_Port:       dbport,
		DB_Name:       dbname,
		DB_User:       dbuser,
		DB_Password:   dbpassword,
		EnableSSLMode: enbleSSLMode,
	}
	configurations = &Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    int(port),
		JWTSecret:   jwtSecretKey,
		DB:          *dbConfig,
	}

}

func GetConfig() Config {
	if configurations == nil {
		loadConfig()
	}
	return *configurations
}
