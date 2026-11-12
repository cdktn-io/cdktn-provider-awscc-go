// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package daxcluster


type DaxClusterSseSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/dax_cluster#sse_enabled DaxCluster#sse_enabled}.
	SseEnabled interface{} `field:"optional" json:"sseEnabled" yaml:"sseEnabled"`
}

