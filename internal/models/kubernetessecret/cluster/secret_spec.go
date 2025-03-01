// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: MPL-2.0

package secret

import (
	"encoding/json"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VmwareTanzuManageV1alpha1ClusterNamespaceSecretSpec Spec of the Secret.
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.secret.Spec
type VmwareTanzuManageV1alpha1ClusterNamespaceSecretSpec struct {

	// Payload of the Secret.
	Data map[string]strfmt.Base64 `json:"data,omitempty"`

	// Type of the Secret.
	SecretType *VmwareTanzuManageV1alpha1ClusterNamespaceSecretType `json:"secretType,omitempty"`
}

// MarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceSecretSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation.
func (m *VmwareTanzuManageV1alpha1ClusterNamespaceSecretSpec) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1ClusterNamespaceSecretSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}

// VmwareTanzuManageV1alpha1ClusterNamespaceSecretType SecretType definition - indicates the kubernetes secret type.
//
//   - SECRET_TYPE_UNSPECIFIED: SECRET_TYPE_UNSPECIFIED, Unspecified secret type (default).
//   - SECRET_TYPE_DOCKERCONFIGJSON: SECRET_TYPE_DOCKERCONFIGJSON, Kubernetes secrets type : kubernetes.io/dockerconfigjson.
//   - SECRET_TYPE_OPAQUE: SECRET_TYPE_OPAQUE, Kubernetes opaque secret type : https://kubernetes.io/docs/concepts/configuration/secret/#opaque-secrets
//
// swagger:model vmware.tanzu.manage.v1alpha1.cluster.namespace.secret.SecretType
type VmwareTanzuManageV1alpha1ClusterNamespaceSecretType string

func NewVmwareTanzuManageV1alpha1ClusterNamespaceSecretType(value VmwareTanzuManageV1alpha1ClusterNamespaceSecretType) *VmwareTanzuManageV1alpha1ClusterNamespaceSecretType {
	return &value
}

const (

	// VmwareTanzuManageV1alpha1ClusterNamespaceSecretTypeSECRETTYPEDOCKERCONFIGJSON captures enum value "SECRET_TYPE_DOCKERCONFIGJSON".
	VmwareTanzuManageV1alpha1ClusterNamespaceSecretTypeSECRETTYPEDOCKERCONFIGJSON VmwareTanzuManageV1alpha1ClusterNamespaceSecretType = "SECRET_TYPE_DOCKERCONFIGJSON"
	// VmwareTanzuManageV1alpha1ClusterNamespaceSecretTypeSECRETTYPEOPAQUE captures enum value "SECRET_TYPE_OPAQUE".
	//nolint:gosec
	VmwareTanzuManageV1alpha1ClusterNamespaceSecretTypeSECRETTYPEOPAQUE VmwareTanzuManageV1alpha1ClusterNamespaceSecretType = "SECRET_TYPE_OPAQUE"
)

// for schema.
var vmwareTanzuManageV1alpha1ClusterNamespaceSecretTypeEnum []interface{}

func init() {
	var res []VmwareTanzuManageV1alpha1ClusterNamespaceSecretType
	if err := json.Unmarshal([]byte(`["SECRET_TYPE_UNSPECIFIED","SECRET_TYPE_DOCKERCONFIGJSON","SECRET_TYPE_OPAQUE"]`), &res); err != nil {
		panic(err)
	}

	for _, v := range res {
		vmwareTanzuManageV1alpha1ClusterNamespaceSecretTypeEnum = append(vmwareTanzuManageV1alpha1ClusterNamespaceSecretTypeEnum, v)
	}
}
