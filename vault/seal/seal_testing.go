package seal

import (
	wrapping "github.com/akadake/go-kms-wrapping"
	"github.com/hashicorp/go-hclog"
)

type TestSealOpts struct {
	Logger     hclog.Logger
	StoredKeys StoredKeysSupport
	Secret     []byte
	Name       string
}

func NewTestSeal(opts *TestSealOpts) *Access {
	if opts == nil {
		opts = new(TestSealOpts)
	}

	return &Access{
		Wrapper:        wrapping.NewTestWrapper(opts.Secret),
		OverriddenType: opts.Name,
	}
}
