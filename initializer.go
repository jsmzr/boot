package boot

import "fmt"

type Initializer interface {
	Init() error
}

var initializers = make(map[string]Initializer)

func RegisterInitializer(name string, initializer Initializer) {
	Log(fmt.Sprintf("Register [%s] checker", name))
	initializers[name] = initializer
}

func runInitializer() error {
	for name, initializer := range initializers {
		Log(fmt.Sprintf("Run [%s] initializer", name))
		if err := initializer.Init(); err != nil {
			return err
		}
	}
	return nil
}
