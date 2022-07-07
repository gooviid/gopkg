package goovimodel

import "time"

type MediaInfo struct {
	CreatedAt             time.Time              `json:"created_at"`
	Mime                  string                 `json:"mime"`
	Size                  int64                  `json:"size"`
	SizeHuman             string                 `json:"size_human"`
	RemoteUrl             string                 `json:"remote_url"`
	RemoteUrlRelative     string                 `json:"remote_url_relative"`
	Video                 *VideoInfo             `json:"video"`
	AudioInfo             []AudioInfo            `json:"audio"`
	ImageInfo             *ImageInfo             `json:"image"`
	AdaptiveStreamingInfo *AdaptiveStreamingInfo `json:"adaptive_streaming"`
}

type VideoInfo struct {
	IsHDR               bool    `json:"is_hdr"`
	IsMultiChannelAudio bool    `json:"is_multi_channel_audio"`
	Width               int     `json:"width"`
	Height              int     `json:"height"`
	FPS                 float64 `json:"fps"`
	CodecName           string  `json:"codec_name"`
	CodecType           string  `json:"codec_type"`
	ResolutionHuman     string  `json:"resolution_human"`
	BitRate             string  `json:"bitrate"`
	DisplayAspectRatio  string  `json:"display_aspect_ratio"`
	AspectRatio         float64 `json:"aspect_ratio"`
	Duration            float64 `json:"duration"`
}

type AdaptiveStreamingInfo struct {
	Type               string  `json:"type"`
	FileType           string  `json:"file_type"`
	IsVideo            bool    `json:"is_video"`
	DisplayAspectRatio string  `json:"display_aspect_ratio"`
	AspectRatio        float64 `json:"aspect_ratio"`
	Duration           float64 `json:"duration"`
	FPS                float64 `json:"fps"`
}

type AudioInfo struct {
	BitRate       string  `json:"bitrate"`
	Frequency     string  `json:"frequency"`
	Type          string  `json:"type"`
	ChannelNumber int     `json:"channel_number"`
	ChannelName   string  `json:"channel_name"`
	Duration      float64 `json:"duration"`
	AudioChannel  int     `json:"audio_channel"`
}

type ImageInfo struct {
	Width              int     `json:"width"`
	Height             int     `json:"height"`
	ImageType          string  `json:"image_type"`
	DisplayAspectRatio string  `json:"display_aspect_ratio"`
	AspectRatio        float64 `json:"aspect_ratio"`
}

type TranscodeError struct {
	ErrorType string    `json:"error_type"`
	Cause     string    `json:"error_cause"`
	ErrorAt   time.Time `json:"error_at"`
}

type TranscodeMeta struct {
	ProcessStartAt    time.Time `json:"process_start_at,omitempty"`
	ProcessEndAt      time.Time `json:"process_end_at,omitempty"`
	ProcessDuration   float32   `json:"process_duration,omitempty"`
	TranscodeStartAt  time.Time `json:"transcode_start_at,omitempty"`
	TranscodeEndAt    time.Time `json:"transcode_end_at,omitempty"`
	TranscodeDuration float32   `json:"transcode_duration,omitempty"`
}

type BillingDetail struct {
	Qty    int32   `json:"qty,omitempty"`
	Method string  `json:"method,omitempty"`
	Total  float32 `json:"total,omitempty"`
	Price  float32 `json:"price,omitempty"`
}
