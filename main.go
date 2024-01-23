package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/diamondburned/gotk4-layer-shell/pkg/gtk4layershell"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func main() {
	app := gtk.NewApplication("com.github.heisfer.borkbar", 0)
	app.Connect("activate", activate)

	if code := app.Run(os.Args); code > 0 {
		os.Exit(1)
	}

}

func activate(app *gtk.Application) {
	if !gtk4layershell.IsSupported() {
		log.Fatalln("Your compositer doesn't support gtk-layer-shell")
	}

	appwin := gtk.NewApplicationWindow(app)
	clock(appwin)

	window := &appwin.Window

	gtk4layershell.InitForWindow(window)

	// set position
	gtk4layershell.SetLayer(window, gtk4layershell.Layer(gtk4layershell.LayerShellLayerBackground))
	gtk4layershell.SetAnchor(window, gtk4layershell.Edge(gtk4layershell.LayerShellEdgeTop), true)
	gtk4layershell.SetAnchor(window, gtk4layershell.Edge(gtk4layershell.LayerShellEdgeLeft), true)
	gtk4layershell.SetAnchor(window, gtk4layershell.Edge(gtk4layershell.LayerShellEdgeRight), true)

	gtk4layershell.SetMargin(window, gtk4layershell.Edge(gtk4layershell.LayerShellEdgeTop), 0)
	gtk4layershell.SetMargin(window, gtk4layershell.Edge(gtk4layershell.LayerShellEdgeLeft), 0)
	gtk4layershell.SetMargin(window, gtk4layershell.Edge(gtk4layershell.LayerShellEdgeRight), 0)
	gtk4layershell.AutoExclusiveZoneEnable(window)

	appwin.Show()
}

func clock(window *gtk.ApplicationWindow) {
	clockLabel := gtk.NewLabel("clock place holder")
	box := gtk.NewBox(gtk.OrientationHorizontal, 4)
	box.Append(clockLabel)

	window.SetChild(box)

	go func() {
		var ix int
		for t := range time.Tick(time.Second) {
			currentTime := t

			ix++

			glib.IdleAdd(func() {
				clockLabel.SetLabel(fmt.Sprintf("Last updated at %s.", currentTime.Format(time.StampMilli)))
			})
		}
	}()
}
