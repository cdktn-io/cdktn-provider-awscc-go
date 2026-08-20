// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package auditmanagerassessment


type AuditmanagerAssessmentAssessmentReportsDestination struct {
	// The URL of the specified Amazon S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/auditmanager_assessment#destination AuditmanagerAssessment#destination}
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
	// The destination type, such as Amazon S3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/auditmanager_assessment#destination_type AuditmanagerAssessment#destination_type}
	DestinationType *string `field:"optional" json:"destinationType" yaml:"destinationType"`
}

