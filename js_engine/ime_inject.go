package js_engine

import (
	"github.com/Dasongzi1366/AutoGo/ime"
	"github.com/dop251/goja"
)

func injectImeMethods(engine *JSEngine) {
	vm := engine.GetVM()

	imeObj := vm.NewObject()
	vm.Set("ime", imeObj)

	imeObj.Set("getClipText", func(call goja.FunctionCall) goja.Value {
		result := ime.GetClipText()
		return vm.ToValue(result)
	})

	imeObj.Set("setClipText", func(call goja.FunctionCall) goja.Value {
		text := call.Argument(0).String()
		result := ime.SetClipText(text)
		return vm.ToValue(result)
	})

	imeObj.Set("keyAction", func(call goja.FunctionCall) goja.Value {
		code := int(call.Argument(0).ToInteger())
		ime.KeyAction(code)
		return goja.Undefined()
	})

	imeObj.Set("inputText", func(call goja.FunctionCall) goja.Value {
		text := call.Argument(0).String()
		displayId := int(call.Argument(1).ToInteger())
		ime.InputText(text, displayId)
		return goja.Undefined()
	})

	imeObj.Set("getIMEList", func(call goja.FunctionCall) goja.Value {
		result := ime.GetIMEList()
		return vm.ToValue(result)
	})

	imeObj.Set("setCurrentIME", func(call goja.FunctionCall) goja.Value {
		packageName := call.Argument(0).String()
		ime.SetCurrentIME(packageName)
		return goja.Undefined()
	})

	// 注册方法到文档
	engine.RegisterMethod("ime.getClipText", "获取剪切板内容", ime.GetClipText, true)
	engine.RegisterMethod("ime.setClipText", "设置剪切板内容", ime.SetClipText, true)
	engine.RegisterMethod("ime.keyAction", "模拟按键", ime.KeyAction, true)
	engine.RegisterMethod("ime.inputText", "输入文本", ime.InputText, true)
	engine.RegisterMethod("ime.getIMEList", "获取输入法列表", ime.GetIMEList, true)
	engine.RegisterMethod("ime.setCurrentIME", "设置当前输入法", ime.SetCurrentIME, true)
}
