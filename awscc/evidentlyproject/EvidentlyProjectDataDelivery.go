// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package evidentlyproject


type EvidentlyProjectDataDelivery struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/evidently_project#log_group EvidentlyProject#log_group}.
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/evidently_project#s3 EvidentlyProject#s3}.
	S3 *EvidentlyProjectDataDeliveryS3 `field:"optional" json:"s3" yaml:"s3"`
}

