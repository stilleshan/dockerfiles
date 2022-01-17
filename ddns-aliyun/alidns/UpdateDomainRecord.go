package dns

import (
	"encoding/json"

	"github.com/denverdino/aliyungo/common"
)

type UpdateDomainRecordArgs struct {
	RecordId string
	RR       string
	Type     string
	Value    string

	//optional
	TTL      json.Number
	Priority json.Number
	Line     string
}

type UpdateDomainRecordResponse struct {
	common.Response
	InstanceId string
	RecordId   string
}

// UpdateDomainRecord
//
// You can read doc at https://docs.aliyun.com/#/pub/dns/api-reference/record-related&UpdateDomainRecord
func (client *Client) UpdateDomainRecord(args *UpdateDomainRecordArgs) (response *UpdateDomainRecordResponse, err error) {
	action := "UpdateDomainRecord"
	response = &UpdateDomainRecordResponse{}
	err = client.Invoke(action, args, response)
	if err == nil {
		return response, nil
	} else {
		return nil, err
	}
}
