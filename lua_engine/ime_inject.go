package lua_engine

import (
	"github.com/Dasongzi1366/AutoGo/ime"
	lua "github.com/yuin/gopher-lua"
)

func injectImeMethods(engine *LuaEngine) {

	engine.RegisterMethod("ime.getClipText", "获取剪切板内容", ime.GetClipText, true)
	engine.RegisterMethod("ime.setClipText", "设置剪切板内容", ime.SetClipText, true)
	engine.RegisterMethod("ime.keyAction", "模拟按键", ime.KeyAction, true)
	engine.RegisterMethod("ime.inputText", "输入文本", ime.InputText, true)
	engine.RegisterMethod("ime.getIMEList", "获取输入法列表", ime.GetIMEList, true)
	engine.RegisterMethod("ime.setCurrentIME", "设置当前输入法", ime.SetCurrentIME, true)

	registerImeLuaFunctions(engine)
}

func registerImeLuaFunctions(engine *LuaEngine) {
	state := engine.GetState()

	state.Register("ime_getClipText", func(L *lua.LState) int {
		result := ime.GetClipText()
		L.Push(lua.LString(result))
		return 1
	})

	state.Register("ime_setClipText", func(L *lua.LState) int {
		text := L.CheckString(1)
		result := ime.SetClipText(text)
		L.Push(lua.LBool(result))
		return 1
	})

	state.Register("ime_keyAction", func(L *lua.LState) int {
		code := L.CheckInt(1)
		ime.KeyAction(code)
		return 0
	})

	state.Register("ime_inputText", func(L *lua.LState) int {
		text := L.CheckString(1)
		displayId := L.CheckInt(2)
		ime.InputText(text, displayId)
		return 0
	})

	state.Register("ime_getIMEList", func(L *lua.LState) int {
		result := ime.GetIMEList()
		tbl := L.NewTable()
		for i, s := range result {
			L.SetTable(tbl, lua.LNumber(i+1), lua.LString(s))
		}
		L.Push(tbl)
		return 1
	})

	state.Register("ime_setCurrentIME", func(L *lua.LState) int {
		packageName := L.CheckString(1)
		ime.SetCurrentIME(packageName)
		return 0
	})
}
