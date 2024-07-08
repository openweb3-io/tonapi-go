/*
REST api to TON blockchain explorer

Provide access to indexed TON blockchain

API version: 2.0.0
Contact: support@tonkeeper.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tonapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// EmulationAPIService EmulationAPI service
type EmulationAPIService service

type ApiDecodeMessageRequest struct {
	ctx context.Context
	ApiService *EmulationAPIService
	decodeMessageRequest *DecodeMessageRequest
}

// bag-of-cells serialized to base64
func (r ApiDecodeMessageRequest) DecodeMessageRequest(decodeMessageRequest DecodeMessageRequest) ApiDecodeMessageRequest {
	r.decodeMessageRequest = &decodeMessageRequest
	return r
}

func (r ApiDecodeMessageRequest) Execute() (*DecodedMessage, *http.Response, error) {
	return r.ApiService.DecodeMessageExecute(r)
}

/*
DecodeMessage Method for DecodeMessage

Decode a given message. Only external incoming messages can be decoded currently.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDecodeMessageRequest
*/
func (a *EmulationAPIService) DecodeMessage(ctx context.Context) ApiDecodeMessageRequest {
	return ApiDecodeMessageRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DecodedMessage
func (a *EmulationAPIService) DecodeMessageExecute(r ApiDecodeMessageRequest) (*DecodedMessage, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DecodedMessage
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmulationAPIService.DecodeMessage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/message/decode"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.decodeMessageRequest == nil {
		return localVarReturnValue, nil, reportError("decodeMessageRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.decodeMessageRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v StatusDefaultResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiEmulateMessageToAccountEventRequest struct {
	ctx context.Context
	ApiService *EmulationAPIService
	accountId string
	decodeMessageRequest *DecodeMessageRequest
	acceptLanguage *string
	ignoreSignatureCheck *bool
}

// bag-of-cells serialized to base64
func (r ApiEmulateMessageToAccountEventRequest) DecodeMessageRequest(decodeMessageRequest DecodeMessageRequest) ApiEmulateMessageToAccountEventRequest {
	r.decodeMessageRequest = &decodeMessageRequest
	return r
}

func (r ApiEmulateMessageToAccountEventRequest) AcceptLanguage(acceptLanguage string) ApiEmulateMessageToAccountEventRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

func (r ApiEmulateMessageToAccountEventRequest) IgnoreSignatureCheck(ignoreSignatureCheck bool) ApiEmulateMessageToAccountEventRequest {
	r.ignoreSignatureCheck = &ignoreSignatureCheck
	return r
}

func (r ApiEmulateMessageToAccountEventRequest) Execute() (*AccountEvent, *http.Response, error) {
	return r.ApiService.EmulateMessageToAccountEventExecute(r)
}

/*
EmulateMessageToAccountEvent Method for EmulateMessageToAccountEvent

Emulate sending message to blockchain

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId account ID
 @return ApiEmulateMessageToAccountEventRequest
*/
func (a *EmulationAPIService) EmulateMessageToAccountEvent(ctx context.Context, accountId string) ApiEmulateMessageToAccountEventRequest {
	return ApiEmulateMessageToAccountEventRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
	}
}

// Execute executes the request
//  @return AccountEvent
func (a *EmulationAPIService) EmulateMessageToAccountEventExecute(r ApiEmulateMessageToAccountEventRequest) (*AccountEvent, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccountEvent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmulationAPIService.EmulateMessageToAccountEvent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/accounts/{account_id}/events/emulate"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.decodeMessageRequest == nil {
		return localVarReturnValue, nil, reportError("decodeMessageRequest is required and must be specified")
	}

	if r.ignoreSignatureCheck != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ignore_signature_check", r.ignoreSignatureCheck, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.acceptLanguage != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept-Language", r.acceptLanguage, "")
	}
	// body params
	localVarPostBody = r.decodeMessageRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v StatusDefaultResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiEmulateMessageToEventRequest struct {
	ctx context.Context
	ApiService *EmulationAPIService
	decodeMessageRequest *DecodeMessageRequest
	acceptLanguage *string
	ignoreSignatureCheck *bool
}

// bag-of-cells serialized to base64
func (r ApiEmulateMessageToEventRequest) DecodeMessageRequest(decodeMessageRequest DecodeMessageRequest) ApiEmulateMessageToEventRequest {
	r.decodeMessageRequest = &decodeMessageRequest
	return r
}

func (r ApiEmulateMessageToEventRequest) AcceptLanguage(acceptLanguage string) ApiEmulateMessageToEventRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

func (r ApiEmulateMessageToEventRequest) IgnoreSignatureCheck(ignoreSignatureCheck bool) ApiEmulateMessageToEventRequest {
	r.ignoreSignatureCheck = &ignoreSignatureCheck
	return r
}

