package models

var (
	// Config - Global Configuration
	Config AppConfig
)

// AppConfig - AppConfig
type AppConfig struct {
	AppPort          string
	LogDir           string
	DataDir          string
	MongoURL         string
	MongoDBName      string
	JWTSecret        string
	MinJWTLength     int
	TimeOutInSeconds int
}
