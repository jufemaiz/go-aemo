package nmi

import "time"

// Role related to a Nmi.
type Role struct {
	FromDateTime time.Time
	ToDateTime   time.Time
	Party        string
	Type         string
}
