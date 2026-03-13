package lua_engine

import (
	"github.com/Dasongzi1366/AutoGo/uiacc"
	lua "github.com/yuin/gopher-lua"
)

func injectUiaccMethods(engine *LuaEngine) {

	engine.RegisterMethod("uiacc.new", "创建一个新的Accessibility对象", uiacc.New, true)
	engine.RegisterMethod("uiacc.text", "设置选择器的text属性", (*uiacc.Uiacc).Text, true)
	engine.RegisterMethod("uiacc.textContains", "设置选择器的textContains属性", (*uiacc.Uiacc).TextContains, true)
	engine.RegisterMethod("uiacc.id", "设置选择器的id属性", (*uiacc.Uiacc).Id, true)
	engine.RegisterMethod("uiacc.className", "设置选择器的className属性", (*uiacc.Uiacc).ClassName, true)
	engine.RegisterMethod("uiacc.clickable", "设置选择器的clickable属性", (*uiacc.Uiacc).Clickable, true)
	engine.RegisterMethod("uiacc.scrollable", "设置选择器的scrollable属性", (*uiacc.Uiacc).Scrollable, true)
	engine.RegisterMethod("uiacc.waitFor", "等待控件出现并返回UiObject对象", (*uiacc.Uiacc).WaitFor, true)
	engine.RegisterMethod("uiacc.findOnce", "查找单个控件并返回UiObject对象", (*uiacc.Uiacc).FindOnce, true)
	engine.RegisterMethod("uiacc.find", "查找所有符合条件的控件并返回UiObject对象数组", (*uiacc.Uiacc).Find, true)
	engine.RegisterMethod("uiacc.release", "释放无障碍对象", (*uiacc.Uiacc).Release, true)
	engine.RegisterMethod("uiacc.objClick", "点击该控件", (*uiacc.UiObject).Click, true)
	engine.RegisterMethod("uiacc.clickCenter", "使用坐标点击该控件的中点", (*uiacc.UiObject).ClickCenter, true)
	engine.RegisterMethod("uiacc.setText", "设置输入框控件的文本内容", (*uiacc.UiObject).SetText, true)
	engine.RegisterMethod("uiacc.getText", "获取控件的文本内容", (*uiacc.UiObject).GetText, true)
	engine.RegisterMethod("uiacc.getDesc", "获取控件的描述内容", (*uiacc.UiObject).GetDesc, true)
	engine.RegisterMethod("uiacc.getBounds", "获取控件在屏幕上的范围", (*uiacc.UiObject).GetBounds, true)

	registerUiaccLuaFunctions(engine)
}

func registerUiaccLuaFunctions(engine *LuaEngine) {
	state := engine.GetState()

	state.Register("uiacc_new", func(L *lua.LState) int {
		displayId := L.CheckInt(1)
		u := uiacc.New(displayId)
		ud := L.NewUserData()
		ud.Value = u
		L.Push(ud)
		return 1
	})

	state.Register("uiacc_text", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		u := ud.Value.(*uiacc.Uiacc)
		value := L.CheckString(2)
		result := u.Text(value)
		ud.Value = result
		L.Push(ud)
		return 1
	})

	state.Register("uiacc_textContains", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		u := ud.Value.(*uiacc.Uiacc)
		value := L.CheckString(2)
		result := u.TextContains(value)
		ud.Value = result
		L.Push(ud)
		return 1
	})

	state.Register("uiacc_id", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		u := ud.Value.(*uiacc.Uiacc)
		value := L.CheckString(2)
		result := u.Id(value)
		ud.Value = result
		L.Push(ud)
		return 1
	})

	state.Register("uiacc_className", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		u := ud.Value.(*uiacc.Uiacc)
		value := L.CheckString(2)
		result := u.ClassName(value)
		ud.Value = result
		L.Push(ud)
		return 1
	})

	state.Register("uiacc_clickable", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		u := ud.Value.(*uiacc.Uiacc)
		value := L.CheckBool(2)
		result := u.Clickable(value)
		ud.Value = result
		L.Push(ud)
		return 1
	})

	state.Register("uiacc_scrollable", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		u := ud.Value.(*uiacc.Uiacc)
		value := L.CheckBool(2)
		result := u.Scrollable(value)
		ud.Value = result
		L.Push(ud)
		return 1
	})

	state.Register("uiacc_waitFor", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		u := ud.Value.(*uiacc.Uiacc)
		timeout := L.CheckInt(2)
		result := u.WaitFor(timeout)
		if result == nil {
			L.Push(lua.LNil)
			return 1
		}
		ud2 := L.NewUserData()
		ud2.Value = result
		L.Push(ud2)
		return 1
	})

	state.Register("uiacc_findOnce", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		u := ud.Value.(*uiacc.Uiacc)
		result := u.FindOnce()
		if result == nil {
			L.Push(lua.LNil)
			return 1
		}
		ud2 := L.NewUserData()
		ud2.Value = result
		L.Push(ud2)
		return 1
	})

	state.Register("uiacc_find", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		u := ud.Value.(*uiacc.Uiacc)
		results := u.Find()
		tbl := L.NewTable()
		for i, obj := range results {
			ud2 := L.NewUserData()
			ud2.Value = obj
			L.SetTable(tbl, lua.LNumber(i+1), ud2)
		}
		L.Push(tbl)
		return 1
	})

	state.Register("uiacc_release", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		u := ud.Value.(*uiacc.Uiacc)
		u.Release()
		return 0
	})

	// UiObject methods
	state.Register("uiobj_click", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		obj := ud.Value.(*uiacc.UiObject)
		result := obj.Click()
		L.Push(lua.LBool(result))
		return 1
	})

	state.Register("uiobj_clickCenter", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		obj := ud.Value.(*uiacc.UiObject)
		result := obj.ClickCenter()
		L.Push(lua.LBool(result))
		return 1
	})

	state.Register("uiobj_setText", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		obj := ud.Value.(*uiacc.UiObject)
		str := L.CheckString(2)
		result := obj.SetText(str)
		L.Push(lua.LBool(result))
		return 1
	})

	state.Register("uiobj_getText", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		obj := ud.Value.(*uiacc.UiObject)
		result := obj.GetText()
		L.Push(lua.LString(result))
		return 1
	})

	state.Register("uiobj_getDesc", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		obj := ud.Value.(*uiacc.UiObject)
		result := obj.GetDesc()
		L.Push(lua.LString(result))
		return 1
	})

	state.Register("uiobj_getBounds", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		obj := ud.Value.(*uiacc.UiObject)
		result := obj.GetBounds()
		tbl := L.NewTable()
		L.SetField(tbl, "Left", lua.LNumber(result.Left))
		L.SetField(tbl, "Top", lua.LNumber(result.Top))
		L.SetField(tbl, "Right", lua.LNumber(result.Right))
		L.SetField(tbl, "Bottom", lua.LNumber(result.Bottom))
		L.SetField(tbl, "CenterX", lua.LNumber(result.CenterX))
		L.SetField(tbl, "CenterY", lua.LNumber(result.CenterY))
		L.SetField(tbl, "Width", lua.LNumber(result.Width))
		L.SetField(tbl, "Height", lua.LNumber(result.Height))
		L.Push(tbl)
		return 1
	})
}
