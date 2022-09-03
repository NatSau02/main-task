package service

import "github.com/NatSau02/main-task"

type RateServicee interface {
	Save(entity.Сalculator) entity.Сalculator
	FindAll() []entity.Сalculator
}

type rateServicee struct {
	rates []entity.Сalculator
}

func New() RateServicee {
	return &rateServicee{
		rates: []entity.Сalculator{},
	}
}

func (service *rateServicee) Save(rate entity.Сalculator) entity.Сalculator {
	service.rates = append(service.rates, rate)
	return rate 
}

func (service *rateServicee) FindAll() []entity.Сalculator {
	return service.rates
}

