# ffmepg-conv-tool
**ffmpeg 格式转换小工具，用于批量转换**

此项目（或视为脚本）来源于一个音频转换的小小的需求。

希望能有一个执行文件，双击后直接将同级目录下的指定后缀文件全部批量转为另一个后缀。搜索了一下，大概知道怎么做了，然后选择了能做交叉编译的 Go 语言实现。

## 使用说明
先了解一下配置项：

```go
package main

const (
    from = "m4a" // 原后缀
    to   = "mp3" // 现后缀
    
    // ffmpegDarwin MacOS 下 ffmpeg 的位置
    ffmpegDarwin = "./ffmpeg/ffmpeg"
    // Linux 同理
    ffmpegLinux = "./ffmpeg/ffmpeg"
    // Windows 同理
    ffmpegWindows = "./ffmpeg/bin/ffmpeg.exe"
)
```

**使用步骤**

1. [下载 `ffmpeg`](http://ffmpeg.org/download.html) ，然后会得到 `ffmpeg` 的执行文件。
2. 在源代码中根据配置项说明配置好自己需要格式转换的后缀，以及 `ffmpwg` 执行文件的相对位置，参考相对位置例如：
   ```
   ├── ffmepg-conv-tool // 本项目目录
   │ ├── ffmpeg         // ffmpeg 文件夹
   │ │ └── ffmpeg       // ffmpeg 执行文件
   │ ├── go.mod         // go mod 文件
   │ ├── main.exe       // Windows 下执行文件
   │ ├── main           // MacOS 或 Linux 下执行文件
   │ ├── main.go        // 全部的源码文件
   │ ├── test.m4a       // 待转换文件放到 main 执行文件同级目录下
   │ └── test.mp3       // 转换完成的文件会在相同位置输出
   ```
3. 编译本项目（脚本），得到执行文件 `main`。
4. 确保相对位置放好后，执行文件 `main` 和待转换文件应该已在同一个目录下。
5. 然后双击 `main` 执行文件即可。

**说明**

源代码里写死了从 `m4a` 到 `mp3` 的转换，所以需要自己手动改源码重新编译，故不提供编译后的文件。
