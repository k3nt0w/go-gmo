package deferred

import "context"

type shippingReportRequest struct {
	ShopInfo    *shopInfo                  `xml:"shopInfo"`
	Transaction *shippingReportTransaction `xml:"transaction" validate:"required"`
}

type shippingReportTransaction struct {
	GMOTransactionID string `xml:"gmoTransactionId" validate:"required"`
	PDCompanyCode    string `xml:"pdcompanycode" validate:"required"`
	SlipNo           string `xml:"slipno" validate:"required"`
	InvoiceIssueDate string `xml:"invoiceIssueDate"`
}

type shippingReportResponse struct {
	Result          string           `xml:"result"`
	Errors          *gmoErrors       `xml:"errors"`
	TransactionInfo *transactionInfo `xml:"transactionInfo"`
}

type transactionInfo struct {
	GMOTransactionID string `xml:"gmoTransactionId"`
}

func (c *Client) PostShippingReport(ctx context.Context, req *ShippingReportRequest) (*ShippingReportResponse, error) {
	if req == nil {
		return nil, errInvalidParameterPassed
	}
	body, err := req.toParam()
	if err != nil {
		return nil, err
	}
	respParam := shippingReportResponse{}
	body.ShopInfo = &shopInfo{
		AuthenticationID: c.AuthenticationID,
		ShopCode:         c.ShopCode,
		ConnectPassword:  c.ConnectPassword,
	}
	status, err := c.doAndUnmarshalXML(ctx, POST, c.APIHost, []string{"auto", "pdrequest.do"}, map[string]string{},
		body, &respParam)
	if err != nil {
		return nil, err
	}
	resp := newShippingReportResponse(&respParam)
	resp.Status = status
	return resp, nil
}
