# Excel-to-Markdown

====================================

## Usage

coming soon ...

## Demo

### Input `.xlsx`

[sample](https://github.com/kyokomi/excel-to-markdown/blob/master/test/excel/sample.xlsx)

![/excel-to-markdown_demo1.png](https://dl.dropbox.com/u/49084962/excel-to-markdown_demo1.png)

### Output `.md`

[sample](https://github.com/kyokomi/excel-to-markdown/blob/master/test/build/sample/sheet1.md)

```
# GoでExcelをMarkdownにする


## 目的とゴール

まずExcelで仕様を管理しているが、変更が多すぎるわりに変更管理がちゃんとできていない。レビューの体制もちゃんと取れないという問題がある。

エンジニアも見やすく、差分管理も楽なMarkdownで書いてくれれば一番良いのだが、この手の新しい取り組みは、よほど儲かっていて時間や人に余裕がある会社しかなかなかできないと思われる。

（弊社がまさにそれで、そんなことやってる場合があったら〜とよく一蹴される。やったほうが後々の効率は良くなるという説明も無駄に等しい）

となると、こちらから歩みよるしかなくて、Excelがつかいたいのは譲るが文章のフォーマットは整えてもらいそれをMarkdownに変換することにした。


## メリット

- 変更履歴を自動的に管理できるようになる
- 仕様書もレビューできるようになる

## デメリット

- フォーマットが矯正されて少しExclelが使いにくくなる
- 図やシェイプなどが使えない（使ってもいいけどMarkdown変換できない）
- 色等の視覚情報での説明ができない
- 画像のリンクを置くのがめんどくさそう

## 対応が必要なフォーマット

|No|フォーマット|ルール|
| --- | --- | --- |
|1|タイトル見出し|シートの1行目が#見出し|
|2|区切りの見出し|空白行の後が##見出し|
|3|リスト表示|1セル目が空白|
|4|表形式の表示|2セル以上の利用|
|5|画像の表示|先頭「http〜の」、先頭「/」とかパスっぽいやつ|
|6|本文（通常）|上記以外|

## ゴーファー君

![https://dl.dropbox.com/u/49084962/gopher.png](https://dl.dropbox.com/u/49084962/gopher.png)
```

## Author

[kyokomi](https://github.com/kyokomi)

## License

[MIT](https://github.com/kyokomi/excel-to-markdown/blob/master/LICENSE)

