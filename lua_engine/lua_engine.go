package lua_engine

import (
	"fmt"
	"sync"
	"time"

	lua "github.com/yuin/gopher-lua"
)

var (
	engine *LuaEngine
	once   sync.Once
)

// GetLuaEngine 获取默认引擎实例（使用默认配置，自动注入所有方法）
func GetLuaEngine() *LuaEngine {
	once.Do(func() {
		engine = &LuaEngine{
			config: DefaultConfig(),
		}
		initRegistry()
		engine.init()
	})
	return engine
}

// GetEngine 获取默认引擎实例（GetLuaEngine 的别名）
func GetEngine() *LuaEngine {
	return GetLuaEngine()
}

// NewLuaEngine 创建新的引擎实例
// config: 引擎配置，传入 nil 使用默认配置
func NewLuaEngine(config *EngineConfig) *LuaEngine {
	cfg := DefaultConfig()
	if config != nil {
		cfg = *config
	}

	e := &LuaEngine{
		config: cfg,
	}
	initRegistry()
	e.init()
	return e
}

// NewEngine 创建新的引擎实例（NewLuaEngine 的别名）
func NewEngine(config *EngineConfig) *LuaEngine {
	return NewLuaEngine(config)
}

func (e *LuaEngine) init() {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.state = lua.NewState(lua.Options{
		SkipOpenLibs:        false,
		IncludeGoStackTrace: true,
	})

	e.registerCoreFunctions()
	if e.config.AutoInjectMethods {
		e.injectAllMethods()
	}
}

func (e *LuaEngine) GetState() *lua.LState {
	return e.state
}

func (e *LuaEngine) registerCoreFunctions() {
	state := e.state

	state.Register("registerMethod", e.registerMethodLua)
	state.Register("unregisterMethod", e.unregisterMethodLua)
	state.Register("listMethods", e.listMethodsLua)
	state.Register("overrideMethod", e.overrideMethodLua)
	state.Register("restoreMethod", e.restoreMethodLua)
	state.Register("createCoroutine", e.createCoroutineLua)
	state.Register("resumeCoroutine", e.resumeCoroutineLua)
	state.Register("yieldCoroutine", e.yieldCoroutineLua)
	state.Register("listCoroutines", e.listCoroutinesLua)
	state.Register("removeCoroutine", e.removeCoroutineLua)
	state.Register("sleep", e.sleepLua)
}

func (e *LuaEngine) injectAllMethods() {
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
func (e *LuaEngine) InjectModule(module string) {
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
func (e *LuaEngine) InjectModules(modules []string) {
	for _, module := range modules {
		e.InjectModule(module)
	}
}

// GetAvailableModules 获取所有可用模块列表
func (e *LuaEngine) GetAvailableModules() []string {
	return []string{
		"app", "device", "motion", "files", "images", "storages",
		"system", "http", "media", "opencv", "ppocr",
		"console", "dotocr", "hud", "ime", "plugin",
		"rhino", "uiacc", "utils", "vdisplay", "yolo", "imgui",
	}
}

// InjectAllMethods 注入所有方法（公开方法，允许手动调用）
func (e *LuaEngine) InjectAllMethods() {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.injectAllMethods()
}

func (e *LuaEngine) RegisterMethod(name, description string, goFunc interface{}, overridable bool) {
	RegisterMethod(name, description, goFunc, overridable)
}

func (e *LuaEngine) ExecuteString(script string) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.state == nil {
		return fmt.Errorf("Lua engine not initialized")
	}

	return e.state.DoString(script)
}

func (e *LuaEngine) ExecuteFile(path string) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.state == nil {
		return fmt.Errorf("Lua engine not initialized")
	}

	return e.state.DoFile(path)
}

func (e *LuaEngine) Close() {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.state != nil {
		e.state.Close()
		e.state = nil
	}
}

func (e *LuaEngine) GetRegistry() *MethodRegistry {
	return GetRegistry()
}

func ExecuteString(script string) error {
	if engine != nil {
		return engine.ExecuteString(script)
	}
	return fmt.Errorf("Lua engine not initialized")
}

func ExecuteFile(path string) error {
	if engine != nil {
		return engine.ExecuteFile(path)
	}
	return fmt.Errorf("Lua engine not initialized")
}

func Close() {
	if engine != nil {
		engine.Close()
	}
}

