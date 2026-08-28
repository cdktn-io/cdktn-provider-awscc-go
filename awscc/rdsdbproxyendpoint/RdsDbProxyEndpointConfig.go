// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rdsdbproxyendpoint

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RdsDbProxyEndpointConfig struct {
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
	// The identifier for the DB proxy endpoint.
	//
	// This name must be unique for all DB proxy endpoints owned by your AWS account in the specified AWS Region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rds_db_proxy_endpoint#db_proxy_endpoint_name RdsDbProxyEndpoint#db_proxy_endpoint_name}
	DbProxyEndpointName *string `field:"required" json:"dbProxyEndpointName" yaml:"dbProxyEndpointName"`
	// The identifier for the proxy.
	//
	// This name must be unique for all proxies owned by your AWS account in the specified AWS Region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rds_db_proxy_endpoint#db_proxy_name RdsDbProxyEndpoint#db_proxy_name}
	DbProxyName *string `field:"required" json:"dbProxyName" yaml:"dbProxyName"`
	// VPC subnet IDs to associate with the new DB proxy endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rds_db_proxy_endpoint#vpc_subnet_ids RdsDbProxyEndpoint#vpc_subnet_ids}
	VpcSubnetIds *[]*string `field:"required" json:"vpcSubnetIds" yaml:"vpcSubnetIds"`
	// The network type of the DB proxy endpoint.
	//
	// The network type determines the IP version that the proxy endpoint supports.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rds_db_proxy_endpoint#endpoint_network_type RdsDbProxyEndpoint#endpoint_network_type}
	EndpointNetworkType *string `field:"optional" json:"endpointNetworkType" yaml:"endpointNetworkType"`
	// An optional set of key-value pairs to associate arbitrary data of your choosing with the DB proxy endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rds_db_proxy_endpoint#tags RdsDbProxyEndpoint#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// A value that indicates whether the DB proxy endpoint can be used for read/write or read-only operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rds_db_proxy_endpoint#target_role RdsDbProxyEndpoint#target_role}
	TargetRole *string `field:"optional" json:"targetRole" yaml:"targetRole"`
	// VPC security group IDs to associate with the new DB proxy endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rds_db_proxy_endpoint#vpc_security_group_ids RdsDbProxyEndpoint#vpc_security_group_ids}
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}

