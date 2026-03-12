-- AutoGo Lua API 文档
-- 自动生成时间: 2026-03-12

--------------------------------------------------
-- 应用管理 (app)
--------------------------------------------------
app_currentPackage() -> string
  获取当前页面应用包名

app_currentActivity() -> string
  获取当前页面应用类名

app_launch(packageName: string) -> boolean
  通过应用包名启动应用

app_openAppSetting(packageName: string) -> boolean
  打开应用的详情页(设置页)

app_viewFile(path: string)
  用其他应用查看文件

app_editFile(path: string)
  用其他应用编辑文件

app_uninstall(packageName: string)
  卸载应用

app_install(path: string)
  安装应用

app_isInstalled(packageName: string) -> boolean
  判断是否已经安装某个应用

app_clear(packageName: string)
  清除应用数据

app_forceStop(packageName: string)
  强制停止应用

app_disable(packageName: string)
  禁用应用

app_ignoreBattOpt(packageName: string)
  忽略电池优化

app_enable(packageName: string)
  启用应用

app_getBrowserPackage() -> string
  获取系统默认浏览器包名

app_openUrl(url: string)
  用浏览器打开网站url

--------------------------------------------------
-- 设备管理 (device)
--------------------------------------------------
device.width -> number (只读)
  设备分辨率宽度

device.height -> number (只读)
  设备分辨率高度

device.sdkInt -> number (只读)
  安卓系统API版本

device.cpuAbi -> string (只读)
  设备的CPU架构

device.getImei() -> string
  返回设备的IMEI

device.getAndroidId() -> string
  返回设备的Android ID

device.getWifiMac() -> string
  获取设备WIFI-MAC

device.getWlanMac() -> string
  获取设备以太网MAC

device.getIp() -> string
  获取设备局域网IP地址

device.getBrightness() -> string
  返回当前的(手动)亮度

device.getBrightnessMode() -> string
  返回当前亮度模式

device.getMusicVolume() -> number
  返回当前媒体音量

device.getNotificationVolume() -> number
  返回当前通知音量

device.getAlarmVolume() -> number
  返回当前闹钟音量

device.getMusicMaxVolume() -> number
  返回媒体音量的最大值

device.getNotificationMaxVolume() -> number
  返回通知音量的最大值

device.getAlarmMaxVolume() -> number
  返回闹钟音量的最大值

device.setMusicVolume(volume: number)
  设置当前媒体音量

device.setNotificationVolume(volume: number)
  设置当前通知音量

device.setAlarmVolume(volume: number)
  设置当前闹钟音量

device.getBattery() -> number
  返回当前电量百分比

device.getBatteryStatus() -> number
  获取电池状态

device.setBatteryStatus(value: number)
  模拟电池状态

device.setBatteryLevel(value: number)
  模拟电池电量百分百

device.getTotalMem() -> number
  返回设备内存总量

device.getAvailMem() -> number
  返回设备当前可用的内存

device.isScreenOn() -> boolean
  返回设备屏幕是否是亮着的

device.isScreenUnlock() -> boolean
  返回屏幕锁是否已经解开

device.setScreenMode(mode: number)
  设置屏幕显示模式

device.wakeUp()
  唤醒设备

device.keepScreenOn()
  保持屏幕常亮

device.vibrate(ms: number)
  使设备震动一段时间

device.cancelVibration()
  如果设备处于震动状态，则取消震动

--------------------------------------------------
-- 触摸操作 (touch)
--------------------------------------------------
touchDown(x: number, y: number, fingerID: number)
  按下屏幕

touchMove(x: number, y: number, fingerID: number)
  移动手指

touchUp(x: number, y: number, fingerID: number)
  抬起手指

click(x: number, y: number, fingerID: number)
  点击

longClick(x: number, y: number, duration: number)
  长按

swipe(x1: number, y1: number, x2: number, y2: number, duration: number)
  滑动

swipe2(x1: number, y1: number, x2: number, y2: number, duration: number)
  滑动(两点)

home()
  按下Home键

