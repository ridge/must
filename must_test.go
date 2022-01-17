package must

import (
	"errors"
	"os"
	"reflect"
	"testing"
	"time"
)

func capturePanic(f func()) (ret interface{}) {
	defer func() {
		ret = recover()
	}()
	f()
	return nil
}

func TestOK(t *testing.T) {
	e := errors.New("")

	if capturePanic(func() { OK(e) }) != e {
		t.Errorf("OK(non-nil-error) did not panic with the passed error")
	}

	if capturePanic(func() { OK(nil) }) != nil {
		t.Errorf("OK(nil) panicked")
	}
}

func TestDo(t *testing.T) {
	e := errors.New("")

	ret := capturePanic(func() {
		Do(func() error {
			return e
		})
	})
	if ret != e {
		t.Errorf("Do({return e}) did not panic with the returned error")
	}

	ret = capturePanic(func() {
		Do(func() error {
			return nil
		})
	})
	if ret != nil {
		t.Errorf("Do({return nil}) panicked")
	}
}

func TestInt(t *testing.T) {
	e := errors.New("")

	if capturePanic(func() { Int(10, e) }) != e {
		t.Errorf("Int(10, non-nil-error) did not panic with the passed error")
	}

	if Int(10, nil) != 10 {
		t.Errorf("Int(10, nil) did not return 10")
	}
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

	if capturePanic(func() { OSFileInfos(fis, e) }) != e {
		t.Errorf("OSFileInfos(fis, non-nil-error) did not panic with the passed error")
	}

	if !reflect.DeepEqual(OSFileInfos(fis, nil), fis) {
		t.Errorf("OSFileInfos(fis, nil) did not return fis")
	}
}
