package js_engine

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/dop251/goja"
)

var (
	engine *JSEngine
	once   sync.Once
)

func GetJSEngine() *JSEngine {
	once.Do(func() {
		engine = &JSEngine{}
		initRegistry()
		engine.init()
	})
	return engine
}

func GetEngine() *JSEngine {
	return GetJSEngine()
}

func (e *JSEngine) init() {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.vm = goja.New()

	e.registerCoreFunctions()
	e.injectAllMethods()
}

func (e *JSEngine) GetVM() *goja.Runtime {
	return e.vm
}

func (e *JSEngine) registerCoreFunctions() {
	vm := e.vm

	vm.Set("registerMethod", e.registerMethodJS)
	vm.Set("unregisterMethod", e.unregisterMethodJS)
	vm.Set("listMethods", e.listMethodsJS)
	vm.Set("overrideMethod", e.overrideMethodJS)
	vm.Set("restoreMethod", e.restoreMethodJS)
	vm.Set("sleep", e.sleepJS)

	consoleObj := vm.NewObject()
	consoleObj.Set("log", e.consoleLogJS)
	vm.Set("console", consoleObj)
}

func (e *JSEngine) consoleLogJS(call goja.FunctionCall) goja.Value {
	args := call.Arguments
	for _, arg := range args {
		fmt.Print(arg.Export(), " ")
	}
	fmt.Println()
	return goja.Undefined()
}

func (e *JSEngine) injectAllMethods() {
	injectAppMethods(e)
	injectDeviceMethods(e)
	injectMotionMethods(e)
	injectFilesMethods(e)
	injectImagesMethods(e)
	injectStoragesMethods(e)
	injectSystemMethods(e)
	injectHttpsMethods(e)
	injectMediaMethods(e)
	injectOpenCVMethods(e)
	injectPpocrMethods(e)
}

func (e *JSEngine) RegisterMethod(name, description string, goFunc interface{}, overridable bool) {
	RegisterMethod(name, description, goFunc, overridable)
}

func (e *JSEngine) ExecuteString(script string) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.vm == nil {
		return fmt.Errorf("JavaScript engine not initialized")
	}

	_, err := e.vm.RunString(script)
	return err
}

func (e *JSEngine) ExecuteFile(path string) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.vm == nil {
		return fmt.Errorf("JavaScript engine not initialized")
	}

	_, err := e.vm.RunString("load('" + path + "')")
	return err
}

func (e *JSEngine) Close() {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.vm = nil
}

func (e *JSEngine) GetRegistry() *MethodRegistry {
	return GetRegistry()
}

func ExecuteString(script string) error {
	if engine != nil {
		return engine.ExecuteString(script)
	}
	return fmt.Errorf("JavaScript engine not initialized")
}

func ExecuteFile(path string) error {
	if engine != nil {
		return engine.ExecuteFile(path)
	}
	return fmt.Errorf("JavaScript engine not initialized")
}

func Close() {
	if engine != nil {
		engine.Close()
	}
}

func (e *JSEngine) registerMethodJS(call goja.FunctionCall) goja.Value {
	name := call.Argument(0).String()
	description := call.Argument(1).String()
	overridable := call.Argument(2).ToBoolean()

	e.RegisterMethod(name, description, nil, overridable)
	return e.vm.ToValue(true)
}

func (e *JSEngine) unregisterMethodJS(call goja.FunctionCall) goja.Value {
	name := call.Argument(0).String()
	result := registry.RemoveMethod(name)
	return e.vm.ToValue(result)
}

func (e *JSEngine) listMethodsJS(call goja.FunctionCall) goja.Value {
	methods := registry.ListMethods()
	arr := e.vm.NewArray()
	for i, method := range methods {
		item := e.vm.NewObject()
		item.Set("name", method.Name)
		item.Set("description", method.Description)
		item.Set("overridable", method.Overridable)
		item.Set("overridden", method.Overridden)
		arr.Set(strconv.Itoa(i), item)
	}
	return arr
}

func (e *JSEngine) overrideMethodJS(call goja.FunctionCall) goja.Value {
	name := call.Argument(0).String()
	fn, ok := goja.AssertFunction(call.Argument(1))
	if !ok {
		panic("overrideMethod: second argument must be a function")
	}
	result := registry.OverrideMethod(name, fn)
	return e.vm.ToValue(result)
}

func (e *JSEngine) restoreMethodJS(call goja.FunctionCall) goja.Value {
	name := call.Argument(0).String()
	result := registry.RestoreMethod(name)
	return e.vm.ToValue(result)
}

func (e *JSEngine) sleepJS(call goja.FunctionCall) goja.Value {
	ms := int(call.Argument(0).ToInteger())
	time.Sleep(time.Duration(ms) * time.Millisecond)
	return goja.Undefined()
}
