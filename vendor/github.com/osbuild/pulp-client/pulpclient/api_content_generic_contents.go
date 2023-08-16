/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pulpclient

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"os"
	"reflect"
)


// ContentGenericContentsAPIService ContentGenericContentsAPI service
type ContentGenericContentsAPIService service

type ContentGenericContentsAPIContentDebGenericContentsCreateRequest struct {
	ctx context.Context
	ApiService *ContentGenericContentsAPIService
	relativePath *string
	repository *string
	artifact *string
	file *os.File
	upload *string
}

// Path where the artifact is located relative to distributions base_path
func (r ContentGenericContentsAPIContentDebGenericContentsCreateRequest) RelativePath(relativePath string) ContentGenericContentsAPIContentDebGenericContentsCreateRequest {
	r.relativePath = &relativePath
	return r
}

// A URI of a repository the new content unit should be associated with.
func (r ContentGenericContentsAPIContentDebGenericContentsCreateRequest) Repository(repository string) ContentGenericContentsAPIContentDebGenericContentsCreateRequest {
	r.repository = &repository
	return r
}

// Artifact file representing the physical content
func (r ContentGenericContentsAPIContentDebGenericContentsCreateRequest) Artifact(artifact string) ContentGenericContentsAPIContentDebGenericContentsCreateRequest {
	r.artifact = &artifact
	return r
}

// An uploaded file that may be turned into the artifact of the content unit.
func (r ContentGenericContentsAPIContentDebGenericContentsCreateRequest) File(file *os.File) ContentGenericContentsAPIContentDebGenericContentsCreateRequest {
	r.file = file
	return r
}

// An uncommitted upload that may be turned into the artifact of the content unit.
func (r ContentGenericContentsAPIContentDebGenericContentsCreateRequest) Upload(upload string) ContentGenericContentsAPIContentDebGenericContentsCreateRequest {
	r.upload = &upload
	return r
}

func (r ContentGenericContentsAPIContentDebGenericContentsCreateRequest) Execute() (*AsyncOperationResponse, *http.Response, error) {
	return r.ApiService.ContentDebGenericContentsCreateExecute(r)
}

/*
ContentDebGenericContentsCreate Create a generic content

Trigger an asynchronous task to create content,optionally create new repository version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ContentGenericContentsAPIContentDebGenericContentsCreateRequest
*/
func (a *ContentGenericContentsAPIService) ContentDebGenericContentsCreate(ctx context.Context) ContentGenericContentsAPIContentDebGenericContentsCreateRequest {
	return ContentGenericContentsAPIContentDebGenericContentsCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AsyncOperationResponse
func (a *ContentGenericContentsAPIService) ContentDebGenericContentsCreateExecute(r ContentGenericContentsAPIContentDebGenericContentsCreateRequest) (*AsyncOperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncOperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentGenericContentsAPIService.ContentDebGenericContentsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/deb/generic_contents/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.relativePath == nil {
		return localVarReturnValue, nil, reportError("relativePath is required and must be specified")
	}
	if strlen(*r.relativePath) < 1 {
		return localVarReturnValue, nil, reportError("relativePath must have at least 1 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data", "application/x-www-form-urlencoded"}

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
	if r.repository != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "repository", r.repository, "")
	}
	if r.artifact != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "artifact", r.artifact, "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "relative_path", r.relativePath, "")
	var fileLocalVarFormFileName string
	var fileLocalVarFileName     string
	var fileLocalVarFileBytes    []byte

	fileLocalVarFormFileName = "file"


	fileLocalVarFile := r.file

	if fileLocalVarFile != nil {
		fbs, _ := io.ReadAll(fileLocalVarFile)

		fileLocalVarFileBytes = fbs
		fileLocalVarFileName = fileLocalVarFile.Name()
		fileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: fileLocalVarFileBytes, fileName: fileLocalVarFileName, formFileName: fileLocalVarFormFileName})
	}
	if r.upload != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "upload", r.upload, "")
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

type ContentGenericContentsAPIContentDebGenericContentsListRequest struct {
	ctx context.Context
	ApiService *ContentGenericContentsAPIService
	limit *int32
	offset *int32
	ordering *[]string
	pulpHrefIn *[]string
	pulpIdIn *[]string
	relativePath *string
	repositoryVersion *string
	repositoryVersionAdded *string
	repositoryVersionRemoved *string
	sha256 *string
	fields *[]string
	excludeFields *[]string
}

