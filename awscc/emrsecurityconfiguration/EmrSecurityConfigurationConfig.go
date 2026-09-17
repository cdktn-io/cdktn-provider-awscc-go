// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package emrsecurityconfiguration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EmrSecurityConfigurationConfig struct {
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
	// The security configuration details in JSON format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/emr_security_configuration#security_configuration EmrSecurityConfiguration#security_configuration}
	SecurityConfiguration *string `field:"required" json:"securityConfiguration" yaml:"securityConfiguration"`
	// The name of the security configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/emr_security_configuration#name EmrSecurityConfiguration#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

