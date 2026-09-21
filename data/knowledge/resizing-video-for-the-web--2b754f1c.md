---
title: "Resizing Video for the Web"
notion_id: 2b754f1c-7d23-81f4-a756-c8417d18a4c7
notion_url: https://app.notion.com/p/Resizing-Video-for-the-Web-2b754f1c7d2381f4a756c8417d18a4c7
last_edited: 2025-11-26T18:02:00.000Z
source_url: https://joeldare.com/resizing-video-for-the-web
tags: ["English", "Web Development", "Video", "Automation", "Article", "Tutorial", "joeldare.com"]
---
I often need to resize video for use on the web.

Especially video created by recording my screen with Quicktime on MacOS, which tends to be rather large by default.

Here’s a quick ffmpeg command you can use to resize to something more reasonable.

- `I input.mov` specifies your source video file.
- `vcodec libx264` uses the H.264 video codec, which has good we support.
- `crf 23` sets a constant rate factor, lower for better quality and a larger file. The range is 0–51 with 23 being a good default.
- `preset fast` is a good trade-off between encoding time and compression. The options are: ultrafast, superfast, veryfast, faster, fast, medium (default), slow, slower, veryslow. Use slow or slower for better compression/quality at the cost of speed.
- `acodec aac` uses AAC audio, which is widely supported.
- `b:a 128k` sets the audio bitrate.
- `movflags +faststart` optimizes the file for web streaming by moving metadata to the beginning of the file.

Want to see my personal updates, tech projects, and business ideas? Join the mailing list.

JoelDare.com © Dare Companies Dotcom LLC
