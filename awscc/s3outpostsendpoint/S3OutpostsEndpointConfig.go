// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3outpostsendpoint

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type S3OutpostsEndpointConfig struct {
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
	// The id of the customer outpost on which the bucket resides.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/s3outposts_endpoint#outpost_id S3OutpostsEndpoint#outpost_id}
	OutpostId *string `field:"required" json:"outpostId" yaml:"outpostId"`
	// The ID of the security group to use with the endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/s3outposts_endpoint#security_group_id S3OutpostsEndpoint#security_group_id}
	SecurityGroupId *string `field:"required" json:"securityGroupId" yaml:"securityGroupId"`
	// The ID of the subnet in the selected VPC. The subnet must belong to the Outpost.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/s3outposts_endpoint#subnet_id S3OutpostsEndpoint#subnet_id}
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// The type of access for the on-premise network connectivity for the Outpost endpoint.
	//
	// To access endpoint from an on-premises network, you must specify the access type and provide the customer owned Ipv4 pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/s3outposts_endpoint#access_type S3OutpostsEndpoint#access_type}
	AccessType *string `field:"optional" json:"accessType" yaml:"accessType"`
	// The ID of the customer-owned IPv4 pool for the Endpoint.
	//
	// IP addresses will be allocated from this pool for the endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/s3outposts_endpoint#customer_owned_ipv_4_pool S3OutpostsEndpoint#customer_owned_ipv_4_pool}
	CustomerOwnedIpv4Pool *string `field:"optional" json:"customerOwnedIpv4Pool" yaml:"customerOwnedIpv4Pool"`
	// The failure reason, if any, for a create or delete endpoint operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/s3outposts_endpoint#failed_reason S3OutpostsEndpoint#failed_reason}
	FailedReason *S3OutpostsEndpointFailedReason `field:"optional" json:"failedReason" yaml:"failedReason"`
}

