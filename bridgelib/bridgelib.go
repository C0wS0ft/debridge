package bridgelib

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type Bridgeprovider interface {
	Estimate(BridgeRequestEstimate) error
	Transaction(BridgeRequestTransaction) error
}

var provider Bridgeprovider

func queryGetInt(query url.Values, param string) (int, error) {
	value, present := query[param]

	if !present {
		return 0, errors.New("missing value: " + param)
	}

	s2i, err := strconv.Atoi(value[0])

	if err != nil {
		return 0, err
	}

	return s2i, nil
}

func Estimate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Estimate")

	if provider == nil {
		http.Error(w, "Bridgelib not initialized, use Init() to set provider", http.StatusBadRequest)
		return
	}

	query := r.URL.Query()

	srcChainId, err := queryGetInt(query, "srcChainId")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	srcChainTokenIn, present := query["srcChainTokenIn"]

	if !present {
		http.Error(w, "Missing srcChainTokenIn", http.StatusBadRequest)
		return
	}

	srcChainTokenInAmount, err := queryGetInt(query, "srcChainTokenInAmount")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dstChainId, err := queryGetInt(query, "dstChainId")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dstChainTokenOut, present := query["dstChainTokenOut"]

	if !present {
		http.Error(w, "Missing dstChainTokenOut", http.StatusBadRequest)
		return
	}

	request := BridgeRequestEstimate{
		srcChainId:            srcChainId,
		srcChainTokenIn:       srcChainTokenIn[0],
		srcChainTokenInAmount: srcChainTokenInAmount,
		dstChainId:            dstChainId,
		dstChainTokenOut:      dstChainTokenOut[0],
	}

	err = provider.Estimate(request)

	if err != nil {
		fmt.Println(err)
	}
}

func Transaction(w http.ResponseWriter, r *http.Request) {

	if provider == nil {
		http.Error(w, "Bridgelib not initialized, use Init() to set provider", http.StatusBadRequest)
		return
	}

	query := r.URL.Query()

	srcChainId, err := queryGetInt(query, "srcChainId")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	srcChainTokenIn, present := query["srcChainTokenIn"]

	if !present {
		http.Error(w, "Missing srcChainTokenIn", http.StatusBadRequest)
		return
	}

	srcChainTokenInAmount, err := queryGetInt(query, "srcChainTokenInAmount")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dstChainId, err := queryGetInt(query, "dstChainId")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dstChainTokenOut, present := query["dstChainTokenOut"]

	if !present {
		http.Error(w, "Missing dstChainTokenOut", http.StatusBadRequest)
		return
	}

	dstChainTokenOutRecipient, present := query["dstChainTokenOutRecipient"]

	if !present {
		http.Error(w, "Missing dstChainTokenOutRecipient", http.StatusBadRequest)
		return
	}

	request := BridgeRequestTransaction{
		srcChainId:                srcChainId,
		srcChainTokenIn:           srcChainTokenIn[0],
		srcChainTokenInAmount:     srcChainTokenInAmount,
		dstChainId:                dstChainId,
		dstChainTokenOut:          dstChainTokenOut[0],
		dstChainTokenOutRecipient: dstChainTokenOutRecipient[0],
	}

	err = provider.Transaction(request)

	if err != nil {
		fmt.Println(err)
	}
}

func Init() {
	provider = &DeswapDebridgeFinance{url: "https://deswap.debridge.finance/v1.0", apiKey: ""}
}
