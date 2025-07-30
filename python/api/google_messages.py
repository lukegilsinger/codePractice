import requests

x = requests.get('https://messages.google.com/web')

print(x.text)