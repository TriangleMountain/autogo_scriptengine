package js_engine

import (
	"reflect"

	"github.com/dop251/goja"
	"github.com/Dasongzi1366/AutoGo/imgui"
)

// injectImguiMethods 注入 imgui GUI 库方法
func injectImguiMethods(e *JSEngine) {
	vm := e.vm

	// 创建 imgui 对象
	imguiObj := vm.NewObject()

	// ==================== 基础函数 ====================

	// 初始化和生命周期
	imguiObj.Set("init", func(call goja.FunctionCall) goja.Value {
		err := imgui.Init()
		if err != nil {
			panic(vm.NewGoError(err))
		}
		return goja.Undefined()
	})

	imguiObj.Set("run", func(call goja.FunctionCall) goja.Value {
		if len(call.Arguments) > 0 {
			fn := call.Arguments[0]
			if callback, ok := goja.AssertFunction(fn); ok {
				imgui.Run(func() {
					callback(nil)
				})
			}
		}
		return goja.Undefined()
	})

	imguiObj.Set("close", func(call goja.FunctionCall) goja.Value {
		imgui.Close()
		return goja.Undefined()
	})

	// 版本信息
	imguiObj.Set("version", imgui.Version())

	// ==================== 窗口函数 ====================

	imguiObj.Set("begin", func(call goja.FunctionCall) goja.Value {
		name := call.Argument(0).String()
		var p_open *bool
		var flags imgui.WindowFlags

		if len(call.Arguments) > 1 && !goja.IsUndefined(call.Arguments[1]) {
			open := call.Argument(1).ToBoolean()
			p_open = &open
		}
		if len(call.Arguments) > 2 && !goja.IsUndefined(call.Arguments[2]) {
			flags = imgui.WindowFlags(call.Argument(2).ToInteger())
		}

		result := imgui.BeginV(name, p_open, flags)
		if p_open != nil {
			vm.Set("__imgui_window_open", *p_open)
		}
		return vm.ToValue(result)
	})

	imguiObj.Set("end", func(call goja.FunctionCall) goja.Value {
		imgui.End()
		return goja.Undefined()
	})

	// 子窗口
	imguiObj.Set("beginChild", func(call goja.FunctionCall) goja.Value {
		str_id := call.Argument(0).String()
		var size imgui.Vec2
		var child_flags imgui.ChildFlags
		var window_flags imgui.WindowFlags

		if len(call.Arguments) > 1 && !goja.IsUndefined(call.Arguments[1]) {
			sz := call.Argument(1).ToObject(vm)
			size = imgui.Vec2{
				X: float32(sz.Get("x").ToFloat()),
				Y: float32(sz.Get("y").ToFloat()),
			}
		}
		if len(call.Arguments) > 2 && !goja.IsUndefined(call.Arguments[2]) {
			child_flags = imgui.ChildFlags(call.Argument(2).ToInteger())
		}
		if len(call.Arguments) > 3 && !goja.IsUndefined(call.Arguments[3]) {
			window_flags = imgui.WindowFlags(call.Argument(3).ToInteger())
		}

		return vm.ToValue(imgui.BeginChildStrV(str_id, size, child_flags, window_flags))
	})

	imguiObj.Set("endChild", func(call goja.FunctionCall) goja.Value {
		imgui.EndChild()
		return goja.Undefined()
	})

	// ==================== 基础控件 ====================

	// 按钮
	imguiObj.Set("button", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		var size imgui.Vec2
		if len(call.Arguments) > 1 && !goja.IsUndefined(call.Arguments[1]) {
			sz := call.Argument(1).ToObject(vm)
			size = imgui.Vec2{
				X: float32(sz.Get("x").ToFloat()),
				Y: float32(sz.Get("y").ToFloat()),
			}
		}
		return vm.ToValue(imgui.ButtonV(label, size))
	})

	imguiObj.Set("smallButton", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		return vm.ToValue(imgui.SmallButton(label))
	})

	imguiObj.Set("invisibleButton", func(call goja.FunctionCall) goja.Value {
		str_id := call.Argument(0).String()
		var size imgui.Vec2
		var flags imgui.ButtonFlags

		if len(call.Arguments) > 1 && !goja.IsUndefined(call.Arguments[1]) {
			sz := call.Argument(1).ToObject(vm)
			size = imgui.Vec2{
				X: float32(sz.Get("x").ToFloat()),
				Y: float32(sz.Get("y").ToFloat()),
			}
		}
		if len(call.Arguments) > 2 && !goja.IsUndefined(call.Arguments[2]) {
			flags = imgui.ButtonFlags(call.Argument(2).ToInteger())
		}
		return vm.ToValue(imgui.InvisibleButtonV(str_id, size, flags))
	})

	// 文本
	imguiObj.Set("text", func(call goja.FunctionCall) goja.Value {
		text := call.Argument(0).String()
		imgui.Text(text)
		return goja.Undefined()
	})

	imguiObj.Set("textColored", func(call goja.FunctionCall) goja.Value {
		col := call.Argument(0).ToObject(vm)
		color := imgui.Vec4{
			X: float32(col.Get("x").ToFloat()),
			Y: float32(col.Get("y").ToFloat()),
			Z: float32(col.Get("z").ToFloat()),
			W: float32(col.Get("w").ToFloat()),
		}
		text := call.Argument(1).String()
		imgui.TextColored(color, text)
		return goja.Undefined()
	})

	imguiObj.Set("textDisabled", func(call goja.FunctionCall) goja.Value {
		text := call.Argument(0).String()
		imgui.TextDisabled(text)
		return goja.Undefined()
	})

	imguiObj.Set("textWrapped", func(call goja.FunctionCall) goja.Value {
		text := call.Argument(0).String()
		imgui.TextWrapped(text)
		return goja.Undefined()
	})

	imguiObj.Set("bulletText", func(call goja.FunctionCall) goja.Value {
		text := call.Argument(0).String()
		imgui.BulletText(text)
		return goja.Undefined()
	})

	imguiObj.Set("bullet", func(call goja.FunctionCall) goja.Value {
		imgui.Bullet()
		return goja.Undefined()
	})

	// 复选框
	imguiObj.Set("checkbox", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := call.Argument(1).ToBoolean()
		result := imgui.Checkbox(label, &v)
		return vm.ToValue(map[string]interface{}{"checked": v, "changed": result})
	})

	// 单选按钮
	imguiObj.Set("radioButton", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		active := call.Argument(1).ToBoolean()
		return vm.ToValue(imgui.RadioButtonBool(label, active))
	})

	imguiObj.Set("radioButtonIntPtr", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := int32(call.Argument(1).ToInteger())
		v_button := int32(call.Argument(2).ToInteger())
		result := imgui.RadioButtonIntPtr(label, &v, v_button)
		return vm.ToValue(map[string]interface{}{"value": v, "changed": result})
	})

	// 进度条
	imguiObj.Set("progressBar", func(call goja.FunctionCall) goja.Value {
		fraction := float32(call.Argument(0).ToFloat())
		var size imgui.Vec2
		var overlay string

		if len(call.Arguments) > 1 && !goja.IsUndefined(call.Arguments[1]) {
			sz := call.Argument(1).ToObject(vm)
			size = imgui.Vec2{
				X: float32(sz.Get("x").ToFloat()),
				Y: float32(sz.Get("y").ToFloat()),
			}
		}
		if len(call.Arguments) > 2 && !goja.IsUndefined(call.Arguments[2]) {
			overlay = call.Argument(2).String()
		}
		imgui.ProgressBarV(fraction, size, overlay)
		return goja.Undefined()
	})

	// 可选择项
	imguiObj.Set("selectable", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		var selected bool
		var flags imgui.SelectableFlags
		var size imgui.Vec2

		if len(call.Arguments) > 1 && !goja.IsUndefined(call.Arguments[1]) {
			selected = call.Argument(1).ToBoolean()
		}
		if len(call.Arguments) > 2 && !goja.IsUndefined(call.Arguments[2]) {
			flags = imgui.SelectableFlags(call.Argument(2).ToInteger())
		}
		if len(call.Arguments) > 3 && !goja.IsUndefined(call.Arguments[3]) {
			sz := call.Argument(3).ToObject(vm)
			size = imgui.Vec2{
				X: float32(sz.Get("x").ToFloat()),
				Y: float32(sz.Get("y").ToFloat()),
			}
		}
		return vm.ToValue(imgui.SelectableBoolV(label, selected, flags, size))
	})

	// 组合框
	imguiObj.Set("beginCombo", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		preview_value := call.Argument(1).String()
		var flags imgui.ComboFlags
		if len(call.Arguments) > 2 && !goja.IsUndefined(call.Arguments[2]) {
			flags = imgui.ComboFlags(call.Argument(2).ToInteger())
		}
		return vm.ToValue(imgui.BeginComboV(label, preview_value, flags))
	})

	imguiObj.Set("endCombo", func(call goja.FunctionCall) goja.Value {
		imgui.EndCombo()
		return goja.Undefined()
	})

	// 列表框
	imguiObj.Set("beginListBox", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		var size imgui.Vec2
		if len(call.Arguments) > 1 && !goja.IsUndefined(call.Arguments[1]) {
			sz := call.Argument(1).ToObject(vm)
			size = imgui.Vec2{
				X: float32(sz.Get("x").ToFloat()),
				Y: float32(sz.Get("y").ToFloat()),
			}
		}
		return vm.ToValue(imgui.BeginListBoxV(label, size))
	})

	imguiObj.Set("endListBox", func(call goja.FunctionCall) goja.Value {
		imgui.EndListBox()
		return goja.Undefined()
	})

	// ==================== 输入控件 ====================

	// 输入文本
	imguiObj.Set("inputText", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		buf := call.Argument(1).String()
		var flags imgui.InputTextFlags
		if len(call.Arguments) > 2 && !goja.IsUndefined(call.Arguments[2]) {
			flags = imgui.InputTextFlags(call.Argument(2).ToInteger())
		}
		result := imgui.InputTextWithHint(label, "", &buf, flags, nil)
		return vm.ToValue(map[string]interface{}{"value": buf, "changed": result})
	})

	imguiObj.Set("inputTextWithHint", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		hint := call.Argument(1).String()
		buf := call.Argument(2).String()
		var flags imgui.InputTextFlags
		if len(call.Arguments) > 3 && !goja.IsUndefined(call.Arguments[3]) {
			flags = imgui.InputTextFlags(call.Argument(3).ToInteger())
		}
		result := imgui.InputTextWithHint(label, hint, &buf, flags, nil)
		return vm.ToValue(map[string]interface{}{"value": buf, "changed": result})
	})

	imguiObj.Set("inputTextMultiline", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		buf := call.Argument(1).String()
		var size imgui.Vec2
		var flags imgui.InputTextFlags

		if len(call.Arguments) > 2 && !goja.IsUndefined(call.Arguments[2]) {
			sz := call.Argument(2).ToObject(vm)
			size = imgui.Vec2{
				X: float32(sz.Get("x").ToFloat()),
				Y: float32(sz.Get("y").ToFloat()),
			}
		}
		if len(call.Arguments) > 3 && !goja.IsUndefined(call.Arguments[3]) {
			flags = imgui.InputTextFlags(call.Argument(3).ToInteger())
		}
		result := imgui.InputTextMultiline(label, &buf, size, flags, nil)
		return vm.ToValue(map[string]interface{}{"value": buf, "changed": result})
	})

	// 输入数字
	imguiObj.Set("inputInt", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := int32(call.Argument(1).ToInteger())
		var step, step_fast int32
		var flags imgui.InputTextFlags

		if len(call.Arguments) > 2 && !goja.IsUndefined(call.Arguments[2]) {
			step = int32(call.Argument(2).ToInteger())
		}
		if len(call.Arguments) > 3 && !goja.IsUndefined(call.Arguments[3]) {
			step_fast = int32(call.Argument(3).ToInteger())
		}
		if len(call.Arguments) > 4 && !goja.IsUndefined(call.Arguments[4]) {
			flags = imgui.InputTextFlags(call.Argument(4).ToInteger())
		}
		result := imgui.InputIntV(label, &v, step, step_fast, flags)
		return vm.ToValue(map[string]interface{}{"value": v, "changed": result})
	})

	imguiObj.Set("inputFloat", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := float32(call.Argument(1).ToFloat())
		var step, step_fast float32
		var format string
		var flags imgui.InputTextFlags

		if len(call.Arguments) > 2 && !goja.IsUndefined(call.Arguments[2]) {
			step = float32(call.Argument(2).ToFloat())
		}
		if len(call.Arguments) > 3 && !goja.IsUndefined(call.Arguments[3]) {
			step_fast = float32(call.Argument(3).ToFloat())
		}
		if len(call.Arguments) > 4 && !goja.IsUndefined(call.Arguments[4]) {
			format = call.Argument(4).String()
		}
		if len(call.Arguments) > 5 && !goja.IsUndefined(call.Arguments[5]) {
			flags = imgui.InputTextFlags(call.Argument(5).ToInteger())
		}
		result := imgui.InputFloatV(label, &v, step, step_fast, format, flags)
		return vm.ToValue(map[string]interface{}{"value": v, "changed": result})
	})

	imguiObj.Set("inputDouble", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := call.Argument(1).ToFloat()
		var step, step_fast float64
		var format string
		var flags imgui.InputTextFlags

		if len(call.Arguments) > 2 && !goja.IsUndefined(call.Arguments[2]) {
			step = call.Argument(2).ToFloat()
		}
		if len(call.Arguments) > 3 && !goja.IsUndefined(call.Arguments[3]) {
			step_fast = call.Argument(3).ToFloat()
		}
		if len(call.Arguments) > 4 && !goja.IsUndefined(call.Arguments[4]) {
			format = call.Argument(4).String()
		}
		if len(call.Arguments) > 5 && !goja.IsUndefined(call.Arguments[5]) {
			flags = imgui.InputTextFlags(call.Argument(5).ToInteger())
		}
		result := imgui.InputDoubleV(label, &v, step, step_fast, format, flags)
		return vm.ToValue(map[string]interface{}{"value": v, "changed": result})
	})

	// ==================== 滑块控件 ====================

	imguiObj.Set("sliderInt", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := int32(call.Argument(1).ToInteger())
		v_min := int32(call.Argument(2).ToInteger())
		v_max := int32(call.Argument(3).ToInteger())
		var format string
		var flags imgui.SliderFlags

		if len(call.Arguments) > 4 && !goja.IsUndefined(call.Arguments[4]) {
			format = call.Argument(4).String()
		}
		if len(call.Arguments) > 5 && !goja.IsUndefined(call.Arguments[5]) {
			flags = imgui.SliderFlags(call.Argument(5).ToInteger())
		}
		result := imgui.SliderIntV(label, &v, v_min, v_max, format, flags)
		return vm.ToValue(map[string]interface{}{"value": v, "changed": result})
	})

	imguiObj.Set("sliderFloat", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := float32(call.Argument(1).ToFloat())
		v_min := float32(call.Argument(2).ToFloat())
		v_max := float32(call.Argument(3).ToFloat())
		var format string
		var flags imgui.SliderFlags

		if len(call.Arguments) > 4 && !goja.IsUndefined(call.Arguments[4]) {
			format = call.Argument(4).String()
		}
		if len(call.Arguments) > 5 && !goja.IsUndefined(call.Arguments[5]) {
			flags = imgui.SliderFlags(call.Argument(5).ToInteger())
		}
		result := imgui.SliderFloatV(label, &v, v_min, v_max, format, flags)
		return vm.ToValue(map[string]interface{}{"value": v, "changed": result})
	})

	imguiObj.Set("sliderAngle", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v_rad := float32(call.Argument(1).ToFloat())
		var v_degrees_min, v_degrees_max float32
		var format string
		var flags imgui.SliderFlags

		if len(call.Arguments) > 2 && !goja.IsUndefined(call.Arguments[2]) {
			v_degrees_min = float32(call.Argument(2).ToFloat())
		}
		if len(call.Arguments) > 3 && !goja.IsUndefined(call.Arguments[3]) {
			v_degrees_max = float32(call.Argument(3).ToFloat())
		}
		if len(call.Arguments) > 4 && !goja.IsUndefined(call.Arguments[4]) {
			format = call.Argument(4).String()
		}
		if len(call.Arguments) > 5 && !goja.IsUndefined(call.Arguments[5]) {
			flags = imgui.SliderFlags(call.Argument(5).ToInteger())
		}
		result := imgui.SliderAngleV(label, &v_rad, v_degrees_min, v_degrees_max, format, flags)
		return vm.ToValue(map[string]interface{}{"value": v_rad, "changed": result})
	})

	// ==================== 拖拽控件 ====================

	imguiObj.Set("dragInt", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := int32(call.Argument(1).ToInteger())
		var v_speed float32
		var v_min, v_max int32
		var format string
		var flags imgui.SliderFlags

		if len(call.Arguments) > 2 && !goja.IsUndefined(call.Arguments[2]) {
			v_speed = float32(call.Argument(2).ToFloat())
		}
		if len(call.Arguments) > 3 && !goja.IsUndefined(call.Arguments[3]) {
			v_min = int32(call.Argument(3).ToInteger())
		}
		if len(call.Arguments) > 4 && !goja.IsUndefined(call.Arguments[4]) {
			v_max = int32(call.Argument(4).ToInteger())
		}
		if len(call.Arguments) > 5 && !goja.IsUndefined(call.Arguments[5]) {
			format = call.Argument(5).String()
		}
		if len(call.Arguments) > 6 && !goja.IsUndefined(call.Arguments[6]) {
			flags = imgui.SliderFlags(call.Argument(6).ToInteger())
		}
		result := imgui.DragIntV(label, &v, v_speed, v_min, v_max, format, flags)
		return vm.ToValue(map[string]interface{}{"value": v, "changed": result})
	})

	imguiObj.Set("dragFloat", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		v := float32(call.Argument(1).ToFloat())
		var v_speed, v_min, v_max float32
		var format string
		var flags imgui.SliderFlags

		if len(call.Arguments) > 2 && !goja.IsUndefined(call.Arguments[2]) {
			v_speed = float32(call.Argument(2).ToFloat())
		}
		if len(call.Arguments) > 3 && !goja.IsUndefined(call.Arguments[3]) {
			v_min = float32(call.Argument(3).ToFloat())
		}
		if len(call.Arguments) > 4 && !goja.IsUndefined(call.Arguments[4]) {
			v_max = float32(call.Argument(4).ToFloat())
		}
		if len(call.Arguments) > 5 && !goja.IsUndefined(call.Arguments[5]) {
			format = call.Argument(5).String()
		}
		if len(call.Arguments) > 6 && !goja.IsUndefined(call.Arguments[6]) {
			flags = imgui.SliderFlags(call.Argument(6).ToInteger())
		}
		result := imgui.DragFloatV(label, &v, v_speed, v_min, v_max, format, flags)
		return vm.ToValue(map[string]interface{}{"value": v, "changed": result})
	})

	// ==================== 颜色控件 ====================

	imguiObj.Set("colorEdit3", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		col := call.Argument(1).ToObject(vm)
		colArr := [3]float32{
			float32(col.Get("0").ToFloat()),
			float32(col.Get("1").ToFloat()),
			float32(col.Get("2").ToFloat()),
		}
		var flags imgui.ColorEditFlags
		if len(call.Arguments) > 2 && !goja.IsUndefined(call.Arguments[2]) {
			flags = imgui.ColorEditFlags(call.Argument(2).ToInteger())
		}
		result := imgui.ColorEdit3V(label, &colArr, flags)
		return vm.ToValue(map[string]interface{}{
			"value":   []float32{colArr[0], colArr[1], colArr[2]},
			"changed": result,
		})
	})

	imguiObj.Set("colorEdit4", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		col := call.Argument(1).ToObject(vm)
		colArr := [4]float32{
			float32(col.Get("0").ToFloat()),
			float32(col.Get("1").ToFloat()),
			float32(col.Get("2").ToFloat()),
			float32(col.Get("3").ToFloat()),
		}
		var flags imgui.ColorEditFlags
		if len(call.Arguments) > 2 && !goja.IsUndefined(call.Arguments[2]) {
			flags = imgui.ColorEditFlags(call.Argument(2).ToInteger())
		}
		result := imgui.ColorEdit4V(label, &colArr, flags)
		return vm.ToValue(map[string]interface{}{
			"value":   []float32{colArr[0], colArr[1], colArr[2], colArr[3]},
			"changed": result,
		})
	})

	imguiObj.Set("colorButton", func(call goja.FunctionCall) goja.Value {
		desc_id := call.Argument(0).String()
		col := call.Argument(1).ToObject(vm)
		color := imgui.Vec4{
			X: float32(col.Get("x").ToFloat()),
			Y: float32(col.Get("y").ToFloat()),
			Z: float32(col.Get("z").ToFloat()),
			W: float32(col.Get("w").ToFloat()),
		}
		var flags imgui.ColorEditFlags
		var size imgui.Vec2

		if len(call.Arguments) > 2 && !goja.IsUndefined(call.Arguments[2]) {
			flags = imgui.ColorEditFlags(call.Argument(2).ToInteger())
		}
		if len(call.Arguments) > 3 && !goja.IsUndefined(call.Arguments[3]) {
			sz := call.Argument(3).ToObject(vm)
			size = imgui.Vec2{
				X: float32(sz.Get("x").ToFloat()),
				Y: float32(sz.Get("y").ToFloat()),
			}
		}
		return vm.ToValue(imgui.ColorButtonV(desc_id, color, flags, size))
	})

	// ==================== 树形控件 ====================

	imguiObj.Set("treeNode", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		var flags imgui.TreeNodeFlags
		if len(call.Arguments) > 1 && !goja.IsUndefined(call.Arguments[1]) {
			flags = imgui.TreeNodeFlags(call.Argument(1).ToInteger())
		}
		return vm.ToValue(imgui.TreeNodeExStrV(label, flags))
	})

	imguiObj.Set("treePush", func(call goja.FunctionCall) goja.Value {
		var str_id string
		if len(call.Arguments) > 0 && !goja.IsUndefined(call.Arguments[0]) {
			str_id = call.Argument(0).String()
		}
		imgui.TreePushStr(str_id)
		return goja.Undefined()
	})

	imguiObj.Set("treePop", func(call goja.FunctionCall) goja.Value {
		imgui.TreePop()
		return goja.Undefined()
	})

	imguiObj.Set("collapsingHeader", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		var flags imgui.TreeNodeFlags
		if len(call.Arguments) > 1 && !goja.IsUndefined(call.Arguments[1]) {
			flags = imgui.TreeNodeFlags(call.Argument(1).ToInteger())
		}
		return vm.ToValue(imgui.CollapsingHeaderTreeNodeFlagsV(label, flags))
	})

	// ==================== 折叠头部 ====================

	imguiObj.Set("setNextItemOpen", func(call goja.FunctionCall) goja.Value {
		is_open := call.Argument(0).ToBoolean()
		var cond imgui.Cond
		if len(call.Arguments) > 1 && !goja.IsUndefined(call.Arguments[1]) {
			cond = imgui.Cond(call.Argument(1).ToInteger())
		}
		imgui.SetNextItemOpenV(is_open, cond)
		return goja.Undefined()
	})

	// ==================== 布局函数 ====================

	imguiObj.Set("separator", func(call goja.FunctionCall) goja.Value {
		imgui.Separator()
		return goja.Undefined()
	})

	imguiObj.Set("sameLine", func(call goja.FunctionCall) goja.Value {
		var offset_from_start_x float32
		var spacing float32 = -1

		if len(call.Arguments) > 0 && !goja.IsUndefined(call.Arguments[0]) {
			offset_from_start_x = float32(call.Argument(0).ToFloat())
		}
		if len(call.Arguments) > 1 && !goja.IsUndefined(call.Arguments[1]) {
			spacing = float32(call.Argument(1).ToFloat())
		}
		imgui.SameLineV(offset_from_start_x, spacing)
		return goja.Undefined()
	})

	imguiObj.Set("newLine", func(call goja.FunctionCall) goja.Value {
		imgui.NewLine()
		return goja.Undefined()
	})

	imguiObj.Set("spacing", func(call goja.FunctionCall) goja.Value {
		imgui.Spacing()
		return goja.Undefined()
	})

	imguiObj.Set("dummy", func(call goja.FunctionCall) goja.Value {
		size := call.Argument(0).ToObject(vm)
		imgui.Dummy(imgui.Vec2{
			X: float32(size.Get("x").ToFloat()),
			Y: float32(size.Get("y").ToFloat()),
		})
		return goja.Undefined()
	})

	imguiObj.Set("indent", func(call goja.FunctionCall) goja.Value {
		var indent_w float32
		if len(call.Arguments) > 0 && !goja.IsUndefined(call.Arguments[0]) {
			indent_w = float32(call.Argument(0).ToFloat())
			imgui.IndentV(indent_w)
		} else {
			imgui.Indent()
		}
		return goja.Undefined()
	})

	imguiObj.Set("unindent", func(call goja.FunctionCall) goja.Value {
		var indent_w float32
		if len(call.Arguments) > 0 && !goja.IsUndefined(call.Arguments[0]) {
			indent_w = float32(call.Argument(0).ToFloat())
			imgui.UnindentV(indent_w)
		} else {
			imgui.Unindent()
		}
		return goja.Undefined()
	})

	imguiObj.Set("beginGroup", func(call goja.FunctionCall) goja.Value {
		imgui.BeginGroup()
		return goja.Undefined()
	})

	imguiObj.Set("endGroup", func(call goja.FunctionCall) goja.Value {
		imgui.EndGroup()
		return goja.Undefined()
	})

	imguiObj.Set("getCursorPos", func(call goja.FunctionCall) goja.Value {
		pos := imgui.CursorPos()
		return vm.ToValue(map[string]float32{"x": pos.X, "y": pos.Y})
	})

	imguiObj.Set("setCursorPos", func(call goja.FunctionCall) goja.Value {
		pos := call.Argument(0).ToObject(vm)
		imgui.SetCursorPos(imgui.Vec2{
			X: float32(pos.Get("x").ToFloat()),
			Y: float32(pos.Get("y").ToFloat()),
		})
		return goja.Undefined()
	})

	imguiObj.Set("getCursorScreenPos", func(call goja.FunctionCall) goja.Value {
		pos := imgui.CursorScreenPos()
		return vm.ToValue(map[string]float32{"x": pos.X, "y": pos.Y})
	})

	imguiObj.Set("setCursorScreenPos", func(call goja.FunctionCall) goja.Value {
		pos := call.Argument(0).ToObject(vm)
		imgui.SetCursorScreenPos(imgui.Vec2{
			X: float32(pos.Get("x").ToFloat()),
			Y: float32(pos.Get("y").ToFloat()),
		})
		return goja.Undefined()
	})

	imguiObj.Set("getCursorStartPos", func(call goja.FunctionCall) goja.Value {
		pos := imgui.CursorStartPos()
		return vm.ToValue(map[string]float32{"x": pos.X, "y": pos.Y})
	})

	// ==================== 菜单和工具栏 ====================

	imguiObj.Set("beginMenuBar", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.BeginMenuBar())
	})

	imguiObj.Set("endMenuBar", func(call goja.FunctionCall) goja.Value {
		imgui.EndMenuBar()
		return goja.Undefined()
	})

	imguiObj.Set("beginMainMenuBar", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.BeginMainMenuBar())
	})

	imguiObj.Set("endMainMenuBar", func(call goja.FunctionCall) goja.Value {
		imgui.EndMainMenuBar()
		return goja.Undefined()
	})

	imguiObj.Set("beginMenu", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		var enabled bool = true
		if len(call.Arguments) > 1 && !goja.IsUndefined(call.Arguments[1]) {
			enabled = call.Argument(1).ToBoolean()
		}
		return vm.ToValue(imgui.BeginMenuV(label, enabled))
	})

	imguiObj.Set("endMenu", func(call goja.FunctionCall) goja.Value {
		imgui.EndMenu()
		return goja.Undefined()
	})

	imguiObj.Set("menuItem", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		var shortcut string
		var selected, enabled bool = false, true

		if len(call.Arguments) > 1 && !goja.IsUndefined(call.Arguments[1]) {
			shortcut = call.Argument(1).String()
		}
		if len(call.Arguments) > 2 && !goja.IsUndefined(call.Arguments[2]) {
			selected = call.Argument(2).ToBoolean()
		}
		if len(call.Arguments) > 3 && !goja.IsUndefined(call.Arguments[3]) {
			enabled = call.Argument(3).ToBoolean()
		}
		return vm.ToValue(imgui.MenuItemBoolV(label, shortcut, selected, enabled))
	})

	// ==================== 弹窗 ====================

	imguiObj.Set("beginPopup", func(call goja.FunctionCall) goja.Value {
		str_id := call.Argument(0).String()
		var flags imgui.WindowFlags
		if len(call.Arguments) > 1 && !goja.IsUndefined(call.Arguments[1]) {
			flags = imgui.WindowFlags(call.Argument(1).ToInteger())
		}
		return vm.ToValue(imgui.BeginPopupV(str_id, flags))
	})

	imguiObj.Set("endPopup", func(call goja.FunctionCall) goja.Value {
		imgui.EndPopup()
		return goja.Undefined()
	})

	imguiObj.Set("openPopup", func(call goja.FunctionCall) goja.Value {
		str_id := call.Argument(0).String()
		var flags imgui.PopupFlags
		if len(call.Arguments) > 1 && !goja.IsUndefined(call.Arguments[1]) {
			flags = imgui.PopupFlags(call.Argument(1).ToInteger())
		}
		imgui.OpenPopupStrV(str_id, flags)
		return goja.Undefined()
	})

	imguiObj.Set("closeCurrentPopup", func(call goja.FunctionCall) goja.Value {
		imgui.CloseCurrentPopup()
		return goja.Undefined()
	})

	imguiObj.Set("beginPopupModal", func(call goja.FunctionCall) goja.Value {
		name := call.Argument(0).String()
		var p_open *bool
		var flags imgui.WindowFlags

		if len(call.Arguments) > 1 && !goja.IsUndefined(call.Arguments[1]) {
			open := call.Argument(1).ToBoolean()
			p_open = &open
		}
		if len(call.Arguments) > 2 && !goja.IsUndefined(call.Arguments[2]) {
			flags = imgui.WindowFlags(call.Argument(2).ToInteger())
		}
		result := imgui.BeginPopupModalV(name, p_open, flags)
		if p_open != nil {
			vm.Set("__imgui_popup_open", *p_open)
		}
		return vm.ToValue(result)
	})

	imguiObj.Set("beginPopupContextItem", func(call goja.FunctionCall) goja.Value {
		var str_id string
		var popup_flags imgui.PopupFlags

		if len(call.Arguments) > 0 && !goja.IsUndefined(call.Arguments[0]) {
			str_id = call.Argument(0).String()
		}
		if len(call.Arguments) > 1 && !goja.IsUndefined(call.Arguments[1]) {
			popup_flags = imgui.PopupFlags(call.Argument(1).ToInteger())
		}
		return vm.ToValue(imgui.BeginPopupContextItemV(str_id, popup_flags))
	})

	imguiObj.Set("beginPopupContextWindow", func(call goja.FunctionCall) goja.Value {
		var str_id string
		var popup_flags imgui.PopupFlags

		if len(call.Arguments) > 0 && !goja.IsUndefined(call.Arguments[0]) {
			str_id = call.Argument(0).String()
		}
		if len(call.Arguments) > 1 && !goja.IsUndefined(call.Arguments[1]) {
			popup_flags = imgui.PopupFlags(call.Argument(1).ToInteger())
		}
		return vm.ToValue(imgui.BeginPopupContextWindowV(str_id, popup_flags))
	})

	// ==================== Tab 栏 ====================

	imguiObj.Set("beginTabBar", func(call goja.FunctionCall) goja.Value {
		str_id := call.Argument(0).String()
		var flags imgui.TabBarFlags
		if len(call.Arguments) > 1 && !goja.IsUndefined(call.Arguments[1]) {
			flags = imgui.TabBarFlags(call.Argument(1).ToInteger())
		}
		return vm.ToValue(imgui.BeginTabBarV(str_id, flags))
	})

	imguiObj.Set("endTabBar", func(call goja.FunctionCall) goja.Value {
		imgui.EndTabBar()
		return goja.Undefined()
	})

	imguiObj.Set("beginTabItem", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		var p_open *bool
		var flags imgui.TabItemFlags

		if len(call.Arguments) > 1 && !goja.IsUndefined(call.Arguments[1]) {
			open := call.Argument(1).ToBoolean()
			p_open = &open
		}
		if len(call.Arguments) > 2 && !goja.IsUndefined(call.Arguments[2]) {
			flags = imgui.TabItemFlags(call.Argument(2).ToInteger())
		}
		result := imgui.BeginTabItemV(label, p_open, flags)
		return vm.ToValue(result)
	})

	imguiObj.Set("endTabItem", func(call goja.FunctionCall) goja.Value {
		imgui.EndTabItem()
		return goja.Undefined()
	})

	imguiObj.Set("setTabItemClosed", func(call goja.FunctionCall) goja.Value {
		tab_or_docked_window_label := call.Argument(0).String()
		imgui.SetTabItemClosed(tab_or_docked_window_label)
		return goja.Undefined()
	})

	// ==================== 表格 ====================

	imguiObj.Set("beginTable", func(call goja.FunctionCall) goja.Value {
		str_id := call.Argument(0).String()
		columns := int32(call.Argument(1).ToInteger())
		var flags imgui.TableFlags
		var outer_size imgui.Vec2
		var inner_width float32

		if len(call.Arguments) > 2 && !goja.IsUndefined(call.Arguments[2]) {
			flags = imgui.TableFlags(call.Argument(2).ToInteger())
		}
		if len(call.Arguments) > 3 && !goja.IsUndefined(call.Arguments[3]) {
			sz := call.Argument(3).ToObject(vm)
			outer_size = imgui.Vec2{
				X: float32(sz.Get("x").ToFloat()),
				Y: float32(sz.Get("y").ToFloat()),
			}
		}
		if len(call.Arguments) > 4 && !goja.IsUndefined(call.Arguments[4]) {
			inner_width = float32(call.Argument(4).ToFloat())
		}
		return vm.ToValue(imgui.BeginTableV(str_id, columns, flags, outer_size, inner_width))
	})

	imguiObj.Set("endTable", func(call goja.FunctionCall) goja.Value {
		imgui.EndTable()
		return goja.Undefined()
	})

	imguiObj.Set("tableNextRow", func(call goja.FunctionCall) goja.Value {
		var row_flags imgui.TableRowFlags
		var min_row_height float32

		if len(call.Arguments) > 0 && !goja.IsUndefined(call.Arguments[0]) {
			row_flags = imgui.TableRowFlags(call.Argument(0).ToInteger())
		}
		if len(call.Arguments) > 1 && !goja.IsUndefined(call.Arguments[1]) {
			min_row_height = float32(call.Argument(1).ToFloat())
		}
		imgui.TableNextRowV(row_flags, min_row_height)
		return goja.Undefined()
	})

	imguiObj.Set("tableNextColumn", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.TableNextColumn())
	})

	imguiObj.Set("tableSetColumnIndex", func(call goja.FunctionCall) goja.Value {
		column_n := int32(call.Argument(0).ToInteger())
		return vm.ToValue(imgui.TableSetColumnIndex(column_n))
	})

	imguiObj.Set("tableSetupColumn", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		var flags imgui.TableColumnFlags
		var init_width_or_weight float32
		var user_id imgui.ID

		if len(call.Arguments) > 1 && !goja.IsUndefined(call.Arguments[1]) {
			flags = imgui.TableColumnFlags(call.Argument(1).ToInteger())
		}
		if len(call.Arguments) > 2 && !goja.IsUndefined(call.Arguments[2]) {
			init_width_or_weight = float32(call.Argument(2).ToFloat())
		}
		imgui.TableSetupColumnV(label, flags, init_width_or_weight, user_id)
		return goja.Undefined()
	})

	imguiObj.Set("tableSetupScrollFreeze", func(call goja.FunctionCall) goja.Value {
		cols := int32(call.Argument(0).ToInteger())
		rows := int32(call.Argument(1).ToInteger())
		imgui.TableSetupScrollFreeze(cols, rows)
		return goja.Undefined()
	})

	imguiObj.Set("tableHeadersRow", func(call goja.FunctionCall) goja.Value {
		imgui.TableHeadersRow()
		return goja.Undefined()
	})

	imguiObj.Set("tableHeader", func(call goja.FunctionCall) goja.Value {
		label := call.Argument(0).String()
		imgui.TableHeader(label)
		return goja.Undefined()
	})

	// ==================== Tooltip ====================

	imguiObj.Set("beginTooltip", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.BeginTooltip())
	})

	imguiObj.Set("endTooltip", func(call goja.FunctionCall) goja.Value {
		imgui.EndTooltip()
		return goja.Undefined()
	})

	imguiObj.Set("setTooltip", func(call goja.FunctionCall) goja.Value {
		text := call.Argument(0).String()
		imgui.SetTooltip(text)
		return goja.Undefined()
	})

	// ==================== 拖放 ====================

	imguiObj.Set("beginDragDropSource", func(call goja.FunctionCall) goja.Value {
		var flags imgui.DragDropFlags
		if len(call.Arguments) > 0 && !goja.IsUndefined(call.Arguments[0]) {
			flags = imgui.DragDropFlags(call.Argument(0).ToInteger())
		}
		return vm.ToValue(imgui.BeginDragDropSourceV(flags))
	})

	imguiObj.Set("endDragDropSource", func(call goja.FunctionCall) goja.Value {
		imgui.EndDragDropSource()
		return goja.Undefined()
	})

	imguiObj.Set("beginDragDropTarget", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.BeginDragDropTarget())
	})

	imguiObj.Set("endDragDropTarget", func(call goja.FunctionCall) goja.Value {
		imgui.EndDragDropTarget()
		return goja.Undefined()
	})

	// ==================== 禁用 ====================

	imguiObj.Set("beginDisabled", func(call goja.FunctionCall) goja.Value {
		var disabled bool = true
		if len(call.Arguments) > 0 && !goja.IsUndefined(call.Arguments[0]) {
			disabled = call.Argument(0).ToBoolean()
		}
		imgui.BeginDisabledV(disabled)
		return goja.Undefined()
	})

	imguiObj.Set("endDisabled", func(call goja.FunctionCall) goja.Value {
		imgui.EndDisabled()
		return goja.Undefined()
	})

	// ==================== 样式函数 ====================

	imguiObj.Set("pushStyleColor", func(call goja.FunctionCall) goja.Value {
		idx := imgui.Col(call.Argument(0).ToInteger())
		col := call.Argument(1).ToObject(vm)
		color := imgui.Vec4{
			X: float32(col.Get("x").ToFloat()),
			Y: float32(col.Get("y").ToFloat()),
			Z: float32(col.Get("z").ToFloat()),
			W: float32(col.Get("w").ToFloat()),
		}
		imgui.PushStyleColorVec4(idx, color)
		return goja.Undefined()
	})

	imguiObj.Set("popStyleColor", func(call goja.FunctionCall) goja.Value {
		var count int32 = 1
		if len(call.Arguments) > 0 && !goja.IsUndefined(call.Arguments[0]) {
			count = int32(call.Argument(0).ToInteger())
		}
		imgui.PopStyleColorV(count)
		return goja.Undefined()
	})

	imguiObj.Set("pushStyleVar", func(call goja.FunctionCall) goja.Value {
		idx := imgui.StyleVar(call.Argument(0).ToInteger())
		if len(call.Arguments) > 1 {
			arg1 := call.Argument(1)
			if arg1.ExportType().Kind() == reflect.Float32 || arg1.ExportType().Kind() == reflect.Float64 {
				// 单个浮点值
				imgui.PushStyleVarFloat(idx, float32(arg1.ToFloat()))
			} else {
				// Vec2 值
				val := arg1.ToObject(vm)
				imgui.PushStyleVarVec2(idx, imgui.Vec2{
					X: float32(val.Get("x").ToFloat()),
					Y: float32(val.Get("y").ToFloat()),
				})
			}
		}
		return goja.Undefined()
	})

	imguiObj.Set("popStyleVar", func(call goja.FunctionCall) goja.Value {
		var count int32 = 1
		if len(call.Arguments) > 0 && !goja.IsUndefined(call.Arguments[0]) {
			count = int32(call.Argument(0).ToInteger())
		}
		imgui.PopStyleVarV(count)
		return goja.Undefined()
	})

	// ==================== 字体 ====================

	imguiObj.Set("pushFont", func(call goja.FunctionCall) goja.Value {
		// 简化实现，需要字体管理支持
		return goja.Undefined()
	})

	imguiObj.Set("popFont", func(call goja.FunctionCall) goja.Value {
		imgui.PopFont()
		return goja.Undefined()
	})

	// ==================== 窗口设置 ====================

	imguiObj.Set("setNextWindowPos", func(call goja.FunctionCall) goja.Value {
		pos := call.Argument(0).ToObject(vm)
		var cond imgui.Cond
		var pivot imgui.Vec2

		if len(call.Arguments) > 1 && !goja.IsUndefined(call.Arguments[1]) {
			cond = imgui.Cond(call.Argument(1).ToInteger())
		}
		if len(call.Arguments) > 2 && !goja.IsUndefined(call.Arguments[2]) {
			pv := call.Argument(2).ToObject(vm)
			pivot = imgui.Vec2{
				X: float32(pv.Get("x").ToFloat()),
				Y: float32(pv.Get("y").ToFloat()),
			}
		}
		imgui.SetNextWindowPosV(imgui.Vec2{
			X: float32(pos.Get("x").ToFloat()),
			Y: float32(pos.Get("y").ToFloat()),
		}, cond, pivot)
		return goja.Undefined()
	})

	imguiObj.Set("setNextWindowSize", func(call goja.FunctionCall) goja.Value {
		size := call.Argument(0).ToObject(vm)
		var cond imgui.Cond

		if len(call.Arguments) > 1 && !goja.IsUndefined(call.Arguments[1]) {
			cond = imgui.Cond(call.Argument(1).ToInteger())
		}
		imgui.SetNextWindowSizeV(imgui.Vec2{
			X: float32(size.Get("x").ToFloat()),
			Y: float32(size.Get("y").ToFloat()),
		}, cond)
		return goja.Undefined()
	})

	imguiObj.Set("setNextWindowSizeConstraints", func(call goja.FunctionCall) goja.Value {
		size_min := call.Argument(0).ToObject(vm)
		size_max := call.Argument(1).ToObject(vm)
		imgui.SetNextWindowSizeConstraintsV(
			imgui.Vec2{X: float32(size_min.Get("x").ToFloat()), Y: float32(size_min.Get("y").ToFloat())},
			imgui.Vec2{X: float32(size_max.Get("x").ToFloat()), Y: float32(size_max.Get("y").ToFloat())},
			nil, 0,
		)
		return goja.Undefined()
	})

	imguiObj.Set("setNextWindowContentSize", func(call goja.FunctionCall) goja.Value {
		size := call.Argument(0).ToObject(vm)
		imgui.SetNextWindowContentSize(imgui.Vec2{
			X: float32(size.Get("x").ToFloat()),
			Y: float32(size.Get("y").ToFloat()),
		})
		return goja.Undefined()
	})

	imguiObj.Set("setNextWindowCollapsed", func(call goja.FunctionCall) goja.Value {
		collapsed := call.Argument(0).ToBoolean()
		var cond imgui.Cond
		if len(call.Arguments) > 1 && !goja.IsUndefined(call.Arguments[1]) {
			cond = imgui.Cond(call.Argument(1).ToInteger())
		}
		imgui.SetNextWindowCollapsedV(collapsed, cond)
		return goja.Undefined()
	})

	imguiObj.Set("setNextWindowFocus", func(call goja.FunctionCall) goja.Value {
		imgui.SetNextWindowFocus()
		return goja.Undefined()
	})

	imguiObj.Set("setWindowPos", func(call goja.FunctionCall) goja.Value {
		pos := call.Argument(0).ToObject(vm)
		var cond imgui.Cond
		if len(call.Arguments) > 1 && !goja.IsUndefined(call.Arguments[1]) {
			cond = imgui.Cond(call.Argument(1).ToInteger())
		}
		imgui.SetWindowPosVec2V(imgui.Vec2{
			X: float32(pos.Get("x").ToFloat()),
			Y: float32(pos.Get("y").ToFloat()),
		}, cond)
		return goja.Undefined()
	})

	imguiObj.Set("setWindowSize", func(call goja.FunctionCall) goja.Value {
		size := call.Argument(0).ToObject(vm)
		var cond imgui.Cond
		if len(call.Arguments) > 1 && !goja.IsUndefined(call.Arguments[1]) {
			cond = imgui.Cond(call.Argument(1).ToInteger())
		}
		imgui.SetWindowSizeVec2V(imgui.Vec2{
			X: float32(size.Get("x").ToFloat()),
			Y: float32(size.Get("y").ToFloat()),
		}, cond)
		return goja.Undefined()
	})

	imguiObj.Set("setWindowCollapsed", func(call goja.FunctionCall) goja.Value {
		collapsed := call.Argument(0).ToBoolean()
		var cond imgui.Cond
		if len(call.Arguments) > 1 && !goja.IsUndefined(call.Arguments[1]) {
			cond = imgui.Cond(call.Argument(1).ToInteger())
		}
		imgui.SetWindowCollapsedBoolV(collapsed, cond)
		return goja.Undefined()
	})

	imguiObj.Set("setWindowFocus", func(call goja.FunctionCall) goja.Value {
		imgui.SetWindowFocus()
		return goja.Undefined()
	})

	// ==================== 窗口信息获取 ====================

	imguiObj.Set("getWindowPos", func(call goja.FunctionCall) goja.Value {
		pos := imgui.WindowPos()
		return vm.ToValue(map[string]float32{"x": pos.X, "y": pos.Y})
	})

	imguiObj.Set("getWindowSize", func(call goja.FunctionCall) goja.Value {
		size := imgui.WindowSize()
		return vm.ToValue(map[string]float32{"x": size.X, "y": size.Y})
	})

	imguiObj.Set("getWindowWidth", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.WindowWidth())
	})

	imguiObj.Set("getWindowHeight", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.WindowHeight())
	})

	// ==================== 内容区域 ====================

	imguiObj.Set("getContentRegionAvail", func(call goja.FunctionCall) goja.Value {
		size := imgui.ContentRegionAvail()
		return vm.ToValue(map[string]float32{"x": size.X, "y": size.Y})
	})

	// ==================== 项目状态查询 ====================

	imguiObj.Set("isItemHovered", func(call goja.FunctionCall) goja.Value {
		var flags imgui.HoveredFlags
		if len(call.Arguments) > 0 && !goja.IsUndefined(call.Arguments[0]) {
			flags = imgui.HoveredFlags(call.Argument(0).ToInteger())
		}
		return vm.ToValue(imgui.IsItemHoveredV(flags))
	})

	imguiObj.Set("isItemActive", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.IsItemActive())
	})

	imguiObj.Set("isItemFocused", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.IsItemFocused())
	})

	imguiObj.Set("isItemClicked", func(call goja.FunctionCall) goja.Value {
		var mouse_button imgui.MouseButton
		if len(call.Arguments) > 0 && !goja.IsUndefined(call.Arguments[0]) {
			mouse_button = imgui.MouseButton(call.Argument(0).ToInteger())
		}
		return vm.ToValue(imgui.IsItemClickedV(mouse_button))
	})

	imguiObj.Set("isItemVisible", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.IsItemVisible())
	})

	imguiObj.Set("isAnyItemHovered", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.IsAnyItemHovered())
	})

	imguiObj.Set("isWindowHovered", func(call goja.FunctionCall) goja.Value {
		var flags imgui.HoveredFlags
		if len(call.Arguments) > 0 && !goja.IsUndefined(call.Arguments[0]) {
			flags = imgui.HoveredFlags(call.Argument(0).ToInteger())
		}
		return vm.ToValue(imgui.IsWindowHoveredV(flags))
	})

	imguiObj.Set("isWindowFocused", func(call goja.FunctionCall) goja.Value {
		var flags imgui.FocusedFlags
		if len(call.Arguments) > 0 && !goja.IsUndefined(call.Arguments[0]) {
			flags = imgui.FocusedFlags(call.Argument(0).ToInteger())
		}
		return vm.ToValue(imgui.IsWindowFocusedV(flags))
	})

	// ==================== 鼠标状态 ====================

	imguiObj.Set("isMouseDown", func(call goja.FunctionCall) goja.Value {
		button := imgui.MouseButton(call.Argument(0).ToInteger())
		return vm.ToValue(imgui.IsMouseDown(button))
	})

	imguiObj.Set("isMouseClicked", func(call goja.FunctionCall) goja.Value {
		button := imgui.MouseButton(call.Argument(0).ToInteger())
		var repeat bool
		if len(call.Arguments) > 1 && !goja.IsUndefined(call.Arguments[1]) {
			repeat = call.Argument(1).ToBoolean()
		}
		return vm.ToValue(imgui.IsMouseClickedBoolV(button, repeat))
	})

	imguiObj.Set("isMouseReleased", func(call goja.FunctionCall) goja.Value {
		button := imgui.MouseButton(call.Argument(0).ToInteger())
		return vm.ToValue(imgui.IsMouseReleased(button))
	})

	imguiObj.Set("isMouseDoubleClicked", func(call goja.FunctionCall) goja.Value {
		button := imgui.MouseButton(call.Argument(0).ToInteger())
		return vm.ToValue(imgui.IsMouseDoubleClicked(button))
	})

	imguiObj.Set("isMouseDragging", func(call goja.FunctionCall) goja.Value {
		button := imgui.MouseButton(call.Argument(0).ToInteger())
		var lock_threshold float32
		if len(call.Arguments) > 1 && !goja.IsUndefined(call.Arguments[1]) {
			lock_threshold = float32(call.Argument(1).ToFloat())
		}
		return vm.ToValue(imgui.IsMouseDraggingV(button, lock_threshold))
	})

	imguiObj.Set("isMouseHoveringRect", func(call goja.FunctionCall) goja.Value {
		r_min := call.Argument(0).ToObject(vm)
		r_max := call.Argument(1).ToObject(vm)
		var clip bool = true
		if len(call.Arguments) > 2 && !goja.IsUndefined(call.Arguments[2]) {
			clip = call.Argument(2).ToBoolean()
		}
		return vm.ToValue(imgui.IsMouseHoveringRectV(
			imgui.Vec2{X: float32(r_min.Get("x").ToFloat()), Y: float32(r_min.Get("y").ToFloat())},
			imgui.Vec2{X: float32(r_max.Get("x").ToFloat()), Y: float32(r_max.Get("y").ToFloat())},
			clip,
		))
	})

	imguiObj.Set("getMousePos", func(call goja.FunctionCall) goja.Value {
		io := imgui.CurrentIO()
		pos := io.MousePos()
		return vm.ToValue(map[string]float32{"x": pos.X, "y": pos.Y})
	})

	// ==================== 键盘状态 ====================

	imguiObj.Set("isKeyDown", func(call goja.FunctionCall) goja.Value {
		key := imgui.Key(call.Argument(0).ToInteger())
		return vm.ToValue(imgui.IsKeyDown(key))
	})

	imguiObj.Set("isKeyPressed", func(call goja.FunctionCall) goja.Value {
		key := imgui.Key(call.Argument(0).ToInteger())
		var repeat bool = true
		if len(call.Arguments) > 1 && !goja.IsUndefined(call.Arguments[1]) {
			repeat = call.Argument(1).ToBoolean()
		}
		return vm.ToValue(imgui.IsKeyPressedBoolV(key, repeat))
	})

	imguiObj.Set("isKeyReleased", func(call goja.FunctionCall) goja.Value {
		key := imgui.Key(call.Argument(0).ToInteger())
		return vm.ToValue(imgui.IsKeyReleased(key))
	})

	// ==================== 焦点控制 ====================

	imguiObj.Set("setItemDefaultFocus", func(call goja.FunctionCall) goja.Value {
		imgui.SetItemDefaultFocus()
		return goja.Undefined()
	})

	imguiObj.Set("setKeyboardFocusHere", func(call goja.FunctionCall) goja.Value {
		var offset int32
		if len(call.Arguments) > 0 && !goja.IsUndefined(call.Arguments[0]) {
			offset = int32(call.Argument(0).ToInteger())
		}
		imgui.SetKeyboardFocusHereV(offset)
		return goja.Undefined()
	})

	// ==================== 时间和帧计数 ====================

	imguiObj.Set("getTime", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.Time())
	})

	imguiObj.Set("getFrameCount", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.FrameCount())
	})

	// ==================== 字体相关 ====================

	imguiObj.Set("getFontSize", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.FontSize())
	})

	imguiObj.Set("calcTextSize", func(call goja.FunctionCall) goja.Value {
		text := call.Argument(0).String()
		var hide_text_after_double_hash bool
		var wrap_width float32

		if len(call.Arguments) > 1 && !goja.IsUndefined(call.Arguments[1]) {
			hide_text_after_double_hash = call.Argument(1).ToBoolean()
		}
		if len(call.Arguments) > 2 && !goja.IsUndefined(call.Arguments[2]) {
			wrap_width = float32(call.Argument(2).ToFloat())
		}
		size := imgui.CalcTextSizeV(text, hide_text_after_double_hash, wrap_width)
		return vm.ToValue(map[string]float32{"x": size.X, "y": size.Y})
	})

	// ==================== 文本行高度 ====================

	imguiObj.Set("getTextLineHeight", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.TextLineHeight())
	})

	imguiObj.Set("getTextLineHeightWithSpacing", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.TextLineHeightWithSpacing())
	})

	imguiObj.Set("getFrameHeight", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.FrameHeight())
	})

	imguiObj.Set("getFrameHeightWithSpacing", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.FrameHeightWithSpacing())
	})

	// ==================== 鼠标光标 ====================

	imguiObj.Set("setMouseCursor", func(call goja.FunctionCall) goja.Value {
		cursor_type := imgui.MouseCursor(call.Argument(0).ToInteger())
		imgui.SetMouseCursor(cursor_type)
		return goja.Undefined()
	})

	// ==================== 颜色转换 ====================

	imguiObj.Set("colorConvertU32ToFloat4", func(call goja.FunctionCall) goja.Value {
		in := uint32(call.Argument(0).ToInteger())
		col := imgui.ColorConvertU32ToFloat4(in)
		return vm.ToValue(map[string]float32{"x": col.X, "y": col.Y, "z": col.Z, "w": col.W})
	})

	imguiObj.Set("colorConvertFloat4ToU32", func(call goja.FunctionCall) goja.Value {
		col := call.Argument(0).ToObject(vm)
		in := imgui.Vec4{
			X: float32(col.Get("x").ToFloat()),
			Y: float32(col.Get("y").ToFloat()),
			Z: float32(col.Get("z").ToFloat()),
			W: float32(col.Get("w").ToFloat()),
		}
		return vm.ToValue(imgui.ColorConvertFloat4ToU32(in))
	})

	imguiObj.Set("colorConvertRGBtoHSV", func(call goja.FunctionCall) goja.Value {
		r := float32(call.Argument(0).ToFloat())
		g := float32(call.Argument(1).ToFloat())
		b := float32(call.Argument(2).ToFloat())
		var h, s, v float32
		imgui.ColorConvertRGBtoHSV(r, g, b, &h, &s, &v)
		return vm.ToValue(map[string]float32{"h": h, "s": s, "v": v})
	})

	imguiObj.Set("colorConvertHSVtoRGB", func(call goja.FunctionCall) goja.Value {
		h := float32(call.Argument(0).ToFloat())
		s := float32(call.Argument(1).ToFloat())
		v := float32(call.Argument(2).ToFloat())
		var r, g, b float32
		imgui.ColorConvertHSVtoRGB(h, s, v, &r, &g, &b)
		return vm.ToValue(map[string]float32{"r": r, "g": g, "b": b})
	})

	// ==================== ID 操作 ====================

	imguiObj.Set("pushID", func(call goja.FunctionCall) goja.Value {
		str_id := call.Argument(0).String()
		imgui.PushIDStr(str_id)
		return goja.Undefined()
	})

	imguiObj.Set("popID", func(call goja.FunctionCall) goja.Value {
		imgui.PopID()
		return goja.Undefined()
	})

	imguiObj.Set("getID", func(call goja.FunctionCall) goja.Value {
		str_id := call.Argument(0).String()
		return vm.ToValue(uint(imgui.IDStr(str_id)))
	})

	// ==================== 列布局（旧版） ====================

	imguiObj.Set("columns", func(call goja.FunctionCall) goja.Value {
		var count int32 = 1
		var id string
		var border bool = true

		if len(call.Arguments) > 0 && !goja.IsUndefined(call.Arguments[0]) {
			count = int32(call.Argument(0).ToInteger())
		}
		if len(call.Arguments) > 1 && !goja.IsUndefined(call.Arguments[1]) {
			id = call.Argument(1).String()
		}
		if len(call.Arguments) > 2 && !goja.IsUndefined(call.Arguments[2]) {
			border = call.Argument(2).ToBoolean()
		}
		imgui.ColumnsV(count, id, border)
		return goja.Undefined()
	})

	imguiObj.Set("nextColumn", func(call goja.FunctionCall) goja.Value {
		imgui.NextColumn()
		return goja.Undefined()
	})

	imguiObj.Set("getColumnIndex", func(call goja.FunctionCall) goja.Value {
		return vm.ToValue(imgui.ColumnIndex())
	})

	imguiObj.Set("getColumnWidth", func(call goja.FunctionCall) goja.Value {
		var column_index int32 = -1
		if len(call.Arguments) > 0 && !goja.IsUndefined(call.Arguments[0]) {
			column_index = int32(call.Argument(0).ToInteger())
		}
		return vm.ToValue(imgui.ColumnWidthV(column_index))
	})

	imguiObj.Set("setColumnWidth", func(call goja.FunctionCall) goja.Value {
		column_index := int32(call.Argument(0).ToInteger())
		width := float32(call.Argument(1).ToFloat())
		imgui.SetColumnWidth(column_index, width)
		return goja.Undefined()
	})

	// ==================== 常量定义 ====================

	// 窗口标志
	imguiObj.Set("WindowFlags", map[string]int{
		"None":                    int(imgui.WindowFlagsNone),
		"NoTitleBar":              int(imgui.WindowFlagsNoTitleBar),
		"NoResize":                int(imgui.WindowFlagsNoResize),
		"NoMove":                  int(imgui.WindowFlagsNoMove),
		"NoScrollbar":             int(imgui.WindowFlagsNoScrollbar),
		"NoScrollWithMouse":       int(imgui.WindowFlagsNoScrollWithMouse),
		"NoCollapse":              int(imgui.WindowFlagsNoCollapse),
		"AlwaysAutoResize":        int(imgui.WindowFlagsAlwaysAutoResize),
		"NoBackground":            int(imgui.WindowFlagsNoBackground),
		"NoSavedSettings":         int(imgui.WindowFlagsNoSavedSettings),
		"NoMouseInputs":           int(imgui.WindowFlagsNoMouseInputs),
		"MenuBar":                 int(imgui.WindowFlagsMenuBar),
		"HorizontalScrollbar":     int(imgui.WindowFlagsHorizontalScrollbar),
		"NoFocusOnAppearing":      int(imgui.WindowFlagsNoFocusOnAppearing),
		"NoBringToFrontOnFocus":   int(imgui.WindowFlagsNoBringToFrontOnFocus),
		"AlwaysVerticalScrollbar": int(imgui.WindowFlagsAlwaysVerticalScrollbar),
		"AlwaysHorizontalScrollbar": int(imgui.WindowFlagsAlwaysHorizontalScrollbar),
		"NoNavInputs":             int(imgui.WindowFlagsNoNavInputs),
		"NoNavFocus":              int(imgui.WindowFlagsNoNavFocus),
		"UnsavedDocument":         int(imgui.WindowFlagsUnsavedDocument),
		"NoDocking":               int(imgui.WindowFlagsNoDocking),
	})

	// 颜色索引
	imguiObj.Set("Col", map[string]int{
		"Text":                  int(imgui.ColText),
		"TextDisabled":          int(imgui.ColTextDisabled),
		"WindowBg":              int(imgui.ColWindowBg),
		"ChildBg":               int(imgui.ColChildBg),
		"PopupBg":               int(imgui.ColPopupBg),
		"Border":                int(imgui.ColBorder),
		"BorderShadow":          int(imgui.ColBorderShadow),
		"FrameBg":               int(imgui.ColFrameBg),
		"FrameBgHovered":        int(imgui.ColFrameBgHovered),
		"FrameBgActive":         int(imgui.ColFrameBgActive),
		"TitleBg":               int(imgui.ColTitleBg),
		"TitleBgActive":         int(imgui.ColTitleBgActive),
		"TitleBgCollapsed":      int(imgui.ColTitleBgCollapsed),
		"MenuBarBg":             int(imgui.ColMenuBarBg),
		"ScrollbarBg":           int(imgui.ColScrollbarBg),
		"ScrollbarGrab":         int(imgui.ColScrollbarGrab),
		"ScrollbarGrabHovered":  int(imgui.ColScrollbarGrabHovered),
		"ScrollbarGrabActive":   int(imgui.ColScrollbarGrabActive),
		"CheckMark":             int(imgui.ColCheckMark),
		"SliderGrab":            int(imgui.ColSliderGrab),
		"SliderGrabActive":      int(imgui.ColSliderGrabActive),
		"Button":                int(imgui.ColButton),
		"ButtonHovered":         int(imgui.ColButtonHovered),
		"ButtonActive":          int(imgui.ColButtonActive),
		"Header":                int(imgui.ColHeader),
		"HeaderHovered":         int(imgui.ColHeaderHovered),
		"HeaderActive":          int(imgui.ColHeaderActive),
		"Separator":             int(imgui.ColSeparator),
		"SeparatorHovered":      int(imgui.ColSeparatorHovered),
		"SeparatorActive":       int(imgui.ColSeparatorActive),
		"ResizeGrip":            int(imgui.ColResizeGrip),
		"ResizeGripHovered":     int(imgui.ColResizeGripHovered),
		"ResizeGripActive":      int(imgui.ColResizeGripActive),
		"Tab":                   int(imgui.ColTab),
		"TabHovered":            int(imgui.ColTabHovered),
		"TabSelected":           int(imgui.ColTabSelected),
		"TabSelectedOverline":   int(imgui.ColTabSelectedOverline),
		"TabDimmed":             int(imgui.ColTabDimmed),
		"TabDimmedSelected":     int(imgui.ColTabDimmedSelected),
		"TabDimmedSelectedOverline": int(imgui.ColTabDimmedSelectedOverline),
		"PlotLines":             int(imgui.ColPlotLines),
		"PlotLinesHovered":      int(imgui.ColPlotLinesHovered),
		"PlotHistogram":         int(imgui.ColPlotHistogram),
		"PlotHistogramHovered":  int(imgui.ColPlotHistogramHovered),
		"TableHeaderBg":         int(imgui.ColTableHeaderBg),
		"TableBorderStrong":     int(imgui.ColTableBorderStrong),
		"TableBorderLight":      int(imgui.ColTableBorderLight),
		"TableRowBg":            int(imgui.ColTableRowBg),
		"TableRowBgAlt":         int(imgui.ColTableRowBgAlt),
		"TextSelectedBg":        int(imgui.ColTextSelectedBg),
		"DragDropTarget":        int(imgui.ColDragDropTarget),
		"NavWindowingDimBg":     int(imgui.ColNavWindowingDimBg),
		"ModalWindowDimBg":      int(imgui.ColModalWindowDimBg),
	})

	// 鼠标按钮
	imguiObj.Set("MouseButton", map[string]int{
		"Left":   int(imgui.MouseButtonLeft),
		"Right":  int(imgui.MouseButtonRight),
		"Middle": int(imgui.MouseButtonMiddle),
	})

	// 键盘键
	imguiObj.Set("Key", map[string]int{
		"None":          int(imgui.KeyNone),
		"Tab":           int(imgui.KeyTab),
		"LeftArrow":     int(imgui.KeyLeftArrow),
		"RightArrow":    int(imgui.KeyRightArrow),
		"UpArrow":       int(imgui.KeyUpArrow),
		"DownArrow":     int(imgui.KeyDownArrow),
		"PageUp":        int(imgui.KeyPageUp),
		"PageDown":      int(imgui.KeyPageDown),
		"Home":          int(imgui.KeyHome),
		"End":           int(imgui.KeyEnd),
		"Insert":        int(imgui.KeyInsert),
		"Delete":        int(imgui.KeyDelete),
		"Backspace":     int(imgui.KeyBackspace),
		"Space":         int(imgui.KeySpace),
		"Enter":         int(imgui.KeyEnter),
		"Escape":        int(imgui.KeyEscape),
		"LeftCtrl":      int(imgui.KeyLeftCtrl),
		"LeftShift":     int(imgui.KeyLeftShift),
		"LeftAlt":       int(imgui.KeyLeftAlt),
		"LeftSuper":     int(imgui.KeyLeftSuper),
		"RightCtrl":     int(imgui.KeyRightCtrl),
		"RightShift":    int(imgui.KeyRightShift),
		"RightAlt":      int(imgui.KeyRightAlt),
		"RightSuper":    int(imgui.KeyRightSuper),
		"Menu":          int(imgui.KeyMenu),
		"0":             int(imgui.Key0),
		"1":             int(imgui.Key1),
		"2":             int(imgui.Key2),
		"3":             int(imgui.Key3),
		"4":             int(imgui.Key4),
		"5":             int(imgui.Key5),
		"6":             int(imgui.Key6),
		"7":             int(imgui.Key7),
		"8":             int(imgui.Key8),
		"9":             int(imgui.Key9),
		"A":             int(imgui.KeyA),
		"B":             int(imgui.KeyB),
		"C":             int(imgui.KeyC),
		"D":             int(imgui.KeyD),
		"E":             int(imgui.KeyE),
		"F":             int(imgui.KeyF),
		"G":             int(imgui.KeyG),
		"H":             int(imgui.KeyH),
		"I":             int(imgui.KeyI),
		"J":             int(imgui.KeyJ),
		"K":             int(imgui.KeyK),
		"L":             int(imgui.KeyL),
		"M":             int(imgui.KeyM),
		"N":             int(imgui.KeyN),
		"O":             int(imgui.KeyO),
		"P":             int(imgui.KeyP),
		"Q":             int(imgui.KeyQ),
		"R":             int(imgui.KeyR),
		"S":             int(imgui.KeyS),
		"T":             int(imgui.KeyT),
		"U":             int(imgui.KeyU),
		"V":             int(imgui.KeyV),
		"W":             int(imgui.KeyW),
		"X":             int(imgui.KeyX),
		"Y":             int(imgui.KeyY),
		"Z":             int(imgui.KeyZ),
		"F1":            int(imgui.KeyF1),
		"F2":            int(imgui.KeyF2),
		"F3":            int(imgui.KeyF3),
		"F4":            int(imgui.KeyF4),
		"F5":            int(imgui.KeyF5),
		"F6":            int(imgui.KeyF6),
		"F7":            int(imgui.KeyF7),
		"F8":            int(imgui.KeyF8),
		"F9":            int(imgui.KeyF9),
		"F10":           int(imgui.KeyF10),
		"F11":           int(imgui.KeyF11),
		"F12":           int(imgui.KeyF12),
	})

	// 条件
	imguiObj.Set("Cond", map[string]int{
		"None":        int(imgui.CondNone),
		"Always":      int(imgui.CondAlways),
		"Once":        int(imgui.CondOnce),
		"FirstUseEver": int(imgui.CondFirstUseEver),
		"Appearing":   int(imgui.CondAppearing),
	})

	// 注册到全局对象
	vm.Set("imgui", imguiObj)
}
