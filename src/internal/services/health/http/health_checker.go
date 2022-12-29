package http

import (
	"fmt"
	"health-checker/internal/services/health"
	"net/http"
)

type Service struct {
	conf Configuration
}

type Configuration struct {
	URL string
}

func NewService(c Configuration) *Service {
	return &Service{
		conf: c,
	}
}

func (s *Service) Check() (*health.ServiceHealth, error) {
	resp, err := http.Get(s.conf.URL)
	if err != nil {
		return nil, err
	}

	sh := health.NewServiceHealth(resp.StatusCode == http.StatusOK,
		s.conf.URL,
		fmt.Sprintf("%d", resp.StatusCode))

	return sh, err
}
