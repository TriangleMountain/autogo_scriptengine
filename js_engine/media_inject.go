package js_engine

import (
	"github.com/Dasongzi1366/AutoGo/media"
	"github.com/dop251/goja"
)

func injectMediaMethods(engine *JSEngine) {
	vm := engine.GetVM()

	mediaObj := vm.NewObject()
	vm.Set("media", mediaObj)

	mediaObj.Set("scanFile", func(call goja.FunctionCall) goja.Value {
		path := call.Argument(0).String()
		media.ScanFile(path)
		return goja.Undefined()
	})

	engine.RegisterMethod("media.scanFile", "扫描路径path的媒体文件", func(path string) { media.ScanFile(path) }, true)
}
