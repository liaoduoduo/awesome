package main

import (
	"fmt"
	"github.com/liaoduoduo/awesome/sdkInit"
	"github.com/liaoduoduo/awesome/service"
	"os"
)

const (
	configFile = "config.yaml"
	initialized = false
	ComCC = "comcc"
)

func main() {

	initInfo := &sdkInit.InitInfo{

		ChannelID: "kevinkongyixueyuan",
		ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/liaoduoduo/awesome/fixtures/artifacts/channel.tx",

		OrgAdmin:"Admin",
		OrgName:"Org1",
		OrdererOrgName: "orderer.kevin.kongyixueyuan.com",

		ChaincodeID: ComCC,
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath: "github.com/liaoduoduo/awesome/chaincode/",
		UserName:"User1",
	}

	sdk, err := sdkInit.SetupSDK(configFile, initialized)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	defer sdk.Close()

	err = sdkInit.CreateChannel(sdk, initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	channelClient, err := sdkInit.InstallAndInstantiateCC(sdk, initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(channelClient)

	serviceSetup := service.ServiceSetup{
		ChaincodeID:ComCC,
		Client:channelClient,
	}

	com := service.Commodity{
		Name:     "廖多多",
		EntityID: "123",
		FileHash: "QmdKjne7dhQ99GxMZ5DqcN3FqYB2TrQN3a8hefFiouAC2a",
		FileType: "jpg",
		Desc:     "2021年7月，广东省广州市某小区公寓房内有多人聚众吸毒照片",
		Company:  "广州市花都区花山镇禁毒办",
	}

	msg, err := serviceSetup.SaveCom(com)
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println("信息发布成功, 交易编号为: " + msg)
	}
}