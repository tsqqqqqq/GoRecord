package winrt

import (
	"github.com/go-ole/go-ole"
	"unsafe"
)

type VideoEncodingPropertiesItem struct {
	ole.IInspectable
}

type VideoEncodingPropertiesVtbl struct {
	ole.IInspectableVtbl
	Width            uintptr // 获取或设置视频帧的宽度。
	Height           uintptr // 获取或设置视频帧的高度。
	FrameRate        uintptr // 获取或设置视频帧的帧速率。
	Bitrate          uintptr // 获取或设置视频帧的比特率。
	PixelAspectRatio uintptr // 获取或设置视频帧的像素宽高比。
	ProfileId        uintptr // 获取视频编码的配置文件 ID。
	Properties       uintptr // 获取视频编码属性的其他属性。
	Subtype          uintptr // 获取视频编码的子类型。
	Level            uintptr // 获取视频编码的级别。
	Kind             uintptr // 获取视频编码的类型。
	SphericalVideo   uintptr // 获取或设置视频帧是否为球面视频。
	StereoMode       uintptr // 获取或设置视频帧的立体声模式。
}

func (v *VideoEncodingPropertiesItem) VTable() *VideoEncodingPropertiesVtbl {
	return (*VideoEncodingPropertiesVtbl)(unsafe.Pointer(v.RawVTable))
}
