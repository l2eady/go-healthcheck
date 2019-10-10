# go-healthcheck

program to check website and response time when given the CSV list, calls the Healthcheck Report API to send the statistic of each website

## What is this repository for ##

* Healthy Check websites from domain name in csv list

## CSV Example ##

```csv
http://www.google.com,
http://microsoft.com,
http://apple.com,
http://youtube.com,
http://www.blogger.com,
http://support.google.com,
http://play.google.com,
http://docs.google.com,
```
## Project Structure ##

├───configs
└───internal
    └───app
        ├───healthycheck
        │   └───mocks
        ├───lhttp
        │   └───mocks
        └───models

## Testing instructions ##

```bash
go clean -testcache && go test ./... -coverprofile=coverage.out & go tool cover -html=coverage.out
```

## Running instructions ##

|       Param Name       | Required |  Type  | Default Value |              Description               |
| :--------------------: | :------: | :----: | :-----------: | :------------------------------------: |
|        filename        |   true   | string |  example.csv  |     csv filename for healthy check     |
| ping_timeout_in_second |  false   |  int   |       2       |      HTTP Timeout for ping domain      |
|       max_worker       |  false   |  int   |      50       | Maximum of worker for concurrency ping |
|         stage          |  false   | string |     local     |        set working environment         |


```bash
go run main.go -filename example.csv -max_worker 100 -ping_timeout_in_second 2
```

## Deployment instructions ##

```bash
sh deployment.sh
```

the package will move to ${GOPATH}/bin/go-healthycheck with configs folder

## How to contact ##

m.khemcharoen@gmail.com
