// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package omicsconfiguration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OmicsConfigurationConfig struct {
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
	// User-friendly name for the configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/omics_configuration#name OmicsConfiguration#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Required run-specific configurations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/omics_configuration#run_configurations OmicsConfiguration#run_configurations}
	RunConfigurations *OmicsConfigurationRunConfigurations `field:"required" json:"runConfigurations" yaml:"runConfigurations"`
	// Optional description for the configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/omics_configuration#description OmicsConfiguration#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Tags for the configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/omics_configuration#tags OmicsConfiguration#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

