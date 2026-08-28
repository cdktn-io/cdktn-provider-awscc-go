// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2vpcendpointservicepermissions

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2VpcEndpointServicePermissionsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_vpc_endpoint_service_permissions#service_id Ec2VpcEndpointServicePermissions#service_id}.
	ServiceId *string `field:"required" json:"serviceId" yaml:"serviceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_vpc_endpoint_service_permissions#allowed_principals Ec2VpcEndpointServicePermissions#allowed_principals}.
	AllowedPrincipals *[]*string `field:"optional" json:"allowedPrincipals" yaml:"allowedPrincipals"`
}

