package lua_engine

import (
	"github.com/Dasongzi1366/AutoGo/plugin"
	lua "github.com/yuin/gopher-lua"
)

func injectPluginMethods(engine *LuaEngine) {

	engine.RegisterMethod("plugin.newContext", "创建Context参数", plugin.NewContext, true)
	engine.RegisterMethod("plugin.newAssetManager", "创建AssetManager参数", plugin.NewAssetManager, true)
	engine.RegisterMethod("plugin.loadApk", "加载外部APK", plugin.LoadApk, true)
	engine.RegisterMethod("plugin.newInstance", "创建类实例", (*plugin.ClassLoader).NewInstance, true)
	engine.RegisterMethod("plugin.callString", "调用返回String的方法", (*plugin.Instance).CallString, true)
	engine.RegisterMethod("plugin.callInt", "调用返回int的方法", (*plugin.Instance).CallInt, true)
	engine.RegisterMethod("plugin.callLong", "调用返回long的方法", (*plugin.Instance).CallLong, true)
	engine.RegisterMethod("plugin.callFloat", "调用返回float的方法", (*plugin.Instance).CallFloat, true)
	engine.RegisterMethod("plugin.callDouble", "调用返回double的方法", (*plugin.Instance).CallDouble, true)
	engine.RegisterMethod("plugin.callBool", "调用返回boolean的方法", (*plugin.Instance).CallBool, true)
	engine.RegisterMethod("plugin.callVoid", "调用无返回值的方法", (*plugin.Instance).CallVoid, true)
	engine.RegisterMethod("plugin.releaseInstance", "释放实例", (*plugin.Instance).Release, true)
	engine.RegisterMethod("plugin.releaseClassLoader", "释放类加载器", (*plugin.ClassLoader).Release, true)

	registerPluginLuaFunctions(engine)
}

func registerPluginLuaFunctions(engine *LuaEngine) {
	state := engine.GetState()

	state.Register("plugin_newContext", func(L *lua.LState) int {
		ctx := plugin.NewContext()
		ud := L.NewUserData()
		ud.Value = ctx
		L.Push(ud)
		return 1
	})

	state.Register("plugin_newAssetManager", func(L *lua.LState) int {
		am := plugin.NewAssetManager()
		ud := L.NewUserData()
		ud.Value = am
		L.Push(ud)
		return 1
	})

	state.Register("plugin_loadApk", func(L *lua.LState) int {
		apkPath := L.CheckString(1)
		cl := plugin.LoadApk(apkPath)
		ud := L.NewUserData()
		ud.Value = cl
		L.Push(ud)
		return 1
	})

	state.Register("plugin_newInstance", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		cl := ud.Value.(*plugin.ClassLoader)
		className := L.CheckString(2)
		n := L.GetTop()
		var args []interface{}
		for i := 3; i <= n; i++ {
			args = append(args, L.Get(i).String())
		}
		inst := cl.NewInstance(className, args...)
		ud2 := L.NewUserData()
		ud2.Value = inst
		L.Push(ud2)
		return 1
	})

	state.Register("plugin_callString", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		inst := ud.Value.(*plugin.Instance)
		methodName := L.CheckString(2)
		n := L.GetTop()
		var args []interface{}
		for i := 3; i <= n; i++ {
			args = append(args, L.Get(i).String())
		}
		result, _ := inst.CallString(methodName, args...)
		L.Push(lua.LString(result))
		return 1
	})

	state.Register("plugin_callInt", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		inst := ud.Value.(*plugin.Instance)
		methodName := L.CheckString(2)
		n := L.GetTop()
		var args []interface{}
		for i := 3; i <= n; i++ {
			args = append(args, L.Get(i).String())
		}
		result, _ := inst.CallInt(methodName, args...)
		L.Push(lua.LNumber(result))
		return 1
	})

	state.Register("plugin_callLong", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		inst := ud.Value.(*plugin.Instance)
		methodName := L.CheckString(2)
		n := L.GetTop()
		var args []interface{}
		for i := 3; i <= n; i++ {
			args = append(args, L.Get(i).String())
		}
		result, _ := inst.CallLong(methodName, args...)
		L.Push(lua.LNumber(result))
		return 1
	})

	state.Register("plugin_callFloat", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		inst := ud.Value.(*plugin.Instance)
		methodName := L.CheckString(2)
		n := L.GetTop()
		var args []interface{}
		for i := 3; i <= n; i++ {
			args = append(args, L.Get(i).String())
		}
		result, _ := inst.CallFloat(methodName, args...)
		L.Push(lua.LNumber(result))
		return 1
	})

	state.Register("plugin_callDouble", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		inst := ud.Value.(*plugin.Instance)
		methodName := L.CheckString(2)
		n := L.GetTop()
		var args []interface{}
		for i := 3; i <= n; i++ {
			args = append(args, L.Get(i).String())
		}
		result, _ := inst.CallDouble(methodName, args...)
		L.Push(lua.LNumber(result))
		return 1
	})

	state.Register("plugin_callBool", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		inst := ud.Value.(*plugin.Instance)
		methodName := L.CheckString(2)
		n := L.GetTop()
		var args []interface{}
		for i := 3; i <= n; i++ {
			args = append(args, L.Get(i).String())
		}
		result, _ := inst.CallBool(methodName, args...)
		L.Push(lua.LBool(result))
		return 1
	})

	state.Register("plugin_callVoid", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		inst := ud.Value.(*plugin.Instance)
		methodName := L.CheckString(2)
		n := L.GetTop()
		var args []interface{}
		for i := 3; i <= n; i++ {
			args = append(args, L.Get(i).String())
		}
		inst.CallVoid(methodName, args...)
		return 0
	})

	state.Register("plugin_releaseInstance", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		inst := ud.Value.(*plugin.Instance)
		inst.Release()
		return 0
	})

	state.Register("plugin_releaseClassLoader", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		cl := ud.Value.(*plugin.ClassLoader)
		cl.Release()
		return 0
	})
}
