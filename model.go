package goovi

import (
	"time"

	"github.com/gooviid/gopkg/goovimodel"
	"github.com/gooviid/gopkg/goovireq"
)

type WebhookData struct {
	IsSuccess bool                       `json:"is_success"`
	QueueID   string                     `json:"queue_id"`
	Error     *goovimodel.TranscodeError `json:"error"`
	Media     []goovimodel.MediaInfo     `json:"media"`
}

type TranscodeQueue struct {
	QueueId          string                     `json:"queue_id"`
	QueueStatus      string                     `json:"queue_status"`
	ChargeStatus     string                     `json:"charge_status"`
	TotalPrice       float32                    `json:"total_price"`
	CreatedAt        time.Time                  `json:"created_at"`
	ResolutionTarget []string                   `json:"resolution_target"`
	FileTarget       []string                   `json:"file_target"`
	RequestDuration  int32                      `json:"request_duration"`
	Source           goovireq.SourceTranscode   `json:"source"`
	Destination      goovireq.SourceTranscode   `json:"destination"`
	BillingDetail    []goovimodel.BillingDetail `json:"billing_detail"`
	Media            goovimodel.MediaInfo       `json:"media"`
	Error            *goovimodel.TranscodeError `json:"error"`
}
