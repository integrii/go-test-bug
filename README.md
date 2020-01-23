# go-test-bug
Demonstrates a bug in argument passing in go test.

You would expect to be able to run `go test` with the flag package, but instead you get the following:

```
flag provided but not defined: -test.timeout
Usage of /var/folders/nt/y0clsxwx4wjf7w8b9kcb7wgc0000gn/T/go-build978967543/b001/go-test-bug.test:
  -f	An example bool flag
exit status 2
FAIL	_/Users/ericgreer/git/go-test-bug	0.005s
```
