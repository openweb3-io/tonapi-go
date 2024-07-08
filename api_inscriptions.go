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


// InscriptionsAPIService InscriptionsAPI service
type InscriptionsAPIService service

type ApiGetAccountInscriptionsRequest struct {
	ctx context.Context
	ApiService *InscriptionsAPIService
	accountId string
	limit *int32
	offset *int32
}

func (r ApiGetAccountInscriptionsRequest) Limit(limit int32) ApiGetAccountInscriptionsRequest {
	r.limit = &limit
	return r
}

func (r ApiGetAccountInscriptionsRequest) Offset(offset int32) ApiGetAccountInscriptionsRequest {
	r.offset = &offset
	return r
}

func (r ApiGetAccountInscriptionsRequest) Execute() (*InscriptionBalances, *http.Response, error) {
	return r.ApiService.GetAccountInscriptionsExecute(r)
}

/*
GetAccountInscriptions Method for GetAccountInscriptions

Get all inscriptions by owner address. It's experimental API and can be dropped in the future.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId account ID
 @return ApiGetAccountInscriptionsRequest
*/
func (a *InscriptionsAPIService) GetAccountInscriptions(ctx context.Context, accountId string) ApiGetAccountInscriptionsRequest {
	return ApiGetAccountInscriptionsRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
	}
}

// Execute executes the request
//  @return InscriptionBalances
func (a *InscriptionsAPIService) GetAccountInscriptionsExecute(r ApiGetAccountInscriptionsRequest) (*InscriptionBalances, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InscriptionBalances
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InscriptionsAPIService.GetAccountInscriptions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/experimental/accounts/{account_id}/inscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	} else {
		var defaultValue int32 = 1000
		r.limit = &defaultValue
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	} else {
		var defaultValue int32 = 0
		r.offset = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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

type ApiGetAccountInscriptionsHistoryRequest struct {
	ctx context.Context
	ApiService *InscriptionsAPIService
	accountId string
	acceptLanguage *string
	beforeLt *int64
	limit *int32
}

func (r ApiGetAccountInscriptionsHistoryRequest) AcceptLanguage(acceptLanguage string) ApiGetAccountInscriptionsHistoryRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

// omit this parameter to get last events
func (r ApiGetAccountInscriptionsHistoryRequest) BeforeLt(beforeLt int64) ApiGetAccountInscriptionsHistoryRequest {
	r.beforeLt = &beforeLt
	return r
}

func (r ApiGetAccountInscriptionsHistoryRequest) Limit(limit int32) ApiGetAccountInscriptionsHistoryRequest {
	r.limit = &limit
	return r
}

func (r ApiGetAccountInscriptionsHistoryRequest) Execute() (*AccountEvents, *http.Response, error) {
	return r.ApiService.GetAccountInscriptionsHistoryExecute(r)
}

/*
GetAccountInscriptionsHistory Method for GetAccountInscriptionsHistory

Get the transfer inscriptions history for account. It's experimental API and can be dropped in the future.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId account ID
 @return ApiGetAccountInscriptionsHistoryRequest
*/
func (a *InscriptionsAPIService) GetAccountInscriptionsHistory(ctx context.Context, accountId string) ApiGetAccountInscriptionsHistoryRequest {
	return ApiGetAccountInscriptionsHistoryRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
	}
}

// Execute executes the request
//  @return AccountEvents
func (a *InscriptionsAPIService) GetAccountInscriptionsHistoryExecute(r ApiGetAccountInscriptionsHistoryRequest) (*AccountEvents, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccountEvents
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InscriptionsAPIService.GetAccountInscriptionsHistory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/experimental/accounts/{account_id}/inscriptions/history"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.beforeLt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "before_lt", r.beforeLt, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	} else {
		var defaultValue int32 = 100
		r.limit = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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

type ApiGetAccountInscriptionsHistoryByTickerRequest struct {
	ctx context.Context
	ApiService *InscriptionsAPIService
	accountId string
	ticker string
	acceptLanguage *string
	beforeLt *int64
	limit *int32
}

func (r ApiGetAccountInscriptionsHistoryByTickerRequest) AcceptLanguage(acceptLanguage string) ApiGetAccountInscriptionsHistoryByTickerRequest {
	r.acceptLanguage = &acceptLanguage
	return r
}

// omit this parameter to get last events
func (r ApiGetAccountInscriptionsHistoryByTickerRequest) BeforeLt(beforeLt int64) ApiGetAccountInscriptionsHistoryByTickerRequest {
	r.beforeLt = &beforeLt
	return r
}

func (r ApiGetAccountInscriptionsHistoryByTickerRequest) Limit(limit int32) ApiGetAccountInscriptionsHistoryByTickerRequest {
	r.limit = &limit
	return r
}

func (r ApiGetAccountInscriptionsHistoryByTickerRequest) Execute() (*AccountEvents, *http.Response, error) {
	return r.ApiService.GetAccountInscriptionsHistoryByTickerExecute(r)
}

/*
GetAccountInscriptionsHistoryByTicker Method for GetAccountInscriptionsHistoryByTicker

Get the transfer inscriptions history for account. It's experimental API and can be dropped in the future.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId account ID
 @param ticker
 @return ApiGetAccountInscriptionsHistoryByTickerRequest
*/
func (a *InscriptionsAPIService) GetAccountInscriptionsHistoryByTicker(ctx context.Context, accountId string, ticker string) ApiGetAccountInscriptionsHistoryByTickerRequest {
	return ApiGetAccountInscriptionsHistoryByTickerRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		ticker: ticker,
	}
}

