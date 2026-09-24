// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazoneconnection


type DatazoneConnectionPropsLakehouseProperties struct {
	// Specifies whether Glue lineage sync is enabled for the lakehouse connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datazone_connection#glue_lineage_sync_enabled DatazoneConnection#glue_lineage_sync_enabled}
	GlueLineageSyncEnabled interface{} `field:"optional" json:"glueLineageSyncEnabled" yaml:"glueLineageSyncEnabled"`
}

