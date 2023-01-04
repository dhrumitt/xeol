package rpmmodularity

import (
	"fmt"

	"github.com/noqcks/xeol/xeol/pkg/qualifier"
	"github.com/noqcks/xeol/xeol/pkg/qualifier/rpmmodularity"
)

type Qualifier struct {
	Kind   string `json:"kind" mapstructure:"kind"`                         // Kind of qualifier
	Module string `json:"module,omitempty" mapstructure:"module,omitempty"` // Modularity label
}

func (q Qualifier) Parse() qualifier.Qualifier {
	return rpmmodularity.New(q.Module)
}

func (q Qualifier) String() string {
	return fmt.Sprintf("kind: %s, module: %q", q.Kind, q.Module)
}
