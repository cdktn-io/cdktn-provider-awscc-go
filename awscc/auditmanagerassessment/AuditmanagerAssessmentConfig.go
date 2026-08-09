// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package auditmanagerassessment

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AuditmanagerAssessmentConfig struct {
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
	// The destination in which evidence reports are stored for the specified assessment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/auditmanager_assessment#assessment_reports_destination AuditmanagerAssessment#assessment_reports_destination}
	AssessmentReportsDestination *AuditmanagerAssessmentAssessmentReportsDestination `field:"optional" json:"assessmentReportsDestination" yaml:"assessmentReportsDestination"`
	// The AWS account associated with the assessment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/auditmanager_assessment#aws_account AuditmanagerAssessment#aws_account}
	AwsAccount *AuditmanagerAssessmentAwsAccount `field:"optional" json:"awsAccount" yaml:"awsAccount"`
	// The list of delegations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/auditmanager_assessment#delegations AuditmanagerAssessment#delegations}
	Delegations interface{} `field:"optional" json:"delegations" yaml:"delegations"`
	// The description of the specified assessment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/auditmanager_assessment#description AuditmanagerAssessment#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The identifier for the specified framework.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/auditmanager_assessment#framework_id AuditmanagerAssessment#framework_id}
	FrameworkId *string `field:"optional" json:"frameworkId" yaml:"frameworkId"`
	// The name of the related assessment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/auditmanager_assessment#name AuditmanagerAssessment#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The list of roles for the specified assessment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/auditmanager_assessment#roles AuditmanagerAssessment#roles}
	Roles interface{} `field:"optional" json:"roles" yaml:"roles"`
	// The wrapper that contains the AWS accounts and AWS services in scope for the assessment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/auditmanager_assessment#scope AuditmanagerAssessment#scope}
	Scope *AuditmanagerAssessmentScope `field:"optional" json:"scope" yaml:"scope"`
	// The status of the specified assessment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/auditmanager_assessment#status AuditmanagerAssessment#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The tags associated with the assessment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/auditmanager_assessment#tags AuditmanagerAssessment#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

