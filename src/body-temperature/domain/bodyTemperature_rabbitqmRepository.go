package domain

type IBodyTemperatureRabbitMQ interface {
	Save(BodyTemperature *BodyTemperature) error
}