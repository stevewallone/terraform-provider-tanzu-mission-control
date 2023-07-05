/*
Copyright 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package akscluster

import (
	"net/url"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/client/transport"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	aksmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/akscluster"
)

const (
	apiVersionAndGroup             = "/v1alpha1/aksclusters"
	queryParamKeyForce             = "force"
	queryParamKeyCredentialName    = "fullName.credentialName" //nolint:gosec
	queryParamKeySubscriptionID    = "fullName.subscriptionId"
	queryParamKeyResourceGroupName = "fullName.resourceGroupName"
)

// New creates a new aks cluster resource service API client.
func New(transport *transport.Client) ClientService {
	return &Client{Client: transport}
}

/*
Client for aks cluster resource service API.
*/
type Client struct {
	*transport.Client
}

// ClientService is the interface for Client methods.
type ClientService interface {
	AksClusterResourceServiceCreate(request *aksmodel.VmwareTanzuManageV1alpha1AksclusterCreateAksClusterRequest) (*aksmodel.VmwareTanzuManageV1alpha1AksclusterCreateAksClusterResponse, error)

	AksClusterResourceServiceGet(fn *aksmodel.VmwareTanzuManageV1alpha1AksclusterFullName) (*aksmodel.VmwareTanzuManageV1alpha1AksclusterGetAksClusterResponse, error)

	AksClusterResourceServiceUpdate(request *aksmodel.VmwareTanzuManageV1alpha1AksclusterUpdateAksClusterRequest) (*aksmodel.VmwareTanzuManageV1alpha1AksclusterUpdateAksClusterResponse, error)

	AksClusterResourceServiceDelete(fn *aksmodel.VmwareTanzuManageV1alpha1AksclusterFullName, force string) error
}

/*
AksClusterResourceServiceCreate creates an aks cluster.
*/
func (c *Client) AksClusterResourceServiceCreate(request *aksmodel.VmwareTanzuManageV1alpha1AksclusterCreateAksClusterRequest) (*aksmodel.VmwareTanzuManageV1alpha1AksclusterCreateAksClusterResponse, error) {
	response := &aksmodel.VmwareTanzuManageV1alpha1AksclusterCreateAksClusterResponse{}
	err := c.Create(apiVersionAndGroup, request, response)

	return response, err
}

/*
AksClusterResourceServiceGet gets an aks cluster.
*/
func (c *Client) AksClusterResourceServiceGet(fn *aksmodel.VmwareTanzuManageV1alpha1AksclusterFullName) (*aksmodel.VmwareTanzuManageV1alpha1AksclusterGetAksClusterResponse, error) {
	queryParams := url.Values{}
	if fn.CredentialName != "" {
		queryParams.Add(queryParamKeyCredentialName, fn.CredentialName)
	}

	if fn.SubscriptionID != "" {
		queryParams.Add(queryParamKeySubscriptionID, fn.SubscriptionID)
	}

	if fn.ResourceGroupName != "" {
		queryParams.Add(queryParamKeyResourceGroupName, fn.ResourceGroupName)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.Name).AppendQueryParams(queryParams).String()
	clusterResponse := &aksmodel.VmwareTanzuManageV1alpha1AksclusterGetAksClusterResponse{}

	err := c.Get(requestURL, clusterResponse)

	return clusterResponse, err
}

/*
AksClusterResourceServiceUpdate updates overwrite an aks cluster.
*/
func (c *Client) AksClusterResourceServiceUpdate(request *aksmodel.VmwareTanzuManageV1alpha1AksclusterUpdateAksClusterRequest) (*aksmodel.VmwareTanzuManageV1alpha1AksclusterUpdateAksClusterResponse, error) {
	response := &aksmodel.VmwareTanzuManageV1alpha1AksclusterUpdateAksClusterResponse{}
	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, request.AksCluster.FullName.Name).String()
	err := c.Update(requestURL, request, response)

	return response, err
}

/*
AksClusterResourceServiceDelete deletes an aks cluster.
*/
func (c *Client) AksClusterResourceServiceDelete(fn *aksmodel.VmwareTanzuManageV1alpha1AksclusterFullName, force string) error {
	queryParams := url.Values{
		queryParamKeyForce: []string{force},
	}
	if fn.CredentialName != "" {
		queryParams.Add(queryParamKeyCredentialName, fn.CredentialName)
	}

	if fn.SubscriptionID != "" {
		queryParams.Add(queryParamKeySubscriptionID, fn.SubscriptionID)
	}

	if fn.ResourceGroupName != "" {
		queryParams.Add(queryParamKeyResourceGroupName, fn.ResourceGroupName)
	}

	requestURL := helper.ConstructRequestURL(apiVersionAndGroup, fn.Name).AppendQueryParams(queryParams).String()

	return c.Delete(requestURL)
}
