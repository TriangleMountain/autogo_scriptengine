package lua_engine

import (
	"github.com/Dasongzi1366/AutoGo/files"
	lua "github.com/yuin/gopher-lua"
)

func injectFilesMethods(engine *LuaEngine) {

	engine.RegisterMethod("files.isDir", "返回路径path是否是文件夹", files.IsDir, true)
	engine.RegisterMethod("files.isFile", "返回路径path是否是文件", files.IsFile, true)
	engine.RegisterMethod("files.isEmptyDir", "返回文件夹path是否为空文件夹", files.IsEmptyDir, true)
	engine.RegisterMethod("files.create", "创建一个文件或文件夹", files.Create, true)
	engine.RegisterMethod("files.exists", "返回在路径path处的文件是否存在", files.Exists, true)
	engine.RegisterMethod("files.ensureDir", "确保路径path所在的文件夹存在", files.EnsureDir, true)
	engine.RegisterMethod("files.read", "读取文本文件path的所有内容并返回", files.Read, true)
	engine.RegisterMethod("files.readBytes", "读取文件path的所有内容并返回", files.ReadBytes, true)
	engine.RegisterMethod("files.write", "把text写入到文件path中", func(path, text string) { files.Write(path, text) }, true)
	engine.RegisterMethod("files.writeBytes", "把bytes写入到文件path中", func(path string, bytes []byte) { files.WriteBytes(path, bytes) }, true)
	engine.RegisterMethod("files.append", "把text追加到文件path的末尾", func(path, text string) { files.Append(path, text) }, true)
	engine.RegisterMethod("files.appendBytes", "把bytes追加到文件path的末尾", func(path string, bytes []byte) { files.AppendBytes(path, bytes) }, true)
	engine.RegisterMethod("files.copy", "复制文件", files.Copy, true)
	engine.RegisterMethod("files.move", "移动文件", files.Move, true)
	engine.RegisterMethod("files.rename", "重命名文件", files.Rename, true)
	engine.RegisterMethod("files.remove", "删除文件或文件夹", files.Remove, true)
	engine.RegisterMethod("files.getName", "返回文件的文件名", files.GetName, true)
	engine.RegisterMethod("files.getNameWithoutExtension", "返回不含拓展名的文件的文件名", files.GetNameWithoutExtension, true)
	engine.RegisterMethod("files.getExtension", "返回文件的拓展名", files.GetExtension, true)
	engine.RegisterMethod("files.path", "返回相对路径对应的绝对路径", files.Path, true)
	engine.RegisterMethod("files.listDir", "列出文件夹path下的所有文件和文件夹", files.ListDir, true)

	registerFilesLuaFunctions(engine)
}

func registerFilesLuaFunctions(engine *LuaEngine) {
	state := engine.GetState()

	state.Register("files_isDir", func(L *lua.LState) int {
		path := L.CheckString(1)
		result := files.IsDir(path)
		L.Push(lua.LBool(result))
		return 1
	})

	state.Register("files_isFile", func(L *lua.LState) int {
		path := L.CheckString(1)
		result := files.IsFile(path)
		L.Push(lua.LBool(result))
		return 1
	})

	state.Register("files_isEmptyDir", func(L *lua.LState) int {
		path := L.CheckString(1)
		result := files.IsEmptyDir(path)
		L.Push(lua.LBool(result))
		return 1
	})

	state.Register("files_create", func(L *lua.LState) int {
		path := L.CheckString(1)
		result := files.Create(path)
		L.Push(lua.LBool(result))
		return 1
	})

	state.Register("files_exists", func(L *lua.LState) int {
		path := L.CheckString(1)
		result := files.Exists(path)
		L.Push(lua.LBool(result))
		return 1
	})

	state.Register("files_ensureDir", func(L *lua.LState) int {
		path := L.CheckString(1)
		result := files.EnsureDir(path)
		L.Push(lua.LBool(result))
		return 1
	})

	state.Register("files_read", func(L *lua.LState) int {
		path := L.CheckString(1)
		result := files.Read(path)
		L.Push(lua.LString(result))
		return 1
	})

	state.Register("files_readBytes", func(L *lua.LState) int {
		path := L.CheckString(1)
		result := files.ReadBytes(path)
		if result != nil {
			L.Push(lua.LString(result))
		} else {
			L.Push(lua.LNil)
		}
		return 1
	})

	state.Register("files_write", func(L *lua.LState) int {
		path := L.CheckString(1)
		text := L.CheckString(2)
		files.Write(path, text)
		return 0
	})

	state.Register("files_writeBytes", func(L *lua.LState) int {
		path := L.CheckString(1)
		bytes := L.CheckString(2)
		files.WriteBytes(path, []byte(bytes))
		return 0
	})

	state.Register("files_append", func(L *lua.LState) int {
		path := L.CheckString(1)
		text := L.CheckString(2)
		files.Append(path, text)
		return 0
	})

	state.Register("files_appendBytes", func(L *lua.LState) int {
		path := L.CheckString(1)
		bytes := L.CheckString(2)
		files.AppendBytes(path, []byte(bytes))
		return 0
	})

	state.Register("files_copy", func(L *lua.LState) int {
		fromPath := L.CheckString(1)
		toPath := L.CheckString(2)
		result := files.Copy(fromPath, toPath)
		L.Push(lua.LBool(result))
		return 1
	})

	state.Register("files_move", func(L *lua.LState) int {
		fromPath := L.CheckString(1)
		toPath := L.CheckString(2)
		result := files.Move(fromPath, toPath)
		L.Push(lua.LBool(result))
		return 1
	})

	state.Register("files_rename", func(L *lua.LState) int {
		path := L.CheckString(1)
		newName := L.CheckString(2)
		result := files.Rename(path, newName)
		L.Push(lua.LBool(result))
		return 1
	})

	state.Register("files_remove", func(L *lua.LState) int {
		path := L.CheckString(1)
		result := files.Remove(path)
		L.Push(lua.LBool(result))
		return 1
	})

	state.Register("files_getName", func(L *lua.LState) int {
		path := L.CheckString(1)
		result := files.GetName(path)
		L.Push(lua.LString(result))
		return 1
	})

	state.Register("files_getNameWithoutExtension", func(L *lua.LState) int {
		path := L.CheckString(1)
		result := files.GetNameWithoutExtension(path)
		L.Push(lua.LString(result))
		return 1
	})

	state.Register("files_getExtension", func(L *lua.LState) int {
		path := L.CheckString(1)
		result := files.GetExtension(path)
		L.Push(lua.LString(result))
		return 1
	})

	state.Register("files_path", func(L *lua.LState) int {
		relativePath := L.CheckString(1)
		result := files.Path(relativePath)
		L.Push(lua.LString(result))
		return 1
	})

	state.Register("files_listDir", func(L *lua.LState) int {
		path := L.CheckString(1)
		result := files.ListDir(path)
		table := L.NewTable()
		for i, item := range result {
			L.SetTable(table, lua.LNumber(i+1), lua.LString(item))
		}
		L.Push(table)
		return 1
	})
}
