package service

type ServiceType int

const (
	ServiceTypeUnknown ServiceType = iota
	ServiceTypeIndexer
	ServiceTypeRelayer
)

type IService interface {
}