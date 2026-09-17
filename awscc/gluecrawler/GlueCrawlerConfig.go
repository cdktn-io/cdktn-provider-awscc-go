// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluecrawler

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueCrawlerConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The Amazon Resource Name (ARN) of an IAM role that's used to access customer resources, such as Amazon Simple Storage Service (Amazon S3) data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_crawler#role GlueCrawler#role}
	Role *string `field:"required" json:"role" yaml:"role"`
	// Specifies data stores to crawl.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_crawler#targets GlueCrawler#targets}
	Targets *GlueCrawlerTargets `field:"required" json:"targets" yaml:"targets"`
	// A list of UTF-8 strings that specify the names of custom classifiers that are associated with the crawler.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_crawler#classifiers GlueCrawler#classifiers}
	Classifiers *[]*string `field:"optional" json:"classifiers" yaml:"classifiers"`
	// Crawler configuration information. This versioned JSON string allows users to specify aspects of a crawler's behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_crawler#configuration GlueCrawler#configuration}
	Configuration *string `field:"optional" json:"configuration" yaml:"configuration"`
	// The name of the SecurityConfiguration structure to be used by this crawler.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_crawler#crawler_security_configuration GlueCrawler#crawler_security_configuration}
	CrawlerSecurityConfiguration *string `field:"optional" json:"crawlerSecurityConfiguration" yaml:"crawlerSecurityConfiguration"`
	// The name of the database in which the crawler's output is stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_crawler#database_name GlueCrawler#database_name}
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// A description of the crawler.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_crawler#description GlueCrawler#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies AWS Lake Formation configuration settings for the crawler.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_crawler#lake_formation_configuration GlueCrawler#lake_formation_configuration}
	LakeFormationConfiguration *GlueCrawlerLakeFormationConfiguration `field:"optional" json:"lakeFormationConfiguration" yaml:"lakeFormationConfiguration"`
	// The name of the crawler.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_crawler#name GlueCrawler#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// When crawling an Amazon S3 data source after the first crawl is complete, specifies whether to crawl the entire dataset again or to crawl only folders that were added since the last crawler run.
	//
	// For more information, see Incremental Crawls in AWS Glue in the developer guide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_crawler#recrawl_policy GlueCrawler#recrawl_policy}
	RecrawlPolicy *GlueCrawlerRecrawlPolicy `field:"optional" json:"recrawlPolicy" yaml:"recrawlPolicy"`
	// A scheduling object using a cron statement to schedule an event.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_crawler#schedule GlueCrawler#schedule}
	Schedule *GlueCrawlerSchedule `field:"optional" json:"schedule" yaml:"schedule"`
	// The policy that specifies update and delete behaviors for the crawler.
	//
	// The policy tells the crawler what to do in the event that it detects a change in a table that already exists in the customer's database at the time of the crawl. The SchemaChangePolicy does not affect whether or how new tables and partitions are added. New tables and partitions are always created regardless of the SchemaChangePolicy on a crawler. The SchemaChangePolicy consists of two components, UpdateBehavior and DeleteBehavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_crawler#schema_change_policy GlueCrawler#schema_change_policy}
	SchemaChangePolicy *GlueCrawlerSchemaChangePolicy `field:"optional" json:"schemaChangePolicy" yaml:"schemaChangePolicy"`
	// The prefix added to the names of tables that are created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_crawler#table_prefix GlueCrawler#table_prefix}
	TablePrefix *string `field:"optional" json:"tablePrefix" yaml:"tablePrefix"`
	// The tags to use with this crawler.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_crawler#tags GlueCrawler#tags}
	Tags *string `field:"optional" json:"tags" yaml:"tags"`
}

