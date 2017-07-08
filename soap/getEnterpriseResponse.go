package soap

import "encoding/xml"

type GetEnterpriseResponse struct {
	XMLName xml.Name `xml:"http://service.enterprise.soa.gttown.com/ getEnterpriseResponse"`

	enterpriseId int32
}
