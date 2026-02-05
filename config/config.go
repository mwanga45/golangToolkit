package config

import "os"

// Config holds application configuration from environment variables.
type Config struct {
	App  AppConfig
	DB   DBConfig
	JWT  JWTConfig
}

// AppConfig holds app-level settings.
type AppConfig struct {
	Port string
	Env  string
}

// DBConfig holds database connection settings.
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

// JWTConfig holds JWT/auth settings.
type JWTConfig struct {
	Secret string
}

// Load reads configuration from environment variables.
// Call godotenv.Load() in main before using this.
func Load() *Config {
	return &Config{
		App: AppConfig{
			Port: getEnv("APP_PORT", "3000"),
			Env:  getEnv("APP_ENV", "development"),
		},
		DB: DBConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", ""),
			Name:     getEnv("DB_NAME", "project_db"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		JWT: JWTConfig{
			Secret: getEnv("JWT_SECRET", "change-me"),
		},
	}
}

// DSN returns the PostgreSQL connection string for GORM.
func (c *DBConfig) DSN() string {
	return "host=" + c.Host +
		" port=" + c.Port +
		" user=" + c.User +
		" password=" + c.Password +
		" dbname=" + c.Name +
		" sslmode=" + c.SSLMode
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
