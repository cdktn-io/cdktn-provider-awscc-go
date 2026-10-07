// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazoneconnection


type DatazoneConnectionPropsIamProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/datazone_connection#glue_lineage_sync_enabled DatazoneConnection#glue_lineage_sync_enabled}.
	GlueLineageSyncEnabled interface{} `field:"optional" json:"glueLineageSyncEnabled" yaml:"glueLineageSyncEnabled"`
	// The ARN of the IAM role to associate with the connection as the project user role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/datazone_connection#role_arn DatazoneConnection#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

