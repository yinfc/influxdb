package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/influxdata/chronograf/models"
)

/*GetSourcesIDUsersOK An array of users

swagger:response getSourcesIdUsersOK
*/
type GetSourcesIDUsersOK struct {

	// In: body
	Payload *models.Users `json:"body,omitempty"`
}

// NewGetSourcesIDUsersOK creates GetSourcesIDUsersOK with default headers values
func NewGetSourcesIDUsersOK() *GetSourcesIDUsersOK {
	return &GetSourcesIDUsersOK{}
}

// WithPayload adds the payload to the get sources Id users o k response
func (o *GetSourcesIDUsersOK) WithPayload(payload *models.Users) *GetSourcesIDUsersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get sources Id users o k response
func (o *GetSourcesIDUsersOK) SetPayload(payload *models.Users) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSourcesIDUsersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetSourcesIDUsersDefault Unexpected internal service error

swagger:response getSourcesIdUsersDefault
*/
type GetSourcesIDUsersDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetSourcesIDUsersDefault creates GetSourcesIDUsersDefault with default headers values
func NewGetSourcesIDUsersDefault(code int) *GetSourcesIDUsersDefault {
	if code <= 0 {
		code = 500
	}

	return &GetSourcesIDUsersDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get sources ID users default response
func (o *GetSourcesIDUsersDefault) WithStatusCode(code int) *GetSourcesIDUsersDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get sources ID users default response
func (o *GetSourcesIDUsersDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get sources ID users default response
func (o *GetSourcesIDUsersDefault) WithPayload(payload *models.Error) *GetSourcesIDUsersDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get sources ID users default response
func (o *GetSourcesIDUsersDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetSourcesIDUsersDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
