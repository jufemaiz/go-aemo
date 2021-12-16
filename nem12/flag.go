package nem12

import (
	"fmt"
)

// Flag is an interface for flags.
type Flag interface {
	fmt.GoStringer
	Identifier() string
	Description() (string, error)
}
