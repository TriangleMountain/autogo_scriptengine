package js_engine

import (
	"github.com/dop251/goja"
)

func injectOpenCVMethods(engine *JSEngine) {
	vm := engine.GetVM()

	opencvObj := vm.NewObject()
	vm.Set("opencv", opencvObj)

	opencvObj.Set("findImage", func(call goja.FunctionCall) goja.Value {
		_ = int(call.Argument(0).ToInteger())
		_ = int(call.Argument(1).ToInteger())
		_ = int(call.Argument(2).ToInteger())
		_ = int(call.Argument(3).ToInteger())
		_ = call.Argument(4).Export().(*[]byte)
		_ = call.Argument(5).ToBoolean()
		_ = float32(call.Argument(6).ToFloat())
		_ = float32(call.Argument(7).ToFloat())
		_ = 0
		if len(call.Arguments) > 8 {
			_ = int(call.Argument(8).ToInteger())
		}

		arr := vm.NewArray()
		arr.Set("0", -1)
		arr.Set("1", -1)
		return arr
	})

	engine.RegisterMethod("opencv.findImage", "在指定区域内查找匹配的图片模板", func(x1, y1, x2, y2 int, template *[]byte, isGray bool, scalingFactor, sim float32, displayId int) (int, int) {
		return -1, -1
	}, true)
}
