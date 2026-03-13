package lua_engine

import (
	"github.com/yuin/gopher-lua"
	"github.com/Dasongzi1366/AutoGo/imgui"
)

// injectImguiMethods 注入 imgui GUI 库方法
func injectImguiMethods(e *LuaEngine) {
	L := e.GetState()

	// 创建 imgui 表
	imguiTable := L.NewTable()

	// ==================== 基础函数 ====================

	// 初始化和生命周期
	L.SetField(imguiTable, "init", L.NewFunction(func(L *lua.LState) int {
		err := imgui.Init()
		if err != nil {
			L.RaiseError(err.Error())
			return 0
		}
		return 0
	}))

	L.SetField(imguiTable, "run", L.NewFunction(func(L *lua.LState) int {
		fn := L.CheckFunction(1)
		imgui.Run(func() {
			L.Push(fn)
			L.PCall(0, 0, nil)
		})
		return 0
	}))

	L.SetField(imguiTable, "close", L.NewFunction(func(L *lua.LState) int {
		imgui.Close()
		return 0
	}))

	// 版本信息
	L.SetField(imguiTable, "version", lua.LString(imgui.Version()))

	// ==================== 窗口函数 ====================

	L.SetField(imguiTable, "begin", L.NewFunction(func(L *lua.LState) int {
		name := L.CheckString(1)
		var p_open *bool
		var flags imgui.WindowFlags

		if L.GetTop() >= 2 {
			open := L.CheckBool(2)
			p_open = &open
		}
		if L.GetTop() >= 3 {
			flags = imgui.WindowFlags(L.CheckInt(3))
		}

		result := imgui.BeginV(name, p_open, flags)
		if p_open != nil {
			L.Push(lua.LBool(result))
			L.Push(lua.LBool(*p_open))
			return 2
		}
		L.Push(lua.LBool(result))
		return 1
	}))

	L.SetField(imguiTable, "end", L.NewFunction(func(L *lua.LState) int {
		imgui.End()
		return 0
	}))

	// 子窗口
	L.SetField(imguiTable, "beginChild", L.NewFunction(func(L *lua.LState) int {
		str_id := L.CheckString(1)
		var size imgui.Vec2
		var child_flags imgui.ChildFlags
		var window_flags imgui.WindowFlags

		if L.GetTop() >= 2 {
			t := L.CheckTable(2)
			size.X = float32(L.GetField(t, "x").(lua.LNumber))
			size.Y = float32(L.GetField(t, "y").(lua.LNumber))
		}
		if L.GetTop() >= 3 {
			child_flags = imgui.ChildFlags(L.CheckInt(3))
		}
		if L.GetTop() >= 4 {
			window_flags = imgui.WindowFlags(L.CheckInt(4))
		}

		L.Push(lua.LBool(imgui.BeginChildStrV(str_id, size, child_flags, window_flags)))
		return 1
	}))

	L.SetField(imguiTable, "endChild", L.NewFunction(func(L *lua.LState) int {
		imgui.EndChild()
		return 0
	}))

	// ==================== 基础控件 ====================

	// 按钮
	L.SetField(imguiTable, "button", L.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		var size imgui.Vec2
		if L.GetTop() >= 2 {
			t := L.CheckTable(2)
			size.X = float32(L.GetField(t, "x").(lua.LNumber))
			size.Y = float32(L.GetField(t, "y").(lua.LNumber))
		}
		L.Push(lua.LBool(imgui.ButtonV(label, size)))
		return 1
	}))

	L.SetField(imguiTable, "smallButton", L.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		L.Push(lua.LBool(imgui.SmallButton(label)))
		return 1
	}))

	L.SetField(imguiTable, "invisibleButton", L.NewFunction(func(L *lua.LState) int {
		str_id := L.CheckString(1)
		var size imgui.Vec2
		var flags imgui.ButtonFlags

		if L.GetTop() >= 2 {
			t := L.CheckTable(2)
			size.X = float32(L.GetField(t, "x").(lua.LNumber))
			size.Y = float32(L.GetField(t, "y").(lua.LNumber))
		}
		if L.GetTop() >= 3 {
			flags = imgui.ButtonFlags(L.CheckInt(3))
		}
		L.Push(lua.LBool(imgui.InvisibleButtonV(str_id, size, flags)))
		return 1
	}))

	// 文本
	L.SetField(imguiTable, "text", L.NewFunction(func(L *lua.LState) int {
		text := L.CheckString(1)
		imgui.Text(text)
		return 0
	}))

	L.SetField(imguiTable, "textColored", L.NewFunction(func(L *lua.LState) int {
		t := L.CheckTable(1)
		color := imgui.Vec4{
			X: float32(L.GetField(t, "x").(lua.LNumber)),
			Y: float32(L.GetField(t, "y").(lua.LNumber)),
			Z: float32(L.GetField(t, "z").(lua.LNumber)),
			W: float32(L.GetField(t, "w").(lua.LNumber)),
		}
		text := L.CheckString(2)
		imgui.TextColored(color, text)
		return 0
	}))

	L.SetField(imguiTable, "textDisabled", L.NewFunction(func(L *lua.LState) int {
		text := L.CheckString(1)
		imgui.TextDisabled(text)
		return 0
	}))

	L.SetField(imguiTable, "textWrapped", L.NewFunction(func(L *lua.LState) int {
		text := L.CheckString(1)
		imgui.TextWrapped(text)
		return 0
	}))

	L.SetField(imguiTable, "bulletText", L.NewFunction(func(L *lua.LState) int {
		text := L.CheckString(1)
		imgui.BulletText(text)
		return 0
	}))

	L.SetField(imguiTable, "bullet", L.NewFunction(func(L *lua.LState) int {
		imgui.Bullet()
		return 0
	}))

	// 复选框
	L.SetField(imguiTable, "checkbox", L.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		v := L.CheckBool(2)
		result := imgui.Checkbox(label, &v)
		L.Push(lua.LBool(result))
		L.Push(lua.LBool(v))
		return 2
	}))

	// 单选按钮
	L.SetField(imguiTable, "radioButton", L.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		active := L.CheckBool(2)
		L.Push(lua.LBool(imgui.RadioButtonBool(label, active)))
		return 1
	}))

	L.SetField(imguiTable, "radioButtonInt", L.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		v := int32(L.CheckInt(2))
		v_button := int32(L.CheckInt(3))
		result := imgui.RadioButtonIntPtr(label, &v, v_button)
		L.Push(lua.LBool(result))
		L.Push(lua.LNumber(v))
		return 2
	}))

	// 进度条
	L.SetField(imguiTable, "progressBar", L.NewFunction(func(L *lua.LState) int {
		fraction := float32(L.CheckNumber(1))
		var size imgui.Vec2
		var overlay string

		if L.GetTop() >= 2 {
			t := L.CheckTable(2)
			size.X = float32(L.GetField(t, "x").(lua.LNumber))
			size.Y = float32(L.GetField(t, "y").(lua.LNumber))
		}
		if L.GetTop() >= 3 {
			overlay = L.CheckString(3)
		}
		imgui.ProgressBarV(fraction, size, overlay)
		return 0
	}))

	// 可选择项
	L.SetField(imguiTable, "selectable", L.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		var selected bool
		var flags imgui.SelectableFlags
		var size imgui.Vec2

		if L.GetTop() >= 2 {
			selected = L.CheckBool(2)
		}
		if L.GetTop() >= 3 {
			flags = imgui.SelectableFlags(L.CheckInt(3))
		}
		if L.GetTop() >= 4 {
			t := L.CheckTable(4)
			size.X = float32(L.GetField(t, "x").(lua.LNumber))
			size.Y = float32(L.GetField(t, "y").(lua.LNumber))
		}
		L.Push(lua.LBool(imgui.SelectableBoolV(label, selected, flags, size)))
		return 1
	}))

	// 组合框
	L.SetField(imguiTable, "beginCombo", L.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		preview_value := L.CheckString(2)
		var flags imgui.ComboFlags
		if L.GetTop() >= 3 {
			flags = imgui.ComboFlags(L.CheckInt(3))
		}
		L.Push(lua.LBool(imgui.BeginComboV(label, preview_value, flags)))
		return 1
	}))

	L.SetField(imguiTable, "endCombo", L.NewFunction(func(L *lua.LState) int {
		imgui.EndCombo()
		return 0
	}))

	// 列表框
	L.SetField(imguiTable, "beginListBox", L.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		var size imgui.Vec2
		if L.GetTop() >= 2 {
			t := L.CheckTable(2)
			size.X = float32(L.GetField(t, "x").(lua.LNumber))
			size.Y = float32(L.GetField(t, "y").(lua.LNumber))
		}
		L.Push(lua.LBool(imgui.BeginListBoxV(label, size)))
		return 1
	}))

	L.SetField(imguiTable, "endListBox", L.NewFunction(func(L *lua.LState) int {
		imgui.EndListBox()
		return 0
	}))

	// ==================== 输入控件 ====================

	// 输入文本
	L.SetField(imguiTable, "inputText", L.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		buf := L.CheckString(2)
		var flags imgui.InputTextFlags
		if L.GetTop() >= 3 {
			flags = imgui.InputTextFlags(L.CheckInt(3))
		}
		result := imgui.InputTextWithHint(label, "", &buf, flags, nil)
		L.Push(lua.LBool(result))
		L.Push(lua.LString(buf))
		return 2
	}))

	L.SetField(imguiTable, "inputTextWithHint", L.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		hint := L.CheckString(2)
		buf := L.CheckString(3)
		var flags imgui.InputTextFlags
		if L.GetTop() >= 4 {
			flags = imgui.InputTextFlags(L.CheckInt(4))
		}
		result := imgui.InputTextWithHint(label, hint, &buf, flags, nil)
		L.Push(lua.LBool(result))
		L.Push(lua.LString(buf))
		return 2
	}))

	L.SetField(imguiTable, "inputTextMultiline", L.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		buf := L.CheckString(2)
		var size imgui.Vec2
		var flags imgui.InputTextFlags

		if L.GetTop() >= 3 {
			t := L.CheckTable(3)
			size.X = float32(L.GetField(t, "x").(lua.LNumber))
			size.Y = float32(L.GetField(t, "y").(lua.LNumber))
		}
		if L.GetTop() >= 4 {
			flags = imgui.InputTextFlags(L.CheckInt(4))
		}
		result := imgui.InputTextMultiline(label, &buf, size, flags, nil)
		L.Push(lua.LBool(result))
		L.Push(lua.LString(buf))
		return 2
	}))

	// 输入数字
	L.SetField(imguiTable, "inputInt", L.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		v := int32(L.CheckInt(2))
		var step, step_fast int32
		var flags imgui.InputTextFlags

		if L.GetTop() >= 3 {
			step = int32(L.CheckInt(3))
		}
		if L.GetTop() >= 4 {
			step_fast = int32(L.CheckInt(4))
		}
		if L.GetTop() >= 5 {
			flags = imgui.InputTextFlags(L.CheckInt(5))
		}
		result := imgui.InputIntV(label, &v, step, step_fast, flags)
		L.Push(lua.LBool(result))
		L.Push(lua.LNumber(v))
		return 2
	}))

	L.SetField(imguiTable, "inputFloat", L.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		v := float32(L.CheckNumber(2))
		var step, step_fast float32
		var format string
		var flags imgui.InputTextFlags

		if L.GetTop() >= 3 {
			step = float32(L.CheckNumber(3))
		}
		if L.GetTop() >= 4 {
			step_fast = float32(L.CheckNumber(4))
		}
		if L.GetTop() >= 5 {
			format = L.CheckString(5)
		}
		if L.GetTop() >= 6 {
			flags = imgui.InputTextFlags(L.CheckInt(6))
		}
		result := imgui.InputFloatV(label, &v, step, step_fast, format, flags)
		L.Push(lua.LBool(result))
		L.Push(lua.LNumber(v))
		return 2
	}))

	L.SetField(imguiTable, "inputDouble", L.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		v := float64(L.CheckNumber(2))
		var step, step_fast float64
		var format string
		var flags imgui.InputTextFlags

		if L.GetTop() >= 3 {
			step = float64(L.CheckNumber(3))
		}
		if L.GetTop() >= 4 {
			step_fast = float64(L.CheckNumber(4))
		}
		if L.GetTop() >= 5 {
			format = L.CheckString(5)
		}
		if L.GetTop() >= 6 {
			flags = imgui.InputTextFlags(L.CheckInt(6))
		}
		result := imgui.InputDoubleV(label, &v, step, step_fast, format, flags)
		L.Push(lua.LBool(result))
		L.Push(lua.LNumber(v))
		return 2
	}))

	// ==================== 滑块控件 ====================

	L.SetField(imguiTable, "sliderInt", L.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		v := int32(L.CheckInt(2))
		v_min := int32(L.CheckInt(3))
		v_max := int32(L.CheckInt(4))
		var format string
		var flags imgui.SliderFlags

		if L.GetTop() >= 5 {
			format = L.CheckString(5)
		}
		if L.GetTop() >= 6 {
			flags = imgui.SliderFlags(L.CheckInt(6))
		}
		result := imgui.SliderIntV(label, &v, v_min, v_max, format, flags)
		L.Push(lua.LBool(result))
		L.Push(lua.LNumber(v))
		return 2
	}))

	L.SetField(imguiTable, "sliderFloat", L.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		v := float32(L.CheckNumber(2))
		v_min := float32(L.CheckNumber(3))
		v_max := float32(L.CheckNumber(4))
		var format string
		var flags imgui.SliderFlags

		if L.GetTop() >= 5 {
			format = L.CheckString(5)
		}
		if L.GetTop() >= 6 {
			flags = imgui.SliderFlags(L.CheckInt(6))
		}
		result := imgui.SliderFloatV(label, &v, v_min, v_max, format, flags)
		L.Push(lua.LBool(result))
		L.Push(lua.LNumber(v))
		return 2
	}))

	L.SetField(imguiTable, "sliderAngle", L.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		v_rad := float32(L.CheckNumber(2))
		var v_degrees_min, v_degrees_max float32
		var format string
		var flags imgui.SliderFlags

		if L.GetTop() >= 3 {
			v_degrees_min = float32(L.CheckNumber(3))
		}
		if L.GetTop() >= 4 {
			v_degrees_max = float32(L.CheckNumber(4))
		}
		if L.GetTop() >= 5 {
			format = L.CheckString(5)
		}
		if L.GetTop() >= 6 {
			flags = imgui.SliderFlags(L.CheckInt(6))
		}
		result := imgui.SliderAngleV(label, &v_rad, v_degrees_min, v_degrees_max, format, flags)
		L.Push(lua.LBool(result))
		L.Push(lua.LNumber(v_rad))
		return 2
	}))

	// ==================== 拖拽控件 ====================

	L.SetField(imguiTable, "dragInt", L.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		v := int32(L.CheckInt(2))
		var v_speed float32
		var v_min, v_max int32
		var format string
		var flags imgui.SliderFlags

		if L.GetTop() >= 3 {
			v_speed = float32(L.CheckNumber(3))
		}
		if L.GetTop() >= 4 {
			v_min = int32(L.CheckInt(4))
		}
		if L.GetTop() >= 5 {
			v_max = int32(L.CheckInt(5))
		}
		if L.GetTop() >= 6 {
			format = L.CheckString(6)
		}
		if L.GetTop() >= 7 {
			flags = imgui.SliderFlags(L.CheckInt(7))
		}
		result := imgui.DragIntV(label, &v, v_speed, v_min, v_max, format, flags)
		L.Push(lua.LBool(result))
		L.Push(lua.LNumber(v))
		return 2
	}))

	L.SetField(imguiTable, "dragFloat", L.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		v := float32(L.CheckNumber(2))
		var v_speed, v_min, v_max float32
		var format string
		var flags imgui.SliderFlags

		if L.GetTop() >= 3 {
			v_speed = float32(L.CheckNumber(3))
		}
		if L.GetTop() >= 4 {
			v_min = float32(L.CheckNumber(4))
		}
		if L.GetTop() >= 5 {
			v_max = float32(L.CheckNumber(5))
		}
		if L.GetTop() >= 6 {
			format = L.CheckString(6)
		}
		if L.GetTop() >= 7 {
			flags = imgui.SliderFlags(L.CheckInt(7))
		}
		result := imgui.DragFloatV(label, &v, v_speed, v_min, v_max, format, flags)
		L.Push(lua.LBool(result))
		L.Push(lua.LNumber(v))
		return 2
	}))

	// ==================== 颜色控件 ====================

	L.SetField(imguiTable, "colorEdit3", L.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		t := L.CheckTable(2)
		colArr := [3]float32{
			float32(L.GetField(t, "1").(lua.LNumber)),
			float32(L.GetField(t, "2").(lua.LNumber)),
			float32(L.GetField(t, "3").(lua.LNumber)),
		}
		var flags imgui.ColorEditFlags
		if L.GetTop() >= 3 {
			flags = imgui.ColorEditFlags(L.CheckInt(3))
		}
		result := imgui.ColorEdit3V(label, &colArr, flags)
		L.Push(lua.LBool(result))
		retTable := L.NewTable()
		L.SetField(retTable, "1", lua.LNumber(colArr[0]))
		L.SetField(retTable, "2", lua.LNumber(colArr[1]))
		L.SetField(retTable, "3", lua.LNumber(colArr[2]))
		L.Push(retTable)
		return 2
	}))

	L.SetField(imguiTable, "colorEdit4", L.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		t := L.CheckTable(2)
		colArr := [4]float32{
			float32(L.GetField(t, "1").(lua.LNumber)),
			float32(L.GetField(t, "2").(lua.LNumber)),
			float32(L.GetField(t, "3").(lua.LNumber)),
			float32(L.GetField(t, "4").(lua.LNumber)),
		}
		var flags imgui.ColorEditFlags
		if L.GetTop() >= 3 {
			flags = imgui.ColorEditFlags(L.CheckInt(3))
		}
		result := imgui.ColorEdit4V(label, &colArr, flags)
		L.Push(lua.LBool(result))
		retTable := L.NewTable()
		L.SetField(retTable, "1", lua.LNumber(colArr[0]))
		L.SetField(retTable, "2", lua.LNumber(colArr[1]))
		L.SetField(retTable, "3", lua.LNumber(colArr[2]))
		L.SetField(retTable, "4", lua.LNumber(colArr[3]))
		L.Push(retTable)
		return 2
	}))

	L.SetField(imguiTable, "colorButton", L.NewFunction(func(L *lua.LState) int {
		desc_id := L.CheckString(1)
		t := L.CheckTable(2)
		color := imgui.Vec4{
			X: float32(L.GetField(t, "x").(lua.LNumber)),
			Y: float32(L.GetField(t, "y").(lua.LNumber)),
			Z: float32(L.GetField(t, "z").(lua.LNumber)),
			W: float32(L.GetField(t, "w").(lua.LNumber)),
		}
		var flags imgui.ColorEditFlags
		var size imgui.Vec2

		if L.GetTop() >= 3 {
			flags = imgui.ColorEditFlags(L.CheckInt(3))
		}
		if L.GetTop() >= 4 {
			sz := L.CheckTable(4)
			size.X = float32(L.GetField(sz, "x").(lua.LNumber))
			size.Y = float32(L.GetField(sz, "y").(lua.LNumber))
		}
		L.Push(lua.LBool(imgui.ColorButtonV(desc_id, color, flags, size)))
		return 1
	}))

	// ==================== 树形控件 ====================

	L.SetField(imguiTable, "treeNode", L.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		var flags imgui.TreeNodeFlags
		if L.GetTop() >= 2 {
			flags = imgui.TreeNodeFlags(L.CheckInt(2))
		}
		L.Push(lua.LBool(imgui.TreeNodeExStrV(label, flags)))
		return 1
	}))

	L.SetField(imguiTable, "treePush", L.NewFunction(func(L *lua.LState) int {
		var str_id string
		if L.GetTop() >= 1 {
			str_id = L.CheckString(1)
		}
		imgui.TreePushStr(str_id)
		return 0
	}))

	L.SetField(imguiTable, "treePop", L.NewFunction(func(L *lua.LState) int {
		imgui.TreePop()
		return 0
	}))

	L.SetField(imguiTable, "collapsingHeader", L.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		var flags imgui.TreeNodeFlags
		if L.GetTop() >= 2 {
			flags = imgui.TreeNodeFlags(L.CheckInt(2))
		}
		L.Push(lua.LBool(imgui.CollapsingHeaderTreeNodeFlagsV(label, flags)))
		return 1
	}))

	L.SetField(imguiTable, "setNextItemOpen", L.NewFunction(func(L *lua.LState) int {
		is_open := L.CheckBool(1)
		var cond imgui.Cond
		if L.GetTop() >= 2 {
			cond = imgui.Cond(L.CheckInt(2))
		}
		imgui.SetNextItemOpenV(is_open, cond)
		return 0
	}))

	// ==================== 布局函数 ====================

	L.SetField(imguiTable, "separator", L.NewFunction(func(L *lua.LState) int {
		imgui.Separator()
		return 0
	}))

	L.SetField(imguiTable, "sameLine", L.NewFunction(func(L *lua.LState) int {
		var offset_from_start_x float32
		var spacing float32 = -1

		if L.GetTop() >= 1 {
			offset_from_start_x = float32(L.CheckNumber(1))
		}
		if L.GetTop() >= 2 {
			spacing = float32(L.CheckNumber(2))
		}
		imgui.SameLineV(offset_from_start_x, spacing)
		return 0
	}))

	L.SetField(imguiTable, "newLine", L.NewFunction(func(L *lua.LState) int {
		imgui.NewLine()
		return 0
	}))

	L.SetField(imguiTable, "spacing", L.NewFunction(func(L *lua.LState) int {
		imgui.Spacing()
		return 0
	}))

	L.SetField(imguiTable, "dummy", L.NewFunction(func(L *lua.LState) int {
		t := L.CheckTable(1)
		imgui.Dummy(imgui.Vec2{
			X: float32(L.GetField(t, "x").(lua.LNumber)),
			Y: float32(L.GetField(t, "y").(lua.LNumber)),
		})
		return 0
	}))

	L.SetField(imguiTable, "indent", L.NewFunction(func(L *lua.LState) int {
		if L.GetTop() >= 1 {
			indent_w := float32(L.CheckNumber(1))
			imgui.IndentV(indent_w)
		} else {
			imgui.Indent()
		}
		return 0
	}))

	L.SetField(imguiTable, "unindent", L.NewFunction(func(L *lua.LState) int {
		if L.GetTop() >= 1 {
			indent_w := float32(L.CheckNumber(1))
			imgui.UnindentV(indent_w)
		} else {
			imgui.Unindent()
		}
		return 0
	}))

	L.SetField(imguiTable, "beginGroup", L.NewFunction(func(L *lua.LState) int {
		imgui.BeginGroup()
		return 0
	}))

	L.SetField(imguiTable, "endGroup", L.NewFunction(func(L *lua.LState) int {
		imgui.EndGroup()
		return 0
	}))

	L.SetField(imguiTable, "getCursorPos", L.NewFunction(func(L *lua.LState) int {
		pos := imgui.CursorPos()
		t := L.NewTable()
		L.SetField(t, "x", lua.LNumber(pos.X))
		L.SetField(t, "y", lua.LNumber(pos.Y))
		L.Push(t)
		return 1
	}))

	L.SetField(imguiTable, "setCursorPos", L.NewFunction(func(L *lua.LState) int {
		t := L.CheckTable(1)
		imgui.SetCursorPos(imgui.Vec2{
			X: float32(L.GetField(t, "x").(lua.LNumber)),
			Y: float32(L.GetField(t, "y").(lua.LNumber)),
		})
		return 0
	}))

	L.SetField(imguiTable, "getCursorScreenPos", L.NewFunction(func(L *lua.LState) int {
		pos := imgui.CursorScreenPos()
		t := L.NewTable()
		L.SetField(t, "x", lua.LNumber(pos.X))
		L.SetField(t, "y", lua.LNumber(pos.Y))
		L.Push(t)
		return 1
	}))

	L.SetField(imguiTable, "setCursorScreenPos", L.NewFunction(func(L *lua.LState) int {
		t := L.CheckTable(1)
		imgui.SetCursorScreenPos(imgui.Vec2{
			X: float32(L.GetField(t, "x").(lua.LNumber)),
			Y: float32(L.GetField(t, "y").(lua.LNumber)),
		})
		return 0
	}))

	L.SetField(imguiTable, "getCursorStartPos", L.NewFunction(func(L *lua.LState) int {
		pos := imgui.CursorStartPos()
		t := L.NewTable()
		L.SetField(t, "x", lua.LNumber(pos.X))
		L.SetField(t, "y", lua.LNumber(pos.Y))
		L.Push(t)
		return 1
	}))

	// ==================== 菜单和工具栏 ====================

	L.SetField(imguiTable, "beginMenuBar", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(imgui.BeginMenuBar()))
		return 1
	}))

	L.SetField(imguiTable, "endMenuBar", L.NewFunction(func(L *lua.LState) int {
		imgui.EndMenuBar()
		return 0
	}))

	L.SetField(imguiTable, "beginMainMenuBar", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(imgui.BeginMainMenuBar()))
		return 1
	}))

	L.SetField(imguiTable, "endMainMenuBar", L.NewFunction(func(L *lua.LState) int {
		imgui.EndMainMenuBar()
		return 0
	}))

	L.SetField(imguiTable, "beginMenu", L.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		var enabled bool = true
		if L.GetTop() >= 2 {
			enabled = L.CheckBool(2)
		}
		L.Push(lua.LBool(imgui.BeginMenuV(label, enabled)))
		return 1
	}))

	L.SetField(imguiTable, "endMenu", L.NewFunction(func(L *lua.LState) int {
		imgui.EndMenu()
		return 0
	}))

	L.SetField(imguiTable, "menuItem", L.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		var shortcut string
		var selected, enabled bool = false, true

		if L.GetTop() >= 2 {
			shortcut = L.CheckString(2)
		}
		if L.GetTop() >= 3 {
			selected = L.CheckBool(3)
		}
		if L.GetTop() >= 4 {
			enabled = L.CheckBool(4)
		}
		L.Push(lua.LBool(imgui.MenuItemBoolV(label, shortcut, selected, enabled)))
		return 1
	}))

	// ==================== 弹窗 ====================

	L.SetField(imguiTable, "beginPopup", L.NewFunction(func(L *lua.LState) int {
		str_id := L.CheckString(1)
		var flags imgui.WindowFlags
		if L.GetTop() >= 2 {
			flags = imgui.WindowFlags(L.CheckInt(2))
		}
		L.Push(lua.LBool(imgui.BeginPopupV(str_id, flags)))
		return 1
	}))

	L.SetField(imguiTable, "endPopup", L.NewFunction(func(L *lua.LState) int {
		imgui.EndPopup()
		return 0
	}))

	L.SetField(imguiTable, "openPopup", L.NewFunction(func(L *lua.LState) int {
		str_id := L.CheckString(1)
		var flags imgui.PopupFlags
		if L.GetTop() >= 2 {
			flags = imgui.PopupFlags(L.CheckInt(2))
		}
		imgui.OpenPopupStrV(str_id, flags)
		return 0
	}))

	L.SetField(imguiTable, "closeCurrentPopup", L.NewFunction(func(L *lua.LState) int {
		imgui.CloseCurrentPopup()
		return 0
	}))

	L.SetField(imguiTable, "beginPopupModal", L.NewFunction(func(L *lua.LState) int {
		name := L.CheckString(1)
		var p_open *bool
		var flags imgui.WindowFlags

		if L.GetTop() >= 2 {
			open := L.CheckBool(2)
			p_open = &open
		}
		if L.GetTop() >= 3 {
			flags = imgui.WindowFlags(L.CheckInt(3))
		}
		result := imgui.BeginPopupModalV(name, p_open, flags)
		L.Push(lua.LBool(result))
		if p_open != nil {
			L.Push(lua.LBool(*p_open))
			return 2
		}
		return 1
	}))

	L.SetField(imguiTable, "beginPopupContextItem", L.NewFunction(func(L *lua.LState) int {
		var str_id string
		var popup_flags imgui.PopupFlags

		if L.GetTop() >= 1 {
			str_id = L.CheckString(1)
		}
		if L.GetTop() >= 2 {
			popup_flags = imgui.PopupFlags(L.CheckInt(2))
		}
		L.Push(lua.LBool(imgui.BeginPopupContextItemV(str_id, popup_flags)))
		return 1
	}))

	L.SetField(imguiTable, "beginPopupContextWindow", L.NewFunction(func(L *lua.LState) int {
		var str_id string
		var popup_flags imgui.PopupFlags

		if L.GetTop() >= 1 {
			str_id = L.CheckString(1)
		}
		if L.GetTop() >= 2 {
			popup_flags = imgui.PopupFlags(L.CheckInt(2))
		}
		L.Push(lua.LBool(imgui.BeginPopupContextWindowV(str_id, popup_flags)))
		return 1
	}))

	// ==================== Tab 栏 ====================

	L.SetField(imguiTable, "beginTabBar", L.NewFunction(func(L *lua.LState) int {
		str_id := L.CheckString(1)
		var flags imgui.TabBarFlags
		if L.GetTop() >= 2 {
			flags = imgui.TabBarFlags(L.CheckInt(2))
		}
		L.Push(lua.LBool(imgui.BeginTabBarV(str_id, flags)))
		return 1
	}))

	L.SetField(imguiTable, "endTabBar", L.NewFunction(func(L *lua.LState) int {
		imgui.EndTabBar()
		return 0
	}))

	L.SetField(imguiTable, "beginTabItem", L.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		var p_open *bool
		var flags imgui.TabItemFlags

		if L.GetTop() >= 2 {
			open := L.CheckBool(2)
			p_open = &open
		}
		if L.GetTop() >= 3 {
			flags = imgui.TabItemFlags(L.CheckInt(3))
		}
		result := imgui.BeginTabItemV(label, p_open, flags)
		L.Push(lua.LBool(result))
		return 1
	}))

	L.SetField(imguiTable, "endTabItem", L.NewFunction(func(L *lua.LState) int {
		imgui.EndTabItem()
		return 0
	}))

	L.SetField(imguiTable, "setTabItemClosed", L.NewFunction(func(L *lua.LState) int {
		tab_or_docked_window_label := L.CheckString(1)
		imgui.SetTabItemClosed(tab_or_docked_window_label)
		return 0
	}))

	// ==================== 表格 ====================

	L.SetField(imguiTable, "beginTable", L.NewFunction(func(L *lua.LState) int {
		str_id := L.CheckString(1)
		columns := int32(L.CheckInt(2))
		var flags imgui.TableFlags
		var outer_size imgui.Vec2
		var inner_width float32

		if L.GetTop() >= 3 {
			flags = imgui.TableFlags(L.CheckInt(3))
		}
		if L.GetTop() >= 4 {
			t := L.CheckTable(4)
			outer_size.X = float32(L.GetField(t, "x").(lua.LNumber))
			outer_size.Y = float32(L.GetField(t, "y").(lua.LNumber))
		}
		if L.GetTop() >= 5 {
			inner_width = float32(L.CheckNumber(5))
		}
		L.Push(lua.LBool(imgui.BeginTableV(str_id, columns, flags, outer_size, inner_width)))
		return 1
	}))

	L.SetField(imguiTable, "endTable", L.NewFunction(func(L *lua.LState) int {
		imgui.EndTable()
		return 0
	}))

	L.SetField(imguiTable, "tableNextRow", L.NewFunction(func(L *lua.LState) int {
		var row_flags imgui.TableRowFlags
		var min_row_height float32

		if L.GetTop() >= 1 {
			row_flags = imgui.TableRowFlags(L.CheckInt(1))
		}
		if L.GetTop() >= 2 {
			min_row_height = float32(L.CheckNumber(2))
		}
		imgui.TableNextRowV(row_flags, min_row_height)
		return 0
	}))

	L.SetField(imguiTable, "tableNextColumn", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(imgui.TableNextColumn()))
		return 1
	}))

	L.SetField(imguiTable, "tableSetColumnIndex", L.NewFunction(func(L *lua.LState) int {
		column_n := int32(L.CheckInt(1))
		L.Push(lua.LBool(imgui.TableSetColumnIndex(column_n)))
		return 1
	}))

	L.SetField(imguiTable, "tableSetupColumn", L.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		var flags imgui.TableColumnFlags
		var init_width_or_weight float32

		if L.GetTop() >= 2 {
			flags = imgui.TableColumnFlags(L.CheckInt(2))
		}
		if L.GetTop() >= 3 {
			init_width_or_weight = float32(L.CheckNumber(3))
		}
		imgui.TableSetupColumnV(label, flags, init_width_or_weight, 0)
		return 0
	}))

	L.SetField(imguiTable, "tableSetupScrollFreeze", L.NewFunction(func(L *lua.LState) int {
		cols := int32(L.CheckInt(1))
		rows := int32(L.CheckInt(2))
		imgui.TableSetupScrollFreeze(cols, rows)
		return 0
	}))

	L.SetField(imguiTable, "tableHeadersRow", L.NewFunction(func(L *lua.LState) int {
		imgui.TableHeadersRow()
		return 0
	}))

	L.SetField(imguiTable, "tableHeader", L.NewFunction(func(L *lua.LState) int {
		label := L.CheckString(1)
		imgui.TableHeader(label)
		return 0
	}))

	// ==================== Tooltip ====================

	L.SetField(imguiTable, "beginTooltip", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(imgui.BeginTooltip()))
		return 1
	}))

	L.SetField(imguiTable, "endTooltip", L.NewFunction(func(L *lua.LState) int {
		imgui.EndTooltip()
		return 0
	}))

	L.SetField(imguiTable, "setTooltip", L.NewFunction(func(L *lua.LState) int {
		text := L.CheckString(1)
		imgui.SetTooltip(text)
		return 0
	}))

	// ==================== 拖放 ====================

	L.SetField(imguiTable, "beginDragDropSource", L.NewFunction(func(L *lua.LState) int {
		var flags imgui.DragDropFlags
		if L.GetTop() >= 1 {
			flags = imgui.DragDropFlags(L.CheckInt(1))
		}
		L.Push(lua.LBool(imgui.BeginDragDropSourceV(flags)))
		return 1
	}))

	L.SetField(imguiTable, "endDragDropSource", L.NewFunction(func(L *lua.LState) int {
		imgui.EndDragDropSource()
		return 0
	}))

	L.SetField(imguiTable, "beginDragDropTarget", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(imgui.BeginDragDropTarget()))
		return 1
	}))

	L.SetField(imguiTable, "endDragDropTarget", L.NewFunction(func(L *lua.LState) int {
		imgui.EndDragDropTarget()
		return 0
	}))

	// ==================== 禁用 ====================

	L.SetField(imguiTable, "beginDisabled", L.NewFunction(func(L *lua.LState) int {
		var disabled bool = true
		if L.GetTop() >= 1 {
			disabled = L.CheckBool(1)
		}
		imgui.BeginDisabledV(disabled)
		return 0
	}))

	L.SetField(imguiTable, "endDisabled", L.NewFunction(func(L *lua.LState) int {
		imgui.EndDisabled()
		return 0
	}))

	// ==================== 样式函数 ====================

	L.SetField(imguiTable, "pushStyleColor", L.NewFunction(func(L *lua.LState) int {
		idx := imgui.Col(L.CheckInt(1))
		t := L.CheckTable(2)
		color := imgui.Vec4{
			X: float32(L.GetField(t, "x").(lua.LNumber)),
			Y: float32(L.GetField(t, "y").(lua.LNumber)),
			Z: float32(L.GetField(t, "z").(lua.LNumber)),
			W: float32(L.GetField(t, "w").(lua.LNumber)),
		}
		imgui.PushStyleColorVec4(idx, color)
		return 0
	}))

	L.SetField(imguiTable, "popStyleColor", L.NewFunction(func(L *lua.LState) int {
		var count int32 = 1
		if L.GetTop() >= 1 {
			count = int32(L.CheckInt(1))
		}
		imgui.PopStyleColorV(count)
		return 0
	}))

	L.SetField(imguiTable, "pushStyleVar", L.NewFunction(func(L *lua.LState) int {
		idx := imgui.StyleVar(L.CheckInt(1))
		if L.GetTop() >= 2 {
			if L.Get(2).Type() == lua.LTNumber {
				imgui.PushStyleVarFloat(idx, float32(L.CheckNumber(2)))
			} else {
				t := L.CheckTable(2)
				imgui.PushStyleVarVec2(idx, imgui.Vec2{
					X: float32(L.GetField(t, "x").(lua.LNumber)),
					Y: float32(L.GetField(t, "y").(lua.LNumber)),
				})
			}
		}
		return 0
	}))

	L.SetField(imguiTable, "popStyleVar", L.NewFunction(func(L *lua.LState) int {
		var count int32 = 1
		if L.GetTop() >= 1 {
			count = int32(L.CheckInt(1))
		}
		imgui.PopStyleVarV(count)
		return 0
	}))

	// ==================== 字体 ====================

	L.SetField(imguiTable, "popFont", L.NewFunction(func(L *lua.LState) int {
		imgui.PopFont()
		return 0
	}))

	// ==================== 窗口设置 ====================

	L.SetField(imguiTable, "setNextWindowPos", L.NewFunction(func(L *lua.LState) int {
		t := L.CheckTable(1)
		var cond imgui.Cond
		var pivot imgui.Vec2

		if L.GetTop() >= 2 {
			cond = imgui.Cond(L.CheckInt(2))
		}
		if L.GetTop() >= 3 {
			pv := L.CheckTable(3)
			pivot.X = float32(L.GetField(pv, "x").(lua.LNumber))
			pivot.Y = float32(L.GetField(pv, "y").(lua.LNumber))
		}
		imgui.SetNextWindowPosV(imgui.Vec2{
			X: float32(L.GetField(t, "x").(lua.LNumber)),
			Y: float32(L.GetField(t, "y").(lua.LNumber)),
		}, cond, pivot)
		return 0
	}))

	L.SetField(imguiTable, "setNextWindowSize", L.NewFunction(func(L *lua.LState) int {
		t := L.CheckTable(1)
		var cond imgui.Cond

		if L.GetTop() >= 2 {
			cond = imgui.Cond(L.CheckInt(2))
		}
		imgui.SetNextWindowSizeV(imgui.Vec2{
			X: float32(L.GetField(t, "x").(lua.LNumber)),
			Y: float32(L.GetField(t, "y").(lua.LNumber)),
		}, cond)
		return 0
	}))

	L.SetField(imguiTable, "setNextWindowSizeConstraints", L.NewFunction(func(L *lua.LState) int {
		t1 := L.CheckTable(1)
		t2 := L.CheckTable(2)
		imgui.SetNextWindowSizeConstraintsV(
			imgui.Vec2{X: float32(L.GetField(t1, "x").(lua.LNumber)), Y: float32(L.GetField(t1, "y").(lua.LNumber))},
			imgui.Vec2{X: float32(L.GetField(t2, "x").(lua.LNumber)), Y: float32(L.GetField(t2, "y").(lua.LNumber))},
			nil, 0,
		)
		return 0
	}))

	L.SetField(imguiTable, "setNextWindowContentSize", L.NewFunction(func(L *lua.LState) int {
		t := L.CheckTable(1)
		imgui.SetNextWindowContentSize(imgui.Vec2{
			X: float32(L.GetField(t, "x").(lua.LNumber)),
			Y: float32(L.GetField(t, "y").(lua.LNumber)),
		})
		return 0
	}))

	L.SetField(imguiTable, "setNextWindowCollapsed", L.NewFunction(func(L *lua.LState) int {
		collapsed := L.CheckBool(1)
		var cond imgui.Cond
		if L.GetTop() >= 2 {
			cond = imgui.Cond(L.CheckInt(2))
		}
		imgui.SetNextWindowCollapsedV(collapsed, cond)
		return 0
	}))

	L.SetField(imguiTable, "setNextWindowFocus", L.NewFunction(func(L *lua.LState) int {
		imgui.SetNextWindowFocus()
		return 0
	}))

	L.SetField(imguiTable, "setWindowPos", L.NewFunction(func(L *lua.LState) int {
		t := L.CheckTable(1)
		var cond imgui.Cond
		if L.GetTop() >= 2 {
			cond = imgui.Cond(L.CheckInt(2))
		}
		imgui.SetWindowPosVec2V(imgui.Vec2{
			X: float32(L.GetField(t, "x").(lua.LNumber)),
			Y: float32(L.GetField(t, "y").(lua.LNumber)),
		}, cond)
		return 0
	}))

	L.SetField(imguiTable, "setWindowSize", L.NewFunction(func(L *lua.LState) int {
		t := L.CheckTable(1)
		var cond imgui.Cond
		if L.GetTop() >= 2 {
			cond = imgui.Cond(L.CheckInt(2))
		}
		imgui.SetWindowSizeVec2V(imgui.Vec2{
			X: float32(L.GetField(t, "x").(lua.LNumber)),
			Y: float32(L.GetField(t, "y").(lua.LNumber)),
		}, cond)
		return 0
	}))

	L.SetField(imguiTable, "setWindowCollapsed", L.NewFunction(func(L *lua.LState) int {
		collapsed := L.CheckBool(1)
		var cond imgui.Cond
		if L.GetTop() >= 2 {
			cond = imgui.Cond(L.CheckInt(2))
		}
		imgui.SetWindowCollapsedBoolV(collapsed, cond)
		return 0
	}))

	L.SetField(imguiTable, "setWindowFocus", L.NewFunction(func(L *lua.LState) int {
		imgui.SetWindowFocus()
		return 0
	}))

	// ==================== 窗口信息获取 ====================

	L.SetField(imguiTable, "getWindowPos", L.NewFunction(func(L *lua.LState) int {
		pos := imgui.WindowPos()
		t := L.NewTable()
		L.SetField(t, "x", lua.LNumber(pos.X))
		L.SetField(t, "y", lua.LNumber(pos.Y))
		L.Push(t)
		return 1
	}))

	L.SetField(imguiTable, "getWindowSize", L.NewFunction(func(L *lua.LState) int {
		size := imgui.WindowSize()
		t := L.NewTable()
		L.SetField(t, "x", lua.LNumber(size.X))
		L.SetField(t, "y", lua.LNumber(size.Y))
		L.Push(t)
		return 1
	}))

	L.SetField(imguiTable, "getWindowWidth", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LNumber(imgui.WindowWidth()))
		return 1
	}))

	L.SetField(imguiTable, "getWindowHeight", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LNumber(imgui.WindowHeight()))
		return 1
	}))

	// ==================== 内容区域 ====================

	L.SetField(imguiTable, "getContentRegionAvail", L.NewFunction(func(L *lua.LState) int {
		size := imgui.ContentRegionAvail()
		t := L.NewTable()
		L.SetField(t, "x", lua.LNumber(size.X))
		L.SetField(t, "y", lua.LNumber(size.Y))
		L.Push(t)
		return 1
	}))

	// ==================== 项目状态查询 ====================

	L.SetField(imguiTable, "isItemHovered", L.NewFunction(func(L *lua.LState) int {
		var flags imgui.HoveredFlags
		if L.GetTop() >= 1 {
			flags = imgui.HoveredFlags(L.CheckInt(1))
		}
		L.Push(lua.LBool(imgui.IsItemHoveredV(flags)))
		return 1
	}))

	L.SetField(imguiTable, "isItemActive", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(imgui.IsItemActive()))
		return 1
	}))

	L.SetField(imguiTable, "isItemFocused", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(imgui.IsItemFocused()))
		return 1
	}))

	L.SetField(imguiTable, "isItemClicked", L.NewFunction(func(L *lua.LState) int {
		var mouse_button imgui.MouseButton
		if L.GetTop() >= 1 {
			mouse_button = imgui.MouseButton(L.CheckInt(1))
		}
		L.Push(lua.LBool(imgui.IsItemClickedV(mouse_button)))
		return 1
	}))

	L.SetField(imguiTable, "isItemVisible", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(imgui.IsItemVisible()))
		return 1
	}))

	L.SetField(imguiTable, "isAnyItemHovered", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LBool(imgui.IsAnyItemHovered()))
		return 1
	}))

	L.SetField(imguiTable, "isWindowHovered", L.NewFunction(func(L *lua.LState) int {
		var flags imgui.HoveredFlags
		if L.GetTop() >= 1 {
			flags = imgui.HoveredFlags(L.CheckInt(1))
		}
		L.Push(lua.LBool(imgui.IsWindowHoveredV(flags)))
		return 1
	}))

	L.SetField(imguiTable, "isWindowFocused", L.NewFunction(func(L *lua.LState) int {
		var flags imgui.FocusedFlags
		if L.GetTop() >= 1 {
			flags = imgui.FocusedFlags(L.CheckInt(1))
		}
		L.Push(lua.LBool(imgui.IsWindowFocusedV(flags)))
		return 1
	}))

	// ==================== 鼠标状态 ====================

	L.SetField(imguiTable, "isMouseDown", L.NewFunction(func(L *lua.LState) int {
		button := imgui.MouseButton(L.CheckInt(1))
		L.Push(lua.LBool(imgui.IsMouseDown(button)))
		return 1
	}))

	L.SetField(imguiTable, "isMouseClicked", L.NewFunction(func(L *lua.LState) int {
		button := imgui.MouseButton(L.CheckInt(1))
		var repeat bool
		if L.GetTop() >= 2 {
			repeat = L.CheckBool(2)
		}
		L.Push(lua.LBool(imgui.IsMouseClickedBoolV(button, repeat)))
		return 1
	}))

	L.SetField(imguiTable, "isMouseReleased", L.NewFunction(func(L *lua.LState) int {
		button := imgui.MouseButton(L.CheckInt(1))
		L.Push(lua.LBool(imgui.IsMouseReleased(button)))
		return 1
	}))

	L.SetField(imguiTable, "isMouseDoubleClicked", L.NewFunction(func(L *lua.LState) int {
		button := imgui.MouseButton(L.CheckInt(1))
		L.Push(lua.LBool(imgui.IsMouseDoubleClicked(button)))
		return 1
	}))

	L.SetField(imguiTable, "isMouseDragging", L.NewFunction(func(L *lua.LState) int {
		button := imgui.MouseButton(L.CheckInt(1))
		var lock_threshold float32
		if L.GetTop() >= 2 {
			lock_threshold = float32(L.CheckNumber(2))
		}
		L.Push(lua.LBool(imgui.IsMouseDraggingV(button, lock_threshold)))
		return 1
	}))

	L.SetField(imguiTable, "isMouseHoveringRect", L.NewFunction(func(L *lua.LState) int {
		t1 := L.CheckTable(1)
		t2 := L.CheckTable(2)
		var clip bool = true
		if L.GetTop() >= 3 {
			clip = L.CheckBool(3)
		}
		L.Push(lua.LBool(imgui.IsMouseHoveringRectV(
			imgui.Vec2{X: float32(L.GetField(t1, "x").(lua.LNumber)), Y: float32(L.GetField(t1, "y").(lua.LNumber))},
			imgui.Vec2{X: float32(L.GetField(t2, "x").(lua.LNumber)), Y: float32(L.GetField(t2, "y").(lua.LNumber))},
			clip,
		)))
		return 1
	}))

	L.SetField(imguiTable, "getMousePos", L.NewFunction(func(L *lua.LState) int {
		io := imgui.CurrentIO()
		pos := io.MousePos()
		t := L.NewTable()
		L.SetField(t, "x", lua.LNumber(pos.X))
		L.SetField(t, "y", lua.LNumber(pos.Y))
		L.Push(t)
		return 1
	}))

	// ==================== 键盘状态 ====================

	L.SetField(imguiTable, "isKeyDown", L.NewFunction(func(L *lua.LState) int {
		key := imgui.Key(L.CheckInt(1))
		L.Push(lua.LBool(imgui.IsKeyDown(key)))
		return 1
	}))

	L.SetField(imguiTable, "isKeyPressed", L.NewFunction(func(L *lua.LState) int {
		key := imgui.Key(L.CheckInt(1))
		var repeat bool = true
		if L.GetTop() >= 2 {
			repeat = L.CheckBool(2)
		}
		L.Push(lua.LBool(imgui.IsKeyPressedBoolV(key, repeat)))
		return 1
	}))

	L.SetField(imguiTable, "isKeyReleased", L.NewFunction(func(L *lua.LState) int {
		key := imgui.Key(L.CheckInt(1))
		L.Push(lua.LBool(imgui.IsKeyReleased(key)))
		return 1
	}))

	// ==================== 焦点控制 ====================

	L.SetField(imguiTable, "setItemDefaultFocus", L.NewFunction(func(L *lua.LState) int {
		imgui.SetItemDefaultFocus()
		return 0
	}))

	L.SetField(imguiTable, "setKeyboardFocusHere", L.NewFunction(func(L *lua.LState) int {
		var offset int32
		if L.GetTop() >= 1 {
			offset = int32(L.CheckInt(1))
		}
		imgui.SetKeyboardFocusHereV(offset)
		return 0
	}))

	// ==================== 时间和帧计数 ====================

	L.SetField(imguiTable, "getTime", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LNumber(imgui.Time()))
		return 1
	}))

	L.SetField(imguiTable, "getFrameCount", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LNumber(imgui.FrameCount()))
		return 1
	}))

	// ==================== 字体相关 ====================

	L.SetField(imguiTable, "getFontSize", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LNumber(imgui.FontSize()))
		return 1
	}))

	L.SetField(imguiTable, "calcTextSize", L.NewFunction(func(L *lua.LState) int {
		text := L.CheckString(1)
		var hide_text_after_double_hash bool
		var wrap_width float32

		if L.GetTop() >= 2 {
			hide_text_after_double_hash = L.CheckBool(2)
		}
		if L.GetTop() >= 3 {
			wrap_width = float32(L.CheckNumber(3))
		}
		size := imgui.CalcTextSizeV(text, hide_text_after_double_hash, wrap_width)
		t := L.NewTable()
		L.SetField(t, "x", lua.LNumber(size.X))
		L.SetField(t, "y", lua.LNumber(size.Y))
		L.Push(t)
		return 1
	}))

	// ==================== 文本行高度 ====================

	L.SetField(imguiTable, "getTextLineHeight", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LNumber(imgui.TextLineHeight()))
		return 1
	}))

	L.SetField(imguiTable, "getTextLineHeightWithSpacing", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LNumber(imgui.TextLineHeightWithSpacing()))
		return 1
	}))

	L.SetField(imguiTable, "getFrameHeight", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LNumber(imgui.FrameHeight()))
		return 1
	}))

	L.SetField(imguiTable, "getFrameHeightWithSpacing", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LNumber(imgui.FrameHeightWithSpacing()))
		return 1
	}))

	// ==================== 鼠标光标 ====================

	L.SetField(imguiTable, "setMouseCursor", L.NewFunction(func(L *lua.LState) int {
		cursor_type := imgui.MouseCursor(L.CheckInt(1))
		imgui.SetMouseCursor(cursor_type)
		return 0
	}))

	// ==================== 颜色转换 ====================

	L.SetField(imguiTable, "colorConvertU32ToFloat4", L.NewFunction(func(L *lua.LState) int {
		in := uint32(L.CheckInt(1))
		col := imgui.ColorConvertU32ToFloat4(in)
		t := L.NewTable()
		L.SetField(t, "x", lua.LNumber(col.X))
		L.SetField(t, "y", lua.LNumber(col.Y))
		L.SetField(t, "z", lua.LNumber(col.Z))
		L.SetField(t, "w", lua.LNumber(col.W))
		L.Push(t)
		return 1
	}))

	L.SetField(imguiTable, "colorConvertFloat4ToU32", L.NewFunction(func(L *lua.LState) int {
		t := L.CheckTable(1)
		in := imgui.Vec4{
			X: float32(L.GetField(t, "x").(lua.LNumber)),
			Y: float32(L.GetField(t, "y").(lua.LNumber)),
			Z: float32(L.GetField(t, "z").(lua.LNumber)),
			W: float32(L.GetField(t, "w").(lua.LNumber)),
		}
		L.Push(lua.LNumber(imgui.ColorConvertFloat4ToU32(in)))
		return 1
	}))

	L.SetField(imguiTable, "colorConvertRGBtoHSV", L.NewFunction(func(L *lua.LState) int {
		r := float32(L.CheckNumber(1))
		g := float32(L.CheckNumber(2))
		b := float32(L.CheckNumber(3))
		var h, s, v float32
		imgui.ColorConvertRGBtoHSV(r, g, b, &h, &s, &v)
		t := L.NewTable()
		L.SetField(t, "h", lua.LNumber(h))
		L.SetField(t, "s", lua.LNumber(s))
		L.SetField(t, "v", lua.LNumber(v))
		L.Push(t)
		return 1
	}))

	L.SetField(imguiTable, "colorConvertHSVtoRGB", L.NewFunction(func(L *lua.LState) int {
		h := float32(L.CheckNumber(1))
		s := float32(L.CheckNumber(2))
		v := float32(L.CheckNumber(3))
		var r, g, b float32
		imgui.ColorConvertHSVtoRGB(h, s, v, &r, &g, &b)
		t := L.NewTable()
		L.SetField(t, "r", lua.LNumber(r))
		L.SetField(t, "g", lua.LNumber(g))
		L.SetField(t, "b", lua.LNumber(b))
		L.Push(t)
		return 1
	}))

	// ==================== ID 操作 ====================

	L.SetField(imguiTable, "pushID", L.NewFunction(func(L *lua.LState) int {
		str_id := L.CheckString(1)
		imgui.PushIDStr(str_id)
		return 0
	}))

	L.SetField(imguiTable, "popID", L.NewFunction(func(L *lua.LState) int {
		imgui.PopID()
		return 0
	}))

	L.SetField(imguiTable, "getID", L.NewFunction(func(L *lua.LState) int {
		str_id := L.CheckString(1)
		L.Push(lua.LNumber(uint(imgui.IDStr(str_id))))
		return 1
	}))

	// ==================== 列布局（旧版） ====================

	L.SetField(imguiTable, "columns", L.NewFunction(func(L *lua.LState) int {
		var count int32 = 1
		var id string
		var border bool = true

		if L.GetTop() >= 1 {
			count = int32(L.CheckInt(1))
		}
		if L.GetTop() >= 2 {
			id = L.CheckString(2)
		}
		if L.GetTop() >= 3 {
			border = L.CheckBool(3)
		}
		imgui.ColumnsV(count, id, border)
		return 0
	}))

	L.SetField(imguiTable, "nextColumn", L.NewFunction(func(L *lua.LState) int {
		imgui.NextColumn()
		return 0
	}))

	L.SetField(imguiTable, "getColumnIndex", L.NewFunction(func(L *lua.LState) int {
		L.Push(lua.LNumber(imgui.ColumnIndex()))
		return 1
	}))

	L.SetField(imguiTable, "getColumnWidth", L.NewFunction(func(L *lua.LState) int {
		var column_index int32 = -1
		if L.GetTop() >= 1 {
			column_index = int32(L.CheckInt(1))
		}
		L.Push(lua.LNumber(imgui.ColumnWidthV(column_index)))
		return 1
	}))

	L.SetField(imguiTable, "setColumnWidth", L.NewFunction(func(L *lua.LState) int {
		column_index := int32(L.CheckInt(1))
		width := float32(L.CheckNumber(2))
		imgui.SetColumnWidth(column_index, width)
		return 0
	}))

	// ==================== 常量定义 ====================

	// 窗口标志
	windowFlags := L.NewTable()
	L.SetField(windowFlags, "None", lua.LNumber(imgui.WindowFlagsNone))
	L.SetField(windowFlags, "NoTitleBar", lua.LNumber(imgui.WindowFlagsNoTitleBar))
	L.SetField(windowFlags, "NoResize", lua.LNumber(imgui.WindowFlagsNoResize))
	L.SetField(windowFlags, "NoMove", lua.LNumber(imgui.WindowFlagsNoMove))
	L.SetField(windowFlags, "NoScrollbar", lua.LNumber(imgui.WindowFlagsNoScrollbar))
	L.SetField(windowFlags, "NoScrollWithMouse", lua.LNumber(imgui.WindowFlagsNoScrollWithMouse))
	L.SetField(windowFlags, "NoCollapse", lua.LNumber(imgui.WindowFlagsNoCollapse))
	L.SetField(windowFlags, "AlwaysAutoResize", lua.LNumber(imgui.WindowFlagsAlwaysAutoResize))
	L.SetField(windowFlags, "NoBackground", lua.LNumber(imgui.WindowFlagsNoBackground))
	L.SetField(windowFlags, "NoSavedSettings", lua.LNumber(imgui.WindowFlagsNoSavedSettings))
	L.SetField(windowFlags, "NoMouseInputs", lua.LNumber(imgui.WindowFlagsNoMouseInputs))
	L.SetField(windowFlags, "MenuBar", lua.LNumber(imgui.WindowFlagsMenuBar))
	L.SetField(windowFlags, "HorizontalScrollbar", lua.LNumber(imgui.WindowFlagsHorizontalScrollbar))
	L.SetField(windowFlags, "NoFocusOnAppearing", lua.LNumber(imgui.WindowFlagsNoFocusOnAppearing))
	L.SetField(windowFlags, "NoBringToFrontOnFocus", lua.LNumber(imgui.WindowFlagsNoBringToFrontOnFocus))
	L.SetField(windowFlags, "AlwaysVerticalScrollbar", lua.LNumber(imgui.WindowFlagsAlwaysVerticalScrollbar))
	L.SetField(windowFlags, "AlwaysHorizontalScrollbar", lua.LNumber(imgui.WindowFlagsAlwaysHorizontalScrollbar))
	L.SetField(windowFlags, "NoNavInputs", lua.LNumber(imgui.WindowFlagsNoNavInputs))
	L.SetField(windowFlags, "NoNavFocus", lua.LNumber(imgui.WindowFlagsNoNavFocus))
	L.SetField(windowFlags, "UnsavedDocument", lua.LNumber(imgui.WindowFlagsUnsavedDocument))
	L.SetField(windowFlags, "NoDocking", lua.LNumber(imgui.WindowFlagsNoDocking))
	L.SetField(imguiTable, "WindowFlags", windowFlags)

	// 颜色索引
	colTable := L.NewTable()
	L.SetField(colTable, "Text", lua.LNumber(imgui.ColText))
	L.SetField(colTable, "TextDisabled", lua.LNumber(imgui.ColTextDisabled))
	L.SetField(colTable, "WindowBg", lua.LNumber(imgui.ColWindowBg))
	L.SetField(colTable, "ChildBg", lua.LNumber(imgui.ColChildBg))
	L.SetField(colTable, "PopupBg", lua.LNumber(imgui.ColPopupBg))
	L.SetField(colTable, "Border", lua.LNumber(imgui.ColBorder))
	L.SetField(colTable, "BorderShadow", lua.LNumber(imgui.ColBorderShadow))
	L.SetField(colTable, "FrameBg", lua.LNumber(imgui.ColFrameBg))
	L.SetField(colTable, "FrameBgHovered", lua.LNumber(imgui.ColFrameBgHovered))
	L.SetField(colTable, "FrameBgActive", lua.LNumber(imgui.ColFrameBgActive))
	L.SetField(colTable, "TitleBg", lua.LNumber(imgui.ColTitleBg))
	L.SetField(colTable, "TitleBgActive", lua.LNumber(imgui.ColTitleBgActive))
	L.SetField(colTable, "TitleBgCollapsed", lua.LNumber(imgui.ColTitleBgCollapsed))
	L.SetField(colTable, "MenuBarBg", lua.LNumber(imgui.ColMenuBarBg))
	L.SetField(colTable, "ScrollbarBg", lua.LNumber(imgui.ColScrollbarBg))
	L.SetField(colTable, "ScrollbarGrab", lua.LNumber(imgui.ColScrollbarGrab))
	L.SetField(colTable, "ScrollbarGrabHovered", lua.LNumber(imgui.ColScrollbarGrabHovered))
	L.SetField(colTable, "ScrollbarGrabActive", lua.LNumber(imgui.ColScrollbarGrabActive))
	L.SetField(colTable, "CheckMark", lua.LNumber(imgui.ColCheckMark))
	L.SetField(colTable, "SliderGrab", lua.LNumber(imgui.ColSliderGrab))
	L.SetField(colTable, "SliderGrabActive", lua.LNumber(imgui.ColSliderGrabActive))
	L.SetField(colTable, "Button", lua.LNumber(imgui.ColButton))
	L.SetField(colTable, "ButtonHovered", lua.LNumber(imgui.ColButtonHovered))
	L.SetField(colTable, "ButtonActive", lua.LNumber(imgui.ColButtonActive))
	L.SetField(colTable, "Header", lua.LNumber(imgui.ColHeader))
	L.SetField(colTable, "HeaderHovered", lua.LNumber(imgui.ColHeaderHovered))
	L.SetField(colTable, "HeaderActive", lua.LNumber(imgui.ColHeaderActive))
	L.SetField(colTable, "Separator", lua.LNumber(imgui.ColSeparator))
	L.SetField(colTable, "SeparatorHovered", lua.LNumber(imgui.ColSeparatorHovered))
	L.SetField(colTable, "SeparatorActive", lua.LNumber(imgui.ColSeparatorActive))
	L.SetField(colTable, "ResizeGrip", lua.LNumber(imgui.ColResizeGrip))
	L.SetField(colTable, "ResizeGripHovered", lua.LNumber(imgui.ColResizeGripHovered))
	L.SetField(colTable, "ResizeGripActive", lua.LNumber(imgui.ColResizeGripActive))
	L.SetField(colTable, "Tab", lua.LNumber(imgui.ColTab))
	L.SetField(colTable, "TabHovered", lua.LNumber(imgui.ColTabHovered))
	L.SetField(colTable, "TabSelected", lua.LNumber(imgui.ColTabSelected))
	L.SetField(colTable, "TabSelectedOverline", lua.LNumber(imgui.ColTabSelectedOverline))
	L.SetField(colTable, "TabDimmed", lua.LNumber(imgui.ColTabDimmed))
	L.SetField(colTable, "TabDimmedSelected", lua.LNumber(imgui.ColTabDimmedSelected))
	L.SetField(colTable, "TabDimmedSelectedOverline", lua.LNumber(imgui.ColTabDimmedSelectedOverline))
	L.SetField(colTable, "PlotLines", lua.LNumber(imgui.ColPlotLines))
	L.SetField(colTable, "PlotLinesHovered", lua.LNumber(imgui.ColPlotLinesHovered))
	L.SetField(colTable, "PlotHistogram", lua.LNumber(imgui.ColPlotHistogram))
	L.SetField(colTable, "PlotHistogramHovered", lua.LNumber(imgui.ColPlotHistogramHovered))
	L.SetField(colTable, "TableHeaderBg", lua.LNumber(imgui.ColTableHeaderBg))
	L.SetField(colTable, "TableBorderStrong", lua.LNumber(imgui.ColTableBorderStrong))
	L.SetField(colTable, "TableBorderLight", lua.LNumber(imgui.ColTableBorderLight))
	L.SetField(colTable, "TableRowBg", lua.LNumber(imgui.ColTableRowBg))
	L.SetField(colTable, "TableRowBgAlt", lua.LNumber(imgui.ColTableRowBgAlt))
	L.SetField(colTable, "TextSelectedBg", lua.LNumber(imgui.ColTextSelectedBg))
	L.SetField(colTable, "DragDropTarget", lua.LNumber(imgui.ColDragDropTarget))
	L.SetField(colTable, "NavWindowingHighlight", lua.LNumber(imgui.ColNavWindowingHighlight))
	L.SetField(colTable, "NavWindowingDimBg", lua.LNumber(imgui.ColNavWindowingDimBg))
	L.SetField(colTable, "ModalWindowDimBg", lua.LNumber(imgui.ColModalWindowDimBg))
	L.SetField(imguiTable, "Col", colTable)

	// 鼠标按钮
	mouseButton := L.NewTable()
	L.SetField(mouseButton, "Left", lua.LNumber(imgui.MouseButtonLeft))
	L.SetField(mouseButton, "Right", lua.LNumber(imgui.MouseButtonRight))
	L.SetField(mouseButton, "Middle", lua.LNumber(imgui.MouseButtonMiddle))
	L.SetField(imguiTable, "MouseButton", mouseButton)

	// 键盘键
	keyTable := L.NewTable()
	L.SetField(keyTable, "None", lua.LNumber(imgui.KeyNone))
	L.SetField(keyTable, "Tab", lua.LNumber(imgui.KeyTab))
	L.SetField(keyTable, "LeftArrow", lua.LNumber(imgui.KeyLeftArrow))
	L.SetField(keyTable, "RightArrow", lua.LNumber(imgui.KeyRightArrow))
	L.SetField(keyTable, "UpArrow", lua.LNumber(imgui.KeyUpArrow))
	L.SetField(keyTable, "DownArrow", lua.LNumber(imgui.KeyDownArrow))
	L.SetField(keyTable, "PageUp", lua.LNumber(imgui.KeyPageUp))
	L.SetField(keyTable, "PageDown", lua.LNumber(imgui.KeyPageDown))
	L.SetField(keyTable, "Home", lua.LNumber(imgui.KeyHome))
	L.SetField(keyTable, "End", lua.LNumber(imgui.KeyEnd))
	L.SetField(keyTable, "Insert", lua.LNumber(imgui.KeyInsert))
	L.SetField(keyTable, "Delete", lua.LNumber(imgui.KeyDelete))
	L.SetField(keyTable, "Backspace", lua.LNumber(imgui.KeyBackspace))
	L.SetField(keyTable, "Space", lua.LNumber(imgui.KeySpace))
	L.SetField(keyTable, "Enter", lua.LNumber(imgui.KeyEnter))
	L.SetField(keyTable, "Escape", lua.LNumber(imgui.KeyEscape))
	L.SetField(keyTable, "LeftCtrl", lua.LNumber(imgui.KeyLeftCtrl))
	L.SetField(keyTable, "LeftShift", lua.LNumber(imgui.KeyLeftShift))
	L.SetField(keyTable, "LeftAlt", lua.LNumber(imgui.KeyLeftAlt))
	L.SetField(keyTable, "LeftSuper", lua.LNumber(imgui.KeyLeftSuper))
	L.SetField(keyTable, "RightCtrl", lua.LNumber(imgui.KeyRightCtrl))
	L.SetField(keyTable, "RightShift", lua.LNumber(imgui.KeyRightShift))
	L.SetField(keyTable, "RightAlt", lua.LNumber(imgui.KeyRightAlt))
	L.SetField(keyTable, "RightSuper", lua.LNumber(imgui.KeyRightSuper))
	L.SetField(keyTable, "Menu", lua.LNumber(imgui.KeyMenu))
	L.SetField(keyTable, "F1", lua.LNumber(imgui.KeyF1))
	L.SetField(keyTable, "F2", lua.LNumber(imgui.KeyF2))
	L.SetField(keyTable, "F3", lua.LNumber(imgui.KeyF3))
	L.SetField(keyTable, "F4", lua.LNumber(imgui.KeyF4))
	L.SetField(keyTable, "F5", lua.LNumber(imgui.KeyF5))
	L.SetField(keyTable, "F6", lua.LNumber(imgui.KeyF6))
	L.SetField(keyTable, "F7", lua.LNumber(imgui.KeyF7))
	L.SetField(keyTable, "F8", lua.LNumber(imgui.KeyF8))
	L.SetField(keyTable, "F9", lua.LNumber(imgui.KeyF9))
	L.SetField(keyTable, "F10", lua.LNumber(imgui.KeyF10))
	L.SetField(keyTable, "F11", lua.LNumber(imgui.KeyF11))
	L.SetField(keyTable, "F12", lua.LNumber(imgui.KeyF12))
	L.SetField(imguiTable, "Key", keyTable)

	// 条件
	condTable := L.NewTable()
	L.SetField(condTable, "None", lua.LNumber(imgui.CondNone))
	L.SetField(condTable, "Always", lua.LNumber(imgui.CondAlways))
	L.SetField(condTable, "Once", lua.LNumber(imgui.CondOnce))
	L.SetField(condTable, "FirstUseEver", lua.LNumber(imgui.CondFirstUseEver))
	L.SetField(condTable, "Appearing", lua.LNumber(imgui.CondAppearing))
	L.SetField(imguiTable, "Cond", condTable)

	// 注册到全局
	L.SetGlobal("imgui", imguiTable)
}
