// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package auditmanagerassessment


type AuditmanagerAssessmentDelegations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/auditmanager_assessment#assessment_id AuditmanagerAssessment#assessment_id}.
	AssessmentId *string `field:"optional" json:"assessmentId" yaml:"assessmentId"`
	// The name of the related assessment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/auditmanager_assessment#assessment_name AuditmanagerAssessment#assessment_name}
	AssessmentName *string `field:"optional" json:"assessmentName" yaml:"assessmentName"`
	// The comment related to the delegation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/auditmanager_assessment#comment AuditmanagerAssessment#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The identifier for the specified control set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/auditmanager_assessment#control_set_id AuditmanagerAssessment#control_set_id}
	ControlSetId *string `field:"optional" json:"controlSetId" yaml:"controlSetId"`
	// The IAM user or role that performed the action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/auditmanager_assessment#created_by AuditmanagerAssessment#created_by}
	CreatedBy *string `field:"optional" json:"createdBy" yaml:"createdBy"`
	// The sequence of characters that identifies when the event occurred.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/auditmanager_assessment#creation_time AuditmanagerAssessment#creation_time}
	CreationTime *float64 `field:"optional" json:"creationTime" yaml:"creationTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/auditmanager_assessment#id AuditmanagerAssessment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The sequence of characters that identifies when the event occurred.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/auditmanager_assessment#last_updated AuditmanagerAssessment#last_updated}
	LastUpdated *float64 `field:"optional" json:"lastUpdated" yaml:"lastUpdated"`
	// The Amazon Resource Name (ARN) of the IAM user or role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/auditmanager_assessment#role_arn AuditmanagerAssessment#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The IAM role type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/auditmanager_assessment#role_type AuditmanagerAssessment#role_type}
	RoleType *string `field:"optional" json:"roleType" yaml:"roleType"`
	// The status of the delegation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/auditmanager_assessment#status AuditmanagerAssessment#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

