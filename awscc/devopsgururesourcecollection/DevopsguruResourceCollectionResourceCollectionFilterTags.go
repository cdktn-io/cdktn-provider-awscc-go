// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsgururesourcecollection


type DevopsguruResourceCollectionResourceCollectionFilterTags struct {
	// A Tag key for DevOps Guru app boundary.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/devopsguru_resource_collection#app_boundary_key DevopsguruResourceCollection#app_boundary_key}
	AppBoundaryKey *string `field:"optional" json:"appBoundaryKey" yaml:"appBoundaryKey"`
	// Tag values of DevOps Guru app boundary.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/devopsguru_resource_collection#tag_values DevopsguruResourceCollection#tag_values}
	TagValues *[]*string `field:"optional" json:"tagValues" yaml:"tagValues"`
}

