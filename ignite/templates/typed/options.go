package typed

import (
	"github.com/ignite-hq/cli/ignite/pkg/multiformatname"
	"github.com/ignite-hq/cli/ignite/templates/field"
)

// Options ...
type Options struct {
	AppName      string
	AppPath      string
	ModuleName   string
	ModulePath   string
	OwnerName    string
	TypeName     multiformatname.Name
	MsgSigner    multiformatname.Name
	Fields       field.Fields
	Indexes      field.Fields
	NoMessage    bool
	NoSimulation bool
	IsIBC        bool
}

// Validate that options are usable
func (opts *Options) Validate() error {
	return nil
}
