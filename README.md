go-to-thrift
========

**Go-to-thrift** is a compiler that compiles golang file to thrift IDL.

Installation
------------

Note: before executing the following commands, **make sure your `GOPATH` environment is properly set**.

Using `go install`:

```shell
GO111MODULE=on go install github.com/qwsc/go-to-thrift@latest
```

Usage
-----


To compile a golang file to thrift IDL with the default setting, you can just run:

```shell
go-to-thrift the-golang-file.go
```
