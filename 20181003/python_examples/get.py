import requests

response = requests.get('https://api.tech.gemcook.org/v1/gems')
print(response.status_code)
print(response.text)