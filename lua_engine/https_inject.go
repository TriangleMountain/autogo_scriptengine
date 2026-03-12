package lua_engine

import (
	"github.com/Dasongzi1366/AutoGo/https"
	lua "github.com/yuin/gopher-lua"
)

func injectHttpsMethods(engine *LuaEngine) {

	engine.RegisterMethod("http.get", "发送GET请求", func(url string, timeout int) (int, []byte) { return https.Get(url, timeout) }, true)
	engine.RegisterMethod("http.postMultipart", "发送Multipart POST请求", func(url, fileName string, fileData []byte, timeout int) (int, []byte) {
		return https.PostMultipart(url, fileName, fileData, timeout)
	}, true)

	registerHttpsLuaFunctions(engine)
}

func registerHttpsLuaFunctions(engine *LuaEngine) {
	state := engine.GetState()

	state.Register("http_get", func(L *lua.LState) int {
		url := L.CheckString(1)
		timeout := L.CheckInt(2)
		code, data := https.Get(url, timeout)
		L.Push(lua.LNumber(code))
		if data != nil {
			L.Push(lua.LString(data))
		} else {
			L.Push(lua.LNil)
		}
		return 2
	})

	state.Register("http_postMultipart", func(L *lua.LState) int {
		url := L.CheckString(1)
		fileName := L.CheckString(2)
		fileData := L.CheckString(3)
		timeout := L.CheckInt(4)
		code, responseData := https.PostMultipart(url, fileName, []byte(fileData), timeout)
		L.Push(lua.LNumber(code))
		if responseData != nil {
			L.Push(lua.LString(responseData))
		} else {
			L.Push(lua.LNil)
		}
		return 2
	})
}
