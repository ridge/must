# Go library to shorten handling of impossible conditions
[![Go Reference](https://pkg.go.dev/badge/github.com/ridge/must.svg)](https://pkg.go.dev/github.com/ridge/must)

    must.OK(os.Unsetenv("FOO"))
    bs := must.OK1(json.Marshal(dataStructureDefinedInCode))
    defer must.Do(f.Close)

is

    if err := os.Unsetenv("FOO"); err != nil {
         panic(err)
    }
    bs, err := json.Marshal(dataStructureDefinedInCode)
    if err != nil {
        panic(err)
    }
    defer func() {
        if err := f.Close(); err != nil {
            panic(err)
        }
    }()

## Why panic?

Go error handling style is practical for majority of errors. However not all
errors are meaningful and actionable by the caller, so it does not make sense to
surface them.

Go tacitly acknowledges it by providing functions `regex.MustCompile` and
`template.Must` in the standard library. This library expands on the same idea.

## Where is it useful?

The library is proven to be useful in the following situations:

- Programming errors in all kinds of code, similar to `regex.MustCompile`,
- Filesystem-related errors in a build system,
- Fatal errors in server-side SaaS code.

The last point needs an explanation.

Unlike CLI tools where errors ought to be handled in a way meaningful to a user
who might not know the internals of the application, server-side SaaS
applications are run by the SREs or developers who do not need the extra
wrapping, and instead require fatal errors to be delivered reliably and with all
gory details.

Passing errors up the stack in a regular Go way risks losing the error (by
forgetting to wrap it, or by swallowing it completely), and prevents the
original stacktrace from being surfaced. `panic`ing at the place of fatal error
is reliable and precise.

## Is it a good idea to use `must` everywhere?

Definitely not. Go rules for error handling are well thought-out. Reserve `must`
for programming errors and impossible conditions.

## Legal

Copyright Tectonic Networks, Inc.

Licensed under [Apache 2.0](LICENSE) license.

Authors:
- [Mikhail Gusarov](https://github.com/dottedmag)
- [Alexey Feldgendler](https://github.com/feldgendler)
