package record

// windows 录屏
// 1. 获取窗口句柄
// 2. 获取窗口大小
// 3. 获取窗口位置
// 4. 获取窗口内容
// 5. 保存为视频

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	modkernel32     = syscall.MustLoadDLL("kernel32.dll")
	user32          = syscall.MustLoadDLL("user32.dll")
	graphicsCapture = syscall.MustLoadDLL("GraphicsCapture.dll")

	graphicsObj        = graphicsCapture.MustFindProc("DllGetActivationFactory")
	procLoadLibraryExW = modkernel32.MustFindProc("LoadLibraryExW")
	procGetProcAddress = modkernel32.MustFindProc("GetProcAddress")

	procGetDpiForWindow = user32.MustFindProc("GetDpiForWindow")
)

func Record(hwnd uintptr) int {
	//moduser32, err := loadLibraryExw("user32.dll", 0, 0) // 加载 user32.dll	// 0x7FFA4C1B0000
	//if err != nil {
	//	fmt.Println(err)
	//	panic(err)
	//}
	//procGetDpiForWindow, err := getProcAddress(moduser32, "GetDpiForWindow") // 获取 GetDpiForWindow 函数地址	// 0x7FFA4C1B0A80
	//if err != nil {
	//	fmt.Println(err)
	//	panic(err)
	//}

	//graphicsCapture, err := loadLibraryExw("GraphicsCapture.dll", 0, 0)
	//if err != nil {
	//	fmt.Println(err)
	//	panic(err)
	//}
	//getObj, err := getProcAddress(graphicsCapture, "DllGetClassObject")
	//if err != nil {
	//	fmt.Println(err)
	//	panic(err)
	//}
	fmt.Println("------------------->", graphicsObj)

	var dpi int
	r1, _, e1 := syscall.Syscall(procGetDpiForWindow.Addr(), 1, uintptr(hwnd), uintptr(unsafe.Pointer(&dpi)), 0)
	if r1 == 0 {
		if e1 != 0 {
			fmt.Println(e1)
		} else {
			e1 = syscall.EINVAL
		}
	}
	fmt.Println(r1)
	return dpi
}

// 解释一下这个函数
// 1. 通过 syscall.UTF16PtrFromString 将字符串转换为 UTF16 编码的指针
// 2. 通过 procLoadLibraryExW.Call 调用 LoadLibraryExW 函数
// 3. 通过 handle == 0 判断是否调用成功
func loadLibraryExw(name string, hfile uintptr, flags uint32) (uintptr, error) {
	lpName, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return 0, err
	}

	handle, _, err := procLoadLibraryExW.Call(
		uintptr(unsafe.Pointer(lpName)),
		hfile,
		uintptr(flags),
	)
	if handle == 0 {
		return 0, err
	}
	return handle, nil
}

// getProcAddress 解释一下这个函数
// 1. 通过 syscall.UTF16PtrFromString 将字符串转换为 UTF16 编码的指针
// 2. 通过 procGetProcAddress.Call 调用 GetProcAddress 函数
// 3. 通过 addr == 0 判断是否调用成功
func getProcAddress(hModule uintptr, procName string) (uintptr, error) {
	lpProcName, err := syscall.UTF16PtrFromString(procName)
	if err != nil {
		return 0, err
	}
	addr, _, err := procGetProcAddress.Call(
		hModule,
		uintptr(unsafe.Pointer(lpProcName)),
	)
	if addr == 0 {
		return 0, err
	}
	return addr, nil
}
