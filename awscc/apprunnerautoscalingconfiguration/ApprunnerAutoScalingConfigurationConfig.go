// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apprunnerautoscalingconfiguration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ApprunnerAutoScalingConfigurationConfig struct {
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
	// The customer-provided auto scaling configuration name.
	//
	// When you use it for the first time in an AWS Region, App Runner creates revision number 1 of this name. When you use the same name in subsequent calls, App Runner creates incremental revisions of the configuration. The auto scaling configuration name can be used in multiple revisions of a configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apprunner_auto_scaling_configuration#auto_scaling_configuration_name ApprunnerAutoScalingConfiguration#auto_scaling_configuration_name}
	AutoScalingConfigurationName *string `field:"optional" json:"autoScalingConfigurationName" yaml:"autoScalingConfigurationName"`
	// The maximum number of concurrent requests that an instance processes.
	//
	// If the number of concurrent requests exceeds this limit, App Runner scales the service up to use more instances to process the requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apprunner_auto_scaling_configuration#max_concurrency ApprunnerAutoScalingConfiguration#max_concurrency}
	MaxConcurrency *float64 `field:"optional" json:"maxConcurrency" yaml:"maxConcurrency"`
	// The maximum number of instances that an App Runner service scales up to.
	//
	// At most MaxSize instances actively serve traffic for your service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apprunner_auto_scaling_configuration#max_size ApprunnerAutoScalingConfiguration#max_size}
	MaxSize *float64 `field:"optional" json:"maxSize" yaml:"maxSize"`
	// The minimum number of instances that App Runner provisions for a service.
	//
	// The service always has at least MinSize provisioned instances. Some of them actively serve traffic. The rest of them (provisioned and inactive instances) are a cost-effective compute capacity reserve and are ready to be quickly activated. You pay for memory usage of all the provisioned instances. You pay for CPU usage of only the active subset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apprunner_auto_scaling_configuration#min_size ApprunnerAutoScalingConfiguration#min_size}
	MinSize *float64 `field:"optional" json:"minSize" yaml:"minSize"`
	// A list of metadata items that you can associate with your auto scaling configuration resource.
	//
	// A tag is a key-value pair.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apprunner_auto_scaling_configuration#tags ApprunnerAutoScalingConfiguration#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

