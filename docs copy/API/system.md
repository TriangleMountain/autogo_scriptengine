# system - 系统功能
---
提供与系统相关的功能，包括进程管理和资源查询。

## GetPid
<hr style="margin: 0;">

获取指定进程的 PID。

- `processName` {string} 进程名。如果为空则返回当前进程的 PID。

```go
pid := system.GetPid("com.example.app")
fmt.Println("Process ID:", pid)
```

## GetMemoryUsage
<hr style="margin: 0;">

获取指定进程的内存使用量。

- `pid` {int} 进程的 PID。如果为 0 则获取当前进程的内存使用量。

```go
memoryUsage := system.GetMemoryUsage(12345)
fmt.Println("Memory Usage (KB):", memoryUsage)
```

## GetCpuUsage
<hr style="margin: 0;">

获取指定进程的 CPU 使用率。

- `pid` {int} 进程的 PID。如果为 0 则获取当前进程的 CPU 使用率。

```go
cpuUsage := system.GetCpuUsage(12345)
fmt.Println("CPU Usage (%):", cpuUsage)
```

## RestartSelf
<hr style="margin: 0;">

重启当前脚本进程。

```go
system.RestartSelf()
```

## SetBootStart
<hr style="margin: 0;">

设置脚本开机自动运行，需要root权限。

```go
system.SetBootStart(true)
```