[![Gitpod ready-to-code](https://img.shields.io/badge/Gitpod-ready--to--code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/monch1962/go-mssql)

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=monch1962_go-mssql&metric=alert_status)](https://sonarcloud.io/dashboard?id=monch1962_go-mssql)
[![Build Status](https://dev.azure.com/monch1962/monch1962/_apis/build/status/monch1962.go-mssql?branchName=master)](https://dev.azure.com/monch1962/monch1962/_build/latest?definitionId=9&branchName=master)
# go-mssql
This is a slightly unusual MS-SQL Server client in that you configure your connection details & the SQL statements you want to execute inside a YAML file. Running the app will then connect to that SQL Server and execute those statements. You will then get the following data returned in CSV format:
- statement ID
- response time in microseconds
- number of records returned
- SQL statement

Within the YAML file you can also request to iterate through the defined SQL statements multiple times, as well as specify the concurrency (i.e. how many SQL statements to be executing simultaneously).

To run it, all you need are 2 files: the compiled Go executable and the `database.yaml` file, which can be tweaked for your specific environment and use case.

This makes it a good fit for highly-specific tasks such as benchmarking response times and latency from end-user systems, which is the reason I created it in the first place. The intended use case is to run it from different PCs over different network links, and compare the latencies. This information can then be used to inform possible infrastructure upgrades.

## Run as executable
`$ go build`

`$ ./sqlclient > results.csv`

## Run as Docker container
...


## To test within Docker
To spin up a local MS SQL Server instance inside Docker:

`$ docker pull mcr.microsoft.com/mssql/server:2019-latest`

`$ docker run -d --name sql_server_demo -e 'ACCEPT_EULA=Y' -e 'SA_PASSWORD=reallyStrongPwd123' -p 1433:1433 mcr.microsoft.com/mssql/server:2019-latest`
