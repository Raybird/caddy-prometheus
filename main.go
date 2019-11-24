package main

import (
	"github.com/caddyserver/caddy/caddy/caddymain"
	// plug in plugins here, for example:
	// _ "import/path/here"
	_ "github.com/miekg/caddy-prometheus"
	_ "github.com/nicolasazrak/caddy-cache"
	_ "github.com/captncraig/caddy-realip"
	_ "github.com/captncraig/cors"
	_ "github.com/epicagency/caddy-expires"

)

func main() {
	// optional: disable telemetry
	// caddymain.EnableTelemetry = false
	caddymain.Run()
}
