// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceServiceConnectConfigurationAccessLogConfiguration struct {
	// The format for Service Connect access log output.
	//
	// Choose TEXT for human-readable logs or JSON for structured data that integrates well with log analysis tools.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ecs_service#format EcsService#format}
	Format *string `field:"optional" json:"format" yaml:"format"`
	// Specifies whether to include query parameters in Service Connect access logs.
	//
	// When enabled, query parameters from HTTP requests are included in the access logs. Consider security and privacy implications when enabling this feature, as query parameters may contain sensitive information such as request IDs and tokens. By default, this parameter is ``DISABLED``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ecs_service#include_query_parameters EcsService#include_query_parameters}
	IncludeQueryParameters *string `field:"optional" json:"includeQueryParameters" yaml:"includeQueryParameters"`
}

