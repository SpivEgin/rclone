package iothub

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
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// DpsCertificateClient is the API for using the Azure IoT Hub Device Provisioning Service features.
type DpsCertificateClient struct {
	BaseClient
}

// NewDpsCertificateClient creates an instance of the DpsCertificateClient client.
func NewDpsCertificateClient(subscriptionID string) DpsCertificateClient {
	return NewDpsCertificateClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDpsCertificateClientWithBaseURI creates an instance of the DpsCertificateClient client.
func NewDpsCertificateClientWithBaseURI(baseURI string, subscriptionID string) DpsCertificateClient {
	return DpsCertificateClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate add new certificate or update an existing certificate.
// Parameters:
// resourceGroupName - resource group identifier.
// provisioningServiceName - the name of the provisioning service.
// certificateName - the name of the certificate create or update.
// certificateDescription - the certificate body.
// ifMatch - eTag of the certificate. This is required to update an existing certificate, and ignored while
// creating a brand new certificate.
func (client DpsCertificateClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, provisioningServiceName string, certificateName string, certificateDescription CertificateBodyDescription, ifMatch string) (result CertificateResponse, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: certificateName,
			Constraints: []validation.Constraint{{Target: "certificateName", Name: validation.MaxLength, Rule: 256, Chain: nil}}}}); err != nil {
		return result, validation.NewError("iothub.DpsCertificateClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, provisioningServiceName, certificateName, certificateDescription, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iothub.DpsCertificateClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "iothub.DpsCertificateClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iothub.DpsCertificateClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client DpsCertificateClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, provisioningServiceName string, certificateName string, certificateDescription CertificateBodyDescription, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"certificateName":         autorest.Encode("path", certificateName),
		"provisioningServiceName": autorest.Encode("path", provisioningServiceName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-08-21-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}/certificates/{certificateName}", pathParameters),
		autorest.WithJSON(certificateDescription),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client DpsCertificateClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client DpsCertificateClient) CreateOrUpdateResponder(resp *http.Response) (result CertificateResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete sends the delete request.
// Parameters:
// resourceGroupName - resource group identifier.
// ifMatch - eTag of the certificate
// provisioningServiceName - the name of the provisioning service.
// certificateName - this is a mandatory field, and is the logical name of the certificate that the
// provisioning service will access by.
// certificatename - this is optional, and it is the Common Name of the certificate.
// certificaterawBytes - raw data within the certificate.
// certificateisVerified - indicates if certificate has been verified by owner of the private key.
// certificatepurpose - a description that mentions the purpose of the certificate.
// certificatecreated - time the certificate is created.
// certificatelastUpdated - time the certificate is last updated.
// certificatehasPrivateKey - indicates if the certificate contains a private key.
// certificatenonce - random number generated to indicate Proof of Possession.
func (client DpsCertificateClient) Delete(ctx context.Context, resourceGroupName string, ifMatch string, provisioningServiceName string, certificateName string, certificatename string, certificaterawBytes []byte, certificateisVerified *bool, certificatepurpose CertificatePurpose, certificatecreated *date.Time, certificatelastUpdated *date.Time, certificatehasPrivateKey *bool, certificatenonce string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(ctx, resourceGroupName, ifMatch, provisioningServiceName, certificateName, certificatename, certificaterawBytes, certificateisVerified, certificatepurpose, certificatecreated, certificatelastUpdated, certificatehasPrivateKey, certificatenonce)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iothub.DpsCertificateClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "iothub.DpsCertificateClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iothub.DpsCertificateClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client DpsCertificateClient) DeletePreparer(ctx context.Context, resourceGroupName string, ifMatch string, provisioningServiceName string, certificateName string, certificatename string, certificaterawBytes []byte, certificateisVerified *bool, certificatepurpose CertificatePurpose, certificatecreated *date.Time, certificatelastUpdated *date.Time, certificatehasPrivateKey *bool, certificatenonce string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"certificateName":         autorest.Encode("path", certificateName),
		"provisioningServiceName": autorest.Encode("path", provisioningServiceName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-08-21-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(certificatename) > 0 {
		queryParameters["certificate.name"] = autorest.Encode("query", certificatename)
	}
	if certificaterawBytes != nil && len(certificaterawBytes) > 0 {
		queryParameters["certificate.rawBytes"] = autorest.Encode("query", certificaterawBytes)
	}
	if certificateisVerified != nil {
		queryParameters["certificate.isVerified"] = autorest.Encode("query", *certificateisVerified)
	}
	if len(string(certificatepurpose)) > 0 {
		queryParameters["certificate.purpose"] = autorest.Encode("query", certificatepurpose)
	}
	if certificatecreated != nil {
		queryParameters["certificate.created"] = autorest.Encode("query", *certificatecreated)
	}
	if certificatelastUpdated != nil {
		queryParameters["certificate.lastUpdated"] = autorest.Encode("query", *certificatelastUpdated)
	}
	if certificatehasPrivateKey != nil {
		queryParameters["certificate.hasPrivateKey"] = autorest.Encode("query", *certificatehasPrivateKey)
	}
	if len(certificatenonce) > 0 {
		queryParameters["certificate.nonce"] = autorest.Encode("query", certificatenonce)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}/certificates/{certificateName}", pathParameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client DpsCertificateClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client DpsCertificateClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GenerateVerificationCode generate verification code for Proof of Possession.
// Parameters:
// certificateName - the mandatory logical name of the certificate, that the provisioning service uses to
// access.
// ifMatch - eTag of the certificate. This is required to update an existing certificate, and ignored while
// creating a brand new certificate.
// resourceGroupName - name of resource group.
// provisioningServiceName - name of provisioning service.
// certificatename - common Name for the certificate.
// certificaterawBytes - raw data of certificate.
// certificateisVerified - indicates if the certificate has been verified by owner of the private key.
// certificatepurpose - description mentioning the purpose of the certificate.
// certificatecreated - certificate creation time.
// certificatelastUpdated - certificate last updated time.
// certificatehasPrivateKey - indicates if the certificate contains private key.
// certificatenonce - random number generated to indicate Proof of Possession.
func (client DpsCertificateClient) GenerateVerificationCode(ctx context.Context, certificateName string, ifMatch string, resourceGroupName string, provisioningServiceName string, certificatename string, certificaterawBytes []byte, certificateisVerified *bool, certificatepurpose CertificatePurpose, certificatecreated *date.Time, certificatelastUpdated *date.Time, certificatehasPrivateKey *bool, certificatenonce string) (result VerificationCodeResponse, err error) {
	req, err := client.GenerateVerificationCodePreparer(ctx, certificateName, ifMatch, resourceGroupName, provisioningServiceName, certificatename, certificaterawBytes, certificateisVerified, certificatepurpose, certificatecreated, certificatelastUpdated, certificatehasPrivateKey, certificatenonce)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iothub.DpsCertificateClient", "GenerateVerificationCode", nil, "Failure preparing request")
		return
	}

	resp, err := client.GenerateVerificationCodeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "iothub.DpsCertificateClient", "GenerateVerificationCode", resp, "Failure sending request")
		return
	}

	result, err = client.GenerateVerificationCodeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iothub.DpsCertificateClient", "GenerateVerificationCode", resp, "Failure responding to request")
	}

	return
}

