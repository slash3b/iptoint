package iptoint_test

import (
	"github.com/slash3b/iptoint"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

func TestIpv4ToInt(t *testing.T) {
	testTable := []struct {
		testName string
		input    string
		output   int
	}{
		{
			testName: "Simple IPv4 conversion",
			input:    "89.28.125.178",
			output:   1495039410,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.testName, func(t *testing.T) {
			assert.Equal(t, iptoint.IpToInt(net.ParseIP(tt.input)), tt.output)
		})
	}
}

func BenchmarkIpv4ToInt(b *testing.B) {

}