// Number of results to return per page.
func (r ContentGenericContentsAPIContentDebGenericContentsListRequest) Limit(limit int32) ContentGenericContentsAPIContentDebGenericContentsListRequest {
	r.limit = &limit
	return r
}

// The initial index from which to return the results.
func (r ContentGenericContentsAPIContentDebGenericContentsListRequest) Offset(offset int32) ContentGenericContentsAPIContentDebGenericContentsListRequest {
	r.offset = &offset
	return r
}

// Ordering  * &#x60;pulp_id&#x60; - Pulp id * &#x60;-pulp_id&#x60; - Pulp id (descending) * &#x60;pulp_created&#x60; - Pulp created * &#x60;-pulp_created&#x60; - Pulp created (descending) * &#x60;pulp_last_updated&#x60; - Pulp last updated * &#x60;-pulp_last_updated&#x60; - Pulp last updated (descending) * &#x60;pulp_type&#x60; - Pulp type * &#x60;-pulp_type&#x60; - Pulp type (descending) * &#x60;upstream_id&#x60; - Upstream id * &#x60;-upstream_id&#x60; - Upstream id (descending) * &#x60;timestamp_of_interest&#x60; - Timestamp of interest * &#x60;-timestamp_of_interest&#x60; - Timestamp of interest (descending) * &#x60;relative_path&#x60; - Relative path * &#x60;-relative_path&#x60; - Relative path (descending) * &#x60;sha256&#x60; - Sha256 * &#x60;-sha256&#x60; - Sha256 (descending) * &#x60;pk&#x60; - Pk * &#x60;-pk&#x60; - Pk (descending)
func (r ContentGenericContentsAPIContentDebGenericContentsListRequest) Ordering(ordering []string) ContentGenericContentsAPIContentDebGenericContentsListRequest {
	r.ordering = &ordering
	return r
}

// Multiple values may be separated by commas.
func (r ContentGenericContentsAPIContentDebGenericContentsListRequest) PulpHrefIn(pulpHrefIn []string) ContentGenericContentsAPIContentDebGenericContentsListRequest {
	r.pulpHrefIn = &pulpHrefIn
	return r
}

// Multiple values may be separated by commas.
func (r ContentGenericContentsAPIContentDebGenericContentsListRequest) PulpIdIn(pulpIdIn []string) ContentGenericContentsAPIContentDebGenericContentsListRequest {
	r.pulpIdIn = &pulpIdIn
	return r
}

// Filter results where relative_path matches value
func (r ContentGenericContentsAPIContentDebGenericContentsListRequest) RelativePath(relativePath string) ContentGenericContentsAPIContentDebGenericContentsListRequest {
	r.relativePath = &relativePath
	return r
}

// Repository Version referenced by HREF
func (r ContentGenericContentsAPIContentDebGenericContentsListRequest) RepositoryVersion(repositoryVersion string) ContentGenericContentsAPIContentDebGenericContentsListRequest {
	r.repositoryVersion = &repositoryVersion
	return r
}

// Repository Version referenced by HREF
func (r ContentGenericContentsAPIContentDebGenericContentsListRequest) RepositoryVersionAdded(repositoryVersionAdded string) ContentGenericContentsAPIContentDebGenericContentsListRequest {
	r.repositoryVersionAdded = &repositoryVersionAdded
	return r
}

// Repository Version referenced by HREF
func (r ContentGenericContentsAPIContentDebGenericContentsListRequest) RepositoryVersionRemoved(repositoryVersionRemoved string) ContentGenericContentsAPIContentDebGenericContentsListRequest {
	r.repositoryVersionRemoved = &repositoryVersionRemoved
	return r
}

// Filter results where sha256 matches value
func (r ContentGenericContentsAPIContentDebGenericContentsListRequest) Sha256(sha256 string) ContentGenericContentsAPIContentDebGenericContentsListRequest {
	r.sha256 = &sha256
	return r
}

// A list of fields to include in the response.
func (r ContentGenericContentsAPIContentDebGenericContentsListRequest) Fields(fields []string) ContentGenericContentsAPIContentDebGenericContentsListRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentGenericContentsAPIContentDebGenericContentsListRequest) ExcludeFields(excludeFields []string) ContentGenericContentsAPIContentDebGenericContentsListRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentGenericContentsAPIContentDebGenericContentsListRequest) Execute() (*PaginateddebGenericContentResponseList, *http.Response, error) {
	return r.ApiService.ContentDebGenericContentsListExecute(r)
}

