// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package auditmanagerassessment


type AuditmanagerAssessmentRoles struct {
	// The Amazon Resource Name (ARN) of the IAM user or role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/auditmanager_assessment#role_arn AuditmanagerAssessment#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The IAM role type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/auditmanager_assessment#role_type AuditmanagerAssessment#role_type}
	RoleType *string `field:"optional" json:"roleType" yaml:"roleType"`
}

