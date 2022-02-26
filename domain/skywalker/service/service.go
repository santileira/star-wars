package service

import "rings/domain/skywalker/domain"

type Service interface {
	GetMessage() (*domain.Response, error)
}
