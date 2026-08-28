// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2ipset

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Wafv2IpSetConfig struct {
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
	// List of IPAddresses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wafv2_ip_set#addresses Wafv2IpSet#addresses}
	Addresses *[]*string `field:"required" json:"addresses" yaml:"addresses"`
	// Type of addresses in the IPSet, use IPV4 for IPV4 IP addresses, IPV6 for IPV6 address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wafv2_ip_set#ip_address_version Wafv2IpSet#ip_address_version}
	IpAddressVersion *string `field:"required" json:"ipAddressVersion" yaml:"ipAddressVersion"`
	// Use CLOUDFRONT for CloudFront IPSet, use REGIONAL for Application Load Balancer and API Gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wafv2_ip_set#scope Wafv2IpSet#scope}
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// Description of the entity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wafv2_ip_set#description Wafv2IpSet#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Name of the IPSet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wafv2_ip_set#name Wafv2IpSet#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wafv2_ip_set#tags Wafv2IpSet#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

