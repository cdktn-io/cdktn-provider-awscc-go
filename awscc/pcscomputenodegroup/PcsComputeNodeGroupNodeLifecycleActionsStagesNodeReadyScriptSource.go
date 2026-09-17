// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pcscomputenodegroup


type PcsComputeNodeGroupNodeLifecycleActionsStagesNodeReadyScriptSource struct {
	// A 64-character hexadecimal SHA-256 digest used to verify script integrity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/pcs_compute_node_group#checksum PcsComputeNodeGroup#checksum}
	Checksum *string `field:"optional" json:"checksum" yaml:"checksum"`
	// The S3 object version ID of the script, when stored in a versioned bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/pcs_compute_node_group#s3_version_id PcsComputeNodeGroup#s3_version_id}
	S3VersionId *string `field:"optional" json:"s3VersionId" yaml:"s3VersionId"`
	// The S3 URI or HTTPS URL where the script is stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/pcs_compute_node_group#script_location PcsComputeNodeGroup#script_location}
	ScriptLocation *string `field:"optional" json:"scriptLocation" yaml:"scriptLocation"`
}

