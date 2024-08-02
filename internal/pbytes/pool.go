//go:build !disablepool
// +build !disablepool

package pbytes

import "github.com/gobwas/pool/pbytes"

var GetLen = pbytes.GetLen

var Put = pbytes.Put
