package itest

import (
	"github.com/btcsuite/btclog/v2"
)

var (
	backend = btclog.NewBackend(newPrefixStdout("itest"))
	log     = backend.Logger("")
)
