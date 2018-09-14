# MySQL接続プログラム起動手順

Go言語でxormを使用してMySQLに接続するサンプルプログラムの起動手順の説明です。

## プログラムのダウンロード手順

git コマンドでプログラムのソースコードをダウンロードしてきて、VSCodeで開くところまでのコマンドの実行例です。

```sh
mkdir -p $GOPATH/src/github.com/gemcook/
cd go/src/github.com/gemcook/
git clone https://github.com/gemcook/study-go/
cd study-go/20180914/
code .
```

## 外部パッケージのインストール

1. 外部パッケージのインストールコマンドを実行します。
    ```sh
    go get -v ./...
    ```

2. 動作させたいプログラムのディレクトリに移動します。
    ```sh
    cd readOne
    ```

3. プログラムを起動します。
    ```sh
    go run main.go
    ```
