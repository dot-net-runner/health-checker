package health

type (
	HealthChecker interface {
		Check() (*ServiceHealth, error)
	}

	ServiceHealth struct {
		IsHealthy bool
		Service   string
		ErrorCode string
	}
)

func NewServiceHealth(isHealthy bool, service string, errorCode string) *ServiceHealth {
	return &ServiceHealth{
		IsHealthy: isHealthy,
		Service:   service,
		ErrorCode: errorCode,
	}
}
