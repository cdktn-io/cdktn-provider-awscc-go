// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluecrawler


type GlueCrawlerTargetsMongoDbTargets struct {
	// The name of the connection to use to connect to the Amazon DocumentDB or MongoDB target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/glue_crawler#connection_name GlueCrawler#connection_name}
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// The path of the Amazon DocumentDB or MongoDB target (database/collection).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/glue_crawler#path GlueCrawler#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

