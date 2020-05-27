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

//go:generate go run ./generator types_gen.go

// OK panics on error
func OK(err error) {
	if err != nil {
		panic(err)
	}
}

// Do calls the function and panics on error
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
