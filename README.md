# ethereum-lake
`ethereum-lake` provide you a one-stop tool to Extract/Transform/Load data from Ethereum like EVM Block chain.  

All you need to do is changing the destination endpoints like `MySQL` / `Postgresql` / `MongoDB` / `TiDB` / `CockroachDB`. 


### 
[![Test](https://github.com/indexer3/ethereum-lake/actions/workflows/lint-and-test.yaml/badge.svg?branch=main)](https://github.com/indexer3/ethereum-lake/actions/workflows/lint-and-test.yml)
[![GoDoc](https://godoc.org/github.com/indexer3/ethereum-lake?status.png)](https://godoc.org/github.com/indexer3/ethereum-lake)
[![GoReportCard](https://goreportcard.com/badge/github.com/indexer3/ethereum-lake)](https://goreportcard.com/report/github.com/indexer3/ethereum-lake)



## Goals 

This project aims to solve the pain points and difficulties encountered when pulling data from EVM-based public chains such as Ethereum for ETL purposes, helping users synchronize on-chain data and supporting data synchronization to various types of OLAP/OLTP/HTAP databases. And Provide users a best way to explore on-chain DeFi data with `portfolio` API with [Chainbase](https://chainbase.com/)'s SQL API or using synchronized data.



#### Key Features 

- Support for multiple databases to meet users' various scenario requirements. OLTP database with `MySQL` / `PostgreSQL` protocol and OLAP database with `ClickHouse` protocol.  
- Support on-chain DeFi protocol portfolio API. 
- Resolving [reorg issue](https://www.alchemy.com/overviews/what-is-a-reorg) - when reorg occurs, the data already consumed by the database will be corrected/modified.  
- Real-time updates - users can obtain the latest data.
- Provide blazing-fast, easy solution to synchronize data.
- Support dual-write on data to enhance the experience of HTAP. If you want to take into account the analysis scenario, you can register `MySQL` and `ClickHouse` in the service at the same time. When you need to index a certain row, you can use `MySQL` or other OLTP databases to query your data while `ClickHouse` when you need to analyze.
- Support `Spark` calculation tools for `ClickHouse` / `PostgreSQL` / `TiDB` to support massive data computing. When your dataset's getting larger and it's hard to execute some `SQL`.   
- Integrated Cache Layer to reduce the rpc call and accelerate process when synchronizing data. 



#### Use As A Client
Requirement Go 1.20+

```shell
go get github.com/indexer3/ethereum-lake
```

```go
package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/indexer3/ethereum-lake/common/client"
	"github.com/indexer3/ethereum-lake/constant"
)

func main() {
	ethCli, err := client.NewNodeClientsWithEndpoints([]string{"https://rpc.ankr.com/eth"})
	if err != nil {
		panic(err)
	}

	wethInETHContractAddress := common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2")
	uniswapInETHTokenContractAddress := common.HexToAddress("0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984")

	ctx := context.Background()

	wethPrice, err := ethCli.TokenPrice(ctx, constant.NetworkEthereum, wethInETHContractAddress, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("ethereum network: WETH price:", wethPrice)

	uniPrice, err := ethCli.TokenPrice(ctx, constant.NetworkEthereum, uniswapInETHTokenContractAddress, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("ethereum network: UNI price:", uniPrice)
}
```



#### DeFi Compatiable   
| Protocol| Client Compatiable | Portfolio Compatiable | 
| -- | -- | -- | 
| AaveV2 | ✅ | x | 
| AaveV3 | x | x | 
| Curve | x | x | 
| Balancer | ✅ | x | 
| Bancor | ✅ | x | 
| Compound | ✅ | x | 
| Convex | ✅ | x | 
| UniswapV2 | x | x | 
| UniswapV3 | x | x | 
| GMX | ✅ | x | 
| 1Inch | ✅ | x | 
| Synthetix | ✅ | x | 
| Stargate | ✅ | x | 
| Tornado | ✅ | x | 
| Yearn | ✅ | x | 



## Contribution
- Suggestions, comments (including criticisms) and contributions are welcome.
- Please follow the PR/issue template provided to ensure that your contributions are clear and easy to understand. 

## License 
[MIT License](./LICENSE)


## Donation  

Account Address (Polygon Is Preferred): 0xc101c69340feb4d0c474bf8fc34f5266f3de8a15 


