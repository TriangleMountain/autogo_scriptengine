package js_engine

import (
	"github.com/Dasongzi1366/AutoGo/rhino"
	"github.com/dop251/goja"
)

func injectRhinoMethods(engine *JSEngine) {
	vm := engine.GetVM()

	rhinoObj := vm.NewObject()
	vm.Set("rhino", rhinoObj)

	rhinoObj.Set("eval", func(call goja.FunctionCall) goja.Value {
		contextId := call.Argument(0).String()
		script := call.Argument(1).String()
		result := rhino.Eval(contextId, script)
		return vm.ToValue(result)
	})

	// 注册方法到文档
	engine.RegisterMethod("rhino.eval", "执行指定的JavaScript脚本并返回结果", rhino.Eval, true)
}
