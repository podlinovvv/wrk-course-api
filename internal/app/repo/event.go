package repo

import (
	"github.com/podlinovvv/wrk-course-api/internal/model"
)

type EventRepo interface {
	Lock(n uint64) ([]model.SubdomainEvent, error)
	Unlock(eventIDs []uint64) error

	Add(event []model.SubdomainEvent) error
	Remove(eventIDs []uint64) error
}
