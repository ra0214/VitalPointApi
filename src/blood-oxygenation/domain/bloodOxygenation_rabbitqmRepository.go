package domain

type IBloodOxygenationRabbitMQ interface {
	Save(BloodOxygenation *BloodOxygenation) error
}