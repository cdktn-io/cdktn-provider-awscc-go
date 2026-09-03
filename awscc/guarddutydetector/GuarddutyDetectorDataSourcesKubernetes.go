// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package guarddutydetector


type GuarddutyDetectorDataSourcesKubernetes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/guardduty_detector#audit_logs GuarddutyDetector#audit_logs}.
	AuditLogs *GuarddutyDetectorDataSourcesKubernetesAuditLogs `field:"optional" json:"auditLogs" yaml:"auditLogs"`
}