back()
  按下返回键

recents()
  按下最近任务键

powerDialog()
  长按电源键

notifications()
  下拉通知栏

quickSettings()
  下拉快捷设置

volumeUp()
  按下音量加键

volumeDown()
  按下音量减键

camera()
  按下相机键

keyAction(code: number)
  按键动作

--------------------------------------------------
-- 文件操作 (files)
--------------------------------------------------
files_isDir(path: string) -> boolean
  返回路径path是否是文件夹

files_isFile(path: string) -> boolean
  返回路径path是否是文件

files_isEmptyDir(path: string) -> boolean
  返回文件夹path是否为空文件夹

files_create(path: string) -> boolean
  创建一个文件或文件夹

files_createWithDirs(path: string) -> boolean
  创建一个文件或文件夹并确保所在文件夹存在

files_exists(path: string) -> boolean
  返回在路径path处的文件是否存在

files_ensureDir(path: string) -> boolean
  确保路径path所在的文件夹存在

files_read(path: string) -> string
  读取文本文件path的所有内容并返回

files_readBytes(path: string) -> string
  读取文件path的所有内容并返回

files_write(path: string, text: string)
  把text写入到文件path中

files_writeBytes(path: string, bytes: string)
  把bytes写入到文件path中

files_append(path: string, text: string)
  把text追加到文件path的末尾

files_appendBytes(path: string, bytes: string)
  把bytes追加到文件path的末尾

files_copy(fromPath: string, toPath: string) -> boolean
  复制文件

files_move(fromPath: string, toPath: string) -> boolean
  移动文件

files_rename(path: string, newName: string) -> boolean
  重命名文件

files_remove(path: string) -> boolean
  删除文件或文件夹

files_getName(path: string) -> string
  返回文件的文件名

files_getNameWithoutExtension(path: string) -> string
  返回不含拓展名的文件的文件名

files_getExtension(path: string) -> string
  返回文件的拓展名

files_path(relativePath: string) -> string
  返回相对路径对应的绝对路径

files_listDir(path: string) -> table
  列出文件夹path下的所有文件和文件夹

--------------------------------------------------
-- 图像处理 (images)
--------------------------------------------------
images_pixel(x: number, y: number) -> string
  获取指定坐标的像素颜色

images_captureScreen(x1: number, y1: number, x2: number, y2: number) -> userdata
  截取屏幕

images_cmpColor(x: number, y: number, colorStr: string, sim: number) -> boolean
  比较颜色

images_findColor(x1: number, y1: number, x2: number, y2: number, colorStr: string, sim: number, dir: number) -> number, number
  查找颜色

images_getColorCountInRegion(x1: number, y1: number, x2: number, y2: number, colorStr: string, sim: number) -> number
  获取区域内指定颜色的数量

images_detectsMultiColors(colors: string, sim: number) -> boolean
  检测多点颜色

images_findMultiColors(x1: number, y1: number, x2: number, y2: number, colors: string, sim: number, dir: number) -> number, number
  查找多点颜色

images_readFromPath(path: string) -> userdata
  从路径读取图片

images_readFromUrl(url: string) -> userdata
  从URL读取图片

images_readFromBase64(base64Str: string) -> userdata
  从Base64读取图片

images_readFromBytes(bytes: string) -> userdata
  从字节数组读取图片

images_save(img: userdata, path: string, quality: number) -> boolean
  保存图片

images_encodeToBase64(img: userdata, format: string, quality: number) -> string
  编码为Base64

images_encodeToBytes(img: userdata, format: string, quality: number) -> string
  编码为字节数组

images_clip(img: userdata, x1: number, y1: number, x2: number, y2: number) -> userdata
  裁剪图片

images_resize(img: userdata, width: number, height: number) -> userdata
  调整图片大小

images_rotate(img: userdata, degree: number) -> userdata
  旋转图片

images_grayscale(img: userdata) -> userdata
  灰度化

images_applyThreshold(img: userdata, threshold: number, maxVal: number, typ: string) -> userdata
  应用阈值

