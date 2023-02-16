package router

import "github.com/Gocyber-world/navigator-demo/router/system"

type RouterGroup struct {
	System system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
