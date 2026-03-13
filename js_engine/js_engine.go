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

// GetJSEngine 获取默认引擎实例（使用默认配置，自动注入所有方法）
func GetJSEngine() *JSEngine {
	once.Do(func() {
		engine = &JSEngine{
			config: DefaultConfig(),
		}
		initRegistry()
		engine.init()
	})
	return engine
}

// GetEngine 获取默认引擎实例（GetJSEngine 的别名）
func GetEngine() *JSEngine {
	return GetJSEngine()
}

// NewJSEngine 创建新的引擎实例
// config: 引擎配置，传入 nil 使用默认配置
func NewJSEngine(config *EngineConfig) *JSEngine {
	cfg := DefaultConfig()
	if config != nil {
		cfg = *config
	}

	e := &JSEngine{
		config: cfg,
	}
	initRegistry()
	e.init()
	return e
}

// NewEngine 创建新的引擎实例（NewJSEngine 的别名）
func NewEngine(config *EngineConfig) *JSEngine {
	return NewJSEngine(config)
}

func (e *JSEngine) init() {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.vm = goja.New()

	e.registerCoreFunctions()
	if e.config.AutoInjectMethods {
		e.injectAllMethods()
	}
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
	// 新增模块
	injectConsoleMethods(e)
	injectDotocrMethods(e)
	injectHudMethods(e)
	injectImeMethods(e)
	injectPluginMethods(e)
	injectRhinoMethods(e)
	injectUiaccMethods(e)
	injectUtilsMethods(e)
	injectVdisplayMethods(e)
	injectYoloMethods(e)
	injectImguiMethods(e)
}

// InjectModule 注入指定模块的方法
// module: 模块名称，支持: app, device, motion, files, images, storages, system, http, media, opencv, ppocr, console, dotocr, hud, ime, plugin, rhino, uiacc, utils, vdisplay, yolo, imgui
func (e *JSEngine) InjectModule(module string) {
	e.mu.Lock()
	defer e.mu.Unlock()

	switch module {
	case "app":
		injectAppMethods(e)
	case "device":
		injectDeviceMethods(e)
	case "motion":
		injectMotionMethods(e)
	case "files":
		injectFilesMethods(e)
	case "images":
		injectImagesMethods(e)
	case "storages":
		injectStoragesMethods(e)
	case "system":
		injectSystemMethods(e)
	case "http":
		injectHttpsMethods(e)
	case "media":
		injectMediaMethods(e)
	case "opencv":
		injectOpenCVMethods(e)
	case "ppocr":
		injectPpocrMethods(e)
	case "console":
		injectConsoleMethods(e)
	case "dotocr":
		injectDotocrMethods(e)
	case "hud":
		injectHudMethods(e)
	case "ime":
		injectImeMethods(e)
	case "plugin":
		injectPluginMethods(e)
	case "rhino":
		injectRhinoMethods(e)
	case "uiacc":
		injectUiaccMethods(e)
	case "utils":
		injectUtilsMethods(e)
	case "vdisplay":
		injectVdisplayMethods(e)
	case "yolo":
		injectYoloMethods(e)
	case "imgui":
		injectImguiMethods(e)
	default:
		panic(fmt.Sprintf("unknown module: %s", module))
	}
}

// InjectModules 注入多个模块的方法
func (e *JSEngine) InjectModules(modules []string) {
	for _, module := range modules {
		e.InjectModule(module)
	}
}

// GetAvailableModules 获取所有可用模块列表
func (e *JSEngine) GetAvailableModules() []string {
	return []string{
		"app", "device", "motion", "files", "images", "storages",
		"system", "http", "media", "opencv", "ppocr",
		"console", "dotocr", "hud", "ime", "plugin",
		"rhino", "uiacc", "utils", "vdisplay", "yolo", "imgui",
	}
}

// InjectAllMethods 注入所有方法（公开方法，允许手动调用）
func (e *JSEngine) InjectAllMethods() {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.injectAllMethods()
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
