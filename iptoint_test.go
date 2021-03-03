package iptoint_test

import (
	"github.com/slash3b/iptoint"
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
			result := iptoint.Ipv4ToInt(tt.input)
			if result != tt.output {
				t.Errorf("actual %d is not equal expected %d \n", result, tt.output)
			}
		})
	}
}

func BenchmarkIpv4ToInt(b *testing.B) {

}
