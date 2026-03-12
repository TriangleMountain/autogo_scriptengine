package lua_engine

import (
	"image"
	"sync"

	"github.com/Dasongzi1366/AutoGo/ppocr"
	lua "github.com/yuin/gopher-lua"
)

var (
	ppocrInstance *ppocr.Ppocr
	ppocrOnce     sync.Once
	ppocrInitErr  error
)

func getPpocrInstance() (*ppocr.Ppocr, error) {
	ppocrOnce.Do(func() {
		ppocrInstance = ppocr.New("v5")
	})
	return ppocrInstance, nil
}

func injectPpocrMethods(engine *LuaEngine) {

	engine.RegisterMethod("ppocr.ocr", "识别屏幕文字", func(x1, y1, x2, y2 int, colorStr string, displayId int) []ppocr.Result {
		instance, _ := getPpocrInstance()
		return instance.Ocr(x1, y1, x2, y2, colorStr, displayId)
	}, true)
	engine.RegisterMethod("ppocr.ocrFromImage", "识别图片文字", func(img *image.NRGBA, colorStr string) []ppocr.Result {
		instance, _ := getPpocrInstance()
		return instance.OcrFromImage(img, colorStr)
	}, true)
	engine.RegisterMethod("ppocr.ocrFromBase64", "识别Base64图片文字", func(b64, colorStr string) []ppocr.Result {
		instance, _ := getPpocrInstance()
		return instance.OcrFromBase64(b64, colorStr)
	}, true)
	engine.RegisterMethod("ppocr.ocrFromPath", "识别文件图片文字", func(path, colorStr string) []ppocr.Result {
		instance, _ := getPpocrInstance()
		return instance.OcrFromPath(path, colorStr)
	}, true)

	registerPpocrLuaFunctions(engine)
}

func registerPpocrLuaFunctions(engine *LuaEngine) {
	state := engine.GetState()

	state.Register("ppocr_ocr", func(L *lua.LState) int {
		instance, _ := getPpocrInstance()
		x1 := L.CheckInt(1)
		y1 := L.CheckInt(2)
		x2 := L.CheckInt(3)
		y2 := L.CheckInt(4)
		colorStr := L.CheckString(5)
		displayId := 0
		if L.GetTop() >= 6 {
			displayId = L.CheckInt(6)
		}
		results := instance.Ocr(x1, y1, x2, y2, colorStr, displayId)
		pushResultsToLua(L, results)
		return 1
	})

	state.Register("ppocr_ocrFromImage", func(L *lua.LState) int {
		instance, _ := getPpocrInstance()
		ud := L.CheckUserData(1)
		img, ok := ud.Value.(*image.NRGBA)
		if !ok {
			L.Push(lua.LNil)
			return 1
		}
		colorStr := L.CheckString(2)
		results := instance.OcrFromImage(img, colorStr)
		pushResultsToLua(L, results)
		return 1
	})

	state.Register("ppocr_ocrFromBase64", func(L *lua.LState) int {
		instance, _ := getPpocrInstance()
		b64 := L.CheckString(1)
		colorStr := L.CheckString(2)
		results := instance.OcrFromBase64(b64, colorStr)
		pushResultsToLua(L, results)
		return 1
	})

	state.Register("ppocr_ocrFromPath", func(L *lua.LState) int {
		instance, _ := getPpocrInstance()
		path := L.CheckString(1)
		colorStr := L.CheckString(2)
		results := instance.OcrFromPath(path, colorStr)
		pushResultsToLua(L, results)
		return 1
	})
}

func pushResultsToLua(L *lua.LState, results []ppocr.Result) {
	tbl := L.NewTable()
	for i, result := range results {
		item := L.NewTable()
		L.SetField(item, "X", lua.LNumber(result.X))
		L.SetField(item, "Y", lua.LNumber(result.Y))
		L.SetField(item, "宽", lua.LNumber(result.Width))
		L.SetField(item, "高", lua.LNumber(result.Height))
		L.SetField(item, "标签", lua.LString(result.Label))
		L.SetField(item, "精度", lua.LNumber(result.Score))
		L.SetField(item, "CenterX", lua.LNumber(result.CenterX))
		L.SetField(item, "CenterY", lua.LNumber(result.CenterY))
		L.SetTable(tbl, lua.LNumber(i+1), item)
	}
	L.Push(tbl)
}
