package nem12

import (
	"fmt"
	"strings"
)

const (
	// TransactionUndefined is for undefined transaction flags.
	TransactionUndefined TransactionCode = iota
	// TransactionAlteration is the transaction code value for 'A', when alteration.
	TransactionAlteration
	// TransactionMeterReconfiguration is the transaction code value for 'C', when meter reconfiguration.
	TransactionMeterReconfiguration
	// TransactionReEnergisation is the transaction code value for 'G', when re-energisation.
	TransactionReEnergisation
	// TransactionDeEnergisation is the transaction code value for 'D', when de-energisation.
	TransactionDeEnergisation
	// TransactionEstimate is the transaction code value for 'E', when estimated read.
	TransactionEstimate
	// TransactionNormalRead is the transaction code value for 'N', when normal read.
	TransactionNormalRead
	// TransactionOther is the transaction code value for 'O', when other.
	TransactionOther
	// TransactionSpecialRead is the transaction code value for 'S', when special read.
	TransactionSpecialRead
	// TransactionRemovalOfMeter is the transaction code value for 'R', when removal of meter.
	TransactionRemovalOfMeter
)

var (
	// transactions lists all transactions.
	transactions = []TransactionCode{ //nolint:gochecknoglobals
		TransactionAlteration,
		TransactionMeterReconfiguration,
		TransactionReEnergisation,
		TransactionDeEnergisation,
		TransactionEstimate,
		TransactionNormalRead,
		TransactionOther,
		TransactionSpecialRead,
		TransactionRemovalOfMeter,
	}

	// TransactionName maps a transaction code to its name.
	TransactionName = map[TransactionCode]string{ //nolint:gochecknoglobals
		TransactionAlteration:           "A",
		TransactionMeterReconfiguration: "C",
		TransactionReEnergisation:       "G",
		TransactionDeEnergisation:       "D",
		TransactionEstimate:             "E",
		TransactionNormalRead:           "N",
		TransactionOther:                "O",
		TransactionSpecialRead:          "S",
		TransactionRemovalOfMeter:       "R",
	}

	// TransactionValue maps a name to its value.
	TransactionValue = map[string]TransactionCode{ //nolint:gochecknoglobals
		"A": TransactionAlteration,
		"C": TransactionMeterReconfiguration,
		"G": TransactionReEnergisation,
		"D": TransactionDeEnergisation,
		"E": TransactionEstimate,
		"N": TransactionNormalRead,
		"O": TransactionOther,
		"S": TransactionSpecialRead,
		"R": TransactionRemovalOfMeter,
	}

	// transactionDescriptions provides the descriptions for the transaction flags.
	transactionDescriptions = map[TransactionCode]string{ //nolint:gochecknoglobals
		TransactionAlteration:           "Alteration. Any action involving the alteration of the metering installation at a Site. This includes a removal of one meter and replacing it with another and all new connections and ‘Add/Alts’ Service Orders.",
		TransactionMeterReconfiguration: "Meter Reconfiguration Service Order. This includes off-peak (Controlled Load) timing changes. This does not apply to the removal of the meter.",
		TransactionReEnergisation:       "Re-energisation Service Order.",
		TransactionDeEnergisation:       "De-energisation Service Order.",
		TransactionEstimate:             "Estimate. For all estimates.",
		TransactionNormalRead:           "Normal Read. Scheduled collection of metering data. Also includes the associated Substitutions.",
		TransactionOther:                "Other. Include Meter Investigation & Miscellaneous Service Orders. This value is used when providing Historical Data and where the TransCode information is unavailable.",
		TransactionSpecialRead:          "Special Read service order.",
		TransactionRemovalOfMeter:       "Remove of Meter. This is used for meter removal or supply abolishment where the meter has been removed and will not be replaced. This excludes situations involving a meter changeover or where a meter is added to an existing configuration (these are considered to be alterations).",
	}
)

// TransactionCode represents the value of the transaction code flag.
type TransactionCode int

// Transactions returns a slice of all the transactions.
func Transactions() []TransactionCode {
	return transactions
}

// NewTransactionCode returns a new transaction flag if valid, and an error if not.
func NewTransactionCode(s string) (TransactionCode, error) {
	if s == "" {
		return TransactionUndefined, ErrTransactionCodeNil
	}

	q, ok := TransactionValue[strings.ToUpper(s)]
	if !ok {
		return q, ErrTransactionCodeInvalid
	}

	return q, nil
}

// Identifier to meet the interface specification for a Flag.
func (t TransactionCode) Identifier() string {
	id, ok := TransactionName[t]
	if !ok {
		return fmt.Sprintf("TransactionCode(%d)", t)
	}

	return id
}

// GoString returns a text representation of the Transaction to satisfy the GoStringer
// interface.
func (t TransactionCode) GoString() string {
	return fmt.Sprintf("TransactionCode(%d)", t)
}

// String returns a text representation of the Transaction.
func (t TransactionCode) String() string {
	s, err := t.Description()
	if err != nil {
		return fmt.Sprintf("\"%s\"", t.Identifier())
	}

	return fmt.Sprintf("\"%s: %s\"", t.Identifier(), s)
}

// Description returns the description of a transaction flag. Error is returned if the
// flag is invalid.
func (t TransactionCode) Description() (string, error) {
	d, ok := transactionDescriptions[t]

	if !ok {
		return "", ErrTransactionCodeInvalid
	}

	return d, nil
}
