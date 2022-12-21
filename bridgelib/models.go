package bridgelib

const (
	Ethereum    = 1
	BSC         = 56
	Polygon     = 137
	ArbitrumOne = 42161
)

type BridgeRequestEstimate struct {
	srcChainId            int    // An ID of a source chain, a chain where the cross-chain swap will start
	srcChainTokenIn       string // An address (on a source chain) of an input token to swap
	srcChainTokenInAmount int    // An amount of input tokens to swap
	dstChainId            int    // An ID of a destination chain, a chain where the cross-chain swap will finish.
	dstChainTokenOut      string // An address (on a destination chain) of a target token.
}

type BridgeRequestTransaction struct {
	srcChainId                int    // An ID of a source chain, a chain where the cross-chain swap will start
	srcChainTokenIn           string // An address (on a source chain) of an input token to swap
	srcChainTokenInAmount     int    // An amount of input tokens to swap
	dstChainId                int    // An ID of a destination chain, a chain where the cross-chain swap will finish.
	dstChainTokenOut          string // An address (on a destination chain) of a target token.
	dstChainTokenOutRecipient string // Address (on the destination chain) where target tokens should be transferred to after the swap
}

type Token struct {
	address  string
	name     string
	symbol   string
	decimals int
}

type TokenWithAmount struct {
	address  string
	name     string
	symbol   string
	decimals int
	amount   string
}

type TokenWithMinAmount struct {
	address   string
	name      string
	symbol    string
	decimals  int
	amount    string
	minAmount string
}

type Tx struct {
	to    string
	data  string
	value string
}

// CrossChainEstimation

//	description: A details of an input token (on a source chain).
type srcChainTokenIn struct {
	address  string
	name     string
	symbol   string
	decimals int
	amount   string
}

/*
	description:	This variable contains optional details of an intermediary token (on a source chain).
	If the route planner decides to perform a swap before bridging, the cross-chain transaction will swap an input token to this token, which in turn will be sent to the deBridge gate.
	Otherwise, an input token (srcChainTokenIn) token acts as an intermediary and will be sent to the deBridge gate as is.
*/
type srcChainTokenOut struct {
	address   string
	name      string
	symbol    string
	decimals  int
	amount    string
	minAmount string
}

/*
	description:
	This variable contains optional details of an intermediary token (on a destination chain).
	If the route planner decides to perform a swap after bridging, the deBridge gate will unlock this particular token, and then the cross-chain transaction will swap it to a target token (dstChainTokenOut).
	Otherwise, a target (dstChainTokenOut) token acts as an intermediary and a target token simultaneously: the deBridge gate will unlock it and send it to a recipient as is.
*/
type dstChainTokenIn struct {
	address   string
	name      string
	symbol    string
	decimals  int
	amount    string
	minAmount string
}

/*
   description:
   	A details of a target token (on a destination chain), including the estimated outcome, and the estimated minimum possible outcome (affected by a slippage constraint).
*/
type dstChainTokenOut struct {
	address   string
	name      string
	symbol    string
	decimals  int
	amount    string
	minAmount string
}

/*
	description:
	The execution fee is a small amount of the intermediary token that incentivizes anyone to execute the transaction on the destination chain.This object contains details of the token representing execution fee currency (equal to either dstChainTokenIn or dstChainTokenOut), a recommended amount calculated by the planner, and an actual amount used during route construction.

*/
type executionFee struct {
	Token
	recommendedAmount string
	actualAmount      string
}

type CrossChainEstimation struct {
	srcChainTokenIn
	srcChainTokenOut
	dstChainTokenIn
	dstChainTokenOut
	executionFee
}

type EstimationReply struct {
	Estimation CrossChainEstimation `json:"estimation"`
}

type TransactionReply struct {
	Estimation CrossChainEstimation `json:"estimation"`
	Tx         Tx                   `json:"tx"`
}
