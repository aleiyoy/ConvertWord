# 常用工具集合
## 命名驼峰下划线转换
```cmd 
➜ go run main.go help word
该子命令支持各种单词格式转换，模式如下：
1：全部转大写
2：全部转小写
3：下划线转大写驼峰
4：下划线转小写驼峰
5：驼峰转下划线

Usage:
word [flags]

Flags:
-h, --help         help for word
-m, --mode int8    请输入单词转换的模式
-s, --str string   请输入单词内容
```

## 时间计算：输入时间返回指定格式时间
```cmd
➜ go run main.go help time
时间格式处理

Usage:
   time [flags]
   time [command]

Available Commands:
  calc        计算所需时间
  now         获取当前时间

Flags:
  -h, --help   help for time

Use " time [command] --help" for more information about a command.

➜ go run main.go time calc --help                     
计算所需时间

Usage:
   time calc [flags]

Flags:
  -c, --calculate string   需要计算的时间，有效单位为时间戳或已格式化后的时间
  -d, --duration string    持续时间，有效时间单位为"ns", "us" (or "µ s"), "ms", "s", "m", "h"
  -h, --help               help for calc



➜ go run main.go time calc -c="2029-09-04 12:02:33" -d=5m
2022/06/29 20:16:48 输出结果：2029-09-04 12:07:33, 1883218053
➜ go run main.go time calc -c="2029-09-04 12:02:33" -d=-2h
2022/06/29 20:17:09 输出结果：2029-09-04 10:02:33, 1883210553

```
