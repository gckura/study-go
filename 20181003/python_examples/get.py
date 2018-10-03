# -- coding: utf-8 --
import requests
# pip --versionでpipコマンドがなかった場合

# 1. `sudo easy_install pip` で pip をインストールする。
# 2. `pip install requests --user` で requests をインストールする。
# 上記の手順で requests というライブラリが使用できるようになる。

# - requests とは -
# Python の HTTP ライブラリ。これ一つでGETもPOSTもできる

# requestsの requests.getでデータを取得する
response = requests.get('https://api.tech.gemcook.org/v1/gems')

# ステータスコードを標準出力に出力する
print(response.status_code)

# レスポンスのテキストを標準出力に出力する
print(response.text)

