// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightvpcconnection

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuicksightVpcConnectionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_vpc_connection#availability_status QuicksightVpcConnection#availability_status}.
	AvailabilityStatus *string `field:"optional" json:"availabilityStatus" yaml:"availabilityStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_vpc_connection#aws_account_id QuicksightVpcConnection#aws_account_id}.
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_vpc_connection#dns_resolvers QuicksightVpcConnection#dns_resolvers}.
	DnsResolvers *[]*string `field:"optional" json:"dnsResolvers" yaml:"dnsResolvers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_vpc_connection#name QuicksightVpcConnection#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_vpc_connection#role_arn QuicksightVpcConnection#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_vpc_connection#security_group_ids QuicksightVpcConnection#security_group_ids}.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_vpc_connection#subnet_ids QuicksightVpcConnection#subnet_ids}.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_vpc_connection#tags QuicksightVpcConnection#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_vpc_connection#vpc_connection_id QuicksightVpcConnection#vpc_connection_id}.
	VpcConnectionId *string `field:"optional" json:"vpcConnectionId" yaml:"vpcConnectionId"`
}

