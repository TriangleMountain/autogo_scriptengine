package lua_engine

import (
	"github.com/Dasongzi1366/AutoGo/vdisplay"
	lua "github.com/yuin/gopher-lua"
)

func injectVdisplayMethods(engine *LuaEngine) {

	engine.RegisterMethod("vdisplay.create", "创建一个虚拟显示设备", vdisplay.Create, true)
	engine.RegisterMethod("vdisplay.getDisplayId", "获取虚拟显示设备的DisplayId", (*vdisplay.Vdisplay).GetDisplayId, true)
	engine.RegisterMethod("vdisplay.launchApp", "启动指定包名的应用到虚拟显示设备内", (*vdisplay.Vdisplay).LaunchApp, true)
	engine.RegisterMethod("vdisplay.setTitle", "设置预览窗口标题", (*vdisplay.Vdisplay).SetTitle, true)
	engine.RegisterMethod("vdisplay.showPreviewWindow", "显示预览窗口", (*vdisplay.Vdisplay).ShowPreviewWindow, true)
	engine.RegisterMethod("vdisplay.hidePreviewWindow", "隐藏预览窗口", (*vdisplay.Vdisplay).HidePreviewWindow, true)
	engine.RegisterMethod("vdisplay.setPreviewWindowSize", "设置预览窗口大小", (*vdisplay.Vdisplay).SetPreviewWindowSize, true)
	engine.RegisterMethod("vdisplay.setPreviewWindowPos", "设置预览窗口位置", (*vdisplay.Vdisplay).SetPreviewWindowPos, true)
	engine.RegisterMethod("vdisplay.destroy", "销毁指定的虚拟显示设备", (*vdisplay.Vdisplay).Destroy, true)

	registerVdisplayLuaFunctions(engine)
}

func registerVdisplayLuaFunctions(engine *LuaEngine) {
	state := engine.GetState()

	state.Register("vdisplay_create", func(L *lua.LState) int {
		width := L.CheckInt(1)
		height := L.CheckInt(2)
		dpi := L.CheckInt(3)
		result := vdisplay.Create(width, height, dpi)
		if result == nil {
			L.Push(lua.LNil)
			return 1
		}
		ud := L.NewUserData()
		ud.Value = result
		L.Push(ud)
		return 1
	})

	state.Register("vdisplay_getDisplayId", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		v := ud.Value.(*vdisplay.Vdisplay)
		result := v.GetDisplayId()
		L.Push(lua.LNumber(result))
		return 1
	})

	state.Register("vdisplay_launchApp", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		v := ud.Value.(*vdisplay.Vdisplay)
		packageName := L.CheckString(2)
		result := v.LaunchApp(packageName)
		L.Push(lua.LBool(result))
		return 1
	})

	state.Register("vdisplay_setTitle", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		v := ud.Value.(*vdisplay.Vdisplay)
		title := L.CheckString(2)
		v.SetTitle(title)
		return 0
	})

	state.Register("vdisplay_showPreviewWindow", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		v := ud.Value.(*vdisplay.Vdisplay)
		rotated := L.CheckBool(2)
		v.ShowPreviewWindow(rotated)
		return 0
	})

	state.Register("vdisplay_hidePreviewWindow", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		v := ud.Value.(*vdisplay.Vdisplay)
		v.HidePreviewWindow()
		return 0
	})

	state.Register("vdisplay_setPreviewWindowSize", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		v := ud.Value.(*vdisplay.Vdisplay)
		width := L.CheckInt(2)
		height := L.CheckInt(3)
		v.SetPreviewWindowSize(width, height)
		return 0
	})

	state.Register("vdisplay_setPreviewWindowPos", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		v := ud.Value.(*vdisplay.Vdisplay)
		x := L.CheckInt(2)
		y := L.CheckInt(3)
		v.SetPreviewWindowPos(x, y)
		return 0
	})

	state.Register("vdisplay_destroy", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		v := ud.Value.(*vdisplay.Vdisplay)
		v.Destroy()
		return 0
	})
}
