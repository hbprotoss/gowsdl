package enterprise

import "encoding/xml"

type GetEnterpriseRequest struct {
	XMLName xml.Name `xml:"ser:getEnterprise"`

	UserId int32  `xml:"userId"`
	Local  string `xml:"local"`
}

func NewGetEnterpriseRequest() *GetEnterpriseRequest {
	return &GetEnterpriseRequest{}
}
