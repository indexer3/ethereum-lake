package service

import "context"

type ServiceType int

const (
	ServiceTypeUnknown ServiceType = iota
	ServiceTypeIndexer
	ServiceTypeRelayer
)

type Service interface {
	Start(ctx context.Context) error
}
