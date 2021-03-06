[![GoDoc](https://godoc.org/github.com/owulveryck/onnx-go?status.svg)](https://godoc.org/github.com/owulveryck/onnx-go/backend/testbackend) 

## About
This package contains utilities for testing a backend.

Test cases are described thanks to this structure

[embedmd]:# (test_structure.go /type TestCase/ /}/)
```go
type TestCase struct {
	Title          string
	ModelB         []byte
	Input          []tensor.Tensor
	ExpectedOutput []tensor.Tensor
}
```

The structure own a method that genererate a test function suitable to run through [`T.Run()`](https://golang.org/pkg/testing/#T.Run)

[embedmd]:# (test_structure.go /func .*RunTest/ /{/)
```go
func (tc *TestCase) RunTest(b backend.ComputationBackend, parallel bool) func(t *testing.T) {
```


A test for a certain OpType can be registered with the following command:

[embedmd]:# (test_structure.go /func Register/ /TestCase\)/)
```go
func Register(optype, testTitle string, constructor func() *TestCase)
```

## Usage in your backend implementation

Here is an example of a test for the Convolution Operator.

Register the tests from ONNX:

[embedmd]:# (example_test.go /_ \"git/ /"$/)
```go
_ "github.com/owulveryck/onnx-go/backend/testbackend/onnx"
```
[embedmd]:# (example_test.go /func.t/ /^\t}/)
```go
func(t *testing.T) {
		for _, tc := range testbackend.GetOpTypeTests("Conv") {
			tc := tc // capture range variable
			t.Run(tc().GetInfo(), tc().RunTest(backend, false))
		}
	}
```


## Tests from ONNX

Test files have been autogenerated from the [ONNX tests data](https://github.com/onnx/onnx/tree/master/onnx/backend/test/data/node) and are exposed via a seperate package.

A void import of the package register all the tests that are accessible through a call to the function `GetOpTypeTests`

[embedmd]:# (example_test.go /_ \"git/ /"$/)
```go
_ "github.com/owulveryck/onnx-go/backend/testbackend/onnx"
```
[embedmd]:# (test_structure.go /func GetOpTypeTests/ /{/)
```go
func GetOpTypeTests(optype string) []func() *TestCase {
```

