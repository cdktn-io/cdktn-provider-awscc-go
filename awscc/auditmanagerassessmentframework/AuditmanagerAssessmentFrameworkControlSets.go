// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package auditmanagerassessmentframework


type AuditmanagerAssessmentFrameworkControlSets struct {
	// The list of controls within the control set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/auditmanager_assessment_framework#controls AuditmanagerAssessmentFramework#controls}
	Controls interface{} `field:"required" json:"controls" yaml:"controls"`
	// The name of the control set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/auditmanager_assessment_framework#name AuditmanagerAssessmentFramework#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

