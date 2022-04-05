# Flight Min Api
![!Theme Image](resources/banner.png)

## Take-Home Programming Assignment for GoLang

A take home assignment in Golang

### Story: 
There are over 100,000 flights a day, with millions of people and cargo being transferred around the world. With so many people and different carrier/agency groups, it can be hard to track where a person might be. In order to determine the flight path of a person, we must sort through all of their flight records.

### Goal: 
To create a simple microservice API that can help us understand and track how a particular person’s flight path may be queried. The API should accept a request that includes a list of flights, which are defined by a source and destination airport code. These flights may not be listed in order and will need to be sorted to find the total flight paths starting and ending airports.

## Install & Run

```
    git clone https://github.com/afa7789/flight-min-api
    cd flight-min-api
    go run .
```

## Tree, Project Structure

```
.
├── cmd
│   └── server.go
├── go.mod
├── go.sum
├── internal
│   ├── domain
│   │   └── flags.go
│   ├── flight
│   │   ├── flight.go
│   │   └── flight_test.go
│   └── server
│       ├── flight_points.go
│       ├── flight_points_test.go
│       ├── server.go
│       └── server_test.go
├── main.go
├── Makefile
└── README.md

5 directories, 13 files
```

## Tools used:

- Markdown , for readme, rofl
- Lint
- Unit Test
- MakeFile
- CI , github actions

## Examples 

Curl to test on insomnia
``` curl
curl --request GET \
  --url http://127.0.0.1:8080/flight_points \
  --header 'Content-Type: application/json' \
  --data '{"paths":[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]] }'
```

Requests being made to '/flight_points':
```
{"paths":[["SFO", "EWR"]]}
output:
["SFO","EWR"]

{"paths":[["ATL", "EWR"], ["SFO", "ATL"]] }
output:
["SFO","EWR"]

{"paths":[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]] }
output:
["SFO","EWR"]
```