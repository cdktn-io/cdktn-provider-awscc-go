// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package backupframework


type BackupFrameworkFrameworkControlsControlScope struct {
	// The ID of the only AWS resource that you want your control scope to contain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/backup_framework#compliance_resource_ids BackupFramework#compliance_resource_ids}
	ComplianceResourceIds *[]*string `field:"optional" json:"complianceResourceIds" yaml:"complianceResourceIds"`
	// Describes whether the control scope includes one or more types of resources, such as `EFS` or `RDS`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/backup_framework#compliance_resource_types BackupFramework#compliance_resource_types}
	ComplianceResourceTypes *[]*string `field:"optional" json:"complianceResourceTypes" yaml:"complianceResourceTypes"`
	// Describes whether the control scope includes resources with one or more tags. Each tag is a key-value pair.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/backup_framework#tags BackupFramework#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

