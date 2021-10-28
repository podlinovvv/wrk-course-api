package sender

import (
	"github.com/podlinovvv/wrk-course-api/internal/model"
)

type EventSender interface {
	Send(subdomain *model.SubdomainEvent) error
}
