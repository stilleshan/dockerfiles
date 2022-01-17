package dns

import (
	"encoding/json"

	"github.com/denverdino/aliyungo/common"
)

// endpoint change to 'http://alidns.aliyuncs.com' then record ttl and priority change to string
type RecordTypeNew struct {
	DomainName string
	RecordId   string
	RR         string
	Type       string
	Value      string
	TTL        json.Number
	Line       string
	Status     string
	Locked     bool
}

type DescribeDomainRecordInfoNewArgs struct {
	RecordId string
}

type DescribeDomainRecordInfoNewResponse struct {
	common.Response
	RecordTypeNew
}

// DescribeDomainRecordInformation
//
// You can read doc at https://docs.aliyun.com/#/pub/dns/api-reference/record-related&DescribeDomainRecordInfo
func (client *Client) DescribeDomainRecordInfoNew(args *DescribeDomainRecordInfoNewArgs) (response *DescribeDomainRecordInfoNewResponse, err error) {
	action := "DescribeDomainRecordInfo"
	response = &DescribeDomainRecordInfoNewResponse{}
	err = client.Invoke(action, args, response)
	if err == nil {
		return response, nil
	} else {
		return nil, err
	}
}
