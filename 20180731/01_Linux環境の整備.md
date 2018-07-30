# Linux環境の整備

AWS WorkSpaces で Amazon Linux 2 を起動した後の設定を行います。

## ターミナル操作のポイント

* 範囲選択してCTRL+SHIFT+C -> コピー
* 範囲選択してCTRL+SHIFT+V -> ペースト
* CTRL+C -> プログラム中断

## 設定変更用エディタ

シンプルエディタ pluma をメニューに登録

Applications -> Accessories -> Pluma Text Editor
右クリックして、　Add this launcher to panel

エディタは `pluma` コマンドでも開けます。

使用例:

```sh
pluma <開くファイル>
```

## キーボードレイアウトの設定

USキーボードの人は設定不要です。

まず日本語キーボードを設定します。

### 英語表示の場合

左下のメニュー の System -> Preferences -> Keyboard
Layouts -> +Add

* Country: Japan
* Variants: Japanese

#### 英語表示で MacBook用のキーボードの場合

Keyboard modelを以下に設定します。

* Vendors: Apple
* MacBook/MacBook Pro (intl)

#### 英語表示で Windows 用のキーボードの場合

Keyboard modelを以下に設定します。

* Vendors: Generic
* Generic 105-key (intl) PC

### 日本語表示の場合

左下のメニュー の システム -> 設定 -> キーボード
レイアウト -> +追加

* 国: 日本
* 系列: 日本語

#### 日本語表示で MacBook 用のキーボードの場合

キーボード形式を以下に設定します。

* ベンダー: Apple
* 型式: MacBook/MacBook Pro (intl)

#### 日本語表示で Windows 用のキーボードの場合

キーボード形式を以下に設定します。

* ベンター: Generic
* 型式: Generic 105-key (intl) PC

## 日本語化

AWS WorkSpaces Linux 2 を日本語で起動し、MacBookから接続している場合は、
設定不要です。

AWS WorkSpaces Linux 2 を日本語で起動し、MacBookから接続している場合は、
最後の「ショートカットキーを設定」だけ実施して下さい。

英語で起動した場合は、下記の設定が必要です。

### ターミナルの日本語化

~/.bashrcファイルを開きます。

```sh
pluma ~/.bashrc
```

以下を足します。

```sh
export LANG=ja_JP.UTF8
```

### 日本語入力設定

ibusデーモンを起動しておきます。

```sh
ibus-daemon -d
```

起動時から、日本語入力が有効になるようにします。

左下のメニュー の System -> Preferences -> Startup Applications
+Add をクリックします。

以下を設定して +Add で登録します。

```sh
Name: ibus
Command: ibus-daemon -d
Comment: ibus daemon startup
```

日本語入力の設定をします。

```sh
ibus-setup
```

入力メソッドから英語 - Englishを削除します。

直接入力モードとひらがな入力モードを切り替えれるようにショートカットキーを設定します。

#### Windows用キーボードの場合

* 全角/半角キーをひらがなの「直接入力モードに変更」のショートカットキーに設定
* 全角/半角キーを直接入力の「ひらがな入力モードに変更」のショートカットキーに設定

#### Mac用キーボードの場合

* 英数キーをひらがなの「直接入力モードに変更」のショートカットキーに設定
* かなキーを直接入力の「ひらがな入力モードに変更」のショートカットキーに設定
