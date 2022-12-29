package app

import (
	"health-checker/internal/services/notification"
	"health-checker/internal/services/health"
	"log"
	"time"

	"github.com/madflojo/tasks"
)

func ScheduleHealthCheck(hc health.HealthChecker, errors chan<- *notification.Message) error {
	scheduler := tasks.New()
	defer scheduler.Stop()
	_, err := scheduler.Add(&tasks.Task{
		Interval: time.Duration(10 * time.Second),
		TaskFunc: func() error {
			return healthCheck(hc, errors)
		},
		ErrFunc: func(e error) {
			log.Fatal(e.Error())
		},
	})

	return err
}

func healthCheck(hc health.HealthChecker, errors chan<- *notification.Message) error {
	sh, he := hc.Check()
	if he != nil {
		return he
	}
	if !sh.IsHealthy {
		msg := notification.NewMessage("Сервис недоступен.", "Сервис ответил с кодом: "+sh.ErrorCode)
		errors <- msg
	}
	return he
}
