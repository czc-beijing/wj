package utils

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/davecgh/go-spew/spew"
	"wj/global"
)

func SendSms() {
	client, err := dysmsapi.NewClientWithAccessKey("cn-beijing", global.Config.Sms.Key, global.Config.Sms.Secrect)
	if err != nil {
		panic(err)
	}
	smsCode := "1212"
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https" // https | http
	request.Domain = "dysmsapi.aliyuncs.com"
	request.Version = "2017-05-25"
	request.ApiName = "SendSms"
	request.QueryParams["RegionId"] = "cn-beijing"
	request.QueryParams["PhoneNumbers"] = "15810575564"                 //手机号
	request.QueryParams["SignName"] = "toolsp平台验证码"                     //阿里云验证过的项目名 自己设置
	request.QueryParams["TemplateCode"] = "SMS_127600024"               //阿里云的短信模板号 自己设置
	request.QueryParams["TemplateParam"] = "{\"code\":" + smsCode + "}" //短信模板中的验证码内容 自己生成   之前试过直接返回，但是失败，加上code成功。
	response, err := client.ProcessCommonRequest(request)
	spew.Dump(client.DoAction(request, response))
	spew.Dump(err)
	if err != nil {
		fmt.Print(err.Error())
	}
}
