package main

import (
	"MonsterMash/ff6library"
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed Monster.yml
var monsterData []byte

//go:embed Item.yml
var itemData []byte

//go:embed MetamorphSets.yml
var morphData []byte

//go:embed skillNames.yml
var skillData []byte

func main() {
	// Create an instance of the app structure
	library := ff6library.NewLibrary(itemData, monsterData, morphData, skillData)
	library.Start()
	app := NewApp(library)

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "MonsterMash",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		OnShutdown:       app.shutdown,
		Bind: []interface{}{
			app,
			library,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
