package nmi

import "time"

// Role related to a Nmi.
type Role struct {
	FromDateTime time.Time `json:"fromDateTime,omitempty"`
	ToDateTime   time.Time `json:"toDateTime,omitempty"`
	Party        string    `json:"party,omitempty"`
	Type         string    `json:"type,omitempty"`
}
