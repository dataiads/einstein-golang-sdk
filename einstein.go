package einstein

import (
	apiclient "github.com/Dataiads/einstein-golang-sdk/client"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

func NewEinsteinClient(xCqClientID string) *apiclient.Einstein {
	transport := httptransport.New(apiclient.DefaultHost, apiclient.DefaultBasePath, apiclient.DefaultSchemes)
	transport.DefaultAuthentication = httptransport.APIKeyAuth("x-cq-client-id", "header", xCqClientID)

	client := apiclient.New(transport, strfmt.Default)

	return client
}
