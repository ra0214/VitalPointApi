package domain

type IStressRabbitMQ interface {
	Save(Stress *Stress) error
}