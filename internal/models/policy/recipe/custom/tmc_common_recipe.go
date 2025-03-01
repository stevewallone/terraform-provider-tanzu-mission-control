// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

// Code generated by go-swagger; DO NOT EDIT.

package policyrecipecustommodel

import (
	"github.com/go-openapi/swag"

	policyrecipecustomcommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/custom/common"
)

// VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCCommonRecipe is model for tmc-block-nodeport-service, tmc-block-resources and tmc-https-ingress recipe types.
//
// The input schema for tmc-block-nodeport-service, tmc-block-resources and tmc-https-ingress recipes.
//
// swagger:model VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCCommonRecipe
type VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCCommonRecipe struct {

	// Audit (dry-run).
	// Creates this policy for dry-run. Violations will be logged but not denied. Defaults to false (deny).
	Audit bool `json:"audit,omitempty"`

	// TargetKubernetesResources is a list of kubernetes api resources on which the policy will be enforced, identified using apiGroups and kinds. You can use 'kubectl api-resources' to view the list of available api resources on your cluster.
	// Required: true
	// Min Items: 1
	TargetKubernetesResources []*policyrecipecustomcommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TargetKubernetesResources `json:"targetKubernetesResources"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCCommonRecipe) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCCommonRecipe) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonPolicySpecCustomV1TMCCommonRecipe
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}
