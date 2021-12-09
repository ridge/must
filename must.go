// Package must is a library to shorten handling of impossible conditions
//
//     must.OK(os.Unsetenv("FOO"))
//     bytes := must.Bytes(json.Marshal(dataStructureDefinedInCode))
//     defer must.Do(f.Close)
//
// is
//
//     if err := os.Unsetenv("FOO"); err != nil {
//         panic(err)
//     }
//
//     bytes, err := json.Marshal(dataStructureDefinedInCode)
//     if err != nil {
//         panic(err)
//     }
//
//     defer func() {
//         if err := f.Close(); err != nil{
//             panic(err)
//         }
//     }()
//
// Go error handling style is practical for majority of errors. However not all
// errors are meaningful and actionable by the caller, so it does not make sense
// to surface them.
//
// Go tacitly acknowledges it by providing functions regex.MustCompile and
// template.Must in the standard library. This library expands on the same idea.
//
package must

import (
	"time"
)

//go:generate go run ./generator types_gen.go

// OK panics on error
func OK(err error) {
	if err != nil {
		panic(err)
	}
}

// Do calls the function and panics on error.
//
// For use primarily in defer statements.
//
// BAD example:
//
//     // f.Close will be called now, must.OK will be called on function exit
//     defer must.OK(f.Close())
//
// GOOD example:
//
//     defer must.Do(f.Close)
//
func Do(fn func() error) {
	OK(fn())
}

// DoRetryable calls the function until no error is returned or when the number of attempts has reached.
// If all attempts fail, DoRetryable panics with the last error received.
//
// DoRetryable is an extension for Do(), used for a variety of functions, example:
//
//	func StoreEncrypted(ctx context.Context, u User, password string) error {
//		must.DoRetryable(func() error {
//			return encrypt(ctx, &password)
//		}, 5) // will try to encrypt 5 times before exhasuted
//
//		u.SetPassword(password)
//		return nil
//	}
//
// 	func encrypt(ctx context.Context, s *string) error {
//		if s* == "" {
//			return errors.new("empty string is not supported")
//		}
//		return vault.Encrypt(ctx, s)
// 	}
func DoRetryable(fn func() error, attempts int) {
	DoRetryableStall(fn, attempts, 0)
}

// DoRetryableStall is the same as DoRetryable, but only that it delays every attempt by the given duration. Example:
//	func StoreEncrypted(ctx context.Context, u User, password string) error {
//		must.DoRetryable(func() error {
//			return encrypt(ctx, &password)
//		}, 5, 1 * time.Second)
//
//		u.SetPassword(password)
//		return nil
//	}
func DoRetryableStall(fn func() error, attempts int, d time.Duration) {
	var err error
	for i := 0; i < attempts; i++ {
		if err = fn(); err == nil {
			return
		}
		time.Sleep(d)
	}
	panic(err)
}
