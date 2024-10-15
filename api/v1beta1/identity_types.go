/*
Copyright 2023 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

// OpenStackIdentityReference is a reference to an infrastructure
// provider identity to be used to provision cluster resources.
type OpenStackIdentityReference struct {
	// Name is the name of a secret in the same namespace as the resource being provisioned.
	// The secret must contain a key named `clouds.yaml` which contains an OpenStack clouds.yaml file.
	// The secret may optionally contain a key named `cacert` containing a PEM-encoded CA certificate.
	// +kubebuilder:validation:Required
	Name string `json:"name"`

	// CloudName specifies the name of the entry in the clouds.yaml file to use.
	// +kubebuilder:validation:Required
	CloudName string `json:"cloudName"`

	// ProviderIDRegion specifies the name of the region to specify in providerID field format `openstack://ProviderIDRegion/instance_uuid`.
	// This option permit to been compatible with multiple OCCM (one per OS cluster).
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="region is immutable"
	// +optional
	ProviderIDRegion string `json:"providerIDRegion,omitempty"`
}

// IdentityRefProvider is an interface for obtaining OpenStack credentials from an API object
// +kubebuilder:object:generate:=false
type IdentityRefProvider interface {
	// GetIdentifyRef returns the object's namespace and IdentityRef if it has an IdentityRef, or nulls if it does not
	GetIdentityRef() (*string, *OpenStackIdentityReference)
}
