package shark

import (
	"integrationtests/prey"
)

type Shark interface {
	Hunt(prey prey.Prey) error
}
