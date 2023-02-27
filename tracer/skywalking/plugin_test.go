package skywalking

import (
	"testing"

	"github.com/spf13/viper"
)

func TestEnabled(t *testing.T) {
	p := SkywalkingPlugin{}
	viper.Set(configPrefix+"enabled", false)
	if p.Enabled() {
		t.Fail()
	}

	viper.Set(configPrefix+"enabled", true)
	if !p.Enabled() {
		t.Fail()
	}
}

func TestOrder(t *testing.T) {
	p := SkywalkingPlugin{}
	order := 1
	viper.Set(configPrefix+"order", order)
	if p.Order() != order {
		t.Fail()
	}
}

func TestLoad(t *testing.T) {
	p := SkywalkingPlugin{}
	viper.Set(configPrefix+"address", "")
	if p.Load() == nil {
		t.Fail()
	}

	viper.Set(configPrefix+"address", "failed")
	_ = p.Load()

}
