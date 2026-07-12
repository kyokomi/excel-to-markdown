# Excel-to-Markdown

[![CI](https://github.com/kyokomi/excel-to-markdown/actions/workflows/ci.yml/badge.svg)](https://github.com/kyokomi/excel-to-markdown/actions/workflows/ci.yml)

## Description

Convert Excel (`.xlsx`) files to GitHub-Flavored Markdown; see:

https://docs.github.com/en/get-started/writing-on-github

## Installation

```
$ go install github.com/kyokomi/excel-to-markdown@latest
```

Or download a pre-built binary from [Releases](https://github.com/kyokomi/excel-to-markdown/releases).

## Usage

### Help

```
$ excel-to-markdown --help
NAME:
   excel-to-markdown - convert Excel (.xlsx) files to GitHub-Flavored Markdown

USAGE:
   excel-to-markdown [global options]

VERSION:
   0.2.0

AUTHOR:
   kyokomi <kyoko1220adword@gmail.com>

GLOBAL OPTIONS:
   --input-dir string, -i string   convert target directory path
   --output-dir string, -o string  dist directory after convert path
   --help, -h                      show help
   --version, -v                   print the version
```

### Running

```
$ excel-to-markdown --input-dir example/excel --output-dir example/build
```

## Format Rules

|No|フォーマット|ルール|
| --- | --- | --- |
|1|タイトル見出し|シートの1行目が#見出し|
|2|区切りの見出し|空白行の後が##見出し|
|3|リスト表示|1セル目が空白|
|4|表形式の表示|2セル以上の利用|
|5|画像の表示|先頭が「http〜」|
|6|本文（通常）|上記以外|

## Demo

### Input `.xlsx`

[example/excel/sample.xlsx](https://github.com/kyokomi/excel-to-markdown/blob/master/example/excel/sample.xlsx)

### Output `.md`

[example/build/sample/sheet1.md](https://github.com/kyokomi/excel-to-markdown/blob/master/example/build/sample/sheet1.md)

## Author

[kyokomi](https://github.com/kyokomi)

## License

[MIT](https://github.com/kyokomi/excel-to-markdown/blob/master/LICENSE)