// Execute executes the request
//  @return AccountEvents
func (a *InscriptionsAPIService) GetAccountInscriptionsHistoryByTickerExecute(r ApiGetAccountInscriptionsHistoryByTickerRequest) (*AccountEvents, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccountEvents
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InscriptionsAPIService.GetAccountInscriptionsHistoryByTicker")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/experimental/accounts/{account_id}/inscriptions/{ticker}/history"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"ticker"+"}", url.PathEscape(parameterValueToString(r.ticker, "ticker")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.beforeLt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "before_lt", r.beforeLt, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	} else {
		var defaultValue int32 = 100
		r.limit = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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

type ApiGetInscriptionOpTemplateRequest struct {
	ctx context.Context
	ApiService *InscriptionsAPIService
	type_ *string
	operation *string
	amount *string
	ticker *string
	who *string
	destination *string
	comment *string
}

func (r ApiGetInscriptionOpTemplateRequest) Type_(type_ string) ApiGetInscriptionOpTemplateRequest {
	r.type_ = &type_
	return r
}

func (r ApiGetInscriptionOpTemplateRequest) Operation(operation string) ApiGetInscriptionOpTemplateRequest {
	r.operation = &operation
	return r
}

func (r ApiGetInscriptionOpTemplateRequest) Amount(amount string) ApiGetInscriptionOpTemplateRequest {
	r.amount = &amount
	return r
}

func (r ApiGetInscriptionOpTemplateRequest) Ticker(ticker string) ApiGetInscriptionOpTemplateRequest {
	r.ticker = &ticker
	return r
}

func (r ApiGetInscriptionOpTemplateRequest) Who(who string) ApiGetInscriptionOpTemplateRequest {
	r.who = &who
	return r
}

func (r ApiGetInscriptionOpTemplateRequest) Destination(destination string) ApiGetInscriptionOpTemplateRequest {
	r.destination = &destination
	return r
}

func (r ApiGetInscriptionOpTemplateRequest) Comment(comment string) ApiGetInscriptionOpTemplateRequest {
	r.comment = &comment
	return r
}

func (r ApiGetInscriptionOpTemplateRequest) Execute() (*GetInscriptionOpTemplate200Response, *http.Response, error) {
	return r.ApiService.GetInscriptionOpTemplateExecute(r)
}

/*
GetInscriptionOpTemplate Method for GetInscriptionOpTemplate

return comment for making operation with inscription. please don't use it if you don't know what you are doing

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetInscriptionOpTemplateRequest
*/
func (a *InscriptionsAPIService) GetInscriptionOpTemplate(ctx context.Context) ApiGetInscriptionOpTemplateRequest {
	return ApiGetInscriptionOpTemplateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetInscriptionOpTemplate200Response
func (a *InscriptionsAPIService) GetInscriptionOpTemplateExecute(r ApiGetInscriptionOpTemplateRequest) (*GetInscriptionOpTemplate200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetInscriptionOpTemplate200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InscriptionsAPIService.GetInscriptionOpTemplate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/experimental/inscriptions/op-template"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.type_ == nil {
		return localVarReturnValue, nil, reportError("type_ is required and must be specified")
	}
	if r.operation == nil {
		return localVarReturnValue, nil, reportError("operation is required and must be specified")
	}
	if r.amount == nil {
		return localVarReturnValue, nil, reportError("amount is required and must be specified")
	}
	if r.ticker == nil {
		return localVarReturnValue, nil, reportError("ticker is required and must be specified")
	}
	if r.who == nil {
		return localVarReturnValue, nil, reportError("who is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "")
	if r.destination != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "destination", r.destination, "")
	}
	if r.comment != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "comment", r.comment, "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "operation", r.operation, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "amount", r.amount, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "ticker", r.ticker, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "who", r.who, "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
