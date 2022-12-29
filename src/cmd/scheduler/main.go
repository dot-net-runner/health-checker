package main

import (
	"health-checker/internal/app"
	"health-checker/internal/services/health"
	"health-checker/internal/services/health/http"
	"health-checker/internal/services/notification"
	"health-checker/internal/services/notification/smtp"
	"log"

	"github.com/spf13/viper"
)

func main() {
	notifier, healthChecker := CreateServices()
	messageQueue := make(chan *notification.Message, 4)
	app.ScheduleHealthCheck(healthChecker, messageQueue)
	app.StartNotification(notifier, messageQueue)
}

func CreateServices() (notification.Nitifier, health.HealthChecker) {
	viper.SetConfigFile("appsettings.json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	smtpConf := smtp.Configuration{
		Host:       viper.GetString("notifications.smtp.Host"),
		Port:       viper.GetString("notifications.smtp.Port"),
		Login:      viper.GetString("notifications.smtp.Login"),
		Password:   viper.GetString("notifications.smtp.Password"),
		Sender:     viper.GetString("notifications.smtp.Sender"),
		Recipients: viper.GetStringSlice("notifications.smtp.Recipients"),
	}

	hcConf := http.Configuration{
		URL: viper.GetString("healthCheck.http.URL"),
	}

	notifier := smtp.NewService(smtpConf)
	healthChecker := http.NewService(hcConf)

	return notifier, healthChecker
}
