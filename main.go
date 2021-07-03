package main

import (
	"log"

	v "github.com/caleb-noodahl/macro-maker/views"
	"github.com/jroimartin/gocui"
)

var (
	views map[string]v.View
)

func main() {
	views = generateViews()
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Fatalln(err)
	}
	defer g.Close()
	//since we're loading from main we'll load the home view
	views["home"].Render(g)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Fatalln(err)
	}
}

func generateViews() map[string]v.View {
	home := v.HomeView{}
	list := map[string]v.View{
		"macro": &v.MacroView{},
	}
	home.ViewList = list
	list["home"] = &home
	return list
}
