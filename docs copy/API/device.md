# device - 设备
---
device模块提供了与设备有关的信息与操作，例如获取设备宽高，内存使用率，IMEI，调整设备亮度、音量等。

以下是 `device` 包中定义的设备信息变量：

| **变量名**         | **类型**   | **说明**                              |
|-----------------|----------|-------------------------------------|
| `CpuAbi`        | `string` | 设备的CPU架构，如"arm64-v8a", "x86", "x86_64"等。 |
| `BuildId`       | `string` | 修订版本号，或者诸如"M4-rc20"的标识。             |
| `Broad`         | `string` | 设备的主板型号。                            |
| `Brand`         | `string` | 与产品或硬件相关的厂商品牌，如"Xiaomi", "Huawei"等。 |
| `Device`        | `string` | 设备在工业设计中的名称。                        |
| `Model`         | `string` | 设备型号。                               |
| `Product`       | `string` | 整个产品的名称。                            |
| `Bootloader`    | `string` | 设备 Bootloader 的版本。                  |
| `Hardware`      | `string` | 设备的硬件名称。                            |
| `Fingerprint`   | `string` | 构建 (build) 的唯一标识码。                  |
| `Serial`        | `string` | 硬件序列号。                              |
| `SdkInt`        | `int`    | 安卓系统 API 版本。例如安卓 4.4 的 sdkInt 为 19。 |
| `Incremental`   | `string` | 设备构建的内部版本号。                         |
| `Release`       | `string` | Android 系统版本号。例如 "5.0", "7.1.1"。    |
| `BaseOS`        | `string` | 设备的基础操作系统版本。                        |
| `SecurityPatch` | `string` | 安全补丁程序级别。                           |
| `Codename`      | `string` | 开发代号，例如发行版是"REL"。                   |

## GetDisplayInfo
<hr style="margin: 0;">

获取指定屏幕的分辨率信息。

- `displayId` {int} 屏幕ID，0表示主屏幕，其他值表示虚拟屏幕

```go
width, height, dpi, rotation := device.GetDisplayInfo(0)
fmt.Printf("屏幕分辨率: %dx%d, DPI: %d, 旋转角度: %d\n", width, height, dpi, rotation)
```

## GetImei
<hr style="margin: 0;">

获取设备的 IMEI 码。

```go
imei := device.GetImei()
```

## GetAndroidId
<hr style="margin: 0;">

获取设备的 Android ID。

```go
androidId := device.GetAndroidId()
```

## GetWifiMac
<hr style="margin: 0;">

获取设备 WIFI 网卡的 MAC 地址。

```go
wifiMac := device.GetWifiMac()
```

## GetWlanMac
<hr style="margin: 0;">

获取设备以太网网卡的 MAC 地址。

```go
wlanMac := device.GetWlanMac()
```

## GetIp
<hr style="margin: 0;">

获取设备局域网 IP 地址。

```go
ip := device.GetIp()
```

## GetNotification
<hr style="margin: 0;">

获取设备当前所有通知消息。

```go
notifications := device.GetNotification()
for _, notification := range notifications {
	fmt.Println("ID:" + notification.Id)
	fmt.Println("包名:" + notification.PackageName)
	fmt.Println("标题:" + notification.Title)
	fmt.Println("内容:" + notification.Text)
	fmt.Println("标签:" + notification.Tag)
	fmt.Println()
}
```

## GetBrightness
<hr style="margin: 0;">

获取当前屏幕亮度值，范围为 0~255。

```go
brightness := device.GetBrightness()
```

## GetBrightnessMode
<hr style="margin: 0;">

获取当前屏幕亮度调节模式，0 为手动调节，1 为自动调节。

```go
mode := device.GetBrightnessMode()
```

## GetMusicVolume
<hr style="margin: 0;">

获取当前媒体音量。

```go
volume := device.GetMusicVolume()
```

## GetNotificationVolume
<hr style="margin: 0;">

获取当前通知音量。

```go
volume := device.GetNotificationVolume()
```

## GetAlarmVolume
<hr style="margin: 0;">

获取当前闹钟音量。

```go
volume := device.GetAlarmVolume()
```

## GetMusicMaxVolume
<hr style="margin: 0;">

获取媒体音量最大值。

```go
maxVolume := device.GetMusicMaxVolume()
```

## GetNotificationMaxVolume
<hr style="margin: 0;">

获取通知音量最大值。

```go
maxVolume := device.GetNotificationMaxVolume()
```

## GetAlarmMaxVolume
<hr style="margin: 0;">

获取闹钟音量最大值。

```go
maxVolume := device.GetAlarmMaxVolume()
```

## SetMusicVolume
<hr style="margin: 0;">

设置媒体音量。

- `volume` {int} 要设置的音量值

```go
device.SetMusicVolume(8)
```

## SetNotificationVolume
<hr style="margin: 0;">

设置通知音量。

- `volume` {int} 要设置的音量值

```go
device.SetNotificationVolume(8)
```

## SetAlarmVolume
<hr style="margin: 0;">

设置闹钟音量。

- `volume` {int} 要设置的音量值

```go
device.SetAlarmVolume(8)
```

## GetBattery
<hr style="margin: 0;">

获取当前电量百分比。

```go
battery := device.GetBattery()
```

## GetBatteryStatus
<hr style="margin: 0;">

获取电池状态。1：没有充电；2：正充电；3：没插充电器；4：不充电； 5：电池充满。

```go
status := device.GetBatteryStatus()
```

## SetBatteryStatus
<hr style="margin: 0;">

模拟设置电池状态。1：没有充电；2：正充电；5：电池充满。

- `value` {int} 要设置的电池状态

```go
device.SetBatteryStatus(2)
```

## SetBatteryLevel
<hr style="margin: 0;">

模拟设置电池电量百分比，范围 0-100。

- `value` {int} 要设置的电池电量百分比

```go
device.SetBatteryLevel(75)
```

## GetTotalMem
<hr style="margin: 0;">

获取设备总内存，单位 KB。

```go
totalMem := device.GetTotalMem()
```

## GetAvailMem
<hr style="margin: 0;">

获取设备当前可用内存，单位 KB。

```go
availMem := device.GetAvailMem()
```

## IsScreenOn
<hr style="margin: 0;">

判断屏幕是否点亮状态。

```go
isOn := device.IsScreenOn()
```

## IsScreenUnlock
<hr style="margin: 0;">

判断屏幕是否已解锁。

```go
isUnlock := device.IsScreenUnlock()
```

## SetDisplayPower
<hr style="margin: 0;">

设置屏幕电源模式，不影响脚本运行。

- `on` {bool} 是否点亮。

```go
device.SetDisplayPower(false)//熄屏挂机
```

## WakeUp
<hr style="margin: 0;">

唤醒设备，包括唤醒 CPU、屏幕等，可以用来点亮屏幕。

```go
device.WakeUp()
```

## Reboot
<hr style="margin: 0;">

重启设备。

```go
device.Reboot()
```

## KeepScreenOn
<hr style="margin: 0;">

保持屏幕常亮。

```go
device.KeepScreenOn()
```

## Vibrate
<hr style="margin: 0;">

使设备震动一段时间（单位毫秒，需要 root 权限）。

- `ms` {int} 要震动的时间（毫秒）。

```go
device.Vibrate(500)
```

## CancelVibration
<hr style="margin: 0;">

如果设备处于震动状态，则取消震动。

```go
device.CancelVibration()
``` 