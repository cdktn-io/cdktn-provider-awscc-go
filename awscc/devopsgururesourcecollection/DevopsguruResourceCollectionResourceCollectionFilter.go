// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsgururesourcecollection


type DevopsguruResourceCollectionResourceCollectionFilter struct {
	// CloudFormation resource for DevOps Guru to monitor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/devopsguru_resource_collection#cloudformation DevopsguruResourceCollection#cloudformation}
	Cloudformation *DevopsguruResourceCollectionResourceCollectionFilterCloudformation `field:"optional" json:"cloudformation" yaml:"cloudformation"`
	// Tagged resources for DevOps Guru to monitor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/devopsguru_resource_collection#tags DevopsguruResourceCollection#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

