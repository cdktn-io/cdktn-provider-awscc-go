// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationRunConfigurationFlinkRunConfiguration struct {
	// When restoring from a snapshot, specifies whether the runtime is allowed to skip a state that cannot be mapped to the new program.
	//
	// Defaults to false. If you update your application without specifying this parameter, AllowNonRestoredState will be set to false, even if it was previously set to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kinesisanalyticsv2_application#allow_non_restored_state Kinesisanalyticsv2Application#allow_non_restored_state}
	AllowNonRestoredState interface{} `field:"optional" json:"allowNonRestoredState" yaml:"allowNonRestoredState"`
}

