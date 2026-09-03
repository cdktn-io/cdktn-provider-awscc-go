// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lightsailcontainer


type LightsailContainerContainerServiceDeployment struct {
	// An object that describes the configuration for the containers of the deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lightsail_container#containers LightsailContainer#containers}
	Containers interface{} `field:"optional" json:"containers" yaml:"containers"`
	// An object that describes the endpoint of the deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lightsail_container#public_endpoint LightsailContainer#public_endpoint}
	PublicEndpoint *LightsailContainerContainerServiceDeploymentPublicEndpoint `field:"optional" json:"publicEndpoint" yaml:"publicEndpoint"`
}

