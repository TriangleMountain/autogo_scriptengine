package lua_engine

import (
	"github.com/Dasongzi1366/AutoGo/motion"
	lua "github.com/yuin/gopher-lua"
)

func injectMotionMethods(engine *LuaEngine) {

	engine.RegisterMethod("touchDown", "按下屏幕", func(x, y, fingerID, displayId int) { motion.TouchDown(x, y, fingerID, displayId) }, true)
	engine.RegisterMethod("touchMove", "移动手指", func(x, y, fingerID, displayId int) { motion.TouchMove(x, y, fingerID, displayId) }, true)
	engine.RegisterMethod("touchUp", "抬起手指", func(x, y, fingerID, displayId int) { motion.TouchUp(x, y, fingerID, displayId) }, true)
	engine.RegisterMethod("click", "点击", func(x, y, fingerID, displayId int) { motion.Click(x, y, fingerID, displayId) }, true)
	engine.RegisterMethod("longClick", "长按", func(x, y, duration, fingerID, displayId int) { motion.LongClick(x, y, duration, fingerID, displayId) }, true)
	engine.RegisterMethod("swipe", "滑动", func(x1, y1, x2, y2, duration, fingerID, displayId int) {
		motion.Swipe(x1, y1, x2, y2, duration, fingerID, displayId)
	}, true)
	engine.RegisterMethod("swipe2", "滑动(两点)", func(x1, y1, x2, y2, duration, fingerID, displayId int) {
		motion.Swipe2(x1, y1, x2, y2, duration, fingerID, displayId)
	}, true)
	engine.RegisterMethod("home", "按下Home键", func(displayId int) { motion.Home(displayId) }, true)
	engine.RegisterMethod("back", "按下返回键", func(displayId int) { motion.Back(displayId) }, true)
	engine.RegisterMethod("recents", "按下最近任务键", func(displayId int) { motion.Recents(displayId) }, true)
	engine.RegisterMethod("powerDialog", "长按电源键", motion.PowerDialog, true)
	engine.RegisterMethod("notifications", "下拉通知栏", motion.Notifications, true)
	engine.RegisterMethod("quickSettings", "下拉快捷设置", motion.QuickSettings, true)
	engine.RegisterMethod("volumeUp", "按下音量加键", func(displayId int) { motion.VolumeUp(displayId) }, true)
	engine.RegisterMethod("volumeDown", "按下音量减键", func(displayId int) { motion.VolumeDown(displayId) }, true)
	engine.RegisterMethod("camera", "按下相机键", motion.Camera, true)
	engine.RegisterMethod("keyAction", "按键动作", func(code, displayId int) { motion.KeyAction(code, displayId) }, true)

	registerMotionLuaFunctions(engine)
}

