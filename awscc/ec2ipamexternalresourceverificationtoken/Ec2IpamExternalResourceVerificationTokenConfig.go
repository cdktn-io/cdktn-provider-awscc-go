// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2ipamexternalresourceverificationtoken

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2IpamExternalResourceVerificationTokenConfig struct {
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
	// The ID of the IPAM that will create the token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_ipam_external_resource_verification_token#ipam_id Ec2IpamExternalResourceVerificationToken#ipam_id}
	IpamId *string `field:"required" json:"ipamId" yaml:"ipamId"`
	// The tags for the token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_ipam_external_resource_verification_token#tags Ec2IpamExternalResourceVerificationToken#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

