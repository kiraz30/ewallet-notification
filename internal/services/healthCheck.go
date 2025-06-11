package services

import "ewallet-notification/internal/interfaces"

type HealthCheck struct {
	HealthCheckRepository interfaces.IHealthCheckRepository
}

func (s *HealthCheck) HealthCheckServices() (string, error) {
	return "service healty", nil
}