func registerMotionLuaFunctions(engine *LuaEngine) {
	state := engine.GetState()

	state.Register("touchDown", func(L *lua.LState) int {
		x := L.CheckInt(1)
		y := L.CheckInt(2)
		fingerID := 0
		displayId := 0

		if L.GetTop() >= 3 {
			fingerID = L.CheckInt(3)
		}
		if L.GetTop() >= 4 {
			displayId = L.CheckInt(4)
		}

		motion.TouchDown(x, y, fingerID, displayId)
		return 0
	})

	state.Register("touchMove", func(L *lua.LState) int {
		x := L.CheckInt(1)
		y := L.CheckInt(2)
		fingerID := 0
		displayId := 0

		if L.GetTop() >= 3 {
			fingerID = L.CheckInt(3)
		}
		if L.GetTop() >= 4 {
			displayId = L.CheckInt(4)
		}

		motion.TouchMove(x, y, fingerID, displayId)
		return 0
	})

	state.Register("touchUp", func(L *lua.LState) int {
		x := L.CheckInt(1)
		y := L.CheckInt(2)
		fingerID := 0
		displayId := 0

		if L.GetTop() >= 3 {
			fingerID = L.CheckInt(3)
		}
		if L.GetTop() >= 4 {
			displayId = L.CheckInt(4)
		}

		motion.TouchUp(x, y, fingerID, displayId)
		return 0
	})

	state.Register("click", func(L *lua.LState) int {
		x := L.CheckInt(1)
		y := L.CheckInt(2)
		fingerID := 0
		displayId := 0

		if L.GetTop() >= 3 {
			fingerID = L.CheckInt(3)
		}
		if L.GetTop() >= 4 {
			displayId = L.CheckInt(4)
		}

		motion.Click(x, y, fingerID, displayId)
		return 0
	})

	state.Register("longClick", func(L *lua.LState) int {
		x := L.CheckInt(1)
		y := L.CheckInt(2)
		duration := 500
		fingerID := 0
		displayId := 0

		if L.GetTop() >= 3 {
			duration = L.CheckInt(3)
		}
		if L.GetTop() >= 4 {
			fingerID = L.CheckInt(4)
		}
		if L.GetTop() >= 5 {
			displayId = L.CheckInt(5)
		}

		motion.LongClick(x, y, duration, fingerID, displayId)
		return 0
	})

	state.Register("swipe", func(L *lua.LState) int {
		x1 := L.CheckInt(1)
		y1 := L.CheckInt(2)
		x2 := L.CheckInt(3)
		y2 := L.CheckInt(4)
		duration := L.CheckInt(5)
		fingerID := L.CheckInt(6)
		displayId := L.CheckInt(7)
		motion.Swipe(x1, y1, x2, y2, duration, fingerID, displayId)
		return 0
	})

	state.Register("swipe2", func(L *lua.LState) int {
		x1 := L.CheckInt(1)
		y1 := L.CheckInt(2)
		x2 := L.CheckInt(3)
		y2 := L.CheckInt(4)
		duration := L.CheckInt(5)
		fingerID := L.CheckInt(6)
		displayId := L.CheckInt(7)
		motion.Swipe2(x1, y1, x2, y2, duration, fingerID, displayId)
		return 0
	})

	state.Register("home", func(L *lua.LState) int {
		displayId := 0
		if L.GetTop() >= 1 {
			displayId = L.CheckInt(1)
		}
		motion.Home(displayId)
		return 0
	})

	state.Register("back", func(L *lua.LState) int {
		displayId := 0
		if L.GetTop() >= 1 {
			displayId = L.CheckInt(1)
		}
		motion.Back(displayId)
		return 0
	})

	state.Register("recents", func(L *lua.LState) int {
		displayId := 0
		if L.GetTop() >= 1 {
			displayId = L.CheckInt(1)
		}
		motion.Recents(displayId)
		return 0
	})

	state.Register("powerDialog", func(L *lua.LState) int {
		motion.PowerDialog()
		return 0
	})

	state.Register("notifications", func(L *lua.LState) int {
		motion.Notifications()
		return 0
	})

	state.Register("quickSettings", func(L *lua.LState) int {
		motion.QuickSettings()
		return 0
	})

	state.Register("volumeUp", func(L *lua.LState) int {
		displayId := 0
		if L.GetTop() >= 1 {
			displayId = L.CheckInt(1)
		}
		motion.VolumeUp(displayId)
		return 0
	})

	state.Register("volumeDown", func(L *lua.LState) int {
		displayId := 0
		if L.GetTop() >= 1 {
			displayId = L.CheckInt(1)
		}
		motion.VolumeDown(displayId)
		return 0
	})

	state.Register("camera", func(L *lua.LState) int {
		motion.Camera()
		return 0
	})

	state.Register("keyAction", func(L *lua.LState) int {
		code := L.CheckInt(1)
		displayId := 0
		if L.GetTop() >= 2 {
			displayId = L.CheckInt(2)
		}
		motion.KeyAction(code, displayId)
		return 0
	})
}
