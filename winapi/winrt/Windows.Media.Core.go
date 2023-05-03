package winrt

type MediaStreamSource struct {
	BufferTime               uintptr // 获取或设置 MediaStreamSource 缓冲的数据量。
	CanSeek                  uintptr // 获取或设置应用程序是否支持在媒体时间线中更改其位置。
	Duration                 uintptr // 获取或设置媒体时间线的持续时间。
	IsLive                   uintptr // 获取或设置一个值，该值指示正在处理的媒体内容是否处于活动状态。
	MaxSupportedPlaybackRate uintptr // 获取或设置媒体流源支持的最大回放速率。
	MediaProtectionManager   uintptr // 获取或设置用于保护媒体内容的数字版权管理 (DRM) 保护管理器。
	MusicProperties          uintptr //	获取或设置音频文件的音乐属性。
	Thumbnail                uintptr // 获取或设置音频文件的缩略图。
	VideoProperties          uintptr // 获取或设置视频文件的视频属性。
}
