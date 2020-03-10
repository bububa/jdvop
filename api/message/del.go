package message

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/bububa/jdvop"
)

type DelRequest struct {
	Id uint64 `json:"id"`
}

func (this *DelRequest) Method() string {
	return "api/message/del"
}

func (this *DelRequest) Values() url.Values {
	values := url.Values{}
	values.Set("id", strconv.FormatUint(this.Id, 10))
	return values
}

func Del(clt *jdvop.Client, msgId uint64) (bool, error) {
	resp, err := clt.Do(&DelRequest{Id: msgId})
	if err != nil {
		return false, err
	}
	var ret bool
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return false, err
	}
	return ret, nil
}
