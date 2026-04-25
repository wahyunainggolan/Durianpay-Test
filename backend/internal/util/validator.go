package util

var validPaymentStatuses = map[string]struct{}{
	"completed":  {},
	"processing": {},
	"failed":     {},
}

func IsValidPaymentStatus(status string) bool {
	_, ok := validPaymentStatuses[status]
	return ok
}
