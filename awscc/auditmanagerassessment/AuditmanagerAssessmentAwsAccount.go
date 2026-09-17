// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package auditmanagerassessment


type AuditmanagerAssessmentAwsAccount struct {
	// The unique identifier for the email account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/auditmanager_assessment#email_address AuditmanagerAssessment#email_address}
	EmailAddress *string `field:"optional" json:"emailAddress" yaml:"emailAddress"`
	// The identifier for the specified AWS account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/auditmanager_assessment#id AuditmanagerAssessment#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of the specified AWS account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/auditmanager_assessment#name AuditmanagerAssessment#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

