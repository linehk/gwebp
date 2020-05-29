# gwebp

[![Build Status](https://travis-ci.org/linehk/gwebp.svg?branch=master)](https://travis-ci.org/linehk/gwebp)
[![codecov](https://codecov.io/gh/linehk/gwebp/branch/master/graph/badge.svg)](https://codecov.io/gh/linehk/gwebp)
[![Go Report Card](https://goreportcard.com/badge/github.com/linehk/gwebp)](https://goreportcard.com/report/github.com/linehk/gwebp)

[English](./README-en.md "English") | 简体中文

gwebp 可以递归地将指定目录下的图片转换成同名的 webp 格式的图片。

## 安装

需要先行安装 webp，如：`brew install webp`，然后：

```bash
go install github.com/linehk/gwebp
```

## 使用

```bash
gwebp -r='.' -e='jpg' -a='-q 50 -lossless' -k=true
```

```bash
gwebp -r='.' -e='png' -a='-sns 70 -f 50 -size 60000' -k=false
```

## 参与贡献

如果你觉得我的参考答案有什么需要改进的地方，欢迎发起 Pull Request。

如果有重大变化，请先打开一个 Issue，讨论您想要改变的内容。

## 开源许可证

[MIT License](./LICENSE "MIT License")