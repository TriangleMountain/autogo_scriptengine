package js_engine

import (
	"github.com/Dasongzi1366/AutoGo/plugin"
	"github.com/dop251/goja"
)

func injectPluginMethods(engine *JSEngine) {
	vm := engine.GetVM()

	pluginObj := vm.NewObject()
	vm.Set("plugin", pluginObj)

	pluginObj.Set("newContext", func(call goja.FunctionCall) goja.Value {
		ctx := plugin.NewContext()
		return vm.ToValue(ctx)
	})

	pluginObj.Set("newAssetManager", func(call goja.FunctionCall) goja.Value {
		am := plugin.NewAssetManager()
		return vm.ToValue(am)
	})

	pluginObj.Set("loadApk", func(call goja.FunctionCall) goja.Value {
		apkPath := call.Argument(0).String()
		cl := plugin.LoadApk(apkPath)
		return vm.ToValue(cl)
	})

	pluginObj.Set("newInstance", func(call goja.FunctionCall) goja.Value {
		cl := call.Argument(0).Export().(*plugin.ClassLoader)
		className := call.Argument(1).String()
		var args []interface{}
		for i := 2; i < len(call.Arguments); i++ {
			args = append(args, call.Argument(i).Export())
		}
		inst := cl.NewInstance(className, args...)
		return vm.ToValue(inst)
	})

	pluginObj.Set("callString", func(call goja.FunctionCall) goja.Value {
		inst := call.Argument(0).Export().(*plugin.Instance)
		methodName := call.Argument(1).String()
		var args []interface{}
		for i := 2; i < len(call.Arguments); i++ {
			args = append(args, call.Argument(i).Export())
		}
		result, _ := inst.CallString(methodName, args...)
		return vm.ToValue(result)
	})

	pluginObj.Set("callInt", func(call goja.FunctionCall) goja.Value {
		inst := call.Argument(0).Export().(*plugin.Instance)
		methodName := call.Argument(1).String()
		var args []interface{}
		for i := 2; i < len(call.Arguments); i++ {
			args = append(args, call.Argument(i).Export())
		}
		result, _ := inst.CallInt(methodName, args...)
		return vm.ToValue(result)
	})

	pluginObj.Set("callLong", func(call goja.FunctionCall) goja.Value {
		inst := call.Argument(0).Export().(*plugin.Instance)
		methodName := call.Argument(1).String()
		var args []interface{}
		for i := 2; i < len(call.Arguments); i++ {
			args = append(args, call.Argument(i).Export())
		}
		result, _ := inst.CallLong(methodName, args...)
		return vm.ToValue(result)
	})

	pluginObj.Set("callFloat", func(call goja.FunctionCall) goja.Value {
		inst := call.Argument(0).Export().(*plugin.Instance)
		methodName := call.Argument(1).String()
		var args []interface{}
		for i := 2; i < len(call.Arguments); i++ {
			args = append(args, call.Argument(i).Export())
		}
		result, _ := inst.CallFloat(methodName, args...)
		return vm.ToValue(result)
	})

	pluginObj.Set("callDouble", func(call goja.FunctionCall) goja.Value {
		inst := call.Argument(0).Export().(*plugin.Instance)
		methodName := call.Argument(1).String()
		var args []interface{}
		for i := 2; i < len(call.Arguments); i++ {
			args = append(args, call.Argument(i).Export())
		}
		result, _ := inst.CallDouble(methodName, args...)
		return vm.ToValue(result)
	})

	pluginObj.Set("callBool", func(call goja.FunctionCall) goja.Value {
		inst := call.Argument(0).Export().(*plugin.Instance)
		methodName := call.Argument(1).String()
		var args []interface{}
		for i := 2; i < len(call.Arguments); i++ {
			args = append(args, call.Argument(i).Export())
		}
		result, _ := inst.CallBool(methodName, args...)
		return vm.ToValue(result)
	})

	pluginObj.Set("callVoid", func(call goja.FunctionCall) goja.Value {
		inst := call.Argument(0).Export().(*plugin.Instance)
		methodName := call.Argument(1).String()
		var args []interface{}
		for i := 2; i < len(call.Arguments); i++ {
			args = append(args, call.Argument(i).Export())
		}
		inst.CallVoid(methodName, args...)
		return goja.Undefined()
	})

	pluginObj.Set("releaseInstance", func(call goja.FunctionCall) goja.Value {
		inst := call.Argument(0).Export().(*plugin.Instance)
		inst.Release()
		return goja.Undefined()
	})

	pluginObj.Set("releaseClassLoader", func(call goja.FunctionCall) goja.Value {
		cl := call.Argument(0).Export().(*plugin.ClassLoader)
		cl.Release()
		return goja.Undefined()
	})

	// 注册方法到文档
	engine.RegisterMethod("plugin.newContext", "创建Context参数", plugin.NewContext, true)
	engine.RegisterMethod("plugin.newAssetManager", "创建AssetManager参数", plugin.NewAssetManager, true)
	engine.RegisterMethod("plugin.loadApk", "加载外部APK", plugin.LoadApk, true)
	engine.RegisterMethod("plugin.newInstance", "创建类实例", (*plugin.ClassLoader).NewInstance, true)
	engine.RegisterMethod("plugin.callString", "调用返回String的方法", (*plugin.Instance).CallString, true)
	engine.RegisterMethod("plugin.callInt", "调用返回int的方法", (*plugin.Instance).CallInt, true)
	engine.RegisterMethod("plugin.callLong", "调用返回long的方法", (*plugin.Instance).CallLong, true)
	engine.RegisterMethod("plugin.callFloat", "调用返回float的方法", (*plugin.Instance).CallFloat, true)
	engine.RegisterMethod("plugin.callDouble", "调用返回double的方法", (*plugin.Instance).CallDouble, true)
	engine.RegisterMethod("plugin.callBool", "调用返回boolean的方法", (*plugin.Instance).CallBool, true)
	engine.RegisterMethod("plugin.callVoid", "调用无返回值的方法", (*plugin.Instance).CallVoid, true)
	engine.RegisterMethod("plugin.releaseInstance", "释放实例", (*plugin.Instance).Release, true)
	engine.RegisterMethod("plugin.releaseClassLoader", "释放类加载器", (*plugin.ClassLoader).Release, true)
}
