package must

import (
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestOK(t *testing.T) {
	e := errors.New("")
	assert.PanicsWithValue(t, e, func() { OK(e) })
	assert.NotPanics(t, func() { OK(nil) })
}

func TestDo(t *testing.T) {
	e := errors.New("")
	assert.PanicsWithValue(t, e, func() {
		Do(func() error {
			return e
		})
	})
	assert.NotPanics(t, func() {
		Do(func() error {
			return nil
		})
	})
}

func TestInt(t *testing.T) {
	e := errors.New("")
	assert.PanicsWithValue(t, e, func() { Int(10, e) })
	assert.Equal(t, 10, Int(10, nil))
}

type fakeFileInfo struct {
}

func (fakeFileInfo) Name() string {
	return ""
}

func (fakeFileInfo) Size() int64 {
	return 0
}

func (fakeFileInfo) Mode() os.FileMode {
	return 0
}

func (fakeFileInfo) ModTime() time.Time {
	return time.Time{}
}

func (fakeFileInfo) IsDir() bool {
	return false
}

func (fakeFileInfo) Sys() interface{} {
	return nil
}

func TestOSFileInfos(t *testing.T) {
	e := errors.New("")
	fis := []os.FileInfo{fakeFileInfo{}}

	assert.PanicsWithValue(t, e, func() {
		OSFileInfos(fis, e)
	})
	assert.Equal(t, fis, OSFileInfos(fis, nil))
}
