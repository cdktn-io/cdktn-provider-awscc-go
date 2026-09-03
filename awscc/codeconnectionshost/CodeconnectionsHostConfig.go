// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codeconnectionshost

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CodeconnectionsHostConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The name of the host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/codeconnections_host#name CodeconnectionsHost#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The endpoint of the infrastructure where your provider type is installed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/codeconnections_host#provider_endpoint CodeconnectionsHost#provider_endpoint}
	ProviderEndpoint *string `field:"required" json:"providerEndpoint" yaml:"providerEndpoint"`
	// The name of the installed provider to be associated with your connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/codeconnections_host#provider_type CodeconnectionsHost#provider_type}
	ProviderType *string `field:"required" json:"providerType" yaml:"providerType"`
	// Tags to apply to the host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/codeconnections_host#tags CodeconnectionsHost#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The VPC configuration provisioned for the host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/codeconnections_host#vpc_configuration CodeconnectionsHost#vpc_configuration}
	VpcConfiguration *CodeconnectionsHostVpcConfiguration `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

