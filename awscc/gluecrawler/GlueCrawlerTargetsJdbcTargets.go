// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluecrawler


type GlueCrawlerTargetsJdbcTargets struct {
	// The name of the connection to use to connect to the JDBC target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_crawler#connection_name GlueCrawler#connection_name}
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// Specify a value of RAWTYPES or COMMENTS to enable additional metadata in table responses.
	//
	// RAWTYPES provides the native-level datatype. COMMENTS provides comments associated with a column or table in the database.
	//
	// If you do not need additional metadata, keep the field empty.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_crawler#enable_additional_metadata GlueCrawler#enable_additional_metadata}
	EnableAdditionalMetadata *[]*string `field:"optional" json:"enableAdditionalMetadata" yaml:"enableAdditionalMetadata"`
	// A list of glob patterns used to exclude from the crawl.
	//
	// For more information, see Catalog Tables with a Crawler.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_crawler#exclusions GlueCrawler#exclusions}
	Exclusions *[]*string `field:"optional" json:"exclusions" yaml:"exclusions"`
	// The path of the JDBC target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_crawler#path GlueCrawler#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

