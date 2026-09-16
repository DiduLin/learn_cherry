package main

import "github.com/cherry-game/cherry"

func main() {
	builder := cherry.Configure("./build_api.json", "build_api-1", false, cherry.Standalone)
	builder.Register(&MyComponent{})
	builder.Startup()
}
