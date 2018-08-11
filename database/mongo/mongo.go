package mongo

const MongoURL = "mongodb://192.168.8.104"

// const MongoURL = "mongodb://192.168.0.102"
const Database = "tkBot"

const AssetCollectin = "Asset"
const TradeRecordCollection = "TradeRecord"
const ExchangeCollection = "Exchanges"
const BalancesCollection = "Balances"
const FundCollection = "Funds"
const OkexDiffHistory = "OkexDiffHistory"

const ErrorNotConnected = "Mongo is not connected"

type DBConfig struct {
	CollectionName string
}

