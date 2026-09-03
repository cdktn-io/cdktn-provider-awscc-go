// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluecrawler


type GlueCrawlerTargetsDeltaTargets struct {
	// The name of the connection to use to connect to the Delta table target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/glue_crawler#connection_name GlueCrawler#connection_name}
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// Specifies whether the crawler will create native tables, to allow integration with query engines that support querying of the Delta transaction log directly.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/glue_crawler#create_native_delta_table GlueCrawler#create_native_delta_table}
	CreateNativeDeltaTable interface{} `field:"optional" json:"createNativeDeltaTable" yaml:"createNativeDeltaTable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/glue_crawler#delta_tables GlueCrawler#delta_tables}.
	DeltaTables *[]*string `field:"optional" json:"deltaTables" yaml:"deltaTables"`
	// Specifies whether to write the manifest files to the Delta table path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/glue_crawler#write_manifest GlueCrawler#write_manifest}
	WriteManifest interface{} `field:"optional" json:"writeManifest" yaml:"writeManifest"`
}

