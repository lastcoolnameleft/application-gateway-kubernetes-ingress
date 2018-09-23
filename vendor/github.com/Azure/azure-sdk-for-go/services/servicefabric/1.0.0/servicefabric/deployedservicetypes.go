package servicefabric

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// DeployedServiceTypesClient is the client for the DeployedServiceTypes methods of the Servicefabric service.
type DeployedServiceTypesClient struct {
	BaseClient
}

// NewDeployedServiceTypesClient creates an instance of the DeployedServiceTypesClient client.
func NewDeployedServiceTypesClient(timeout *int32) DeployedServiceTypesClient {
	return NewDeployedServiceTypesClientWithBaseURI(DefaultBaseURI, timeout)
}

// NewDeployedServiceTypesClientWithBaseURI creates an instance of the DeployedServiceTypesClient client.
func NewDeployedServiceTypesClientWithBaseURI(baseURI string, timeout *int32) DeployedServiceTypesClient {
	return DeployedServiceTypesClient{NewWithBaseURI(baseURI, timeout)}
}

// Get get deployed service types
// Parameters:
// nodeName - the name of the node
// applicationName - the name of the application
func (client DeployedServiceTypesClient) Get(ctx context.Context, nodeName string, applicationName string) (result ListDeployedServiceType, err error) {
	req, err := client.GetPreparer(ctx, nodeName, applicationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.DeployedServiceTypesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.DeployedServiceTypesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.DeployedServiceTypesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client DeployedServiceTypesClient) GetPreparer(ctx context.Context, nodeName string, applicationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"applicationName": applicationName,
		"nodeName":        autorest.Encode("path", nodeName),
	}

	const APIVersion = "1.0.0"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if client.Timeout != nil {
		queryParameters["timeout"] = autorest.Encode("query", *client.Timeout)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Nodes/{nodeName}/$/GetApplications/{applicationName}/$/GetServiceTypes", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DeployedServiceTypesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DeployedServiceTypesClient) GetResponder(resp *http.Response) (result ListDeployedServiceType, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
