# GO / Mongo Atlas API ☁️

[![mongobadge](https://img.shields.io/badge/Mongodb-signup-green.svg)](https://www.mongodb.com/atlas/database)   [![mongobadge](https://img.shields.io/badge/go-mongodrivers-orange.svg)](https://docs.mongodb.com/drivers/go/current) <br/>
![MongoDB](https://img.shields.io/badge/MongoDB-%234ea94b.svg?style=for-the-badge&logo=mongodb&logoColor=white)  ![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)  ![HTML5](https://img.shields.io/badge/html5-%23E34F26.svg?style=for-the-badge&logo=html5&logoColor=white)  ![CSS3](https://img.shields.io/badge/css3-%231572B6.svg?style=for-the-badge&logo=css3&logoColor=white)  ![JavaScript](https://img.shields.io/badge/javascript-%23323330.svg?style=for-the-badge&logo=javascript&logoColor=%23F7DF1E)

Basic API written in GO to a Cloud Mongodb, simply stores JSON information for you to call, also has very basic frontend written in for the post requests. Very simple and decided to open source as i didnt see one that was ready to deploy. Just create Mongodb using the link below and select GO drivers when getting URI for DB.



* Good for small personal projects and quick deployment.
* Anything you feel i missed or could be improved/structured please dm me


## Example requests

### <ins>GET</ins>
```
import requests

url = "http://YOUR_APIURL"


headers = {"Content-Type": "application/json"}

response = requests.request("GET", url, headers=headers)

print(response.text)
```

### <ins>POST</ins>

```
url = "http://YOUR_APIURL"

payload = {
    "isbn": "Some",
    "title": "Title",
    "info": {
        "info1": "And",
        "info2": "Information",
        "info3": "N/A",
        "info4": "N/A"
    }
}
headers = {"Content-Type": "application/json"}

response = requests.request("POST", url, json=payload, headers=headers)

print(response.text)
```

### <ins>PUT</ins>

```
url = "http://YOUR_APIURL/{_id}"

payload = {
    "isbn": "Some",
    "title": "Corrected",
    "info": {
        "info1": "Information",
        "info2": "N/A",
        "info3": "N/A",
        "info4": "N/A"
    }
}
headers = {"Content-Type": "application/json"}

response = requests.request("PUT", url, json=payload, headers=headers)

print(response.text)
```

### <ins>DELETE</ins>

```
url = "http://YOUR_APIURL/{_id}"

headers = {"Content-Type": "application/json"}

response = requests.request("DELETE", url, headers=headers)

print(response.text)
```

### <ins>GET</ins> (Specific ID)

```
url = "http://YOUR_APIURL/{_id}"

headers = {"Content-Type": "application/json"}

response = requests.request("GET", url, headers=headers)

print(response.text)
```

## To do:
- [ ] Build frontend for PUT/DELETE/GET
- [ ] Maybe build on POSTGRES
- [ ] Build queue system to carry out functions

## Useful links
* [go.dev](https://go.dev/)
* [Getting started with heroku/go](https://devcenter.heroku.com/articles/getting-started-with-go)


![Discord](https://img.shields.io/badge/Discord-7289DA?style=for-the-badge&logo=discord&logoColor=white) <br/>```@lame impala#0304```

![toms stats](https://github-readme-stats.vercel.app/api?username=trapmorrissey&show_icons=true&theme=dark)
