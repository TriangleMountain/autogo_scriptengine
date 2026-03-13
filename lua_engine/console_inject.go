package lua_engine

import (
	"github.com/Dasongzi1366/AutoGo/console"
	lua "github.com/yuin/gopher-lua"
)

func injectConsoleMethods(engine *LuaEngine) {

	engine.RegisterMethod("console.new", "创建一个新的Console对象", console.New, true)

	registerConsoleLuaFunctions(engine)
}

func registerConsoleLuaFunctions(engine *LuaEngine) {
	state := engine.GetState()

	state.Register("console_new", func(L *lua.LState) int {
		c := console.New()
		ud := L.NewUserData()
		ud.Value = c
		L.Push(ud)
		return 1
	})

	state.Register("console_setWindowSize", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		c := ud.Value.(*console.Console)
		width := L.CheckInt(2)
		height := L.CheckInt(3)
		c.SetWindowSize(width, height)
		ud2 := L.NewUserData()
		ud2.Value = c
		L.Push(ud2)
		return 1
	})

	state.Register("console_setWindowPosition", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		c := ud.Value.(*console.Console)
		x := L.CheckInt(2)
		y := L.CheckInt(3)
		c.SetWindowPosition(x, y)
		ud2 := L.NewUserData()
		ud2.Value = c
		L.Push(ud2)
		return 1
	})

	state.Register("console_setWindowColor", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		c := ud.Value.(*console.Console)
		color := L.CheckString(2)
		c.SetWindowColor(color)
		ud2 := L.NewUserData()
		ud2.Value = c
		L.Push(ud2)
		return 1
	})

	state.Register("console_setTextColor", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		c := ud.Value.(*console.Console)
		color := L.CheckString(2)
		c.SetTextColor(color)
		ud2 := L.NewUserData()
		ud2.Value = c
		L.Push(ud2)
		return 1
	})

	state.Register("console_setTextSize", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		c := ud.Value.(*console.Console)
		size := L.CheckInt(2)
		c.SetTextSize(size)
		ud2 := L.NewUserData()
		ud2.Value = c
		L.Push(ud2)
		return 1
	})

	state.Register("console_println", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		c := ud.Value.(*console.Console)
		n := L.GetTop()
		var args []any
		for i := 2; i <= n; i++ {
			args = append(args, L.Get(i).String())
		}
		c.Println(args...)
		return 0
	})

	state.Register("console_clear", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		c := ud.Value.(*console.Console)
		c.Clear()
		return 0
	})

	state.Register("console_show", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		c := ud.Value.(*console.Console)
		c.Show()
		return 0
	})

	state.Register("console_hide", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		c := ud.Value.(*console.Console)
		c.Hide()
		return 0
	})

	state.Register("console_isVisible", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		c := ud.Value.(*console.Console)
		result := c.IsVisible()
		L.Push(lua.LBool(result))
		return 1
	})

	state.Register("console_destroy", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		c := ud.Value.(*console.Console)
		c.Destroy()
		return 0
	})
}
