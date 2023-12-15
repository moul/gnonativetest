package main

import (
	"context"
	"fmt"

	framework "github.com/gnolang/gnonative/framework/service"
)

// App struct
type App struct {
	ctx    context.Context
	bridge *framework.Bridge
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	config := framework.NewBridgeConfig()
	config.UseTcpListener = true
	config.RootDir = "."
	config.DisableUdsListener = true
	bridge, err := framework.NewBridge(config)
	if err != nil {
		panic(err)
	}
	a.bridge = bridge
	fmt.Println("GnoNative gRPC server created")
}

// Greet returns a greeting for the given name
//func (a *App) Greet(name string) string {
//	return fmt.Sprintf("Hello %s, It's show time!", name)
//}
