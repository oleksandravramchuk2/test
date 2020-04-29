package main

import (
	"flag"
	"log"

	"github.com/test/metric"
	"github.com/test/pool"

	"github.com/valyala/fasthttp"
)

var (
	addr    = flag.String("addr", ":8080", "TCP address to listen to")
	changer = flag.Bool("changer", true, "Whether to enable every 200 ms one remove and one add request")
)

// App main struct that includes metrics and pool to easy accessing
type App struct {
	pool   *pool.ActiveRequestsPool
	metric *metric.CountersMap
}

func newApp(changer bool) *App {
	log.Println("Create 50 requests...")
	m := metric.NewCounters()
	p := pool.NewPool(m)
	if changer {
		log.Println("Start changer...")
		go p.Changer(m)
	}
	return &App{p, m}

}

func main() {
	flag.Parse()
	a := newApp(*changer)

	apiHandler := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/request":
			requestHandler(ctx, a)
		case "/admin/requests":
			adminHandler(ctx, a)
		default:
			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		}
	}
	h := apiHandler

	log.Printf("Start listening %s...", *addr)
	if err := fasthttp.ListenAndServe(*addr, h); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}

}

// requestHandler returns random value from pool
func requestHandler(ctx *fasthttp.RequestCtx, app *App) {
	data := app.pool.GetValue(app.metric)
	ctx.Write([]byte(data))
}

// adminHandler returns metrics
func adminHandler(ctx *fasthttp.RequestCtx, app *App) {
	ctx.Write([]byte(app.metric.Range()))
}
