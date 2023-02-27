package logrus

import (
	"reflect"
	"testing"

	"github.com/jsmzr/boot/tracer"
	"github.com/sirupsen/logrus"
)

func TestLevels(t *testing.T) {
	hook := SkywalkingLogrusHook{}
	if !reflect.DeepEqual(hook.Levels(), logrus.AllLevels) {
		t.Fail()
	}
}

func TestFire(t *testing.T) {
	hook := SkywalkingLogrusHook{}
	tracer.ThreadLocal.Remove()
	e := logrus.Entry{}

	// req is nil
	if hook.Fire(&e) != nil {
		t.Fail()
	}
	if e.Data["tid"] != nil {
		t.Fail()
	}

	// convert to http.Request failed
	var temp map[string]interface{}
	tracer.ThreadLocal.Set(&temp)
	if hook.Fire(&e) != nil {
		t.Fail()
	}
	if e.Data["tid"] != nil {
		t.Fail()
	}

	tracer.ThreadLocal.Remove()
	// todo mock
}
