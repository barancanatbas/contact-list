# contact-list

### an app for experimenting with docker and docker compose
### the purpose of the application is a small guide api

## RUN

```bash

docker-compose -f "docker-compose.yml" up -d --build 

```

## OUTPUT

```bash
                                                                                                        
[+] Running 2/2
 - Container contactlist-db-1          Started                              2.5s 
 - Container contactlist-1             Started                              6.2s

```

### now started my app, We can send requests with curl

```bash
 curl http://localhost:8080/contacts
```
### output
```json
 {
  "data": [
    {
      "ID": 1,
      "CreatedAt": "2022-03-05T20:41:28.346Z",
      "UpdatedAt": "2022-03-05T20:41:28.346Z",
      "DeletedAt": null,
      "name": "baran can atbaş",
      "Phone": "5554567788",
      "Description": "bu benim hayalimdeki adam"
    },
    {
      "ID": 2,
      "CreatedAt": "2022-03-05T20:56:22.743Z",
      "UpdatedAt": "2022-03-05T20:56:22.743Z",
      "DeletedAt": null,
      "name": "baran can atbaş",
      "Phone": "5554567788",
      "Description": "bu benim hayalimdeki adam"
    }
  ],
  "message": "pong"
}
```
