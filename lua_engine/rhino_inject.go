package lua_engine

import (
	"github.com/Dasongzi1366/AutoGo/rhino"
	lua "github.com/yuin/gopher-lua"
)

func injectRhinoMethods(engine *LuaEngine) {

	engine.RegisterMethod("rhino.eval", "执行指定的JavaScript脚本并返回结果", rhino.Eval, true)

	registerRhinoLuaFunctions(engine)
}

func registerRhinoLuaFunctions(engine *LuaEngine) {
	state := engine.GetState()

	state.Register("rhino_eval", func(L *lua.LState) int {
		contextId := L.CheckString(1)
		script := L.CheckString(2)
		result := rhino.Eval(contextId, script)
		L.Push(lua.LString(result))
		return 1
	})
}
