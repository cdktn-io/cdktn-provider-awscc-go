// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceServiceConnectConfigurationServicesClientAliases struct {
	// The ``dnsName`` is the name that you use in the applications of client tasks to connect to this service.
	//
	// The name must be a valid DNS name but doesn't need to be fully-qualified. The name can include up to 127 characters. The name can include lowercase letters, numbers, underscores (_), hyphens (-), and periods (.). The name can't start with a hyphen.
	//  If this parameter isn't specified, the default value of ``discoveryName.namespace`` is used. If the ``discoveryName`` isn't specified, the port mapping name from the task definition is used in ``portName.namespace``.
	//  To avoid changing your applications in client Amazon ECS services, set this to the same name that the client application uses by default. For example, a few common names are ``database``, ``db``, or the lowercase name of a database, such as ``mysql`` or ``redis``. For more information, see [Service Connect](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-connect.html) in the *Amazon Elastic Container Service Developer Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_service#dns_name EcsService#dns_name}
	DnsName *string `field:"optional" json:"dnsName" yaml:"dnsName"`
	// The listening port number for the Service Connect proxy.
	//
	// This port is available inside of all of the tasks within the same namespace.
	//  To avoid changing your applications in client Amazon ECS services, set this to the same port that the client application uses by default. For more information, see [Service Connect](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-connect.html) in the *Amazon Elastic Container Service Developer Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_service#port EcsService#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The configuration for test traffic routing rules used during blue/green deployments with Amazon ECS Service Connect.
	//
	// This allows you to route a portion of traffic to the new service revision of your service for testing before shifting all production traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ecs_service#test_traffic_rules EcsService#test_traffic_rules}
	TestTrafficRules *EcsServiceServiceConnectConfigurationServicesClientAliasesTestTrafficRules `field:"optional" json:"testTrafficRules" yaml:"testTrafficRules"`
}

