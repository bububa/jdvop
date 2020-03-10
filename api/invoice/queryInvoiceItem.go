package invoice

import (
	"encoding/json"
	"net/url"

	"github.com/bububa/jdvop"
)

type QueryInvoiceItemRequest struct {
	InvoiceId   string `json:"invoiceId"`
	InvoiceCode string `json:"invoiceCode"`
}

func (this *QueryInvoiceItemRequest) Method() string {
	return "api/invoice/queryInvoiceItem"
}

func (this *QueryInvoiceItemRequest) Values() url.Values {
	values := url.Values{}
	values.Set("invoiceId", this.InvoiceId)
	values.Set("invoiceCode", this.InvoiceCode)
	return values
}

func QueryInvoiceItem(clt *jdvop.Client, invoiceId string, invoiceCode string) (*InvoiceItem, error) {
	resp, err := clt.Do(&QueryInvoiceItemRequest{InvoiceId: invoiceId, InvoiceCode: invoiceCode})
	if err != nil {
		return nil, err
	}
	var ret InvoiceItem
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
