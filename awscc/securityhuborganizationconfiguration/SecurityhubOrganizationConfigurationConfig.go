// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhuborganizationconfiguration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecurityhubOrganizationConfigurationConfig struct {
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
	// Whether to automatically enable Security Hub in new member accounts when they join the organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/securityhub_organization_configuration#auto_enable SecurityhubOrganizationConfiguration#auto_enable}
	AutoEnable interface{} `field:"required" json:"autoEnable" yaml:"autoEnable"`
	// Whether to automatically enable Security Hub default standards in new member accounts when they join the organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/securityhub_organization_configuration#auto_enable_standards SecurityhubOrganizationConfiguration#auto_enable_standards}
	AutoEnableStandards *string `field:"optional" json:"autoEnableStandards" yaml:"autoEnableStandards"`
	// Indicates whether the organization uses local or central configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/securityhub_organization_configuration#configuration_type SecurityhubOrganizationConfiguration#configuration_type}
	ConfigurationType *string `field:"optional" json:"configurationType" yaml:"configurationType"`
}

