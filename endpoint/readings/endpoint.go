package readings

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"joi-energy-golang/domain"
)

func makeStoreReadingsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(domain.StoreReadings)
		s.StoreReadings(req.SmartMeterId, req.ElectricityReadings)
		return req, nil
	}
}

func makeGetReadingsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(string)
		err := validateSmartMeterId(req)
		if err != nil {
			return nil, err
		}
		res := s.GetReadings(req)
		return domain.StoreReadings{
			SmartMeterId:        req,
			ElectricityReadings: res,
		}, nil
	}
}

func GetusageCostEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(string)
		err := validateSmartMeterId(req)
		if err != nil {
			return nil, err
		}
		res := s.GetusageCost(req)
		var priceplans []domain.StoreReadings

		data := PricePlandData{
			PricePlan: 0.4,
			Readings:  []float64{0.05, 0.07, 0.8},
		}

		// averageReading := calulateAverage(data.Readings)

		for _, value := range priceplans {
			data.Readings = float64(priceplans.ElectricityReadings)
		}
		averageReading := calulateAverage(data.Readings)
		usageTimehours := 24 * 7
		energyConsumed := averageReading * float64(usageTimehours)
		data.PricePlan * energyConsumed

		// response := map[string]float64{
		// 	"cost": cost,
		// }
		return domain.StoreReadings{
			SmartMeterId:        req,
			ElectricityReadings: res,
		}, nil
	}
}

func calulateAverage(readings []float64) float64 {
	var total float64
	for _, value := range readings {
		total += value
	}
	return total / float64(len(readings))
}
