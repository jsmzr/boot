package boot

import (
	"fmt"
	"sort"
)

type Plugin interface {
	Load() error
	Order() int
	Enabled() bool
}

var plugins = make(map[string]Plugin)

func RegisterPlugin(name string, plugin Plugin) {
	_, ok := plugins[name]
	if ok {
		panic(fmt.Errorf("plugin [%s] already register", name))
	}
	Log(fmt.Sprintf("Register [%s] plugin", name))
	plugins[name] = plugin
}

func PostProccess() error {
	if err := initConfig(); err != nil {
		return err
	}
	pluginList := make([]Plugin, 0, len(plugins))
	for _, item := range plugins {
		pluginList = append(pluginList, item)
	}
	sort.Slice(pluginList, func(i, j int) bool {
		return pluginList[i].Order() < pluginList[j].Order()
	})
	for _, item := range pluginList {
		if !item.Enabled() {
			Log(fmt.Sprintf("Plugin [%T] enabled config is false", item))
			continue
		}
		Log(fmt.Sprintf("Load [%T] plugin", item))
		if err := item.Load(); err != nil {
			return err
		}
	}
	return nil
}
