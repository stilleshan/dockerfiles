package dns

import (
	"encoding/json"
	"log"

	"github.com/denverdino/aliyungo/common"
)

type AddDomainRecordArgs struct {
	DomainName string
	RR         string
	Type       string
	Value      string

	//optional
	TTL  json.Number
	Line string
}

type AddDomainRecordResponse struct {
	common.Response
	InstanceId string
	RecordId   string
}

// AddDomainRecord
//
// You can read doc at https://docs.aliyun.com/#/pub/dns/api-reference/record-related&AddDomainRecord
func (client *Client) AddDomainRecord(args *AddDomainRecordArgs) (response *AddDomainRecordResponse, err error) {
	action := "AddDomainRecord"
	response = &AddDomainRecordResponse{}
	err = client.Invoke(action, args, response)
	if err == nil {
		return response, nil
	} else {
		log.Printf("%s error, %v", action, err)
		return response, err
	}
}