/*
ContentDebGenericContentsList List generic contents

GenericContent is a catch all category for storing files not covered by any other type.

Associated artifacts: Exactly one arbitrary file that does not match any other type.

This is needed to store arbitrary files for use with the verbatim publisher. If you are not
using the verbatim publisher, you may ignore this type.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ContentGenericContentsAPIContentDebGenericContentsListRequest
*/
func (a *ContentGenericContentsAPIService) ContentDebGenericContentsList(ctx context.Context) ContentGenericContentsAPIContentDebGenericContentsListRequest {
	return ContentGenericContentsAPIContentDebGenericContentsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PaginateddebGenericContentResponseList
func (a *ContentGenericContentsAPIService) ContentDebGenericContentsListExecute(r ContentGenericContentsAPIContentDebGenericContentsListRequest) (*PaginateddebGenericContentResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginateddebGenericContentResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentGenericContentsAPIService.ContentDebGenericContentsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pulp/api/v3/content/deb/generic_contents/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.ordering != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ordering", r.ordering, "csv")
	}
	if r.pulpHrefIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_href__in", r.pulpHrefIn, "csv")
	}
	if r.pulpIdIn != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pulp_id__in", r.pulpIdIn, "csv")
	}
	if r.relativePath != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "relative_path", r.relativePath, "")
	}
	if r.repositoryVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version", r.repositoryVersion, "")
	}
	if r.repositoryVersionAdded != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version_added", r.repositoryVersionAdded, "")
	}
	if r.repositoryVersionRemoved != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "repository_version_removed", r.repositoryVersionRemoved, "")
	}
	if r.sha256 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sha256", r.sha256, "")
	}
	if r.fields != nil {
		t := *r.fields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "fields", t, "multi")
		}
	}
	if r.excludeFields != nil {
		t := *r.excludeFields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", t, "multi")
		}
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

type ContentGenericContentsAPIContentDebGenericContentsReadRequest struct {
	ctx context.Context
	ApiService *ContentGenericContentsAPIService
	debGenericContentHref string
	fields *[]string
	excludeFields *[]string
}

// A list of fields to include in the response.
func (r ContentGenericContentsAPIContentDebGenericContentsReadRequest) Fields(fields []string) ContentGenericContentsAPIContentDebGenericContentsReadRequest {
	r.fields = &fields
	return r
}

// A list of fields to exclude from the response.
func (r ContentGenericContentsAPIContentDebGenericContentsReadRequest) ExcludeFields(excludeFields []string) ContentGenericContentsAPIContentDebGenericContentsReadRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r ContentGenericContentsAPIContentDebGenericContentsReadRequest) Execute() (*DebGenericContentResponse, *http.Response, error) {
	return r.ApiService.ContentDebGenericContentsReadExecute(r)
}

/*
ContentDebGenericContentsRead Inspect a generic content

GenericContent is a catch all category for storing files not covered by any other type.

Associated artifacts: Exactly one arbitrary file that does not match any other type.

This is needed to store arbitrary files for use with the verbatim publisher. If you are not
using the verbatim publisher, you may ignore this type.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param debGenericContentHref
 @return ContentGenericContentsAPIContentDebGenericContentsReadRequest
*/
func (a *ContentGenericContentsAPIService) ContentDebGenericContentsRead(ctx context.Context, debGenericContentHref string) ContentGenericContentsAPIContentDebGenericContentsReadRequest {
	return ContentGenericContentsAPIContentDebGenericContentsReadRequest{
		ApiService: a,
		ctx: ctx,
		debGenericContentHref: debGenericContentHref,
	}
}

// Execute executes the request
//  @return DebGenericContentResponse
func (a *ContentGenericContentsAPIService) ContentDebGenericContentsReadExecute(r ContentGenericContentsAPIContentDebGenericContentsReadRequest) (*DebGenericContentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DebGenericContentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContentGenericContentsAPIService.ContentDebGenericContentsRead")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "{deb_generic_content_href}"
	localVarPath = strings.Replace(localVarPath, "{"+"deb_generic_content_href"+"}", parameterValueToString(r.debGenericContentHref, "debGenericContentHref"), -1)  // NOTE: paths aren't escaped because Pulp uses hrefs as path parameters

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		t := *r.fields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "fields", t, "multi")
		}
	}
	if r.excludeFields != nil {
		t := *r.excludeFields
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", t, "multi")
		}
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