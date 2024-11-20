package nmi

import (
	"fmt"
	"regexp"
	"sort"

	"github.com/jufemaiz/go-aemo/region"
)

const (
	// NMICHECKSUMINVALID is returned when a valid checksum cannot be determined.
	NMICHECKSUMINVALID = -1
	// NmiLength is the standard length of a Nmi.
	NmiLength = 10
	// NmiValidPattern the pattern for valid Nmis.
	NmiValidPattern = `^([A-HJ-NP-Z\d]{10})`
	// nilstr nil in string form.
	nilstr = "nil"
)

var (
	// NmiValidRegexp is the regular expression that Nmi strings may contain.
	NmiValidRegexp = regexp.MustCompile(NmiValidPattern)
)

// Nmi is a national meter identifier as per AEMO.
// Ref:
// - https://aemo.com.au/-/media/files/electricity/nem/retail_and_metering/metering-procedures/2016/0610-0008-pdf.pdf
type Nmi struct {
	Identifier                 string `json:"identifier,omitempty"`
	MSATSDetail                string `json:"msatsDetail,omitempty"`
	TransmissionNodeIdentifier TNI    `json:"tni,omitempty"`
	DistributionLossFactorCode DLFC   `json:"dlfc,omitempty"`
	CustomerClassificationCode string `json:"customerClassificationCode,omitempty"`
	CustomerThresholdCode      string `json:"customerThresholdCode,omitempty"`
	JurisdictionCode           string `json:"jurisdictionCode,omitempty"`
	ClassificationCode         string `json:"classificationCode,omitempty"`
	Meters                     Meters `json:"meters,omitempty"`
	DataStreams                string `json:"datastreams,omitempty"`
	// Address                    Address
	// Status                     Status
	// Roles       []*Role
}

// Checksum calculates a Nmi's checksum.
func Checksum(s string) int {
	n, err := NewNmi(s)
	if err != nil {
		return NMICHECKSUMINVALID
	}

	return n.Checksum()
}

// NewNmi returns a string as a Nmi along with an error if not valid.
func NewNmi(s string) (*Nmi, error) {
	n := &Nmi{Identifier: s}
	if err := n.Validate(); err != nil {
		return nil, err
	}

	return n, nil
}

// GoString meets the gostring interface.
func (n *Nmi) GoString() string {
	return fmt.Sprintf("Nmi{Identifier: %q}", n.Identifier)
}

// String meets the stringer interface.
func (n *Nmi) String() string {
	return n.Identifier
}

// Checksum returns the checksum of the provided Nmi.
func (n *Nmi) Checksum() int {
	if !n.Valid() {
		return NMICHECKSUMINVALID
	}

	var c, s int64

	sumDig := func(n int64) int64 {
		var sum int64

		for n != 0 {
			sum += n % 10
			n /= 10
		}

		return sum
	}

	chrs := []rune(n.Identifier)

	for i := len(chrs) - 1; i >= 0; i-- {
		v := int64(chrs[i])
		if i%2 != 0 {
			v *= 2
		}

		s += sumDig(v)
	}

	c = (10 - (s % 10)) % 10

	return int(c)
}

// ChecksumValid returns true if the provided checksum is valid.
func (n *Nmi) ChecksumValid(i int) bool {
	c := n.Checksum()
	if c == NMICHECKSUMINVALID {
		return false
	}

	return i == c
}

// Valid returns true if the Nmi is valid.
func (n *Nmi) Valid() bool {
	if err := n.Validate(); err != nil {
		return false
	}

	return true
}

// Validate checks if the Nmi is valid and returns the errors if not.
func (n *Nmi) Validate() error {
	if len(n.Identifier) != NmiLength {
		return fmt.Errorf("'%s': %w", n, ErrNmiInvalidLength)
	}

	if !NmiValidRegexp.MatchString(n.Identifier) {
		return fmt.Errorf("'%s': %w", n, ErrNmiInvalidChar)
	}

	return nil
}

// AllMeters returns all the meters for a Nmi.
func (n *Nmi) AllMeters() ([]*Meter, error) {
	if n == nil {
		return nil, ErrNmiNil
	}

	meters := []*Meter{}

	for _, m := range n.Meters {
		if m == nil {
			continue
		}

		meters = append(meters, m)
	}

	sort.Slice(meters, func(i, j int) bool {
		return meters[i].Identifier < meters[j].Identifier
	})

	return meters, nil
}

// AddMeter adds a new meter, returning an error if already added.
func (n *Nmi) AddMeter(m *Meter) error {
	if n == nil {
		return ErrNmiNil
	}

	if m == nil {
		return ErrMeterNil
	}

	if m.Identifier == "" {
		return ErrNmiMeterIdentifierEmpty
	}

	if _, ok := n.Meters[m.Identifier]; ok {
		return fmt.Errorf("adding meter '%#v': %w", m, ErrNmiMeterFound)
	}

	n.Meters[m.Identifier] = m

	return nil
}

// RemoveMeter removes a meter, returning an error if not in the list.
func (n *Nmi) RemoveMeter(m *Meter) error {
	if n == nil {
		return ErrNmiNil
	}

	if m == nil {
		return fmt.Errorf("RemoveMeter: %w", ErrMeterNil)
	}

	if m.Identifier == "" {
		return fmt.Errorf("removing meter '%s': %w", m.Identifier, ErrNmiMeterIdentifierEmpty)
	}

	if _, ok := n.Meters[m.Identifier]; !ok {
		return fmt.Errorf("removing meter '%s': %w", m.Identifier, ErrNmiMeterNotFound)
	}

	delete(n.Meters, m.Identifier)

	return nil
}

// Participant returns the pariticipant for the Nmi.
func (n *Nmi) Participant() (Participant, error) {
	if n == nil {
		return ParticipantUndefined, ErrNmiNil
	}

	for _, p := range Participants() {
		for _, pat := range p.Allocations() {
			if pat.Match(n.String()) {
				return p, nil
			}
		}
	}

	return ParticipantUndefined, ErrParticipantInvalid
}

// Region returns the region based on the allocation.
func (n *Nmi) Region() (region.Region, error) {
	if n == nil {
		return region.RegionUndefined, ErrNmiNil
	}

	p, err := n.Participant()
	if err != nil {
		return region.RegionUndefined, fmt.Errorf("'%s': %w", n, err)
	}

	return p.Region(), nil
}
