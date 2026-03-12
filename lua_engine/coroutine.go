package lua_engine

import (
	"fmt"
	"sync"
	"time"

	lua "github.com/yuin/gopher-lua"
)

type LuaCoroutine struct {
	id     int
	thread *lua.LState
	base   *lua.LState
	status string
	engine *LuaEngine
	mu     sync.Mutex
}

type CoroutineManager struct {
	coroutines map[int]*LuaCoroutine
	nextID     int
	mu         sync.RWMutex
}

var (
	coroutineManager *CoroutineManager
	coroutineOnce    sync.Once
)

func initCoroutineManager() {
	coroutineOnce.Do(func() {
		coroutineManager = &CoroutineManager{
			coroutines: make(map[int]*LuaCoroutine),
			nextID:     1,
		}
	})
}

func (e *LuaEngine) CreateCoroutine(script string) (*LuaCoroutine, error) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.state == nil {
		return nil, fmt.Errorf("Lua engine not initialized")
	}

	co, _ := e.state.NewThread()

	fn, err := e.state.LoadString(script)
	if err != nil {
		return nil, fmt.Errorf("Lua load error: %v", err)
	}

	initCoroutineManager()

	coroutineManager.mu.Lock()
	id := coroutineManager.nextID
	coroutineManager.nextID++
	coroutineManager.mu.Unlock()

	coroutine := &LuaCoroutine{
		id:     id,
		thread: co,
		base:   e.state,
		status: "ready",
		engine: e,
	}

	coroutine.thread.Push(fn)

	coroutineManager.mu.Lock()
	coroutineManager.coroutines[id] = coroutine
	coroutineManager.mu.Unlock()

	return coroutine, nil
}

func (e *LuaEngine) GetCoroutine(id int) (*LuaCoroutine, bool) {
	initCoroutineManager()
	coroutineManager.mu.RLock()
	defer coroutineManager.mu.RUnlock()

	coroutine, exists := coroutineManager.coroutines[id]
	return coroutine, exists
}

func (e *LuaEngine) ListCoroutines() []*LuaCoroutine {
	initCoroutineManager()
	coroutineManager.mu.RLock()
	defer coroutineManager.mu.RUnlock()

	coroutines := make([]*LuaCoroutine, 0, len(coroutineManager.coroutines))
	for _, coro := range coroutineManager.coroutines {
		coroutines = append(coroutines, coro)
	}
	return coroutines
}

func (e *LuaEngine) RemoveCoroutine(id int) bool {
	initCoroutineManager()
	coroutineManager.mu.Lock()
	defer coroutineManager.mu.Unlock()

	_, exists := coroutineManager.coroutines[id]
	if !exists {
		return false
	}

	delete(coroutineManager.coroutines, id)
	return true
}

func (c *LuaCoroutine) Resume(args ...interface{}) ([]interface{}, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.status == "dead" {
		return nil, fmt.Errorf("Coroutine is dead")
	}

	c.status = "running"

	nargs := len(args)
	for _, arg := range args {
		pushValue(c.thread, arg)
	}

	fn, ok := c.thread.Get(-nargs - 1).(*lua.LFunction)
	if !ok {
		c.status = "error"
		return nil, fmt.Errorf("Invalid coroutine function")
	}

	luaArgs := make([]lua.LValue, nargs)
	for i := 0; i < nargs; i++ {
		luaArgs[i] = c.thread.Get(-nargs + i)
	}

	state, err, luaValues := c.base.Resume(c.thread, fn, luaArgs...)
	if err != nil {
		c.status = "error"
		return nil, fmt.Errorf("Coroutine error: %v", err)
	}

	values := make([]interface{}, len(luaValues))
	for i, val := range luaValues {
		switch v := val.(type) {
		case lua.LString:
			values[i] = string(v)
		case lua.LNumber:
			values[i] = float64(v)
		case lua.LBool:
			values[i] = bool(v)
		case *lua.LUserData:
			values[i] = v.Value
		default:
			values[i] = nil
		}
	}

	if state == lua.ResumeYield {
		c.status = "suspended"
	} else {
		c.status = "dead"
	}

	return values, nil
}

func (c *LuaCoroutine) Yield(args ...interface{}) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.status != "running" {
		return fmt.Errorf("Coroutine is not running")
	}

	c.status = "suspended"
	return nil
}

func (c *LuaCoroutine) Status() string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.status
}

func (c *LuaCoroutine) ID() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.id
}

func (e *LuaEngine) RunCoroutineAsync(script string, callback func([]interface{}, error)) {
	go func() {
		coroutine, err := e.CreateCoroutine(script)
		if err != nil {
			callback(nil, err)
			return
		}

		results, err := coroutine.Resume()
		callback(results, err)
	}()
}

func (e *LuaEngine) RunCoroutineWithTimeout(script string, timeout time.Duration) ([]interface{}, error) {
	resultChan := make(chan []interface{}, 1)
	errChan := make(chan error, 1)

	e.RunCoroutineAsync(script, func(results []interface{}, err error) {
		if err != nil {
			errChan <- err
		} else {
			resultChan <- results
		}
	})

	select {
	case results := <-resultChan:
		return results, nil
	case err := <-errChan:
		return nil, err
	case <-time.After(timeout):
		return nil, fmt.Errorf("Coroutine execution timeout")
	}
}

func pushValue(L *lua.LState, value interface{}) {
	switch v := value.(type) {
	case string:
		L.Push(lua.LString(v))
	case int:
		L.Push(lua.LNumber(v))
	case int64:
		L.Push(lua.LNumber(v))
	case float64:
		L.Push(lua.LNumber(v))
	case bool:
		L.Push(lua.LBool(v))
	case nil:
		L.Push(lua.LNil)
	default:
		L.Push(lua.LNil)
	}
}
