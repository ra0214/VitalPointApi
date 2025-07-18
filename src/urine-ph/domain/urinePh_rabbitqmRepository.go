package domain

type IUrinePhRabbitMQ interface {
	Save(UrinePh *UrinePh) error
}