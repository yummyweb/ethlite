# ETHLite

A very simple implementation of a light client for the Ethereum blockchain. To run ETHLite, an execution node is required, in order to query and download block headers. Alternatively, more data can be requested from a beacon node, however, that is not necessary to run ETHLite. The goal for this project is to make the simplest possible client that qualifies as a light client. Currently, ETHLite is only tested with Geth in dev mode.
The only external dependency used is `go-ethereum` for the sole purpose of querying the execution client. To make ETHLite less dependent, we could use the HTTP API instead of the client library. This is additional work, and would require reinvention of the wheel to implement all the APIs, but could be considered in the future.

## Dependencies

- Go
- Make
- Geth (Or any other execution client like Besu)

To use Geth in dev mode:
`geth --dev --http --http.api eth,web3,net --datadir dev-chain`