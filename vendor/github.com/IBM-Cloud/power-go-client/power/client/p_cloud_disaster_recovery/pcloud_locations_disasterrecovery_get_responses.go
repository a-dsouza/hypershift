// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_disaster_recovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudLocationsDisasterrecoveryGetReader is a Reader for the PcloudLocationsDisasterrecoveryGet structure.
type PcloudLocationsDisasterrecoveryGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudLocationsDisasterrecoveryGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudLocationsDisasterrecoveryGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudLocationsDisasterrecoveryGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudLocationsDisasterrecoveryGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudLocationsDisasterrecoveryGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudLocationsDisasterrecoveryGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudLocationsDisasterrecoveryGetOK creates a PcloudLocationsDisasterrecoveryGetOK with default headers values
func NewPcloudLocationsDisasterrecoveryGetOK() *PcloudLocationsDisasterrecoveryGetOK {
	return &PcloudLocationsDisasterrecoveryGetOK{}
}

/*
PcloudLocationsDisasterrecoveryGetOK describes a response with status code 200, with default header values.

OK
*/
type PcloudLocationsDisasterrecoveryGetOK struct {
	Payload *models.DisasterRecoveryLocation
}

// IsSuccess returns true when this pcloud locations disasterrecovery get o k response has a 2xx status code
func (o *PcloudLocationsDisasterrecoveryGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud locations disasterrecovery get o k response has a 3xx status code
func (o *PcloudLocationsDisasterrecoveryGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud locations disasterrecovery get o k response has a 4xx status code
func (o *PcloudLocationsDisasterrecoveryGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud locations disasterrecovery get o k response has a 5xx status code
func (o *PcloudLocationsDisasterrecoveryGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud locations disasterrecovery get o k response a status code equal to that given
func (o *PcloudLocationsDisasterrecoveryGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *PcloudLocationsDisasterrecoveryGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/locations/disaster-recovery][%d] pcloudLocationsDisasterrecoveryGetOK  %+v", 200, o.Payload)
}

func (o *PcloudLocationsDisasterrecoveryGetOK) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/locations/disaster-recovery][%d] pcloudLocationsDisasterrecoveryGetOK  %+v", 200, o.Payload)
}

func (o *PcloudLocationsDisasterrecoveryGetOK) GetPayload() *models.DisasterRecoveryLocation {
	return o.Payload
}

func (o *PcloudLocationsDisasterrecoveryGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DisasterRecoveryLocation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudLocationsDisasterrecoveryGetBadRequest creates a PcloudLocationsDisasterrecoveryGetBadRequest with default headers values
func NewPcloudLocationsDisasterrecoveryGetBadRequest() *PcloudLocationsDisasterrecoveryGetBadRequest {
	return &PcloudLocationsDisasterrecoveryGetBadRequest{}
}

/*
PcloudLocationsDisasterrecoveryGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudLocationsDisasterrecoveryGetBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud locations disasterrecovery get bad request response has a 2xx status code
func (o *PcloudLocationsDisasterrecoveryGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud locations disasterrecovery get bad request response has a 3xx status code
func (o *PcloudLocationsDisasterrecoveryGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud locations disasterrecovery get bad request response has a 4xx status code
func (o *PcloudLocationsDisasterrecoveryGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud locations disasterrecovery get bad request response has a 5xx status code
func (o *PcloudLocationsDisasterrecoveryGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud locations disasterrecovery get bad request response a status code equal to that given
func (o *PcloudLocationsDisasterrecoveryGetBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PcloudLocationsDisasterrecoveryGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/locations/disaster-recovery][%d] pcloudLocationsDisasterrecoveryGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudLocationsDisasterrecoveryGetBadRequest) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/locations/disaster-recovery][%d] pcloudLocationsDisasterrecoveryGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudLocationsDisasterrecoveryGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudLocationsDisasterrecoveryGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudLocationsDisasterrecoveryGetUnauthorized creates a PcloudLocationsDisasterrecoveryGetUnauthorized with default headers values
func NewPcloudLocationsDisasterrecoveryGetUnauthorized() *PcloudLocationsDisasterrecoveryGetUnauthorized {
	return &PcloudLocationsDisasterrecoveryGetUnauthorized{}
}

/*
PcloudLocationsDisasterrecoveryGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudLocationsDisasterrecoveryGetUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud locations disasterrecovery get unauthorized response has a 2xx status code
func (o *PcloudLocationsDisasterrecoveryGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud locations disasterrecovery get unauthorized response has a 3xx status code
func (o *PcloudLocationsDisasterrecoveryGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud locations disasterrecovery get unauthorized response has a 4xx status code
func (o *PcloudLocationsDisasterrecoveryGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud locations disasterrecovery get unauthorized response has a 5xx status code
func (o *PcloudLocationsDisasterrecoveryGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud locations disasterrecovery get unauthorized response a status code equal to that given
func (o *PcloudLocationsDisasterrecoveryGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PcloudLocationsDisasterrecoveryGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/locations/disaster-recovery][%d] pcloudLocationsDisasterrecoveryGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudLocationsDisasterrecoveryGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/locations/disaster-recovery][%d] pcloudLocationsDisasterrecoveryGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudLocationsDisasterrecoveryGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudLocationsDisasterrecoveryGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudLocationsDisasterrecoveryGetNotFound creates a PcloudLocationsDisasterrecoveryGetNotFound with default headers values
func NewPcloudLocationsDisasterrecoveryGetNotFound() *PcloudLocationsDisasterrecoveryGetNotFound {
	return &PcloudLocationsDisasterrecoveryGetNotFound{}
}

/*
PcloudLocationsDisasterrecoveryGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudLocationsDisasterrecoveryGetNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud locations disasterrecovery get not found response has a 2xx status code
func (o *PcloudLocationsDisasterrecoveryGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud locations disasterrecovery get not found response has a 3xx status code
func (o *PcloudLocationsDisasterrecoveryGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud locations disasterrecovery get not found response has a 4xx status code
func (o *PcloudLocationsDisasterrecoveryGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud locations disasterrecovery get not found response has a 5xx status code
func (o *PcloudLocationsDisasterrecoveryGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud locations disasterrecovery get not found response a status code equal to that given
func (o *PcloudLocationsDisasterrecoveryGetNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PcloudLocationsDisasterrecoveryGetNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/locations/disaster-recovery][%d] pcloudLocationsDisasterrecoveryGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudLocationsDisasterrecoveryGetNotFound) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/locations/disaster-recovery][%d] pcloudLocationsDisasterrecoveryGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudLocationsDisasterrecoveryGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudLocationsDisasterrecoveryGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudLocationsDisasterrecoveryGetInternalServerError creates a PcloudLocationsDisasterrecoveryGetInternalServerError with default headers values
func NewPcloudLocationsDisasterrecoveryGetInternalServerError() *PcloudLocationsDisasterrecoveryGetInternalServerError {
	return &PcloudLocationsDisasterrecoveryGetInternalServerError{}
}

/*
PcloudLocationsDisasterrecoveryGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudLocationsDisasterrecoveryGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud locations disasterrecovery get internal server error response has a 2xx status code
func (o *PcloudLocationsDisasterrecoveryGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud locations disasterrecovery get internal server error response has a 3xx status code
func (o *PcloudLocationsDisasterrecoveryGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud locations disasterrecovery get internal server error response has a 4xx status code
func (o *PcloudLocationsDisasterrecoveryGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud locations disasterrecovery get internal server error response has a 5xx status code
func (o *PcloudLocationsDisasterrecoveryGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud locations disasterrecovery get internal server error response a status code equal to that given
func (o *PcloudLocationsDisasterrecoveryGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PcloudLocationsDisasterrecoveryGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/locations/disaster-recovery][%d] pcloudLocationsDisasterrecoveryGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudLocationsDisasterrecoveryGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/locations/disaster-recovery][%d] pcloudLocationsDisasterrecoveryGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudLocationsDisasterrecoveryGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudLocationsDisasterrecoveryGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}