package lua_engine

import (
	"github.com/Dasongzi1366/AutoGo/utils"
	lua "github.com/yuin/gopher-lua"
)

func injectUtilsMethods(engine *LuaEngine) {

	engine.RegisterMethod("utils.logI", "记录一条INFO级别的日志", utils.LogI, true)
	engine.RegisterMethod("utils.logE", "记录一条ERROR级别的日志", utils.LogE, true)
	engine.RegisterMethod("utils.shell", "执行shell命令并返回输出", utils.Shell, true)
	engine.RegisterMethod("utils.toast", "显示Toast提示信息", utils.Toast, true)
	engine.RegisterMethod("utils.alert", "显示带标题、内容和按钮的弹窗", utils.Alert, true)
	engine.RegisterMethod("utils.random", "返回指定范围内的随机整数", utils.Random, true)
	engine.RegisterMethod("utils.sleep", "暂停当前线程指定的毫秒数", utils.Sleep, true)
	engine.RegisterMethod("utils.i2s", "将整数转换为字符串", utils.I2s, true)
	engine.RegisterMethod("utils.s2i", "将字符串转换为整数", utils.S2i, true)
	engine.RegisterMethod("utils.f2s", "将浮点数转换为字符串", utils.F2s, true)
	engine.RegisterMethod("utils.s2f", "将字符串转换为浮点数", utils.S2f, true)
	engine.RegisterMethod("utils.b2s", "将布尔值转换为字符串", utils.B2s, true)
	engine.RegisterMethod("utils.s2b", "将字符串转换为布尔值", utils.S2b, true)

	registerUtilsLuaFunctions(engine)
}

func registerUtilsLuaFunctions(engine *LuaEngine) {
	state := engine.GetState()

	state.Register("utils_logI", func(L *lua.LState) int {
		label := L.CheckString(1)
		n := L.GetTop()
		var messages []any
		for i := 2; i <= n; i++ {
			messages = append(messages, L.Get(i).String())
		}
		utils.LogI(label, messages...)
		return 0
	})

	state.Register("utils_logE", func(L *lua.LState) int {
		label := L.CheckString(1)
		n := L.GetTop()
		var messages []any
		for i := 2; i <= n; i++ {
			messages = append(messages, L.Get(i).String())
		}
		utils.LogE(label, messages...)
		return 0
	})

	state.Register("utils_shell", func(L *lua.LState) int {
		cmd := L.CheckString(1)
		result := utils.Shell(cmd)
		L.Push(lua.LString(result))
		return 1
	})

	state.Register("utils_toast", func(L *lua.LState) int {
		message := L.CheckString(1)
		utils.Toast(message)
		return 0
	})

	state.Register("utils_alert", func(L *lua.LState) int {
		title := L.CheckString(1)
		content := L.CheckString(2)
		btn1Text := L.CheckString(3)
		btn2Text := L.CheckString(4)
		result := utils.Alert(title, content, btn1Text, btn2Text)
		L.Push(lua.LNumber(result))
		return 1
	})

	state.Register("utils_random", func(L *lua.LState) int {
		min := L.CheckInt(1)
		max := L.CheckInt(2)
		result := utils.Random(min, max)
		L.Push(lua.LNumber(result))
		return 1
	})

	state.Register("utils_sleep", func(L *lua.LState) int {
		i := L.CheckInt(1)
		utils.Sleep(i)
		return 0
	})

	state.Register("utils_i2s", func(L *lua.LState) int {
		i := L.CheckInt(1)
		result := utils.I2s(i)
		L.Push(lua.LString(result))
		return 1
	})

	state.Register("utils_s2i", func(L *lua.LState) int {
		s := L.CheckString(1)
		result := utils.S2i(s)
		L.Push(lua.LNumber(result))
		return 1
	})

	state.Register("utils_f2s", func(L *lua.LState) int {
		f := float64(L.CheckNumber(1))
		result := utils.F2s(f)
		L.Push(lua.LString(result))
		return 1
	})

	state.Register("utils_s2f", func(L *lua.LState) int {
		s := L.CheckString(1)
		result := utils.S2f(s)
		L.Push(lua.LNumber(result))
		return 1
	})

	state.Register("utils_b2s", func(L *lua.LState) int {
		b := L.CheckBool(1)
		result := utils.B2s(b)
		L.Push(lua.LString(result))
		return 1
	})

	state.Register("utils_s2b", func(L *lua.LState) int {
		s := L.CheckString(1)
		result := utils.S2b(s)
		L.Push(lua.LBool(result))
		return 1
	})
}
