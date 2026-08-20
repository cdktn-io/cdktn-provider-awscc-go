// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package auditmanagerassessmentframework

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AuditmanagerAssessmentFrameworkConfig struct {
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
	// The control sets that are associated with the framework.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/auditmanager_assessment_framework#control_sets AuditmanagerAssessmentFramework#control_sets}
	ControlSets interface{} `field:"required" json:"controlSets" yaml:"controlSets"`
	// The name of the framework.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/auditmanager_assessment_framework#name AuditmanagerAssessmentFramework#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The compliance type that the framework supports, such as CIS or HIPAA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/auditmanager_assessment_framework#compliance_type AuditmanagerAssessmentFramework#compliance_type}
	ComplianceType *string `field:"optional" json:"complianceType" yaml:"complianceType"`
	// The description of the framework.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/auditmanager_assessment_framework#description AuditmanagerAssessmentFramework#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The tags associated with the framework.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/auditmanager_assessment_framework#tags AuditmanagerAssessmentFramework#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

