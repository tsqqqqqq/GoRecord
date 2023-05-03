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
	//handles := make(map[win.HWND]string)
	//cb := syscall.NewCallback(func(h win.HWND, p uintptr) uintptr {
	//	// EnumWindows函数用于枚举桌面窗口
	//	// 第一个参数是回调函数，用于处理枚举到的窗口
	//	// 第二个参数是用户自定义参数，会传递给回调函数
	//	// 返回值是枚举是否成功
	//	// 回调函数的返回值是是否继续枚举，如果为false则停止枚举
	//	var name = strings.TrimSpace(winapi.GetWindowTextString(h))
	//	for item := range notHands {
	//		if strings.Contains(name, notHands[item]) {
	//			return 1
	//		}
	//	}
	//	if name != "" {
	//		// handles 去重
	//		handles[h] = name
	//	}
	//	return 1
	//})
	//winapi.EnumDesktopWindows(0, cb, 0)
	//fmt.Println(handles)
	for {
		// FindWindowEx函数用于查找与指定窗口有关的窗口
		// 第一个参数是父窗口句柄，如果为0则从桌面开始查找
		// 第二个参数是子窗口句柄，如果为0则从父窗口开始查找
		// 第三个参数是窗口类名，如果为nil则从父窗口开始查找
		// 第四个参数是窗口标题，如果为nil则从父窗口开始查找
		// 返回值是找到的窗口句柄，如果为0则表示没有找到

		//
		// FindWindow怎么用
		// 第一个参数是窗口类名，如果为nil则从桌面开始查找
		// 第二个参数是窗口标题，如果为nil则从桌面开始查找
		// 返回值是找到的窗口句柄，如果为0则表示没有找到

		rdHwnd = winapi.FindWindow(nil, winapi.MustUTF16PtrFromString("SevenEightSSR"))
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
