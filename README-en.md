# gwebp

[![Build Status](https://travis-ci.org/linehk/gwebp.svg?branch=master)](https://travis-ci.org/linehk/gwebp)
[![Go Report Card](https://goreportcard.com/badge/github.com/linehk/gwebp)](https://goreportcard.com/report/github.com/linehk/gwebp)

English | [简体中文](./README.md "简体中文")

gwebp can recursively convert images in the specified directory into webp of the same name.

## Installation

Need to install webp first, such as: `brew install webp`, then:

```bash
go install github.com/linehk/gwebp
```

## Usages

```bash
gwebp -r='.' -e='jpg' -a='-q 50 -lossless' -k=true
```

```bash
gwebp -r='.' -e='png' -a='-sns 70 -f 50 -size 60000' -k=false
```

## License

[MIT License](./LICENSE "MIT License")