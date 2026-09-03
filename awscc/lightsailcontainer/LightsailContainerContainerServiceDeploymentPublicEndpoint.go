// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lightsailcontainer


type LightsailContainerContainerServiceDeploymentPublicEndpoint struct {
	// The name of the container for the endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lightsail_container#container_name LightsailContainer#container_name}
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// The port of the container to which traffic is forwarded to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lightsail_container#container_port LightsailContainer#container_port}
	ContainerPort *float64 `field:"optional" json:"containerPort" yaml:"containerPort"`
	// An object that describes the health check configuration of the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lightsail_container#health_check_config LightsailContainer#health_check_config}
	HealthCheckConfig *LightsailContainerContainerServiceDeploymentPublicEndpointHealthCheckConfig `field:"optional" json:"healthCheckConfig" yaml:"healthCheckConfig"`
}

