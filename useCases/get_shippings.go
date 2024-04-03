package usecases

type ShippingDTO struct {
	Name          string  `json:"name"`
	Cost          float64 `json:"cost"`
	EstimatedDays int64   `json:"estimated_days"`
}

func GetShippings() []ShippingDTO {
	return []ShippingDTO{
		{
			Name:          "test",
			Cost:          5.63,
			EstimatedDays: 2,
		},
	}
}