func (r ApiEmulateMessageToEventRequest) Execute() (*Event, *http.Response, error) {
	return r.ApiService.EmulateMessageToEventExecute(r)
}

/*
EmulateMessageToEvent Method for EmulateMessageToEvent

Emulate sending message to blockchain

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEmulateMessageToEventRequest
*/
func (a *EmulationAPIService) EmulateMessageToEvent(ctx context.Context) ApiEmulateMessageToEventRequest {
	return ApiEmulateMessageToEventRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Event
func (a *EmulationAPIService) EmulateMessageToEventExecute(r ApiEmulateMessageToEventRequest) (*Event, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Event
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmulationAPIService.EmulateMessageToEvent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/events/emulate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.decodeMessageRequest == nil {
		return localVarReturnValue, nil, reportError("decodeMessageRequest is required and must be specified")
	}

	if r.ignoreSignatureCheck != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ignore_signature_check", r.ignoreSignatureCheck, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.acceptLanguage != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept-Language", r.acceptLanguage, "")
	}
	// body params
	localVarPostBody = r.decodeMessageRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v StatusDefaultResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiEmulateMessageToTraceRequest struct {
	ctx context.Context
	ApiService *EmulationAPIService
	decodeMessageRequest *DecodeMessageRequest
	ignoreSignatureCheck *bool
}

// bag-of-cells serialized to base64
func (r ApiEmulateMessageToTraceRequest) DecodeMessageRequest(decodeMessageRequest DecodeMessageRequest) ApiEmulateMessageToTraceRequest {
	r.decodeMessageRequest = &decodeMessageRequest
	return r
}

func (r ApiEmulateMessageToTraceRequest) IgnoreSignatureCheck(ignoreSignatureCheck bool) ApiEmulateMessageToTraceRequest {
	r.ignoreSignatureCheck = &ignoreSignatureCheck
	return r
}

func (r ApiEmulateMessageToTraceRequest) Execute() (*Trace, *http.Response, error) {
	return r.ApiService.EmulateMessageToTraceExecute(r)
}

/*
EmulateMessageToTrace Method for EmulateMessageToTrace

Emulate sending message to blockchain

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEmulateMessageToTraceRequest
*/
func (a *EmulationAPIService) EmulateMessageToTrace(ctx context.Context) ApiEmulateMessageToTraceRequest {
	return ApiEmulateMessageToTraceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Trace
func (a *EmulationAPIService) EmulateMessageToTraceExecute(r ApiEmulateMessageToTraceRequest) (*Trace, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Trace
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmulationAPIService.EmulateMessageToTrace")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/traces/emulate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.decodeMessageRequest == nil {
		return localVarReturnValue, nil, reportError("decodeMessageRequest is required and must be specified")
	}

	if r.ignoreSignatureCheck != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ignore_signature_check", r.ignoreSignatureCheck, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.decodeMessageRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v StatusDefaultResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiEmulateMessageToWalletRequest struct {
	ctx context.Context
	ApiService *EmulationAPIService
	emulateMessageToWalletRequest *EmulateMessageToWalletRequest
	acceptLanguage *string
}

// bag-of-cells serialized to base64 and additional parameters to configure emulation
func (r ApiEmulateMessageToWalletRequest) EmulateMessageToWalletRequest(emulateMessageToWalletRequest EmulateMessageToWalletRequest) ApiEmulateMessageToWalletRequest {
	r.emulateMessageToWalletRequest = &emulateMessageToWalletRequest
	return r
}

func (r ApiEmulateMessageToWalletRequest) AcceptLanguage(acceptLanguage string) ApiEmulateMessageToWalletRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

func (r ApiEmulateMessageToWalletRequest) Execute() (*MessageConsequences, *http.Response, error) {
	return r.ApiService.EmulateMessageToWalletExecute(r)
}

/*
EmulateMessageToWallet Method for EmulateMessageToWallet

Emulate sending message to blockchain

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEmulateMessageToWalletRequest
*/
func (a *EmulationAPIService) EmulateMessageToWallet(ctx context.Context) ApiEmulateMessageToWalletRequest {
	return ApiEmulateMessageToWalletRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MessageConsequences
func (a *EmulationAPIService) EmulateMessageToWalletExecute(r ApiEmulateMessageToWalletRequest) (*MessageConsequences, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MessageConsequences
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmulationAPIService.EmulateMessageToWallet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/wallet/emulate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.emulateMessageToWalletRequest == nil {
		return localVarReturnValue, nil, reportError("emulateMessageToWalletRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.acceptLanguage != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept-Language", r.acceptLanguage, "")
	}
	// body params
	localVarPostBody = r.emulateMessageToWalletRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v StatusDefaultResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
