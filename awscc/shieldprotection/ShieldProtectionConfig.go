// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package shieldprotection

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ShieldProtectionConfig struct {
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
	// Friendly name for the Protection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/shield_protection#name ShieldProtection#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ARN (Amazon Resource Name) of the resource to be protected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/shield_protection#resource_arn ShieldProtection#resource_arn}
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// The automatic application layer DDoS mitigation settings for a Protection.
	//
	// This configuration determines whether Shield Advanced automatically manages rules in the web ACL in order to respond to application layer events that Shield Advanced determines to be DDoS attacks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/shield_protection#application_layer_automatic_response_configuration ShieldProtection#application_layer_automatic_response_configuration}
	ApplicationLayerAutomaticResponseConfiguration *ShieldProtectionApplicationLayerAutomaticResponseConfiguration `field:"optional" json:"applicationLayerAutomaticResponseConfiguration" yaml:"applicationLayerAutomaticResponseConfiguration"`
	// The Amazon Resource Names (ARNs) of the health check to associate with the protection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/shield_protection#health_check_arns ShieldProtection#health_check_arns}
	HealthCheckArns *[]*string `field:"optional" json:"healthCheckArns" yaml:"healthCheckArns"`
	// One or more tag key-value pairs for the Protection object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/shield_protection#tags ShieldProtection#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

