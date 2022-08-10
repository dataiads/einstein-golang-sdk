// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"einstein-golang-sdk/models"
)

// GetRecommendersReader is a Reader for the GetRecommenders structure.
type GetRecommendersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRecommendersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRecommendersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRecommendersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRecommendersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRecommendersOK creates a GetRecommendersOK with default headers values
func NewGetRecommendersOK() *GetRecommendersOK {
	return &GetRecommendersOK{}
}

/* GetRecommendersOK describes a response with status code 200, with default header values.

A successful response will contain a list of recommenders available for use in recommendation requests.
*/
type GetRecommendersOK struct {
	Payload *models.RecommendersResponse
}

func (o *GetRecommendersOK) Error() string {
	return fmt.Sprintf("[GET /personalization/recommenders/{siteId}][%d] getRecommendersOK  %+v", 200, o.Payload)
}
func (o *GetRecommendersOK) GetPayload() *models.RecommendersResponse {
	return o.Payload
}

func (o *GetRecommendersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RecommendersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecommendersBadRequest creates a GetRecommendersBadRequest with default headers values
func NewGetRecommendersBadRequest() *GetRecommendersBadRequest {
	return &GetRecommendersBadRequest{}
}

/* GetRecommendersBadRequest describes a response with status code 400, with default header values.

Either there was no OCAPI client id in the request, or the given OCAPI client id was not valid.
*/
type GetRecommendersBadRequest struct {
	Payload *models.Error
}

func (o *GetRecommendersBadRequest) Error() string {
	return fmt.Sprintf("[GET /personalization/recommenders/{siteId}][%d] getRecommendersBadRequest  %+v", 400, o.Payload)
}
func (o *GetRecommendersBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRecommendersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecommendersForbidden creates a GetRecommendersForbidden with default headers values
func NewGetRecommendersForbidden() *GetRecommendersForbidden {
	return &GetRecommendersForbidden{}
}

/* GetRecommendersForbidden describes a response with status code 403, with default header values.

The given site ID is not valid.
*/
type GetRecommendersForbidden struct {
	Payload *models.Error
}

func (o *GetRecommendersForbidden) Error() string {
	return fmt.Sprintf("[GET /personalization/recommenders/{siteId}][%d] getRecommendersForbidden  %+v", 403, o.Payload)
}
func (o *GetRecommendersForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRecommendersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
