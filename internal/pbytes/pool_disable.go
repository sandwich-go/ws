//go:build disablepool
// +build disablepool

package pbytes

var GetLen = func(len int) []byte { return make([]byte, len) }

var Put = func(_ []byte) {}
