// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resiliencehubv2service


type Resiliencehubv2ServiceInputSources struct {
	// Resource configuration for an input source. Provide exactly one field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/resiliencehubv2_service#resource_configuration Resiliencehubv2Service#resource_configuration}
	ResourceConfiguration *Resiliencehubv2ServiceInputSourcesResourceConfiguration `field:"optional" json:"resourceConfiguration" yaml:"resourceConfiguration"`
}

