// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2routeserverendpoint

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2RouteServerEndpointConfig struct {
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
	// Route Server ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_route_server_endpoint#route_server_id Ec2RouteServerEndpoint#route_server_id}
	RouteServerId *string `field:"required" json:"routeServerId" yaml:"routeServerId"`
	// Subnet ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_route_server_endpoint#subnet_id Ec2RouteServerEndpoint#subnet_id}
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_route_server_endpoint#tags Ec2RouteServerEndpoint#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

