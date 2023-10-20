# takuhaiコマンド

宅配物の状況をコマンドライン上で一元管理・追跡できるCLIアプリケーションです．荷物情報のインポート・エクスポート機能も有しています．

# Installation

```
$ git clone https://github.com/lapla-cogito/takuhai
$ make
$ echo PATH=$PATH:$(pwd)/takuhai/bin/takuhai >> ~/.bashrc
```

bash以外（zsh等）をお使いの方は，適宜修正してください．

## How to use

```
$ takuhai -h
A CLI application to track packages you registered.
currently, this application can track:
- SAGAWA TRANSPORTATION CO., LTD.
- YAMATO TRANSPORT CO., LTD.
- Nippon Express Co., LTD.

Usage:
  takuhai [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  dereg       Deregister the specific package
  export      Export package informations
  help        Help about any command
  import      Import package informations from exported yaml file
  reg         Registers a package
  rename      Rename a package
  show        Shows the state of a specific package

Flags:
  -h, --help   help for takuhai

Use "takuhai [command] --help" for more information about a command.
```

## 使用例
### 佐川急便が配送する，追跡番号が1234567890の荷物にhogeという名前を付けて登録する
```
$ takuhai reg --sagawa -t 1234567890 -n hoge
```

--sagawaフラグには短縮系として-sがあります．同様に，--yamatoフラグは-y，--jpostフラグは-jとできます．

## hogeと名前が付けられている荷物をfugaという名前に変更する

```
$ takuhai rename -o hoge -n fuga
```

-oオプションに変更したい荷物名を，-nオプションにどのような名前に変更するのかを指定してください．

### 追跡番号が1234567890の荷物を登録解除する
```
$ takuhai dereg -t 1234567890
```

### hogeと名前を付けた荷物を登録解除する
```
$ takuhai dereg -n hoge
```

### 登録済みの全ての荷物について配送状況を一覧表示する
```
$ takuhai show -a
```

このコマンドは，登録済みの全ての荷物について，最新の状況（輸送中とか引き渡し済みとか）のみを一覧表示します．各荷物について詳細を表示したい場合は，次の例を参照してください．

### 追跡番号が1234567890の荷物について，配送状況の詳細を確認する
```
$ takuhai show -t 1234567890
```

これは追跡番号が登録されたときから現在までの荷物の詳細な動きを表示します．

### hogeと名前が付いた荷物について，配送状況の詳細を確認する
```
$ takuhai show -n hoge
```

同様です．

### 名前がhoge，foo，barの3つの荷物を，別のyamlファイルとしてエクスポートする

```
$ takuhai export -n "hoge foo bar" -p result.yml
```

-pオプションを用いて，エクスポート先のファイルを指定できます．これは存在しない場合はtakuhaiコマンドが自動的に作成します．これを他人に送ることで，その人がimportサブコマンドを用いて荷物の情報をインポートでき，各自で配送状況を追えるようになります．

### result.ymlの内容を取り込む

```
$ takuhai import -p result.yml
```

他人からもらったyamlファイルを，importサブコマンドを用いることで自分の環境にインポートできます．

## 記事

[ここ](https://lapla.dev/posts/takuhai)
