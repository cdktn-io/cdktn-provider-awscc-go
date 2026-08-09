// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceServiceConnectConfigurationServicesClientAliasesTestTrafficRules struct {
	// The HTTP header-based routing rules that determine which requests should be routed to the new service version during blue/green deployment testing.
	//
	// These rules provide fine-grained control over test traffic routing based on request headers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ecs_service#header EcsService#header}
	Header *EcsServiceServiceConnectConfigurationServicesClientAliasesTestTrafficRulesHeader `field:"optional" json:"header" yaml:"header"`
}

