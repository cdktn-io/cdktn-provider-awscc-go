// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluecrawler


type GlueCrawlerTargetsIcebergTargets struct {
	// The name of the connection to use to connect to the Iceberg target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/glue_crawler#connection_name GlueCrawler#connection_name}
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// A list of global patterns used to exclude from the crawl.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/glue_crawler#exclusions GlueCrawler#exclusions}
	Exclusions *[]*string `field:"optional" json:"exclusions" yaml:"exclusions"`
	// The maximum depth of Amazon S3 paths that the crawler can traverse to discover the Iceberg metadata folder in your Amazon S3 path.
	//
	// Used to limit the crawler run time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/glue_crawler#maximum_traversal_depth GlueCrawler#maximum_traversal_depth}
	MaximumTraversalDepth *float64 `field:"optional" json:"maximumTraversalDepth" yaml:"maximumTraversalDepth"`
	// One or more Amazon S3 paths that contains Iceberg metadata folders as s3://bucket/prefix .
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/glue_crawler#paths GlueCrawler#paths}
	Paths *[]*string `field:"optional" json:"paths" yaml:"paths"`
}