// GenerateVerificationCodePreparer prepares the GenerateVerificationCode request.
func (client DpsCertificateClient) GenerateVerificationCodePreparer(ctx context.Context, certificateName string, ifMatch string, resourceGroupName string, provisioningServiceName string, certificatename string, certificaterawBytes []byte, certificateisVerified *bool, certificatepurpose CertificatePurpose, certificatecreated *date.Time, certificatelastUpdated *date.Time, certificatehasPrivateKey *bool, certificatenonce string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"certificateName":         autorest.Encode("path", certificateName),
		"provisioningServiceName": autorest.Encode("path", provisioningServiceName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-08-21-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(certificatename) > 0 {
		queryParameters["certificate.name"] = autorest.Encode("query", certificatename)
	}
	if certificaterawBytes != nil && len(certificaterawBytes) > 0 {
		queryParameters["certificate.rawBytes"] = autorest.Encode("query", certificaterawBytes)
	}
	if certificateisVerified != nil {
		queryParameters["certificate.isVerified"] = autorest.Encode("query", *certificateisVerified)
	}
	if len(string(certificatepurpose)) > 0 {
		queryParameters["certificate.purpose"] = autorest.Encode("query", certificatepurpose)
	}
	if certificatecreated != nil {
		queryParameters["certificate.created"] = autorest.Encode("query", *certificatecreated)
	}
	if certificatelastUpdated != nil {
		queryParameters["certificate.lastUpdated"] = autorest.Encode("query", *certificatelastUpdated)
	}
	if certificatehasPrivateKey != nil {
		queryParameters["certificate.hasPrivateKey"] = autorest.Encode("query", *certificatehasPrivateKey)
	}
	if len(certificatenonce) > 0 {
		queryParameters["certificate.nonce"] = autorest.Encode("query", certificatenonce)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}/certificates/{certificateName}/generateVerificationCode", pathParameters),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GenerateVerificationCodeSender sends the GenerateVerificationCode request. The method will close the
// http.Response Body if it receives an error.
func (client DpsCertificateClient) GenerateVerificationCodeSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GenerateVerificationCodeResponder handles the response to the GenerateVerificationCode request. The method always
// closes the http.Response Body.
func (client DpsCertificateClient) GenerateVerificationCodeResponder(resp *http.Response) (result VerificationCodeResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get get the certificate from the provisioning service.
// Parameters:
// certificateName - name of the certificate to retrieve.
// resourceGroupName - resource group identifier.
// provisioningServiceName - name of the provisioning service the certificate is associated with.
// ifMatch - eTag of the certificate.
func (client DpsCertificateClient) Get(ctx context.Context, certificateName string, resourceGroupName string, provisioningServiceName string, ifMatch string) (result CertificateResponse, err error) {
	req, err := client.GetPreparer(ctx, certificateName, resourceGroupName, provisioningServiceName, ifMatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iothub.DpsCertificateClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "iothub.DpsCertificateClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iothub.DpsCertificateClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client DpsCertificateClient) GetPreparer(ctx context.Context, certificateName string, resourceGroupName string, provisioningServiceName string, ifMatch string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"certificateName":         autorest.Encode("path", certificateName),
		"provisioningServiceName": autorest.Encode("path", provisioningServiceName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-08-21-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}/certificates/{certificateName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(ifMatch) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DpsCertificateClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DpsCertificateClient) GetResponder(resp *http.Response) (result CertificateResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// VerifyCertificate verifies certificate for the provisioning service.
// Parameters:
// certificateName - the mandatory logical name of the certificate, that the provisioning service uses to
// access.
// ifMatch - eTag of the certificate.
// resourceGroupName - resource group name.
// provisioningServiceName - provisioning service name.
// certificatename - common Name for the certificate.
// certificaterawBytes - raw data of certificate.
// certificateisVerified - indicates if the certificate has been verified by owner of the private key.
// certificatepurpose - describe the purpose of the certificate.
// certificatecreated - certificate creation time.
// certificatelastUpdated - certificate last updated time.
// certificatehasPrivateKey - indicates if the certificate contains private key.
// certificatenonce - random number generated to indicate Proof of Possession.
func (client DpsCertificateClient) VerifyCertificate(ctx context.Context, certificateName string, ifMatch string, request VerificationCodeRequest, resourceGroupName string, provisioningServiceName string, certificatename string, certificaterawBytes []byte, certificateisVerified *bool, certificatepurpose CertificatePurpose, certificatecreated *date.Time, certificatelastUpdated *date.Time, certificatehasPrivateKey *bool, certificatenonce string) (result CertificateResponse, err error) {
	req, err := client.VerifyCertificatePreparer(ctx, certificateName, ifMatch, request, resourceGroupName, provisioningServiceName, certificatename, certificaterawBytes, certificateisVerified, certificatepurpose, certificatecreated, certificatelastUpdated, certificatehasPrivateKey, certificatenonce)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iothub.DpsCertificateClient", "VerifyCertificate", nil, "Failure preparing request")
		return
	}

	resp, err := client.VerifyCertificateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "iothub.DpsCertificateClient", "VerifyCertificate", resp, "Failure sending request")
		return
	}

	result, err = client.VerifyCertificateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iothub.DpsCertificateClient", "VerifyCertificate", resp, "Failure responding to request")
	}

	return
}

// VerifyCertificatePreparer prepares the VerifyCertificate request.
func (client DpsCertificateClient) VerifyCertificatePreparer(ctx context.Context, certificateName string, ifMatch string, request VerificationCodeRequest, resourceGroupName string, provisioningServiceName string, certificatename string, certificaterawBytes []byte, certificateisVerified *bool, certificatepurpose CertificatePurpose, certificatecreated *date.Time, certificatelastUpdated *date.Time, certificatehasPrivateKey *bool, certificatenonce string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"certificateName":         autorest.Encode("path", certificateName),
		"provisioningServiceName": autorest.Encode("path", provisioningServiceName),
		"resourceGroupName":       autorest.Encode("path", resourceGroupName),
		"subscriptionId":          autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-08-21-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(certificatename) > 0 {
		queryParameters["certificate.name"] = autorest.Encode("query", certificatename)
	}
	if certificaterawBytes != nil && len(certificaterawBytes) > 0 {
		queryParameters["certificate.rawBytes"] = autorest.Encode("query", certificaterawBytes)
	}
	if certificateisVerified != nil {
		queryParameters["certificate.isVerified"] = autorest.Encode("query", *certificateisVerified)
	}
	if len(string(certificatepurpose)) > 0 {
		queryParameters["certificate.purpose"] = autorest.Encode("query", certificatepurpose)
	}
	if certificatecreated != nil {
		queryParameters["certificate.created"] = autorest.Encode("query", *certificatecreated)
	}
	if certificatelastUpdated != nil {
		queryParameters["certificate.lastUpdated"] = autorest.Encode("query", *certificatelastUpdated)
	}
	if certificatehasPrivateKey != nil {
		queryParameters["certificate.hasPrivateKey"] = autorest.Encode("query", *certificatehasPrivateKey)
	}
	if len(certificatenonce) > 0 {
		queryParameters["certificate.nonce"] = autorest.Encode("query", certificatenonce)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Devices/provisioningServices/{provisioningServiceName}/certificates/{certificateName}/verify", pathParameters),
		autorest.WithJSON(request),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("If-Match", autorest.String(ifMatch)))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// VerifyCertificateSender sends the VerifyCertificate request. The method will close the
// http.Response Body if it receives an error.
func (client DpsCertificateClient) VerifyCertificateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// VerifyCertificateResponder handles the response to the VerifyCertificate request. The method always
// closes the http.Response Body.
func (client DpsCertificateClient) VerifyCertificateResponder(resp *http.Response) (result CertificateResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
