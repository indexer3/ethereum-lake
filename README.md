# ethereum-lake
`ethereum-lake` provide you a one-stop tool to Extract/Transform/Load data from Ethereum like EVM Block chain.  

All you need to do is changing the destination endpoints like `MySQL` / `Postgresql` / `MongoDB` / `TiDB` / `CockroachDB`. 


### 
[![Test](https://github.com/indexer3/ethereum-lake/actions/workflows/lint-and-test.yml/badge.svg?branch=main)](https://github.com/indexer3/ethereum-lake/actions/workflows/lint-and-test.yml)
[![GoDoc](https://godoc.org/github.com/indexer3/ethereum-lake?status.png)](https://godoc.org/github.com/indexer3/ethereum-lake)
[![GoReportCard](https://goreportcard.com/badge/github.com/indexer3/ethereum-lake)](https://goreportcard.com/report/github.com/indexer3/ethereum-lake)



## Goals 

This project aims to solve the pain points and difficulties encountered when pulling data from EVM-based public chains such as Ethereum for ETL purposes, helping users synchronize on-chain data and supporting data synchronization to various types of databases such as OLAP/OLTP/HTAP.    



#### Key Features 

- Resolving [reorg issue](https://www.alchemy.com/overviews/what-is-a-reorg) - when reorg occurs, the data already consumed by the database will be corrected/modified.  
- Real-time updates - users can obtain the latest data.
- Provide blazing-fast solution to synchronize data.
- Support for multiple databases to meet users' various scenario requirements. OLTP database with `MySQL` / `PostgreSQL` protocol and OLAP database with `ClickHouse` protocol.  
- Easy to synchronize customized data for users. 
- Support dual-write on data to enhance the experience of HTAP. If you want to take into account the analysis scenario, you can register `MySQL` and `ClickHouse` in the service at the same time. When you need to index a certain row, you can use `MySQL` or other OLTP databases to query your data while `ClickHouse` when you need to analyze.
- Support `Spark` calculation tools for `ClickHouse` / `PostgreSQL` / `TiDB` to support massive data computing. When your dataset's getting larger and it's hard to execute some `SQL`.   
- Intergate Cache Layer to reduce the rpc call and accelerate process when synchronizing data.   








