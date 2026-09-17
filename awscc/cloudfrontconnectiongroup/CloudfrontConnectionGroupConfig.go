// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfrontconnectiongroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CloudfrontConnectionGroupConfig struct {
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
	// The name of the connection group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudfront_connection_group#name CloudfrontConnectionGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the Anycast static IP list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudfront_connection_group#anycast_ip_list_id CloudfrontConnectionGroup#anycast_ip_list_id}
	AnycastIpListId *string `field:"optional" json:"anycastIpListId" yaml:"anycastIpListId"`
	// Whether the connection group is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudfront_connection_group#enabled CloudfrontConnectionGroup#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// IPv6 is enabled for the connection group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudfront_connection_group#ipv_6_enabled CloudfrontConnectionGroup#ipv_6_enabled}
	Ipv6Enabled interface{} `field:"optional" json:"ipv6Enabled" yaml:"ipv6Enabled"`
	// A complex type that contains zero or more ``Tag`` elements.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudfront_connection_group#tags CloudfrontConnectionGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

