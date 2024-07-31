package main

import (
	"MonsterMash/ff6library"
	"context"
)

// App struct
type App struct {
	ctx     context.Context
	library *ff6library.Library
}

// NewApp creates a new App application struct
func NewApp(l *ff6library.Library) *App {
	return &App{library: l}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) shutdown(ctx context.Context) {
	a.library.Stop()
}
