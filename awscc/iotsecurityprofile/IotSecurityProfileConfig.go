// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsecurityprofile

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotSecurityProfileConfig struct {
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
	// A list of metrics whose data is retained (stored).
	//
	// By default, data is retained for any metric used in the profile's behaviors, but it is also retained for any metric specified here.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_security_profile#additional_metrics_to_retain_v2 IotSecurityProfile#additional_metrics_to_retain_v2}
	AdditionalMetricsToRetainV2 interface{} `field:"optional" json:"additionalMetricsToRetainV2" yaml:"additionalMetricsToRetainV2"`
	// Specifies the destinations to which alerts are sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_security_profile#alert_targets IotSecurityProfile#alert_targets}
	AlertTargets interface{} `field:"optional" json:"alertTargets" yaml:"alertTargets"`
	// Specifies the behaviors that, when violated by a device (thing), cause an alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_security_profile#behaviors IotSecurityProfile#behaviors}
	Behaviors interface{} `field:"optional" json:"behaviors" yaml:"behaviors"`
	// A structure containing the mqtt topic for metrics export.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_security_profile#metrics_export_config IotSecurityProfile#metrics_export_config}
	MetricsExportConfig *IotSecurityProfileMetricsExportConfig `field:"optional" json:"metricsExportConfig" yaml:"metricsExportConfig"`
	// A description of the security profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_security_profile#security_profile_description IotSecurityProfile#security_profile_description}
	SecurityProfileDescription *string `field:"optional" json:"securityProfileDescription" yaml:"securityProfileDescription"`
	// A unique identifier for the security profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_security_profile#security_profile_name IotSecurityProfile#security_profile_name}
	SecurityProfileName *string `field:"optional" json:"securityProfileName" yaml:"securityProfileName"`
	// Metadata that can be used to manage the security profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_security_profile#tags IotSecurityProfile#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// A set of target ARNs that the security profile is attached to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_security_profile#target_arns IotSecurityProfile#target_arns}
	TargetArns *[]*string `field:"optional" json:"targetArns" yaml:"targetArns"`
}

