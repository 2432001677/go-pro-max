package gopromax

import (
	"io"
	"os"
)

// FileExists check file exists
func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// ClearFile clear text by file path
func ClearFile(path string) error {
	return os.Truncate(path, 0)
}

type FileTailLine struct {
	Str   string // content
	Count int64  // line count
	Err   error  // error
}

// ReadTailN read file tail with n line
func ReadTailN(f *os.File, n int64) *FileTailLine {
	var (
		res          = &FileTailLine{}
		cursor int64 = 0
		char         = make([]byte, 1)
	)

	if n < 0 {
		content, err := io.ReadAll(f)
		if err != nil {
			res.Err = err
			return res
		}
		res.Str = string(content)
		return res
	}

	stat, err := f.Stat()
	if err != nil {
		res.Err = err
		return res
	}

	fSize := stat.Size()

	for {
		cursor -= 1
		f.Seek(cursor, io.SeekEnd)

		f.Read(char)

		if cursor != -1 && (char[0] == 10 || char[0] == 13) {
			res.Count += 1
		}
		if res.Count == n {
			break
		}

		res.Str = string(char) + res.Str

		if cursor == -fSize {
			break
		}
	}
	return res
}
