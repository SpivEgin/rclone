package dtl

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

// VirtualNetworkClient is the azure DevTest Labs REST API version 2015-05-21-preview.
type VirtualNetworkClient struct {
	BaseClient
}

// NewVirtualNetworkClient creates an instance of the VirtualNetworkClient client.
func NewVirtualNetworkClient(subscriptionID string) VirtualNetworkClient {
	return NewVirtualNetworkClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVirtualNetworkClientWithBaseURI creates an instance of the VirtualNetworkClient client.
func NewVirtualNetworkClientWithBaseURI(baseURI string, subscriptionID string) VirtualNetworkClient {
	return VirtualNetworkClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdateResource create or replace an existing virtual network. This operation can take a while to complete.
// Parameters:
// resourceGroupName - the name of the resource group.
// labName - the name of the lab.
// name - the name of the virtual network.
func (client VirtualNetworkClient) CreateOrUpdateResource(ctx context.Context, resourceGroupName string, labName string, name string, virtualNetwork VirtualNetwork) (result VirtualNetworkCreateOrUpdateResourceFuture, err error) {
	req, err := client.CreateOrUpdateResourcePreparer(ctx, resourceGroupName, labName, name, virtualNetwork)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.VirtualNetworkClient", "CreateOrUpdateResource", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateResourceSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.VirtualNetworkClient", "CreateOrUpdateResource", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdateResourcePreparer prepares the CreateOrUpdateResource request.
func (client VirtualNetworkClient) CreateOrUpdateResourcePreparer(ctx context.Context, resourceGroupName string, labName string, name string, virtualNetwork VirtualNetwork) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-21-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualnetworks/{name}", pathParameters),
		autorest.WithJSON(virtualNetwork),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateResourceSender sends the CreateOrUpdateResource request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkClient) CreateOrUpdateResourceSender(req *http.Request) (future VirtualNetworkCreateOrUpdateResourceFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated))
	return
}

// CreateOrUpdateResourceResponder handles the response to the CreateOrUpdateResource request. The method always
// closes the http.Response Body.
func (client VirtualNetworkClient) CreateOrUpdateResourceResponder(resp *http.Response) (result VirtualNetwork, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DeleteResource delete virtual network. This operation can take a while to complete.
// Parameters:
// resourceGroupName - the name of the resource group.
// labName - the name of the lab.
// name - the name of the virtual network.
func (client VirtualNetworkClient) DeleteResource(ctx context.Context, resourceGroupName string, labName string, name string) (result VirtualNetworkDeleteResourceFuture, err error) {
	req, err := client.DeleteResourcePreparer(ctx, resourceGroupName, labName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.VirtualNetworkClient", "DeleteResource", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteResourceSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.VirtualNetworkClient", "DeleteResource", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeleteResourcePreparer prepares the DeleteResource request.
func (client VirtualNetworkClient) DeleteResourcePreparer(ctx context.Context, resourceGroupName string, labName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-21-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualnetworks/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteResourceSender sends the DeleteResource request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkClient) DeleteResourceSender(req *http.Request) (future VirtualNetworkDeleteResourceFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent))
	return
}

// DeleteResourceResponder handles the response to the DeleteResource request. The method always
// closes the http.Response Body.
func (client VirtualNetworkClient) DeleteResourceResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetResource get virtual network.
// Parameters:
// resourceGroupName - the name of the resource group.
// labName - the name of the lab.
// name - the name of the virtual network.
func (client VirtualNetworkClient) GetResource(ctx context.Context, resourceGroupName string, labName string, name string) (result VirtualNetwork, err error) {
	req, err := client.GetResourcePreparer(ctx, resourceGroupName, labName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.VirtualNetworkClient", "GetResource", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetResourceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dtl.VirtualNetworkClient", "GetResource", resp, "Failure sending request")
		return
	}

	result, err = client.GetResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.VirtualNetworkClient", "GetResource", resp, "Failure responding to request")
	}

	return
}

// GetResourcePreparer prepares the GetResource request.
func (client VirtualNetworkClient) GetResourcePreparer(ctx context.Context, resourceGroupName string, labName string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-21-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualnetworks/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetResourceSender sends the GetResource request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkClient) GetResourceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResourceResponder handles the response to the GetResource request. The method always
// closes the http.Response Body.
func (client VirtualNetworkClient) GetResourceResponder(resp *http.Response) (result VirtualNetwork, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list virtual networks.
// Parameters:
// resourceGroupName - the name of the resource group.
// labName - the name of the lab.
// filter - the filter to apply on the operation.
func (client VirtualNetworkClient) List(ctx context.Context, resourceGroupName string, labName string, filter string, top *int32, orderBy string) (result ResponseWithContinuationVirtualNetworkPage, err error) {
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, labName, filter, top, orderBy)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.VirtualNetworkClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.rwcvn.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dtl.VirtualNetworkClient", "List", resp, "Failure sending request")
		return
	}

	result.rwcvn, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.VirtualNetworkClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client VirtualNetworkClient) ListPreparer(ctx context.Context, resourceGroupName string, labName string, filter string, top *int32, orderBy string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-21-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(orderBy) > 0 {
		queryParameters["$orderBy"] = autorest.Encode("query", orderBy)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualnetworks", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client VirtualNetworkClient) ListResponder(resp *http.Response) (result ResponseWithContinuationVirtualNetwork, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client VirtualNetworkClient) listNextResults(lastResults ResponseWithContinuationVirtualNetwork) (result ResponseWithContinuationVirtualNetwork, err error) {
	req, err := lastResults.responseWithContinuationVirtualNetworkPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "dtl.VirtualNetworkClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "dtl.VirtualNetworkClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.VirtualNetworkClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client VirtualNetworkClient) ListComplete(ctx context.Context, resourceGroupName string, labName string, filter string, top *int32, orderBy string) (result ResponseWithContinuationVirtualNetworkIterator, err error) {
	result.page, err = client.List(ctx, resourceGroupName, labName, filter, top, orderBy)
	return
}

// PatchResource modify properties of virtual networks.
// Parameters:
// resourceGroupName - the name of the resource group.
// labName - the name of the lab.
// name - the name of the virtual network.
func (client VirtualNetworkClient) PatchResource(ctx context.Context, resourceGroupName string, labName string, name string, virtualNetwork VirtualNetwork) (result VirtualNetwork, err error) {
	req, err := client.PatchResourcePreparer(ctx, resourceGroupName, labName, name, virtualNetwork)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.VirtualNetworkClient", "PatchResource", nil, "Failure preparing request")
		return
	}

	resp, err := client.PatchResourceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dtl.VirtualNetworkClient", "PatchResource", resp, "Failure sending request")
		return
	}

	result, err = client.PatchResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dtl.VirtualNetworkClient", "PatchResource", resp, "Failure responding to request")
	}

	return
}

// PatchResourcePreparer prepares the PatchResource request.
func (client VirtualNetworkClient) PatchResourcePreparer(ctx context.Context, resourceGroupName string, labName string, name string, virtualNetwork VirtualNetwork) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-21-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualnetworks/{name}", pathParameters),
		autorest.WithJSON(virtualNetwork),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PatchResourceSender sends the PatchResource request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworkClient) PatchResourceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// PatchResourceResponder handles the response to the PatchResource request. The method always
// closes the http.Response Body.
func (client VirtualNetworkClient) PatchResourceResponder(resp *http.Response) (result VirtualNetwork, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
