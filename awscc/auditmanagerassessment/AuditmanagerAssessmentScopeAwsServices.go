// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package auditmanagerassessment


type AuditmanagerAssessmentScopeAwsServices struct {
	// The name of the AWS service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/auditmanager_assessment#service_name AuditmanagerAssessment#service_name}
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
}

