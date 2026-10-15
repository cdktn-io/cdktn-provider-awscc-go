// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package timestreaminfluxdbbackup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type TimestreamInfluxDbBackupConfig struct {
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
	// The id of the DB instance or DB cluster to back up.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/timestream_influx_db_backup#db_resource_id TimestreamInfluxDbBackup#db_resource_id}
	DbResourceId *string `field:"required" json:"dbResourceId" yaml:"dbResourceId"`
	// The name of the backup. Must be unique within the account and region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/timestream_influx_db_backup#name TimestreamInfluxDbBackup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The number of days to retain the backup. Valid values are 1 to 3650.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/timestream_influx_db_backup#retention_days TimestreamInfluxDbBackup#retention_days}
	RetentionDays *float64 `field:"optional" json:"retentionDays" yaml:"retentionDays"`
	// A list of key-value pairs to associate with the backup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/timestream_influx_db_backup#tags TimestreamInfluxDbBackup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

