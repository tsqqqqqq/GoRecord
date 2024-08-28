package main

import (
	"GoRecord/winapi"
	"fmt"
	"github.com/lxn/win"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

// 解释一下下面的代码
// 1. 通过FindWindow找到窗口句柄
// 2. 通过FindWindowEx找到窗口句柄
// 3. 通过GetWindowTextString获取窗口标题
// 4. 通过StartCapture开始捕获
// 5. 通过signal.Notify监听系统信号
// 6. 通过<-sc阻塞主线程
var notHands = []string{"Progr", "Default IME"}

func main() {
	var rdHwnd win.HWND
	for {
		// FindWindowEx函数用于查找与指定窗口有关的窗口
		// 第一个参数是父窗口句柄，如果为0则从桌面开始查找
		// 第二个参数是子窗口句柄，如果为0则从父窗口开始查找
		// 第三个参数是窗口类名，如果为nil则从父窗口开始查找
		// 第四个参数是窗口标题，如果为nil则从父窗口开始查找
		// 返回值是找到的窗口句柄，如果为0则表示没有找到
		rdHwnd = winapi.FindWindow(nil, winapi.MustUTF16PtrFromString("艾可云"))
		if rdHwnd == 0 {
			win.MessageBox(0, winapi.MustUTF16PtrFromString("Could not find window"), winapi.MustUTF16PtrFromString("RDP Relative Input"), win.MB_ICONERROR)
			return
		}
		var name = strings.TrimSpace(winapi.GetWindowTextString(rdHwnd))
		if name != "" {
			//fmt.Println(name)
			break
		}
	}
	var handler = CaptureHandler{}

	err := handler.StartCapture(rdHwnd)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer handler.Close()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
