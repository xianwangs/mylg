package icmp_test

import (
	"testing"

	"github.com/xianwangs/mylg/cli"
	"github.com/xianwangs/mylg/icmp"
)

func TestNewTrace(t *testing.T) {
	cfg, _ := cli.ReadDefaultConfig()
	_, err := icmp.NewTrace("google.com -n -nr -m 30", cfg)
	if err != nil {
		t.Error("unexpected error. expected %v, actual %v", nil, err)
	}
}
