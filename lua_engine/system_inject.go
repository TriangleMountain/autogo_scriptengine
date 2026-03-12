package lua_engine

import (
	"github.com/Dasongzi1366/AutoGo/system"
	lua "github.com/yuin/gopher-lua"
)

func injectSystemMethods(engine *LuaEngine) {

	engine.RegisterMethod("system.getPid", "获取进程ID", func(processName string) int { return system.GetPid(processName) }, true)
	engine.RegisterMethod("system.getMemoryUsage", "获取内存使用", func(pid int) int { return system.GetMemoryUsage(pid) }, true)
	engine.RegisterMethod("system.getCpuUsage", "获取CPU使用率", func(pid int) float64 { return system.GetCpuUsage(pid) }, true)
	engine.RegisterMethod("system.restartSelf", "重启自身", system.RestartSelf, true)
	engine.RegisterMethod("system.setBootStart", "设置开机自启", func(enable bool) { system.SetBootStart(enable) }, true)

	registerSystemLuaFunctions(engine)
}

func registerSystemLuaFunctions(engine *LuaEngine) {
	state := engine.GetState()

	state.Register("system_getPid", func(L *lua.LState) int {
		processName := L.CheckString(1)
		result := system.GetPid(processName)
		L.Push(lua.LNumber(result))
		return 1
	})

	state.Register("system_getMemoryUsage", func(L *lua.LState) int {
		pid := L.CheckInt(1)
		result := system.GetMemoryUsage(pid)
		L.Push(lua.LNumber(result))
		return 1
	})

	state.Register("system_getCpuUsage", func(L *lua.LState) int {
		pid := L.CheckInt(1)
		result := system.GetCpuUsage(pid)
		L.Push(lua.LNumber(result))
		return 1
	})

	state.Register("system_restartSelf", func(L *lua.LState) int {
		system.RestartSelf()
		return 0
	})

	state.Register("system_setBootStart", func(L *lua.LState) int {
		enable := L.CheckBool(1)
		system.SetBootStart(enable)
		return 0
	})
}
