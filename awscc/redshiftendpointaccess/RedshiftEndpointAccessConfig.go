// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package redshiftendpointaccess

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RedshiftEndpointAccessConfig struct {
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
	// A unique identifier for the cluster.
	//
	// You use this identifier to refer to the cluster for any subsequent cluster operations such as deleting or modifying. All alphabetical characters must be lower case, no hypens at the end, no two consecutive hyphens. Cluster name should be unique for all clusters within an AWS account
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/redshift_endpoint_access#cluster_identifier RedshiftEndpointAccess#cluster_identifier}
	ClusterIdentifier *string `field:"required" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// The name of the endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/redshift_endpoint_access#endpoint_name RedshiftEndpointAccess#endpoint_name}
	EndpointName *string `field:"required" json:"endpointName" yaml:"endpointName"`
	// The subnet group name where Amazon Redshift chooses to deploy the endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/redshift_endpoint_access#subnet_group_name RedshiftEndpointAccess#subnet_group_name}
	SubnetGroupName *string `field:"required" json:"subnetGroupName" yaml:"subnetGroupName"`
	// A list of vpc security group ids to apply to the created endpoint access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/redshift_endpoint_access#vpc_security_group_ids RedshiftEndpointAccess#vpc_security_group_ids}
	VpcSecurityGroupIds *[]*string `field:"required" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
	// The AWS account ID of the owner of the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/redshift_endpoint_access#resource_owner RedshiftEndpointAccess#resource_owner}
	ResourceOwner *string `field:"optional" json:"resourceOwner" yaml:"resourceOwner"`
}

