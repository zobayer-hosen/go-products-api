package config

improt(
	"fmt"
	"os"
	"strconv"
)

// config holds all application configuration, loaded from environment variables
type config struct {
	Server ServerConfig
	Database DatabaseConfig
	JWT JWTConfig
}
type ServerConfig struct{
	Port string
	ReadTimeoutSec int
	WriteTimeoutSec int
}

type DatabaseConfig struct{
	DSN string // postgreSQL Data Source Name
}

type JWTConfig struct{
	Secret string
	ExpiryHours int
}

func Load() (*Config,error){
	secret := os.Getenv("JWT_SECRET")
	if secret == ""{
		return nil, fmt.Errorf("config: JWT_SECRET environment variable is required")
	}

	dbDSN := os.Getenv("DATABASE_URL")
	if dbDSN == ""{
		return nil, fmt.Errorf("config: DATABASE_URL environment variable is required")
	}

	readTimeout, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT_SEC"))
	if readTimeout == 0{
		readTimeout = 10
	}
	writeTimeout, _ := strconv.Atoi(os.Getenv("SERVER_WRITE_TIMEOUT_SEC"))
	if writeTimeout == 0{
		writeTimeout =30
	}
	jwtExpiry,_ := strconv.Atoi(os.Getenv("JWT_EXPIRY_HOURS"))
	if jwtExpiry == 0 {
		jwtExpiry = 24
	}
	port := os.Getenv("PORT")
	if port == ""{
		port = "8080"
	}

	return &Config{
		Server:ServerConfig{
			Port: port,
			ReadTimeoutSec: readTimeout,
			WriteTimeoutSec: writeTimeout,
		},
		Database: DatabaseConfig{DSN: dbDSN},
		JWT: JWTConfig{Secret: secret,ExpiryHours: jwtExpiry}
	},nil
}