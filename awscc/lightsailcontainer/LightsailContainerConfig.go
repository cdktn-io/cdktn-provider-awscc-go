// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lightsailcontainer

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LightsailContainerConfig struct {
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
	// The power specification for the container service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lightsail_container#power LightsailContainer#power}
	Power *string `field:"required" json:"power" yaml:"power"`
	// The scale specification for the container service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lightsail_container#scale LightsailContainer#scale}
	Scale *float64 `field:"required" json:"scale" yaml:"scale"`
	// The name for the container service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lightsail_container#service_name LightsailContainer#service_name}
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
	// Describes a container deployment configuration of an Amazon Lightsail container service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lightsail_container#container_service_deployment LightsailContainer#container_service_deployment}
	ContainerServiceDeployment *LightsailContainerContainerServiceDeployment `field:"optional" json:"containerServiceDeployment" yaml:"containerServiceDeployment"`
	// A Boolean value to indicate whether the container service is disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lightsail_container#is_disabled LightsailContainer#is_disabled}
	IsDisabled interface{} `field:"optional" json:"isDisabled" yaml:"isDisabled"`
	// A Boolean value to indicate whether the container service has access to private container image repositories, such as Amazon Elastic Container Registry (Amazon ECR) private repositories.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lightsail_container#private_registry_access LightsailContainer#private_registry_access}
	PrivateRegistryAccess *LightsailContainerPrivateRegistryAccess `field:"optional" json:"privateRegistryAccess" yaml:"privateRegistryAccess"`
	// The public domain names to use with the container service, such as example.com and www.example.com.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lightsail_container#public_domain_names LightsailContainer#public_domain_names}
	PublicDomainNames interface{} `field:"optional" json:"publicDomainNames" yaml:"publicDomainNames"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lightsail_container#tags LightsailContainer#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

