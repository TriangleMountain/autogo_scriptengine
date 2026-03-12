package lua_engine

import (
	"github.com/Dasongzi1366/AutoGo/device"
	lua "github.com/yuin/gopher-lua"
)

func injectDeviceMethods(engine *LuaEngine) {

	engine.RegisterMethod("device.width", "设备分辨率宽度", func(displayId int) int {
		width, _, _, _ := device.GetDisplayInfo(displayId)
		return width
	}, true)
	engine.RegisterMethod("device.height", "设备分辨率高度", func(displayId int) int {
		_, height, _, _ := device.GetDisplayInfo(displayId)
		return height
	}, true)
	engine.RegisterMethod("device.sdkInt", "安卓系统API版本", func() int { return device.SdkInt }, true)
	engine.RegisterMethod("device.cpuAbi", "设备的CPU架构", func() string { return device.CpuAbi }, true)
	engine.RegisterMethod("device.buildId", "修订版本号", func() string { return device.BuildId }, true)
	engine.RegisterMethod("device.broad", "设备的主板型号", func() string { return device.Broad }, true)
	engine.RegisterMethod("device.brand", "与产品或硬件相关的厂商品牌", func() string { return device.Brand }, true)
	engine.RegisterMethod("device.device", "设备在工业设计中的名称", func() string { return device.Device }, true)
	engine.RegisterMethod("device.model", "设备型号", func() string { return device.Model }, true)
	engine.RegisterMethod("device.product", "整个产品的名称", func() string { return device.Product }, true)
	engine.RegisterMethod("device.bootloader", "设备Bootloader的版本", func() string { return device.Bootloader }, true)
	engine.RegisterMethod("device.hardware", "设备的硬件名称", func() string { return device.Hardware }, true)
	engine.RegisterMethod("device.fingerprint", "构建的唯一标识码", func() string { return device.Fingerprint }, true)
	engine.RegisterMethod("device.serial", "硬件序列号", func() string { return device.Serial }, true)
	engine.RegisterMethod("device.incremental", "设备构建的内部版本号", func() string { return device.Incremental }, true)
	engine.RegisterMethod("device.release", "Android系统版本号", func() string { return device.Release }, true)
	engine.RegisterMethod("device.baseOS", "设备的基础操作系统版本", func() string { return device.BaseOS }, true)
	engine.RegisterMethod("device.securityPatch", "安全补丁程序级别", func() string { return device.SecurityPatch }, true)
	engine.RegisterMethod("device.codename", "开发代号", func() string { return device.Codename }, true)
	engine.RegisterMethod("device.getImei", "返回设备的IMEI", device.GetImei, true)
	engine.RegisterMethod("device.getAndroidId", "返回设备的Android ID", device.GetAndroidId, true)
	engine.RegisterMethod("device.getWifiMac", "获取设备WIFI-MAC", device.GetWifiMac, true)
	engine.RegisterMethod("device.getWlanMac", "获取设备以太网MAC", device.GetWlanMac, true)
	engine.RegisterMethod("device.getIp", "获取设备局域网IP地址", device.GetIp, true)
	engine.RegisterMethod("device.getBrightness", "返回当前的(手动)亮度", device.GetBrightness, true)
	engine.RegisterMethod("device.getBrightnessMode", "返回当前亮度模式", device.GetBrightnessMode, true)
	engine.RegisterMethod("device.getMusicVolume", "返回当前媒体音量", device.GetMusicVolume, true)
	engine.RegisterMethod("device.getNotificationVolume", "返回当前通知音量", device.GetNotificationVolume, true)
	engine.RegisterMethod("device.getAlarmVolume", "返回当前闹钟音量", device.GetAlarmVolume, true)
	engine.RegisterMethod("device.getMusicMaxVolume", "返回媒体音量的最大值", device.GetMusicMaxVolume, true)
	engine.RegisterMethod("device.getNotificationMaxVolume", "返回通知音量的最大值", device.GetNotificationMaxVolume, true)
	engine.RegisterMethod("device.getAlarmMaxVolume", "返回闹钟音量的最大值", device.GetAlarmMaxVolume, true)
	engine.RegisterMethod("device.setMusicVolume", "设置当前媒体音量", func(volume int) { device.SetMusicVolume(volume) }, true)
	engine.RegisterMethod("device.setNotificationVolume", "设置当前通知音量", func(volume int) { device.SetNotificationVolume(volume) }, true)
	engine.RegisterMethod("device.setAlarmVolume", "设置当前闹钟音量", func(volume int) { device.SetAlarmVolume(volume) }, true)
	engine.RegisterMethod("device.getBattery", "返回当前电量百分比", device.GetBattery, true)
	engine.RegisterMethod("device.getBatteryStatus", "获取电池状态", device.GetBatteryStatus, true)
	engine.RegisterMethod("device.setBatteryStatus", "模拟电池状态", func(value int) { device.SetBatteryStatus(value) }, true)
	engine.RegisterMethod("device.setBatteryLevel", "模拟电池电量百分百", func(value int) { device.SetBatteryLevel(value) }, true)
	engine.RegisterMethod("device.getTotalMem", "返回设备内存总量", device.GetTotalMem, true)
	engine.RegisterMethod("device.getAvailMem", "返回设备当前可用的内存", device.GetAvailMem, true)
	engine.RegisterMethod("device.isScreenOn", "返回设备屏幕是否是亮着的", device.IsScreenOn, true)
	engine.RegisterMethod("device.isScreenUnlock", "返回屏幕锁是否已经解开", device.IsScreenUnlock, true)
	engine.RegisterMethod("device.wakeUp", "唤醒设备", device.WakeUp, true)
	engine.RegisterMethod("device.keepScreenOn", "保持屏幕常亮", device.KeepScreenOn, true)
	engine.RegisterMethod("device.vibrate", "使设备震动一段时间", func(ms int) { device.Vibrate(ms) }, true)
	engine.RegisterMethod("device.cancelVibration", "如果设备处于震动状态，则取消震动", device.CancelVibration, true)

	registerDeviceLuaFunctions(engine)
}

