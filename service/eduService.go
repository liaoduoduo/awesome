package service
import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

func (t *ServiceSetup) SaveCom(com Commodity) (string, error) {

	eventID := "eventAddCom"
	reg, notifier := regitserEvent(t.Client, t.ChaincodeID, eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)

	// 将com对象序列化成为字节数组
	b, err := json.Marshal(com)
	if err != nil {
		return "", fmt.Errorf("指定的com对象序列化时发生错误")
	}

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "addCom", Args: [][]byte{b, []byte(eventID)}}
	respone, err := t.Client.Execute(req)
	if err != nil {
		return "", err
	}

	number, err := eventResult(notifier, eventID)
	if err != nil {
		return "", err
	}
	t.BlockNumber = number
	return string(respone.TransactionID), nil
}

