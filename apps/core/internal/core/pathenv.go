package core

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// PathWithDirectory returns a PATH value whose first entry is dir.
//
// Every bundled downloader helper (ffmpeg, ffprobe, deno, ...) lives next to
// the downloader binary in the dependency directory, but only N_m3u8DL-RE gets
// an explicit ffmpeg path through the schema. yt-dlp and BBDown resolve ffmpeg
// through PATH alone, so a container image that ships a slim base system
// without a system-wide ffmpeg leaves them unable to merge separate video and
// audio streams — the download "succeeds" but the result stays split across
// several files.
//
// Existing entries keep their order; a duplicate reference to dir is dropped so
// repeated calls stay stable.
func PathWithDirectory(currentPath, dir string) string {
	if strings.TrimSpace(dir) == "" {
		return currentPath
	}
	dir = filepath.Clean(dir)

	separator := string(os.PathListSeparator)
	entries := make([]string, 0, 8)
	entries = append(entries, dir)
	for _, entry := range strings.Split(currentPath, separator) {
		if entry == "" || sameDirectory(entry, dir) {
			continue
		}
		entries = append(entries, entry)
	}
	return strings.Join(entries, separator)
}

// sameDirectory compares two PATH entries the way the host filesystem does.
func sameDirectory(left, right string) bool {
	left = filepath.Clean(left)
	right = filepath.Clean(right)
	if runtime.GOOS == "windows" {
		return strings.EqualFold(left, right)
	}
	return left == right
}
