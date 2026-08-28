// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceServiceConnectConfigurationServices struct {
	// The list of client aliases for this Service Connect service.
	//
	// You use these to assign names that can be used by client applications. The maximum number of client aliases that you can have in this list is 1.
	//  Each alias ("endpoint") is a fully-qualified name and port number that other Amazon ECS tasks ("clients") can use to connect to this service.
	//  Each name and port mapping must be unique within the namespace.
	//  For each ``ServiceConnectService``, you must provide at least one ``clientAlias`` with one ``port``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ecs_service#client_aliases EcsService#client_aliases}
	ClientAliases interface{} `field:"optional" json:"clientAliases" yaml:"clientAliases"`
	// The ``discoveryName`` is the name of the new CMAP service that Amazon ECS creates for this Amazon ECS service.
	//
	// This must be unique within the CMAP namespace. The name can contain up to 64 characters. The name can include lowercase letters, numbers, underscores (_), and hyphens (-). The name can't start with a hyphen.
	//  If the ``discoveryName`` isn't specified, the port mapping name from the task definition is used in ``portName.namespace``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ecs_service#discovery_name EcsService#discovery_name}
	DiscoveryName *string `field:"optional" json:"discoveryName" yaml:"discoveryName"`
	// The port number for the Service Connect proxy to listen on.
	//
	// Use the value of this field to bypass the proxy for traffic on the port number specified in the named ``portMapping`` in the task definition of this application, and then use it in your VPC security groups to allow traffic into the proxy for this Amazon ECS service.
	//  In ``awsvpc`` mode and Fargate, the default value is the container port number. The container port number is in the ``portMapping`` in the task definition. In bridge mode, the default value is the ephemeral port of the Service Connect proxy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ecs_service#ingress_port_override EcsService#ingress_port_override}
	IngressPortOverride *float64 `field:"optional" json:"ingressPortOverride" yaml:"ingressPortOverride"`
	// The ``portName`` must match the name of one of the ``portMappings`` from all the containers in the task definition of this Amazon ECS service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ecs_service#port_name EcsService#port_name}
	PortName *string `field:"optional" json:"portName" yaml:"portName"`
	// A reference to an object that represents the configured timeouts for Service Connect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ecs_service#timeout EcsService#timeout}
	Timeout *EcsServiceServiceConnectConfigurationServicesTimeout `field:"optional" json:"timeout" yaml:"timeout"`
	// A reference to an object that represents a Transport Layer Security (TLS) configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ecs_service#tls EcsService#tls}
	Tls *EcsServiceServiceConnectConfigurationServicesTls `field:"optional" json:"tls" yaml:"tls"`
}

