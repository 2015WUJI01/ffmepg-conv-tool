package main

import (
	"io/ioutil"
	"log"
	"os/exec"
	"runtime"
	"strings"
)

// Configration 配置项：因为懒，所以暂时不想单独多写一个配置文件再解析，除非有需求。
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

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		filename := file.Name()
		if strings.HasSuffix(filename, from) {
			conv(strings.TrimSuffix(filename, from)) //<- "filename."
		}
	}
}

func conv(name string) {
	// ffmpeg 在当前系统中的位置
	var ffmpeg string
	switch runtime.GOOS {
	case "darwin":
		ffmpeg = ffmpegDarwin
	case "windows":
		ffmpeg = ffmpegWindows
	case "linux":
		ffmpeg = ffmpegLinux
	}

	// 拼凑参数集合
	cmdArguments := []string{"-i", name + from, name + to}

	// 调用当前系统 cmd 执行 ffmpeg
	cmd := exec.Command(ffmpeg, cmdArguments...)

	_ = cmd.Run()
}