images_applyAdaptiveThreshold(img: userdata, maxValue: number, adaptiveMethod: string, thresholdType: string, blockSize: number, C: number) -> userdata
  应用自适应阈值

images_applyBinarization(img: userdata, threshold: number) -> userdata
  二值化

--------------------------------------------------
-- 存储管理 (storages)
--------------------------------------------------
storages_get(table: string, key: string) -> string
  从指定表中获取键值

storages_put(table: string, key: string, value: string)
  写入键值对

storages_remove(table: string, key: string)
  删除指定键

storages_contains(table: string, key: string) -> boolean
  判断键是否存在

storages_getAll(table: string) -> table
  获取所有键值对

storages_clear(table: string)
  清空指定表数据

--------------------------------------------------
-- 系统管理 (system)
--------------------------------------------------
system_getPid(processName: string) -> number
  获取进程ID

system_getMemoryUsage(pid: number) -> number
  获取内存使用

system_getCpuUsage(pid: number) -> number
  获取CPU使用率

system_restartSelf()
  重启自身

system_setBootStart(enable: boolean)
  设置开机自启

--------------------------------------------------
-- 网络请求 (http)
--------------------------------------------------
http_get(url: string, timeout: number) -> number, string
  发送GET请求

http_postMultipart(url: string, fileName: string, fileData: string) -> number, string
  发送Multipart POST请求

--------------------------------------------------
-- 媒体管理 (media)
--------------------------------------------------
media_scanFile(path: string)
  扫描路径path的媒体文件

--------------------------------------------------
-- 图像识别 (opencv)
--------------------------------------------------
opencv_findImage(x1: number, y1: number, x2: number, y2: number, template: string, isGray: boolean, scalingFactor: number, sim: number) -> number, number
  在指定区域内查找匹配的图片模板

--------------------------------------------------
-- 文字识别 (ppocr)
--------------------------------------------------
ppocr_ocr(x1: number, y1: number, x2: number, y2: number, colorStr: string) -> table
  识别屏幕文字

ppocr_ocrFromBase64(b64: string, colorStr: string) -> table
  识别Base64图片文字

ppocr_ocrFromPath(path: string, colorStr: string) -> table
  识别文件图片文字

--------------------------------------------------
-- 方法管理 (method)
--------------------------------------------------
registerMethod(name: string, description: string, goFunc: function, overridable: boolean)
  注册新方法到Lua引擎

unregisterMethod(name: string) -> boolean
  从Lua引擎中移除方法

listMethods() -> table
  列出所有已注册的方法

overrideMethod(name: string, luaFunc: function) -> boolean
  用Lua函数重写已注册的方法

restoreMethod(name: string) -> boolean
  恢复被重写的方法

--------------------------------------------------
-- 使用示例
--------------------------------------------------

-- 基本使用
-- 获取当前应用包名
local packageName = app_currentPackage()
print("当前应用: " .. packageName)

-- 点击屏幕
click(500, 1000, 1)

-- 查找颜色
local x, y = images_findColor(0, 0, 1080, 1920, "#FF0000", 0.9, 0)
if x ~= -1 and y ~= -1 then
    print("找到颜色在: " .. x .. ", " .. y)
    click(x, y, 1)
end

-- 识别文字
local results = ppocr_ocr(0, 0, 1080, 1920, "")
for i, result in ipairs(results) do
    print(result["标签"] .. " (" .. result["精度"] .. ")")
end

-- 方法重写
-- 重写 click 方法，添加日志
local originalClick = click

function click(x, y, fingerID)
    print("点击: (" .. x .. ", " .. y .. ")")
    originalClick(x, y, fingerID)
end

-- 或者使用 overrideMethod
overrideMethod("click", function(x, y, fingerID)
    print("点击: (" .. x .. ", " .. y .. ")")
    -- 调用原始实现
end)

-- 注册新方法
registerMethod("myCustomMethod", "我的自定义方法", nil, true)

function myCustomMethod(param)
    print("自定义方法被调用: " .. param)
end