func (e *LuaEngine) registerMethodLua(L *lua.LState) int {
	name := L.CheckString(1)
	description := L.CheckString(2)
	overridable := L.CheckBool(3)

	e.RegisterMethod(name, description, nil, overridable)
	L.Push(lua.LBool(true))
	return 1
}

func (e *LuaEngine) unregisterMethodLua(L *lua.LState) int {
	name := L.CheckString(1)
	result := registry.RemoveMethod(name)
	L.Push(lua.LBool(result))
	return 1
}

func (e *LuaEngine) listMethodsLua(L *lua.LState) int {
	methods := registry.ListMethods()
	tbl := L.NewTable()
	for i, method := range methods {
		item := L.NewTable()
		L.SetField(item, "name", lua.LString(method.Name))
		L.SetField(item, "description", lua.LString(method.Description))
		L.SetField(item, "overridable", lua.LBool(method.Overridable))
		L.SetField(item, "overridden", lua.LBool(method.Overridden))
		L.SetTable(tbl, lua.LNumber(i+1), item)
	}
	L.Push(tbl)
	return 1
}

func (e *LuaEngine) overrideMethodLua(L *lua.LState) int {
	name := L.CheckString(1)
	fn := L.CheckFunction(2)
	result := registry.OverrideMethod(name, fn)
	L.Push(lua.LBool(result))
	return 1
}

func (e *LuaEngine) restoreMethodLua(L *lua.LState) int {
	name := L.CheckString(1)
	result := registry.RestoreMethod(name)
	L.Push(lua.LBool(result))
	return 1
}

func (e *LuaEngine) createCoroutineLua(L *lua.LState) int {
	script := L.CheckString(1)
	coroutine, err := e.CreateCoroutine(script)
	if err != nil {
		L.RaiseError(err.Error())
		return 0
	}
	L.Push(lua.LNumber(coroutine.ID()))
	return 1
}

func (e *LuaEngine) resumeCoroutineLua(L *lua.LState) int {
	id := L.CheckInt(1)
	coroutine, exists := e.GetCoroutine(id)
	if !exists {
		L.RaiseError("Coroutine not found")
		return 0
	}

	results, err := coroutine.Resume()
	if err != nil {
		L.RaiseError(err.Error())
		return 0
	}

	tbl := L.NewTable()
	for i, result := range results {
		switch v := result.(type) {
		case string:
			L.SetTable(tbl, lua.LNumber(i+1), lua.LString(v))
		case int:
			L.SetTable(tbl, lua.LNumber(i+1), lua.LNumber(v))
		case float64:
			L.SetTable(tbl, lua.LNumber(i+1), lua.LNumber(v))
		case bool:
			L.SetTable(tbl, lua.LNumber(i+1), lua.LBool(v))
		default:
			L.SetTable(tbl, lua.LNumber(i+1), lua.LNil)
		}
	}

	L.Push(tbl)
	L.Push(lua.LString(coroutine.Status()))
	return 2
}

func (e *LuaEngine) yieldCoroutineLua(L *lua.LState) int {
	id := L.CheckInt(1)
	coroutine, exists := e.GetCoroutine(id)
	if !exists {
		L.RaiseError("Coroutine not found")
		return 0
	}

	err := coroutine.Yield()
	if err != nil {
		L.RaiseError(err.Error())
		return 0
	}

	L.Push(lua.LBool(true))
	return 1
}

func (e *LuaEngine) listCoroutinesLua(L *lua.LState) int {
	coroutines := e.ListCoroutines()
	tbl := L.NewTable()
	for i, coro := range coroutines {
		item := L.NewTable()
		L.SetField(item, "id", lua.LNumber(coro.ID()))
		L.SetField(item, "status", lua.LString(coro.Status()))
		L.SetTable(tbl, lua.LNumber(i+1), item)
	}
	L.Push(tbl)
	return 1
}

func (e *LuaEngine) removeCoroutineLua(L *lua.LState) int {
	id := L.CheckInt(1)
	result := e.RemoveCoroutine(id)
	L.Push(lua.LBool(result))
	return 1
}

func (e *LuaEngine) sleepLua(L *lua.LState) int {
	ms := L.CheckInt(1)
	time.Sleep(time.Duration(ms) * time.Millisecond)
	return 0
}
