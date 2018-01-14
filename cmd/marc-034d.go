package main

import (
	"flag"
	"fmt"
	"github.com/thisisaaronland/go-marc/http"
	"github.com/thisisaaronland/go-slippy-tiles"
	"github.com/thisisaaronland/go-slippy-tiles/provider"
	"github.com/whosonfirst/go-http-mapzenjs"
	"log"
	gohttp "net/http"
	"os"
)

func main() {

	var host = flag.String("host", "localhost", "The hostname to listen for requests on")
	var port = flag.Int("port", 8080, "The port number to listen for requests on")

	var api_key = flag.String("mapzen-api-key", "mapzen-xxxxxx", "A valid Mapzen API key")

	flag.Parse()

	www_handler, err := http.WWWHandler()

	if err != nil {
		log.Fatal(err)
	}

	opts := mapzenjs.DefaultMapzenJSOptions()
	opts.APIKey = *api_key

	www_mapzenjs_handler, err := mapzenjs.MapzenJSHandler(www_handler, opts)

	if err != nil {
		log.Fatal(err)
	}

	static_handler, err := http.StaticHandler()

	if err != nil {
		log.Fatal(err)
	}

	mapzenjs_assets_handler, err := mapzenjs.MapzenJSAssetsHandler()

	if err != nil {
		log.Fatal(err)
	}

	bbox_handler, err := http.BboxHandler()

	if err != nil {
		log.Fatal(err)
	}

	ping_handler, err := http.PingHandler()

	if err != nil {
		log.Fatal(err)
	}

	mux := gohttp.NewServeMux()

	mux.Handle("/", www_mapzenjs_handler)
	mux.Handle("/javascript/", static_handler)
	mux.Handle("/css/", static_handler)

	mux.Handle("/javascript/mapzen.min.js", mapzenjs_assets_handler)
	mux.Handle("/javascript/tangram.min.js", mapzenjs_assets_handler)
	mux.Handle("/javascript/mapzen.js", mapzenjs_assets_handler)
	mux.Handle("/javascript/tangram.js", mapzenjs_assets_handler)
	mux.Handle("/css/mapzen.js.css", mapzenjs_assets_handler)
	mux.Handle("/tangram/refill-style.zip", mapzenjs_assets_handler)

	mux.Handle("/bbox", bbox_handler)
	mux.Handle("/ping", ping_handler)

	// WIP - tile caching...

	cache := slippytiles.CacheConfig{
		Name: "Disk",
		Path: "./", // PLEASE FIX ME
	}

	mapzen_url := fmt.Sprintf("https://tile.mapzen.com/mapzen/vector/v1/512/all/{z}/{x}/{y}.topojson?api_key=%s", opts.APIKey)

	mapzen_config := slippytiles.LayerConfig{
		URL:     mapzen_url,
		Formats: []string{"topojson"},
	}

	layers := slippytiles.LayersConfig{"mapzen": mapzen_config}

	config := slippytiles.Config{
		Cache:  cache,
		Layers: layers,
	}

	provider, err := provider.NewProxyProvider(&config)

	if err != nil {
		log.Fatal(err)
	}

	tiles_handler := provider.Handler(nil)
	mux.Handle("/tiles", tiles_handler)

	// EO WIP

	address := fmt.Sprintf("%s:%d", *host, *port)
	log.Printf("listening on %s\n", address)

	err = gohttp.ListenAndServe(address, mux)

	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
