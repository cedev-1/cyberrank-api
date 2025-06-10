# cyberrank-api

Simple API to get cyber ranking of a user on Rootme and Tryhackme.
I use it to get my ranking for my Portfolio.

```bash
http://localhost:8080/api/rootme/<rootme_username>
http://localhost:8080/api/rootme/<rootme_username>?detailed=true
http://localhost:8080/api/tryhackme/<tryhackme_username>
```

## Start

```bash
go run cmd/server/main.go
```

## Example

simple rootme

```json
{
  "username": "cedev001",
  "rank": "7090",
  "platform": "root-me"
}
```

detailed rootme

```json
{
  "username": "cedev001",
  "rank": "7090",
  "platform": "root-me",
  "categories": [
    {
      "name": "Web - Client",
      "percentage": 40
    },
    {
      "name": "Programmation",
      "percentage": 20
    },
    {
      "name": "Cryptanalyse",
      "percentage": 12
    },
    {
      "name": "Stéganographie",
      "percentage": 47
    },
    {
      "name": "Web - Serveur",
      "percentage": 39
    },
    {
      "name": "Cracking",
      "percentage": 7
    },
    {
      "name": "Réaliste",
      "percentage": 3
    },
    {
      "name": "Réseau",
      "percentage": 29
    },
    {
      "name": "App - Script",
      "percentage": 12
    },
    {
      "name": "App - Système",
      "percentage": 1
    },
    {
      "name": "Forensic",
      "percentage": 0
    }
  ],
  "overall": {
    "percentage": 17,
    "solved": 103,
    "total": 599
  }
}
```

tryhackme 

```json
{
  "username": "Cedev001",
  "rank": "29402",
  "platform": "tryhackme"
}
```