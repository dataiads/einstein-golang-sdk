// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"einstein/models"
)

// SendClickRecoReader is a Reader for the SendClickReco structure.
type SendClickRecoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SendClickRecoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSendClickRecoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSendClickRecoBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSendClickRecoOK creates a SendClickRecoOK with default headers values
func NewSendClickRecoOK() *SendClickRecoOK {
	return &SendClickRecoOK{}
}

/* SendClickRecoOK describes a response with status code 200, with default header values.

A successful response contains a UUID for the given user.
*/
type SendClickRecoOK struct {
	Payload *models.ActivityResponse
}

func (o *SendClickRecoOK) Error() string {
	return fmt.Sprintf("[POST /activities/{siteId}/clickReco][%d] sendClickRecoOK  %+v", 200, o.Payload)
}
func (o *SendClickRecoOK) GetPayload() *models.ActivityResponse {
	return o.Payload
}

func (o *SendClickRecoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActivityResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendClickRecoBadRequest creates a SendClickRecoBadRequest with default headers values
func NewSendClickRecoBadRequest() *SendClickRecoBadRequest {
	return &SendClickRecoBadRequest{}
}

/* SendClickRecoBadRequest describes a response with status code 400, with default header values.

A validation error response contains an array of invalid parameters.
*/
type SendClickRecoBadRequest struct {
	Payload *models.Error
}

func (o *SendClickRecoBadRequest) Error() string {
	return fmt.Sprintf("[POST /activities/{siteId}/clickReco][%d] sendClickRecoBadRequest  %+v", 400, o.Payload)
}
func (o *SendClickRecoBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *SendClickRecoBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SendClickRecoBody send click reco body
swagger:model SendClickRecoBody
*/
type SendClickRecoBody struct {

	// The UUID that was returned in the recommendation response.
	// Required: true
	RecoUUID *string `json:"__recoUUID"`

	// client Ip
	ClientIP models.ClientIPParam `json:"clientIp,omitempty"`

	// client user agent
	ClientUserAgent models.ClientUserAgentParam `json:"clientUserAgent,omitempty"`

	// cookie Id
	CookieID models.CookieIDParam `json:"cookieId,omitempty"`

	// product
	// Required: true
	Product models.ProductForView `json:"product"`

	// realm
	Realm models.RealmParam `json:"realm,omitempty"`

	// The name of the recommender used in the recommendation request.
	// Required: true
	RecommenderName *string `json:"recommenderName"`

	// user Id
	UserID models.UserIDParam `json:"userId,omitempty"`
}

// Validate validates this send click reco body
func (o *SendClickRecoBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRecoUUID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateClientIP(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateClientUserAgent(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCookieID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateProduct(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRealm(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRecommenderName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SendClickRecoBody) validateRecoUUID(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"__recoUUID", "body", o.RecoUUID); err != nil {
		return err
	}

	return nil
}

func (o *SendClickRecoBody) validateClientIP(formats strfmt.Registry) error {
	if swag.IsZero(o.ClientIP) { // not required
		return nil
	}

	if err := o.ClientIP.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("body" + "." + "clientIp")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("body" + "." + "clientIp")
		}
		return err
	}

	return nil
}

func (o *SendClickRecoBody) validateClientUserAgent(formats strfmt.Registry) error {
	if swag.IsZero(o.ClientUserAgent) { // not required
		return nil
	}

	if err := o.ClientUserAgent.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("body" + "." + "clientUserAgent")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("body" + "." + "clientUserAgent")
		}
		return err
	}

	return nil
}

func (o *SendClickRecoBody) validateCookieID(formats strfmt.Registry) error {
	if swag.IsZero(o.CookieID) { // not required
		return nil
	}

	if err := o.CookieID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("body" + "." + "cookieId")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("body" + "." + "cookieId")
		}
		return err
	}

	return nil
}

func (o *SendClickRecoBody) validateProduct(formats strfmt.Registry) error {

	if o.Product == nil {
		return errors.Required("body"+"."+"product", "body", nil)
	}

	return nil
}

func (o *SendClickRecoBody) validateRealm(formats strfmt.Registry) error {
	if swag.IsZero(o.Realm) { // not required
		return nil
	}

	if err := o.Realm.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("body" + "." + "realm")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("body" + "." + "realm")
		}
		return err
	}

	return nil
}

func (o *SendClickRecoBody) validateRecommenderName(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"recommenderName", "body", o.RecommenderName); err != nil {
		return err
	}

	return nil
}

func (o *SendClickRecoBody) validateUserID(formats strfmt.Registry) error {
	if swag.IsZero(o.UserID) { // not required
		return nil
	}

	if err := o.UserID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("body" + "." + "userId")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("body" + "." + "userId")
		}
		return err
	}

	return nil
}

// ContextValidate validate this send click reco body based on the context it is used
func (o *SendClickRecoBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateClientIP(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateClientUserAgent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateCookieID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRealm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateUserID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SendClickRecoBody) contextValidateClientIP(ctx context.Context, formats strfmt.Registry) error {

	if err := o.ClientIP.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("body" + "." + "clientIp")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("body" + "." + "clientIp")
		}
		return err
	}

	return nil
}

func (o *SendClickRecoBody) contextValidateClientUserAgent(ctx context.Context, formats strfmt.Registry) error {

	if err := o.ClientUserAgent.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("body" + "." + "clientUserAgent")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("body" + "." + "clientUserAgent")
		}
		return err
	}

	return nil
}

func (o *SendClickRecoBody) contextValidateCookieID(ctx context.Context, formats strfmt.Registry) error {

	if err := o.CookieID.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("body" + "." + "cookieId")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("body" + "." + "cookieId")
		}
		return err
	}

	return nil
}

func (o *SendClickRecoBody) contextValidateRealm(ctx context.Context, formats strfmt.Registry) error {

	if err := o.Realm.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("body" + "." + "realm")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("body" + "." + "realm")
		}
		return err
	}

	return nil
}

func (o *SendClickRecoBody) contextValidateUserID(ctx context.Context, formats strfmt.Registry) error {

	if err := o.UserID.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("body" + "." + "userId")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("body" + "." + "userId")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SendClickRecoBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SendClickRecoBody) UnmarshalBinary(b []byte) error {
	var res SendClickRecoBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}