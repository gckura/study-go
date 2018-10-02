# -- coding: utf-8 -- 
import requests

payload = {
    "name":"ペリドット",
    "japanese_name":"かんらん石",
    "english_name":"peridot",
    "hardness":9.6,
    "price":8000,
    "producing_area":"ハワイ",
    "carat":8.23,
    "weight":19.6,
    "color":"緑",
    "memo":"形状:ラウンド",
    "scratched":True,
    "mining_date":"2017-08-23T00:00:00Z"
    }

response = requests.post('https://api.tech.gemcook.org/v1/gems', json=payload)
print(response.status_code)
print(response.text)