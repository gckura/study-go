# DBeaver インストール手順

Windows, macOS, Linuxで使える、GUIのデータベース管理ツールDBeaverのインストールを行います。
WorkSpacesで起動している、Amazon Linux 2 の場合のインストール手順を説明します。

## インストール手順

1. ダウンロードページを開きます。
    <https://dbeaver.io/download/>

2. Linux RPM package 64 (installer) をダウンロードします。

3. Terminalを開いて、rmpパッケージをインストールします。

    ```sh
    sudo rpm -ivh ~/ダウンロード/dbeaver-ce-5.2.0-stable.x86_64.rpm
    ```
    ダウンロードしたバージョンが異なる場合は、適宜変更して下さい。

## 起動手順

1. 左下のメニューから、 アプリケーション > プログラミング > DBeaver Community を開きます。

## データベース接続手順

初回起動時にはデータベース接続情報を入力する「新しい接続を作成する」
画面が自動的に起動します。

1. メニューの「データベース」 > 新しい接続 を開きます。
2. MySQL を選択して、「次へ」をクリックします。
3. MySQLの接続設定に以下の情報を入力します。

    | 項目 | 値 |
    | --- | --- |
    | Server Host | study-go.chmcz1iuszc0.us-east-1.rds.amazonaws.com |
    | User name | golang |
    | Password | study-go-15 |

4. 「テスト接続」をクリックして、問題なければ、「終了」をクリックします。
5. プロジェクトメニューの General > Connections > MySQL - xxxxx > データベース を開きます。
6. db20180914 > テーブル を開いて、 t_gems をダブルクリックします。
7. 「データ」タブを選択すると、テーブルの中身が参照できます。
