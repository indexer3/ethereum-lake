import pandas as pd

from pyspark.sql import SparkSession
from pyspark.sql import Row, column
from pyspark.sql import Window
import pyspark.sql.functions as F
from pyspark.sql.functions import sum, count, col, expr, desc, max, pow, first, format_number, regexp_replace, lpad
from pyspark.sql.types import StringType, StructType, StructField, IntegerType, DecimalType
from pyspark import StorageLevel
import json


erc20Tokens = json.load(open("erc20Tokens.json", "r"))


def getTokensList():
    tokens = []
    for token in erc20Tokens:
        tokens.append(token['token'].lower())
    return tokens

def getTokensListStr():
    tokens = []
    for token in erc20Tokens:
        tokens.append('\''+token['token'].lower() +'\'')
    return tokens

def calculate_eth_ft_whales(spark):
    spark.sql("use ethereum")
    tokenDf = spark.sql("select contract_address, symbol, decimals, usdPrice from erc20_token")
    tokenDf = tokenDf.withColumn("tLowerContract", F.lower(F.col("contractAddress")))
    tokenDf = tokenDf.where(F.col("tLowerContract").isin(getTokensList()))
    tokenDf = tokenDf.drop("contractAddress")
    
    wind = Window.partitionBy(F.col("lContractAddress"), F.col("accountAddress")).orderBy(desc(F.col("blockNumber")))
    
    erc20Df = spark.sql("select LOWER(contractAddress) as lContractAddress, accountAddress, blockNumber, value from ethereum_mainnet.erc20_balances where LOWER(contractAddress) in (" + ",".join(getTokensListStr()) + ")")
    erc20Df = erc20Df.withColumn("maxBlockNumber", F.first(F.col("blockNumber")).over(wind))
    erc20Df = erc20Df.where(F.col("maxBlockNumber") == F.col("blockNumber"))
    erc20Df = erc20Df.drop("maxBlockNumber").drop("blockNumber")
    erc20Df = erc20Df.groupBy(F.col("accountAddress"), F.col("lContractAddress")).agg(F.sum(F.col("value")).alias("sumValue"))

    erc20Df = erc20Df.join(tokenDf, tokenDf.tLowerContract == erc20Df.lContractAddress, "inner")
    
    erc20Df = erc20Df.withColumn("usdAmount", (F.col("sumValue") / F.pow(10, F.col("decimals"))) * F.col("usdPrice"))
    erc20Df = erc20Df.drop("tLowerContract").drop("sumValue").drop("decimals").drop("usdPrice")


    erc20Df = erc20Df.groupBy("accountAddress").agg(F.sum(F.col("usdAmount")).alias("usdAmount"))
    erc20Df = erc20Df.orderBy(desc("usdAmount")).limit(1000)

    erc20Df = erc20Df.withColumn("usdAmount", regexp_replace(format_number(col("usdAmount"), 0), ",", ""))
    erc20Df = erc20Df.withColumn("usdAmount", lpad(col("usdAmount"), 80, '0'))
    # erc20Df = erc20Df.withColumn("usdAmount", col("usdAmount").cast(StringType()))
    erc20Df = erc20Df.withColumn("updatedAt", F.current_timestamp())
    # erc20Df.show(20, False)
    
    erc20Df.write \
    .format("tidb").option("tidb.addr", "addr") \
    .option("tidb.port", "port").option("tidb.user", "root") \
    .option("tidb.password", "password") \
    .option("database", "eth_nft").option("table", "ft_whales") \
    .option("replace", "true").mode("append").save()



if __name__ == "__main__":
    spark = SparkSession \
        .builder.config("spark.memory.offHeap.enabled","true").config("spark.memory.offHeap.size","4g").config("spark.sql.broadcastTimeout", "600") \
        .appName("ETH erc20 whale") \
        .getOrCreate()
    calculate_eth_ft_whales(spark)
    spark.stop()



