## dlgs
[![TravisCI Build Status](https://travis-ci.org/gen2brain/dlgs.svg?branch=master)](https://travis-ci.org/gen2brain/dlgs)
[![AppVeyor Build Status](https://ci.appveyor.com/api/projects/status/53ekhdkai4r12un3?svg=true)](https://ci.appveyor.com/project/gen2brain/dlgs)
[![GoDoc](https://godoc.org/github.com/gen2brain/dlgs?status.svg)](https://godoc.org/github.com/gen2brain/dlgs)
[![Go Report Card](https://goreportcard.com/badge/github.com/gen2brain/dlgs?branch=master)](https://goreportcard.com/report/github.com/gen2brain/dlgs)

`dlgs` is a cross-platform library for displaying dialogs and input boxes.

### Installation

    go get -u github.com/gen2brain/dlgs

### Documentation

Documentation on [GoDoc](https://godoc.org/github.com/gen2brain/dlgs).

### Examples

```go
item, _, err := dlgs.List("List", "Select item from list:", []string{"Bug", "New Feature", "Improvement"})
if err != nil {
    panic(err)
}
```

```go
passwd, _, err := dlgs.Password("Password", "Enter your API key:")
if err != nil {
    panic(err)
}
```

```go
yes, err := dlgs.Question("Question", "Are you sure you want to format this media?", true)
if err != nil {
    panic(err)
}
```

## More

For cross-platform notifications and alerts see [beeep](https://github.com/gen2brain/beeep).
