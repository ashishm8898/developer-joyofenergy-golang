package readings

import (
	"github.com/sirupsen/logrus"

	"joi-energy-golang/domain"
	"joi-energy-golang/repository"
)

type Service interface {
	StoreReadings(smartMeterId string, reading []domain.ElectricityReading)
	GetReadings(smartMeterId string) []domain.ElectricityReading
	GetusageCost(smartMeterId string) []domain.ElectricityReading
}

type PricePlandData struct {
	PricePlan float64
	Readings  []float64
}

type service struct {
	logger        *logrus.Entry
	meterReadings *repository.MeterReadings
}

func NewService(
	logger *logrus.Entry,
	meterReadings *repository.MeterReadings,
	// GetusageCost *repository.GetusageCost,
) Service {
	return &service{
		logger:        logger,
		meterReadings: meterReadings,
		// GetusageCost:  GetusageCost
	}
}

func (s *service) StoreReadings(smartMeterId string, reading []domain.ElectricityReading) {
	s.meterReadings.StoreReadings(smartMeterId, reading)
}

func (s *service) GetReadings(smartMeterId string) []domain.ElectricityReading {
	return s.meterReadings.GetReadings(smartMeterId)
}

func (s *service) GetusageCost(smartMeterId string) []domain.ElectricityReading {
	return s.meterReadings.GetusageCost(smartMeterId)
}
