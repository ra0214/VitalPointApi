package domain

type ISugarOrineRabbitMQ interface {
	Save(SugarOrine *SugarOrine) error
}