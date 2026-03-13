package lua_engine

import (
	"github.com/Dasongzi1366/AutoGo/dotocr"
	lua "github.com/yuin/gopher-lua"
)

func injectDotocrMethods(engine *LuaEngine) {

	engine.RegisterMethod("dotocr.setDict", "设置字库", dotocr.SetDict, true)
	engine.RegisterMethod("dotocr.ocr", "从屏幕指定区域进行OCR识别", func(x1, y1, x2, y2 int, threshold string, colGap, rowGap int, sim float32, mode int, dictName string, displayId int) string {
		return dotocr.Ocr(x1, y1, x2, y2, threshold, colGap, rowGap, sim, mode, dictName, displayId)
	}, true)
	engine.RegisterMethod("dotocr.ocrFromBase64", "从Base64编码的图像字符串进行OCR识别", dotocr.OcrFromBase64, true)
	engine.RegisterMethod("dotocr.ocrFromPath", "从图像文件路径进行OCR识别", dotocr.OcrFromPath, true)
	engine.RegisterMethod("dotocr.findStr", "在屏幕指定区域中查找指定字符串的位置", dotocr.FindStr, true)
	engine.RegisterMethod("dotocr.findStrFromBase64", "在Base64编码的图像中查找指定字符串的位置", dotocr.FindStrFromBase64, true)
	engine.RegisterMethod("dotocr.findStrFromPath", "在图像文件中查找指定字符串的位置", dotocr.FindStrFromPath, true)

	registerDotocrLuaFunctions(engine)
}

func registerDotocrLuaFunctions(engine *LuaEngine) {
	state := engine.GetState()

	state.Register("dotocr_setDict", func(L *lua.LState) int {
		name := L.CheckString(1)
		dict := L.CheckString(2)
		dotocr.SetDict(name, dict)
		return 0
	})

	state.Register("dotocr_ocr", func(L *lua.LState) int {
		x1 := L.CheckInt(1)
		y1 := L.CheckInt(2)
		x2 := L.CheckInt(3)
		y2 := L.CheckInt(4)
		threshold := L.CheckString(5)
		colGap := L.CheckInt(6)
		rowGap := L.CheckInt(7)
		sim := float32(L.CheckNumber(8))
		mode := L.CheckInt(9)
		dictName := L.CheckString(10)
		displayId := L.CheckInt(11)
		result := dotocr.Ocr(x1, y1, x2, y2, threshold, colGap, rowGap, sim, mode, dictName, displayId)
		L.Push(lua.LString(result))
		return 1
	})

	state.Register("dotocr_ocrFromBase64", func(L *lua.LState) int {
		b64 := L.CheckString(1)
		threshold := L.CheckString(2)
		colGap := L.CheckInt(3)
		rowGap := L.CheckInt(4)
		sim := float32(L.CheckNumber(5))
		mode := L.CheckInt(6)
		dictName := L.CheckString(7)
		result := dotocr.OcrFromBase64(b64, threshold, colGap, rowGap, sim, mode, dictName)
		L.Push(lua.LString(result))
		return 1
	})

	state.Register("dotocr_ocrFromPath", func(L *lua.LState) int {
		path := L.CheckString(1)
		threshold := L.CheckString(2)
		colGap := L.CheckInt(3)
		rowGap := L.CheckInt(4)
		sim := float32(L.CheckNumber(5))
		mode := L.CheckInt(6)
		dictName := L.CheckString(7)
		result := dotocr.OcrFromPath(path, threshold, colGap, rowGap, sim, mode, dictName)
		L.Push(lua.LString(result))
		return 1
	})

	state.Register("dotocr_findStr", func(L *lua.LState) int {
		x1 := L.CheckInt(1)
		y1 := L.CheckInt(2)
		x2 := L.CheckInt(3)
		y2 := L.CheckInt(4)
		text := L.CheckString(5)
		threshold := L.CheckString(6)
		colGap := L.CheckInt(7)
		rowGap := L.CheckInt(8)
		sim := float32(L.CheckNumber(9))
		dictName := L.CheckString(10)
		displayId := L.CheckInt(11)
		x, y := dotocr.FindStr(x1, y1, x2, y2, text, threshold, colGap, rowGap, sim, dictName, displayId)
		L.Push(lua.LNumber(x))
		L.Push(lua.LNumber(y))
		return 2
	})

	state.Register("dotocr_findStrFromBase64", func(L *lua.LState) int {
		b64 := L.CheckString(1)
		text := L.CheckString(2)
		threshold := L.CheckString(3)
		colGap := L.CheckInt(4)
		rowGap := L.CheckInt(5)
		sim := float32(L.CheckNumber(6))
		dictName := L.CheckString(7)
		x, y := dotocr.FindStrFromBase64(b64, text, threshold, colGap, rowGap, sim, dictName)
		L.Push(lua.LNumber(x))
		L.Push(lua.LNumber(y))
		return 2
	})

	state.Register("dotocr_findStrFromPath", func(L *lua.LState) int {
		path := L.CheckString(1)
		text := L.CheckString(2)
		threshold := L.CheckString(3)
		colGap := L.CheckInt(4)
		rowGap := L.CheckInt(5)
		sim := float32(L.CheckNumber(6))
		dictName := L.CheckString(7)
		x, y := dotocr.FindStrFromPath(path, text, threshold, colGap, rowGap, sim, dictName)
		L.Push(lua.LNumber(x))
		L.Push(lua.LNumber(y))
		return 2
	})
}
