package service

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"time"
)
type Commodity struct {
	ObjectType	string	`json:"docType"`
	Name string `json:"Name"`
	EntityID string `json:"EntityID"`
	FileHash string `json:"FileHash"`
	FileType string `json:"FileType"`
	Desc string `json:"Desc"`
	Company string `json:"Company"`

	Historys []HistoryItem
}

type HistoryItem struct {
	TxId string
	Commodity Commodity
}

type ServiceSetup struct {
	ChaincodeID	string
	Client	*channel.Client
	BlockNumber uint64
}

func regitserEvent(client *channel.Client, chaincodeID, eventID string) (fab.Registration, <-chan *fab.CCEvent) {

	reg, notifier, err := client.RegisterChaincodeEvent(chaincodeID, eventID)
	if err != nil {
		fmt.Println("注册链码事件失败: %s", err)
	}
	return reg, notifier
}

func eventResult(notifier <-chan *fab.CCEvent, eventID string)(uint64, error) {
	select {
	case ccEvent := <-notifier:
		fmt.Printf("接收到链码事件: %v\n", ccEvent)

		return ccEvent.BlockNumber, nil
	case <-time.After(time.Second * 20):
		return 0, fmt.Errorf("不能根据指定的事件ID接收到相应的链码事件(%s)", eventID)
	}
	return 0, nil
}

