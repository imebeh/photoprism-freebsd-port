--- internal/ffmpeg/convert.go.orig	2023-07-18 22:21:07 UTC
+++ internal/ffmpeg/convert.go
@@ -47,12 +47,12 @@ func AvcConvertCommand(fileName, avcName string, opt O
 			"-qsv_device", "/dev/dri/renderD128",
 			"-i", fileName,
 			"-c:a", "aac",
+			"-ac", "2",
 			"-vf", opt.VideoFilter(FormatRGB32),
 			"-c:v", opt.Encoder.String(),
 			"-map", opt.MapVideo,
 			"-map", opt.MapAudio,
-			"-vsync", "vfr",
-			"-r", "30",
+			"-fps_mode", "vfr",
 			"-b:v", opt.Bitrate,
 			"-bitrate", opt.Bitrate,
 			"-f", "mp4",
@@ -69,11 +69,11 @@ func AvcConvertCommand(fileName, avcName string, opt O
 			"-map", opt.MapVideo,
 			"-map", opt.MapAudio,
 			"-c:a", "aac",
+			"-ac", "2",
 			"-vf", opt.VideoFilter(FormatYUV420P),
 			"-profile", "high",
 			"-level", "51",
-			"-vsync", "vfr",
-			"-r", "30",
+			"-fps_mode", "vfr",
 			"-b:v", opt.Bitrate,
 			"-f", "mp4",
 			"-y",
@@ -86,12 +86,12 @@ func AvcConvertCommand(fileName, avcName string, opt O
 			"-hwaccel", "vaapi",
 			"-i", fileName,
 			"-c:a", "aac",
+			"-ac", "2",
 			"-vf", opt.VideoFilter(FormatNV12),
 			"-c:v", opt.Encoder.String(),
 			"-map", opt.MapVideo,
 			"-map", opt.MapAudio,
-			"-vsync", "vfr",
-			"-r", "30",
+			"-fps_mode", "vfr",
 			"-b:v", opt.Bitrate,
 			"-f", "mp4",
 			"-y",
@@ -109,6 +109,7 @@ func AvcConvertCommand(fileName, avcName string, opt O
 			"-map", opt.MapVideo,
 			"-map", opt.MapAudio,
 			"-c:a", "aac",
+			"-ac", "2",
 			"-preset", "15",
 			"-pixel_format", "yuv420p",
 			"-gpu", "any",
@@ -116,7 +117,6 @@ func AvcConvertCommand(fileName, avcName string, opt O
 			"-rc:v", "constqp",
 			"-cq", "0",
 			"-tune", "2",
-			"-r", "30",
 			"-b:v", opt.Bitrate,
 			"-profile:v", "1",
 			"-level:v", "auto",
@@ -135,13 +135,13 @@ func AvcConvertCommand(fileName, avcName string, opt O
 			"-map", opt.MapVideo,
 			"-map", opt.MapAudio,
 			"-c:a", "aac",
+			"-ac", "2",
 			"-vf", opt.VideoFilter(FormatYUV420P),
 			"-num_output_buffers", "72",
 			"-num_capture_buffers", "64",
 			"-max_muxing_queue_size", "1024",
 			"-crf", "23",
-			"-vsync", "vfr",
-			"-r", "30",
+			"-fps_mode", "vfr",
 			"-b:v", opt.Bitrate,
 			"-f", "mp4",
 			"-y",
@@ -156,11 +156,11 @@ func AvcConvertCommand(fileName, avcName string, opt O
 			"-map", opt.MapVideo,
 			"-map", opt.MapAudio,
 			"-c:a", "aac",
+			"-ac", "2",
 			"-vf", opt.VideoFilter(FormatYUV420P),
 			"-max_muxing_queue_size", "1024",
 			"-crf", "23",
-			"-vsync", "vfr",
-			"-r", "30",
+			"-fps_mode", "vfr",
 			"-b:v", opt.Bitrate,
 			"-f", "mp4",
 			"-y",
