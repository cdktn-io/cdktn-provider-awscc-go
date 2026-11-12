// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectdestinationendpoint

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediaconnectDestinationEndpointConfig struct {
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
	// The configuration settings for a destination endpoint.
	//
	// This is a union type, so exactly one of its members is set, depending on the type of the destination endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_destination_endpoint#configuration MediaconnectDestinationEndpoint#configuration}
	Configuration *MediaconnectDestinationEndpointConfiguration `field:"required" json:"configuration" yaml:"configuration"`
	// The description of the destination endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_destination_endpoint#description MediaconnectDestinationEndpoint#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the Amazon Web Services Region in which the destination endpoint is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_destination_endpoint#region_name MediaconnectDestinationEndpoint#region_name}
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
	// The tags associated with the destination endpoint. Each tag is a key-value pair.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_destination_endpoint#tags MediaconnectDestinationEndpoint#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

