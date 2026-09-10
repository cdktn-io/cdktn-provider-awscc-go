// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package athenaworkgroup


type AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationClassifications struct {
	// The name of the configuration classification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/athena_work_group#name AthenaWorkGroup#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A set of properties specified within a configuration classification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/athena_work_group#properties AthenaWorkGroup#properties}
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
}

