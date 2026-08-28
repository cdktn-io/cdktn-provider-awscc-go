// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package deadlinelicenseendpoint

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DeadlineLicenseEndpointConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/deadline_license_endpoint#security_group_ids DeadlineLicenseEndpoint#security_group_ids}.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/deadline_license_endpoint#subnet_ids DeadlineLicenseEndpoint#subnet_ids}.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/deadline_license_endpoint#vpc_id DeadlineLicenseEndpoint#vpc_id}.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/deadline_license_endpoint#tags DeadlineLicenseEndpoint#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

