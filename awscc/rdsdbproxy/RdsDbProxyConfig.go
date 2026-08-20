// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rdsdbproxy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RdsDbProxyConfig struct {
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
	// The identifier for the proxy.
	//
	// This name must be unique for all proxies owned by your AWS account in the specified AWS Region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/rds_db_proxy#db_proxy_name RdsDbProxy#db_proxy_name}
	DbProxyName *string `field:"required" json:"dbProxyName" yaml:"dbProxyName"`
	// The kinds of databases that the proxy can connect to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/rds_db_proxy#engine_family RdsDbProxy#engine_family}
	EngineFamily *string `field:"required" json:"engineFamily" yaml:"engineFamily"`
	// The Amazon Resource Name (ARN) of the IAM role that the proxy uses to access secrets in AWS Secrets Manager.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/rds_db_proxy#role_arn RdsDbProxy#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// VPC subnet IDs to associate with the new proxy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/rds_db_proxy#vpc_subnet_ids RdsDbProxy#vpc_subnet_ids}
	VpcSubnetIds *[]*string `field:"required" json:"vpcSubnetIds" yaml:"vpcSubnetIds"`
	// The authorization mechanism that the proxy uses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/rds_db_proxy#auth RdsDbProxy#auth}
	Auth interface{} `field:"optional" json:"auth" yaml:"auth"`
	// Whether the proxy includes detailed information about SQL statements in its logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/rds_db_proxy#debug_logging RdsDbProxy#debug_logging}
	DebugLogging interface{} `field:"optional" json:"debugLogging" yaml:"debugLogging"`
	// The default authentication scheme that the proxy uses for client connections to the proxy and connections from the proxy to the underlying database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/rds_db_proxy#default_auth_scheme RdsDbProxy#default_auth_scheme}
	DefaultAuthScheme *string `field:"optional" json:"defaultAuthScheme" yaml:"defaultAuthScheme"`
	// The network type of the DB proxy endpoint.
	//
	// The network type determines the IP version that the proxy endpoint supports.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/rds_db_proxy#endpoint_network_type RdsDbProxy#endpoint_network_type}
	EndpointNetworkType *string `field:"optional" json:"endpointNetworkType" yaml:"endpointNetworkType"`
	// The number of seconds that a connection to the proxy can be inactive before the proxy disconnects it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/rds_db_proxy#idle_client_timeout RdsDbProxy#idle_client_timeout}
	IdleClientTimeout *float64 `field:"optional" json:"idleClientTimeout" yaml:"idleClientTimeout"`
	// A Boolean parameter that specifies whether Transport Layer Security (TLS) encryption is required for connections to the proxy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/rds_db_proxy#require_tls RdsDbProxy#require_tls}
	RequireTls interface{} `field:"optional" json:"requireTls" yaml:"requireTls"`
	// An optional set of key-value pairs to associate arbitrary data of your choosing with the proxy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/rds_db_proxy#tags RdsDbProxy#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The network type that the proxy uses to connect to the target database.
	//
	// The network type determines the IP version that the proxy uses for connections to the database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/rds_db_proxy#target_connection_network_type RdsDbProxy#target_connection_network_type}
	TargetConnectionNetworkType *string `field:"optional" json:"targetConnectionNetworkType" yaml:"targetConnectionNetworkType"`
	// VPC security group IDs to associate with the new proxy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/rds_db_proxy#vpc_security_group_ids RdsDbProxy#vpc_security_group_ids}
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}

