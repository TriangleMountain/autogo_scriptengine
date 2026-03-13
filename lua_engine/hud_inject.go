package lua_engine

import (
	"github.com/Dasongzi1366/AutoGo/hud"
	lua "github.com/yuin/gopher-lua"
)

func injectHudMethods(engine *LuaEngine) {

	engine.RegisterMethod("hud.new", "创建一个新的HUD对象", hud.New, true)
	engine.RegisterMethod("hud.setPosition", "设置HUD的显示位置", func(h *hud.HUD, x1, y1, x2, y2 int) *hud.HUD {
		return h.SetPosition(x1, y1, x2, y2)
	}, true)
	engine.RegisterMethod("hud.setBackgroundColor", "设置HUD的背景颜色", func(h *hud.HUD, color string) *hud.HUD {
		return h.SetBackgroundColor(color)
	}, true)
	engine.RegisterMethod("hud.setTextSize", "设置HUD中文本的字体大小", func(h *hud.HUD, size int) *hud.HUD {
		return h.SetTextSize(size)
	}, true)
	engine.RegisterMethod("hud.setText", "设置HUD中要显示的文本内容", func(h *hud.HUD, items []hud.TextItem) *hud.HUD {
		return h.SetText(items)
	}, true)
	engine.RegisterMethod("hud.show", "显示HUD", func(h *hud.HUD) {
		h.Show()
	}, true)
	engine.RegisterMethod("hud.hide", "隐藏HUD", func(h *hud.HUD) {
		h.Hide()
	}, true)
	engine.RegisterMethod("hud.isVisible", "返回HUD是否可见", func(h *hud.HUD) bool {
		return h.IsVisible()
	}, true)
	engine.RegisterMethod("hud.destroy", "销毁HUD对象，释放资源", func(h *hud.HUD) {
		h.Destroy()
	}, true)

	registerHudLuaFunctions(engine)
}

func registerHudLuaFunctions(engine *LuaEngine) {
	state := engine.GetState()

	state.Register("hud_new", func(L *lua.LState) int {
		h := hud.New()
		ud := L.NewUserData()
		ud.Value = h
		L.Push(ud)
		return 1
	})

	state.Register("hud_setPosition", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		h := ud.Value.(*hud.HUD)
		x1 := L.CheckInt(2)
		y1 := L.CheckInt(3)
		x2 := L.CheckInt(4)
		y2 := L.CheckInt(5)
		h.SetPosition(x1, y1, x2, y2)
		L.Push(ud)
		return 1
	})

	state.Register("hud_setBackgroundColor", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		h := ud.Value.(*hud.HUD)
		color := L.CheckString(2)
		h.SetBackgroundColor(color)
		L.Push(ud)
		return 1
	})

	state.Register("hud_setTextSize", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		h := ud.Value.(*hud.HUD)
		size := L.CheckInt(2)
		h.SetTextSize(size)
		L.Push(ud)
		return 1
	})

	state.Register("hud_setText", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		h := ud.Value.(*hud.HUD)
		tbl := L.CheckTable(2)
		var items []hud.TextItem
		tbl.ForEach(func(k, v lua.LValue) {
			if t, ok := v.(*lua.LTable); ok {
				item := hud.TextItem{}
				item.TextColor = lua.LVAsString(t.RawGetString("TextColor"))
				item.Text = lua.LVAsString(t.RawGetString("Text"))
				items = append(items, item)
			}
		})
		h.SetText(items)
		L.Push(ud)
		return 1
	})

	state.Register("hud_show", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		h := ud.Value.(*hud.HUD)
		h.Show()
		return 0
	})

	state.Register("hud_hide", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		h := ud.Value.(*hud.HUD)
		h.Hide()
		return 0
	})

	state.Register("hud_isVisible", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		h := ud.Value.(*hud.HUD)
		result := h.IsVisible()
		L.Push(lua.LBool(result))
		return 1
	})

	state.Register("hud_destroy", func(L *lua.LState) int {
		ud := L.CheckUserData(1)
		h := ud.Value.(*hud.HUD)
		h.Destroy()
		return 0
	})
}
