// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apprunnerobservabilityconfiguration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ApprunnerObservabilityConfigurationConfig struct {
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
	// A name for the observability configuration.
	//
	// When you use it for the first time in an AWS Region, App Runner creates revision number 1 of this name. When you use the same name in subsequent calls, App Runner creates incremental revisions of the configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/apprunner_observability_configuration#observability_configuration_name ApprunnerObservabilityConfiguration#observability_configuration_name}
	ObservabilityConfigurationName *string `field:"optional" json:"observabilityConfigurationName" yaml:"observabilityConfigurationName"`
	// A list of metadata items that you can associate with your observability configuration resource.
	//
	// A tag is a key-value pair.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/apprunner_observability_configuration#tags ApprunnerObservabilityConfiguration#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The configuration of the tracing feature within this observability configuration.
	//
	// If you don't specify it, App Runner doesn't enable tracing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/apprunner_observability_configuration#trace_configuration ApprunnerObservabilityConfiguration#trace_configuration}
	TraceConfiguration *ApprunnerObservabilityConfigurationTraceConfiguration `field:"optional" json:"traceConfiguration" yaml:"traceConfiguration"`
}

