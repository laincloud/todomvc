// Code generated by go-swagger; DO NOT EDIT.

package todo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/laincloud/todomvc/gen/models"
)

// CreateTodoCreatedCode is the HTTP code returned for type CreateTodoCreated
const CreateTodoCreatedCode int = 201

/*CreateTodoCreated Created

swagger:response createTodoCreated
*/
type CreateTodoCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Todo `json:"body,omitempty"`
}

// NewCreateTodoCreated creates CreateTodoCreated with default headers values
func NewCreateTodoCreated() *CreateTodoCreated {
	return &CreateTodoCreated{}
}

// WithPayload adds the payload to the create todo created response
func (o *CreateTodoCreated) WithPayload(payload *models.Todo) *CreateTodoCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create todo created response
func (o *CreateTodoCreated) SetPayload(payload *models.Todo) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTodoCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateTodoDefault Error

swagger:response createTodoDefault
*/
type CreateTodoDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateTodoDefault creates CreateTodoDefault with default headers values
func NewCreateTodoDefault(code int) *CreateTodoDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateTodoDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create todo default response
func (o *CreateTodoDefault) WithStatusCode(code int) *CreateTodoDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create todo default response
func (o *CreateTodoDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create todo default response
func (o *CreateTodoDefault) WithPayload(payload *models.Error) *CreateTodoDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create todo default response
func (o *CreateTodoDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateTodoDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
