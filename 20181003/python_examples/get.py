# -- coding: utf-8 --
import requests

# 1. `sudo easy_install pip` で pip をインストールする。
# 2. `pip install requests --user` で requests をインストールする。
# 上記の手順で requests というライブラリが使用できるようになる。

response = requests.get('https://api.tech.gemcook.org/v1/gems')
print(response.status_code)
print(response.text)
