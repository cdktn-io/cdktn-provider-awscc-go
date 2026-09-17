// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wafv2loggingconfiguration


type Wafv2LoggingConfigurationRedactedFieldsSingleHeader struct {
	// The name of the query header to inspect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wafv2_logging_configuration#name Wafv2LoggingConfiguration#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

