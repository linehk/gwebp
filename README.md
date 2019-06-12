# gwebp

[![Build Status](https://travis-ci.org/linehk/gwebp.svg?branch=master)](https://travis-ci.org/linehk/gwebp)
[![Go Report Card](https://goreportcard.com/badge/github.com/linehk/gwebp)](https://goreportcard.com/report/github.com/linehk/gwebp)

English | [简体中文](./README-zh_CN.md "简体中文")

gwebp can recursively convert images in the specified directory into webp of the same name.

## Installation

Need to install webp first, such as: `brew install webp`, then:

```bash
git clone https://github.com/linehk/gwebp.git
```

Or:

```bash
go get -u github.com/linehk/gwebp
```

## Usages

```bash
go run main.go
```

```bash
go run main.go -r='.' -e='jpg' -a='-q 50 -lossless' -k=true
```

```bash
go run main.go -r='.' -e='png' -a='-sns 70 -f 50 -size 60000' -k=false
```

## License

[MIT](https://choosealicense.com/licenses/mit/)