package lua_engine

import (
	"github.com/Dasongzi1366/AutoGo/app"
	lua "github.com/yuin/gopher-lua"
)

func injectAppMethods(engine *LuaEngine) {

	engine.RegisterMethod("app.currentPackage", "获取当前页面应用包名", app.CurrentPackage, true)
	engine.RegisterMethod("app.currentActivity", "获取当前页面应用类名", app.CurrentActivity, true)
	engine.RegisterMethod("app.launch", "通过应用包名启动应用", func(packageName string, displayId int) bool {
		return app.Launch(packageName, displayId)
	}, true)
	engine.RegisterMethod("app.openAppSetting", "打开应用的详情页(设置页)", func(packageName string) bool {
		return app.OpenSetting(packageName)
	}, true)
	engine.RegisterMethod("app.viewFile", "用其他应用查看文件", func(path string) {
		app.ViewFile(path)
	}, true)
	engine.RegisterMethod("app.editFile", "用其他应用编辑文件", func(path string) {
		app.EditFile(path)
	}, true)
	engine.RegisterMethod("app.uninstall", "卸载应用", func(packageName string) {
		app.Uninstall(packageName)
	}, true)
	engine.RegisterMethod("app.install", "安装应用", func(path string) {
		app.Install(path)
	}, true)
	engine.RegisterMethod("app.isInstalled", "判断是否已经安装某个应用", func(packageName string) bool {
		return app.IsInstalled(packageName)
	}, true)
	engine.RegisterMethod("app.clear", "清除应用数据", func(packageName string) {
		app.Clear(packageName)
	}, true)
	engine.RegisterMethod("app.forceStop", "强制停止应用", func(packageName string) {
		app.ForceStop(packageName)
	}, true)
	engine.RegisterMethod("app.disable", "禁用应用", func(packageName string) {
		app.Disable(packageName)
	}, true)
	engine.RegisterMethod("app.ignoreBattOpt", "忽略电池优化", func(packageName string) {
		app.IgnoreBattOpt(packageName)
	}, true)
	engine.RegisterMethod("app.enable", "启用应用", func(packageName string) {
		app.Enable(packageName)
	}, true)
	engine.RegisterMethod("app.getBrowserPackage", "获取系统默认浏览器包名", app.GetBrowserPackage, true)
	engine.RegisterMethod("app.openUrl", "用浏览器打开网站url", func(url string) {
		app.OpenUrl(url)
	}, true)
	engine.RegisterMethod("app.startActivity", "根据选项构造一个Intent，并启动该Activity", func(options app.IntentOptions) {
		app.StartActivity(options)
	}, true)
	engine.RegisterMethod("app.sendBroadcast", "根据选项构造一个Intent，并发送该广播", func(options app.IntentOptions) {
		app.SendBroadcast(options)
	}, true)
	engine.RegisterMethod("app.startService", "根据选项构造一个Intent，并启动该服务", func(options app.IntentOptions) {
		app.StartService(options)
	}, true)

	registerAppLuaFunctions(engine)
}

func registerAppLuaFunctions(engine *LuaEngine) {
	state := engine.GetState()

	state.Register("app_currentPackage", func(L *lua.LState) int {
		result := app.CurrentPackage()
		L.Push(lua.LString(result))
		return 1
	})

	state.Register("app_currentActivity", func(L *lua.LState) int {
		result := app.CurrentActivity()
		L.Push(lua.LString(result))
		return 1
	})

	state.Register("app_launch", func(L *lua.LState) int {
		packageName := L.CheckString(1)
		displayId := L.CheckInt(2)
		result := app.Launch(packageName, displayId)
		L.Push(lua.LBool(result))
		return 1
	})

	state.Register("app_openAppSetting", func(L *lua.LState) int {
		packageName := L.CheckString(1)
		result := app.OpenSetting(packageName)
		L.Push(lua.LBool(result))
		return 1
	})

	state.Register("app_viewFile", func(L *lua.LState) int {
		path := L.CheckString(1)
		app.ViewFile(path)
		return 0
	})

	state.Register("app_editFile", func(L *lua.LState) int {
		path := L.CheckString(1)
		app.EditFile(path)
		return 0
	})

	state.Register("app_uninstall", func(L *lua.LState) int {
		packageName := L.CheckString(1)
		app.Uninstall(packageName)
		return 0
	})

	state.Register("app_install", func(L *lua.LState) int {
		path := L.CheckString(1)
		app.Install(path)
		return 0
	})

	state.Register("app_isInstalled", func(L *lua.LState) int {
		packageName := L.CheckString(1)
		result := app.IsInstalled(packageName)
		L.Push(lua.LBool(result))
		return 1
	})

	state.Register("app_clear", func(L *lua.LState) int {
		packageName := L.CheckString(1)
		app.Clear(packageName)
		return 0
	})

	state.Register("app_forceStop", func(L *lua.LState) int {
		packageName := L.CheckString(1)
		app.ForceStop(packageName)
		return 0
	})

	state.Register("app_disable", func(L *lua.LState) int {
		packageName := L.CheckString(1)
		app.Disable(packageName)
		return 0
	})

	state.Register("app_ignoreBattOpt", func(L *lua.LState) int {
		packageName := L.CheckString(1)
		app.IgnoreBattOpt(packageName)
		return 0
	})

	state.Register("app_enable", func(L *lua.LState) int {
		packageName := L.CheckString(1)
		app.Enable(packageName)
		return 0
	})

	state.Register("app_getBrowserPackage", func(L *lua.LState) int {
		result := app.GetBrowserPackage()
		L.Push(lua.LString(result))
		return 1
	})

	state.Register("app_openUrl", func(L *lua.LState) int {
		url := L.CheckString(1)
		app.OpenUrl(url)
		return 0
	})
}
