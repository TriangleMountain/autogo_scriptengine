package js_engine

import (
	"github.com/dop251/goja"
)

func injectPpocrMethods(engine *JSEngine) {
	vm := engine.GetVM()

	ppocrObj := vm.NewObject()
	vm.Set("ppocr", ppocrObj)

	ppocrObj.Set("ocr", func(call goja.FunctionCall) goja.Value {
		_ = int(call.Argument(0).ToInteger())
		_ = int(call.Argument(1).ToInteger())
		_ = int(call.Argument(2).ToInteger())
		_ = int(call.Argument(3).ToInteger())
		_ = call.Argument(4).String()
		_ = 0
		if len(call.Arguments) > 5 {
			_ = int(call.Argument(5).ToInteger())
		}
		arr := vm.NewArray()
		return arr
	})

	ppocrObj.Set("ocrFromImage", func(call goja.FunctionCall) goja.Value {
		_ = call.Argument(0).Export()
		_ = call.Argument(1).String()
		arr := vm.NewArray()
		return arr
	})

	ppocrObj.Set("ocrFromBase64", func(call goja.FunctionCall) goja.Value {
		_ = call.Argument(0).String()
		_ = call.Argument(1).String()
		arr := vm.NewArray()
		return arr
	})

	ppocrObj.Set("ocrFromPath", func(call goja.FunctionCall) goja.Value {
		_ = call.Argument(0).String()
		_ = call.Argument(1).String()
		arr := vm.NewArray()
		return arr
	})

	engine.RegisterMethod("ppocr.ocr", "识别屏幕文字", func(x1, y1, x2, y2 int, colorStr string, displayId int) []map[string]interface{} {
		return []map[string]interface{}{}
	}, true)
	engine.RegisterMethod("ppocr.ocrFromImage", "识别图片文字", func(img interface{}, colorStr string) []map[string]interface{} {
		return []map[string]interface{}{}
	}, true)
	engine.RegisterMethod("ppocr.ocrFromBase64", "识别Base64图片文字", func(b64, colorStr string) []map[string]interface{} {
		return []map[string]interface{}{}
	}, true)
	engine.RegisterMethod("ppocr.ocrFromPath", "识别文件图片文字", func(path, colorStr string) []map[string]interface{} {
		return []map[string]interface{}{}
	}, true)
}
