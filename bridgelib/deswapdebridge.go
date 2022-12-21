package bridgelib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DeswapDebridgeFinance struct {
	url    string
	apiKey string
}

func (a *DeswapDebridgeFinance) Estimate(request BridgeRequestEstimate) error {

	fmt.Println("DeswapDebridgeFinance::Estimate")

	getRequest := fmt.Sprintf("%s/estimation?srcChainId=%d&srcChainTokenIn=%s&srcChainTokenInAmount=%d&dstChainId=%d&dstChainTokenOut=%s",
		a.url,
		request.srcChainId,
		request.srcChainTokenIn,
		request.srcChainTokenInAmount,
		request.dstChainId,
		request.dstChainTokenOut,
	)

	fmt.Println(getRequest)

	resp, err := http.Get(getRequest)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	reqData := new(EstimationReply)
	err = json.Unmarshal(body, &reqData)

	if err != nil {
		return err
	}

	fmt.Println(reqData)

	return nil
}

func (a *DeswapDebridgeFinance) Transaction(request BridgeRequestTransaction) error {

	getRequest := fmt.Sprintf("%s?srcChainId=%d&srcChanTokenIn=%s&srcChainTokenInAmount=%d&dstChainId=%d&dstChainTokenOut=%s&dstChainTokenOutRecipient=%s",
		a.url,
		request.srcChainId,
		request.srcChainTokenIn,
		request.srcChainTokenInAmount,
		request.dstChainId,
		request.dstChainTokenOut,
		request.dstChainTokenOutRecipient,
	)

	resp, err := http.Get(getRequest)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	reqData := new(TransactionReply)
	err = json.Unmarshal(body, &reqData)

	if err != nil {
		return err
	}

	fmt.Println(reqData)

	return nil
}
