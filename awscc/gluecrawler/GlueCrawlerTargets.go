// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluecrawler


type GlueCrawlerTargets struct {
	// Specifies AWS Glue Data Catalog targets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_crawler#catalog_targets GlueCrawler#catalog_targets}
	CatalogTargets interface{} `field:"optional" json:"catalogTargets" yaml:"catalogTargets"`
	// Specifies an array of Delta data store targets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_crawler#delta_targets GlueCrawler#delta_targets}
	DeltaTargets interface{} `field:"optional" json:"deltaTargets" yaml:"deltaTargets"`
	// Specifies Amazon DynamoDB targets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_crawler#dynamo_db_targets GlueCrawler#dynamo_db_targets}
	DynamoDbTargets interface{} `field:"optional" json:"dynamoDbTargets" yaml:"dynamoDbTargets"`
	// Specifies Apache Hudi data store targets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_crawler#hudi_targets GlueCrawler#hudi_targets}
	HudiTargets interface{} `field:"optional" json:"hudiTargets" yaml:"hudiTargets"`
	// Specifies Apache Iceberg data store targets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_crawler#iceberg_targets GlueCrawler#iceberg_targets}
	IcebergTargets interface{} `field:"optional" json:"icebergTargets" yaml:"icebergTargets"`
	// Specifies JDBC targets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_crawler#jdbc_targets GlueCrawler#jdbc_targets}
	JdbcTargets interface{} `field:"optional" json:"jdbcTargets" yaml:"jdbcTargets"`
	// A list of Mongo DB targets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_crawler#mongo_db_targets GlueCrawler#mongo_db_targets}
	MongoDbTargets interface{} `field:"optional" json:"mongoDbTargets" yaml:"mongoDbTargets"`
	// Specifies Amazon Simple Storage Service (Amazon S3) targets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_crawler#s3_targets GlueCrawler#s3_targets}
	S3Targets interface{} `field:"optional" json:"s3Targets" yaml:"s3Targets"`
}

