package apollo

import (
	"testing"

	"github.com/spf13/viper"
)

func TestEnabled(t *testing.T) {
	p := &ApolloPlugin{}
	viper.Set(configPrefix+"enabled", true)
	if !p.Enabled() {
		t.Fail()
	}

	viper.Set(configPrefix+"enabled", false)
	if p.Enabled() {
		t.Fail()
	}
}

func TestOrder(t *testing.T) {
	p := &ApolloPlugin{}
	order := 1
	viper.Set(configPrefix+"order", order)
	if order != p.Order() {
		t.Fail()
	}
}

func TestLoadFailed(t *testing.T) {
	p := &ApolloPlugin{}

	viper.Set(configPrefix+"appId", "")
	viper.Set(configPrefix+"address", "")
	if p.Load() == nil {
		t.Fail()
	}

	viper.Set(configPrefix+"appId", "")
	viper.Set(configPrefix+"address", "foo")
	if p.Load() == nil {
		t.Fail()
	}

	// todo mock

}
