package config

import (
	"github.com/spf13/viper"
)

// keys for database configuration
const (
	DbName = "db.name"
	DbHost = "db.ip"
	DbPort = "db.port"
	DbUser = "db.user"
	DbPass = "db.pass"

	ServerHost = "server.host"
	ServerPort = "server.port"

	EmailSMTPUsername = "email.smtp.username"
	EmailSMTPPassword = "email.smtp.password"
	EmailSMTPHost     = "email.smtp.host"
	EmailSMTPPort     = "email.smtp.port"
	EmailFromAddress  = "email.from.address"
)

func init() {
	// env var for db
	_ = viper.BindEnv(DbName, "DB_NAME")
	_ = viper.BindEnv(DbHost, "DB_HOST")
	_ = viper.BindEnv(DbPort, "DB_PORT")

	// email and SMTP settings
	_ = viper.BindEnv(EmailSMTPUsername, "EMAIL_SMTP_USERNAME")
	_ = viper.BindEnv(EmailSMTPPassword, "EMAIL_SMTP_PASSWORD")
	_ = viper.BindEnv(EmailSMTPHost, "EMAIL_SMTP_HOST")
	_ = viper.BindEnv(EmailSMTPPort, "EMAIL_SMTP_PORT")
	_ = viper.BindEnv(EmailFromAddress, "EMAIL_FROM_ADDRESS")

	// defaults
	viper.SetDefault(DbName, "travel-booking-portal")
	viper.SetDefault(DbHost, "localhost")
	viper.SetDefault(DbPort, "27017")
	// Set default values for email and SMTP settings here
	viper.SetDefault(EmailSMTPUsername, "sayednur171@gmail.com")
	viper.SetDefault(EmailSMTPPassword, "xjh jjlk adfh akd ")
	viper.SetDefault(EmailSMTPHost, "smtp.gmail.com")
	viper.SetDefault(EmailSMTPPort, "587")
	viper.SetDefault(EmailFromAddress, "sayednur171@gmail.com")
}