func registerDeviceLuaFunctions(engine *LuaEngine) {
	state := engine.GetState()

	state.Register("device_getImei", func(L *lua.LState) int {
		result := device.GetImei()
		L.Push(lua.LString(result))
		return 1
	})

	state.Register("device_getAndroidId", func(L *lua.LState) int {
		result := device.GetAndroidId()
		L.Push(lua.LString(result))
		return 1
	})

	state.Register("device_getWifiMac", func(L *lua.LState) int {
		result := device.GetWifiMac()
		L.Push(lua.LString(result))
		return 1
	})

	state.Register("device_getWlanMac", func(L *lua.LState) int {
		result := device.GetWlanMac()
		L.Push(lua.LString(result))
		return 1
	})

	state.Register("device_getIp", func(L *lua.LState) int {
		result := device.GetIp()
		L.Push(lua.LString(result))
		return 1
	})

	state.Register("device_getBrightness", func(L *lua.LState) int {
		result := device.GetBrightness()
		L.Push(lua.LString(result))
		return 1
	})

	state.Register("device_getBrightnessMode", func(L *lua.LState) int {
		result := device.GetBrightnessMode()
		L.Push(lua.LString(result))
		return 1
	})

	state.Register("device_getMusicVolume", func(L *lua.LState) int {
		result := device.GetMusicVolume()
		L.Push(lua.LNumber(result))
		return 1
	})

	state.Register("device_getNotificationVolume", func(L *lua.LState) int {
		result := device.GetNotificationVolume()
		L.Push(lua.LNumber(result))
		return 1
	})

	state.Register("device_getAlarmVolume", func(L *lua.LState) int {
		result := device.GetAlarmVolume()
		L.Push(lua.LNumber(result))
		return 1
	})

	state.Register("device_getMusicMaxVolume", func(L *lua.LState) int {
		result := device.GetMusicMaxVolume()
		L.Push(lua.LNumber(result))
		return 1
	})

	state.Register("device_getNotificationMaxVolume", func(L *lua.LState) int {
		result := device.GetNotificationMaxVolume()
		L.Push(lua.LNumber(result))
		return 1
	})

	state.Register("device_getAlarmMaxVolume", func(L *lua.LState) int {
		result := device.GetAlarmMaxVolume()
		L.Push(lua.LNumber(result))
		return 1
	})

	state.Register("device_setMusicVolume", func(L *lua.LState) int {
		volume := L.CheckInt(1)
		device.SetMusicVolume(volume)
		return 0
	})

	state.Register("device_setNotificationVolume", func(L *lua.LState) int {
		volume := L.CheckInt(1)
		device.SetNotificationVolume(volume)
		return 0
	})

	state.Register("device_setAlarmVolume", func(L *lua.LState) int {
		volume := L.CheckInt(1)
		device.SetAlarmVolume(volume)
		return 0
	})

	state.Register("device_getBattery", func(L *lua.LState) int {
		result := device.GetBattery()
		L.Push(lua.LNumber(result))
		return 1
	})

	state.Register("device_getBatteryStatus", func(L *lua.LState) int {
		result := device.GetBatteryStatus()
		L.Push(lua.LNumber(result))
		return 1
	})

	state.Register("device_setBatteryStatus", func(L *lua.LState) int {
		value := L.CheckInt(1)
		device.SetBatteryStatus(value)
		return 0
	})

	state.Register("device_setBatteryLevel", func(L *lua.LState) int {
		value := L.CheckInt(1)
		device.SetBatteryLevel(value)
		return 0
	})

	state.Register("device_getTotalMem", func(L *lua.LState) int {
		result := device.GetTotalMem()
		L.Push(lua.LNumber(result))
		return 1
	})

	state.Register("device_getAvailMem", func(L *lua.LState) int {
		result := device.GetAvailMem()
		L.Push(lua.LNumber(result))
		return 1
	})

	state.Register("device_isScreenOn", func(L *lua.LState) int {
		result := device.IsScreenOn()
		L.Push(lua.LBool(result))
		return 1
	})

	state.Register("device_isScreenUnlock", func(L *lua.LState) int {
		result := device.IsScreenUnlock()
		L.Push(lua.LBool(result))
		return 1
	})

	state.Register("device_wakeUp", func(L *lua.LState) int {
		device.WakeUp()
		return 0
	})

	state.Register("device_keepScreenOn", func(L *lua.LState) int {
		device.KeepScreenOn()
		return 0
	})

	state.Register("device_vibrate", func(L *lua.LState) int {
		ms := L.CheckInt(1)
		device.Vibrate(ms)
		return 0
	})

	state.Register("device_cancelVibration", func(L *lua.LState) int {
		device.CancelVibration()
		return 0
	})
}
