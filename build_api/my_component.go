package main

import (
	cfacade "github.com/cherry-game/cherry/facade"
	clog "github.com/cherry-game/cherry/logger"
)

type MyComponent struct {
	cfacade.Component
}

func (p *MyComponent) Name() string {
	return "my_component"
}

func (p *MyComponent) Init() {
	clog.Info("MyComponent Init")
}

func (p *MyComponent) OnAfterInit() {
	clog.Info("MyComponent OnAfterInit")
}

func (p *MyComponent) OnBeforeStop() {
	clog.Info("MyComponent OnBeforeStop")
}

func (p *MyComponent) OnStop() {
	clog.Info("MyComponent OnStop")
}
