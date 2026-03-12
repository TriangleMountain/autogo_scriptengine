package lua_engine

import (
	"github.com/Dasongzi1366/AutoGo/media"
	lua "github.com/yuin/gopher-lua"
)

func injectMediaMethods(engine *LuaEngine) {

	engine.RegisterMethod("media.scanFile", "扫描路径path的媒体文件", func(path string) { media.ScanFile(path) }, true)

	registerMediaLuaFunctions(engine)
}

func registerMediaLuaFunctions(engine *LuaEngine) {
	state := engine.GetState()

	state.Register("media_scanFile", func(L *lua.LState) int {
		path := L.CheckString(1)
		media.ScanFile(path)
		return 0
	})
}
