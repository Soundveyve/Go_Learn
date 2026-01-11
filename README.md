# Go_Learn
Repo for learn Golang

Site for learn: https://quii.gitbook.io/learn-go-with-tests/go-fundamentals

Local view doc: go doc NAME_OF_PACKAGE
Example: go doc fmt
ALTERNATIVE:
go install golang.org/x/pkgsite/cmd/pkgsite@latest
and write: pkgsite -open
all doc on: http://localhost:8080/testing

# Naming for test:
(https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world#writing-tests)

Writing a test is just like writing a function, with a few rules

- It needs to be in a file with a name like xxx_test.go

- The test function must start with the word Test

- The test function takes one argument only t *testing.T

- To use the *testing.T type, you need to import "testing", like we did with fmt in the other file