// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package grafanaworkspace


type GrafanaWorkspaceTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/grafana_workspace#key GrafanaWorkspace#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/grafana_workspace#value GrafanaWorkspace#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

