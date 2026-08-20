// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueintegration


type GlueIntegrationIntegrationConfig struct {
	// Enables continuous synchronization for on-demand data extractions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/glue_integration#continuous_sync GlueIntegration#continuous_sync}
	ContinuousSync interface{} `field:"optional" json:"continuousSync" yaml:"continuousSync"`
	// Specifies the frequency at which CDC (Change Data Capture) pulls or incremental loads should occur.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/glue_integration#refresh_interval GlueIntegration#refresh_interval}
	RefreshInterval *string `field:"optional" json:"refreshInterval" yaml:"refreshInterval"`
	// A collection of key-value pairs that specify additional properties for the integration source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/glue_integration#source_properties GlueIntegration#source_properties}
	SourceProperties *map[string]*string `field:"optional" json:"sourceProperties" yaml:"sourceProperties"`
}

