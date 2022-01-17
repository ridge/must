package must

import (
	"errors"
	"testing"
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

func TestOK1(t *testing.T) {
	e := errors.New("")

	if capturePanic(func() { OK1(10, e) }) != e {
		t.Errorf("OK1(..., non-nil-error) did not panic with the passed error")
	}

	if OK1(10, nil) != 10 {
		t.Errorf("OK1(10, nil) did not return first argument")
	}
}

func TestOK2(t *testing.T) {
	e := errors.New("")

	if capturePanic(func() { OK2(10, 'c', e) }) != e {
		t.Errorf("OK2(..., non-nil-error) did not panic with the passed error")
	}

	v1, v2 := OK2(10, 'c', nil)
	if v1 != 10 || v2 != 'c' {
		t.Errorf("OK2(10, 'c', nil) did not return first arguments")
	}
}

func TestOK3(t *testing.T) {
	e := errors.New("")

	if capturePanic(func() { OK3(10, 'c', 7.0, e) }) != e {
		t.Errorf("OK3(..., non-nil-error) did not panic with the passed error")
	}

	v1, v2, v3 := OK3(10, 'c', 7.0, nil)
	if v1 != 10 || v2 != 'c' || v3 != 7.0 {
		t.Errorf("OK3(10, 'c', 7.0f, nil) did not return first arguments")
	}
}

func TestOK4(t *testing.T) {
	e := errors.New("")

	if capturePanic(func() { OK4(10, 'c', 7.0, "h", e) }) != e {
		t.Errorf("OK4(..., non-nil-error) did not panic with the passed error")
	}

	v1, v2, v3, v4 := OK4(10, 'c', 7.0, "h", nil)
	if v1 != 10 || v2 != 'c' || v3 != 7.0 || v4 != "h" {
		t.Errorf("OK4(10, 'c', 7.0, \"h\", nil) did not return first arguments")
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
