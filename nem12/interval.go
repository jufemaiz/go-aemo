package nem12

import (
	"time"
)

// Interval is duration of time from a start to a finish, with a value.
type Interval struct {
	Time           time.Time
	IntervalLength time.Duration
}

// An IntervalValue represents a single meter interval value as presented by an NEM12 file.
type IntervalValue struct {
	Value             float64 // Value of the interval in the unit of measure specified in the NMIDataDetails record.
	QualityFlag       Quality
	MethodFlag        Method
	ReasonCode        Reason    // ReasonCode applicable to this value.
	ReasonDescription string    // Text describing the reason for when ReasonCode equals FreeTextDescription.
	UpdateDateTime    time.Time // Time that the metering data was created or updated as reported by the MDP
	MSATSLoadDateTime time.Time // Time that the metering data was loaded into MSATS
}
