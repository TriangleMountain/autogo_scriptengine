package lua_engine

import (
	"github.com/Dasongzi1366/AutoGo/storages"
	lua "github.com/yuin/gopher-lua"
)

func injectStoragesMethods(engine *LuaEngine) {

	engine.RegisterMethod("storages.get", "从指定表中获取键值", func(table, key string) string { return storages.Get(table, key) }, true)
	engine.RegisterMethod("storages.put", "写入键值对", func(table, key, value string) { storages.Put(table, key, value) }, true)
	engine.RegisterMethod("storages.remove", "删除指定键", func(table, key string) { storages.Remove(table, key) }, true)
	engine.RegisterMethod("storages.contains", "判断键是否存在", func(table, key string) bool { return storages.Contains(table, key) }, true)
	engine.RegisterMethod("storages.getAll", "获取所有键值对", func(table string) map[string]string { return storages.GetAll(table) }, true)
	engine.RegisterMethod("storages.clear", "清空指定表数据", func(table string) { storages.Clear(table) }, true)

	registerStoragesLuaFunctions(engine)
}

func registerStoragesLuaFunctions(engine *LuaEngine) {
	state := engine.GetState()

	state.Register("storages_get", func(L *lua.LState) int {
		table := L.CheckString(1)
		key := L.CheckString(2)
		result := storages.Get(table, key)
		L.Push(lua.LString(result))
		return 1
	})

	state.Register("storages_put", func(L *lua.LState) int {
		table := L.CheckString(1)
		key := L.CheckString(2)
		value := L.CheckString(3)
		storages.Put(table, key, value)
		return 0
	})

	state.Register("storages_remove", func(L *lua.LState) int {
		table := L.CheckString(1)
		key := L.CheckString(2)
		storages.Remove(table, key)
		return 0
	})

	state.Register("storages_contains", func(L *lua.LState) int {
		table := L.CheckString(1)
		key := L.CheckString(2)
		result := storages.Contains(table, key)
		L.Push(lua.LBool(result))
		return 1
	})

	state.Register("storages_getAll", func(L *lua.LState) int {
		table := L.CheckString(1)
		result := storages.GetAll(table)
		tbl := L.NewTable()
		for key, value := range result {
			L.SetTable(tbl, lua.LString(key), lua.LString(value))
		}
		L.Push(tbl)
		return 1
	})

	state.Register("storages_clear", func(L *lua.LState) int {
		table := L.CheckString(1)
		storages.Clear(table)
		return 0
	})
}
