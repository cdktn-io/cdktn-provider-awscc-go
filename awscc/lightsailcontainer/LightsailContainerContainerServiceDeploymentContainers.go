// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lightsailcontainer


type LightsailContainerContainerServiceDeploymentContainers struct {
	// The launch command for the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lightsail_container#command LightsailContainer#command}
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The name of the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lightsail_container#container_name LightsailContainer#container_name}
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// The environment variables of the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lightsail_container#environment LightsailContainer#environment}
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// The name of the image used for the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lightsail_container#image LightsailContainer#image}
	Image *string `field:"optional" json:"image" yaml:"image"`
	// The open firewall ports of the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lightsail_container#ports LightsailContainer#ports}
	Ports interface{} `field:"optional" json:"ports" yaml:"ports"`
}

