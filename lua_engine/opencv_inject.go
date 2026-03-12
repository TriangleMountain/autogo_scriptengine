package lua_engine

import (
	"github.com/Dasongzi1366/AutoGo/opencv"
	lua "github.com/yuin/gopher-lua"
)

func injectOpenCVMethods(engine *LuaEngine) {

	engine.RegisterMethod("opencv.findImage", "在指定区域内查找匹配的图片模板", func(x1, y1, x2, y2 int, template *[]byte, isGray bool, scalingFactor, sim float32, displayId int) (int, int) {
		return opencv.FindImage(x1, y1, x2, y2, template, isGray, scalingFactor, sim, displayId)
	}, true)

	registerOpenCVLuaFunctions(engine)
}

func registerOpenCVLuaFunctions(engine *LuaEngine) {
	state := engine.GetState()

	state.Register("opencv_findImage", func(L *lua.LState) int {
		x1 := L.CheckInt(1)
		y1 := L.CheckInt(2)
		x2 := L.CheckInt(3)
		y2 := L.CheckInt(4)
		templateData := L.CheckString(5)
		template := []byte(templateData)
		isGray := L.CheckBool(6)
		scalingFactor := float32(L.CheckNumber(7))
		sim := float32(L.CheckNumber(8))
		displayId := 0
		if L.GetTop() >= 9 {
			displayId = L.CheckInt(9)
		}
		x, y := opencv.FindImage(x1, y1, x2, y2, &template, isGray, scalingFactor, sim, displayId)
		L.Push(lua.LNumber(x))
		L.Push(lua.LNumber(y))
		return 2
	})
}
