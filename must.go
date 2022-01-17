// Package must is a library to shorten handling of impossible conditions
//
//     must.OK(os.Unsetenv("FOO"))
//     bytes := must.OK1(json.Marshal(dataStructureDefinedInCode))
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

// OK panics on error
func OK(err error) {
	if err != nil {
		panic(err)
	}
}

// OK1 panics on error, returns the first argument otherwise
func OK1[T any](t T, err error) T {
	OK(err)
	return t
}

// OK2 panics on error, returns the first arguments otherwise
func OK2[T1, T2 any](t1 T1, t2 T2, err error) (T1, T2) {
	OK(err)
	return t1, t2
}

// OK3 panics on error, returns the first arguments otherwise
func OK3[T1, T2, T3 any](t1 T1, t2 T2, t3 T3, err error) (T1, T2, T3) {
	OK(err)
	return t1, t2, t3
}

// OK4 panics on error, returns the first arguments otherwise
func OK4[T1, T2, T3, T4 any](t1 T1, t2 T2, t3 T3, t4 T4, err error) (T1, T2, T3, T4) {
	OK(err)
	return t1, t2, t3, t4
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
