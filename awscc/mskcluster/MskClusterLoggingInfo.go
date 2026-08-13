// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskcluster


type MskClusterLoggingInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/msk_cluster#authorizer_logs MskCluster#authorizer_logs}.
	AuthorizerLogs *MskClusterLoggingInfoAuthorizerLogs `field:"optional" json:"authorizerLogs" yaml:"authorizerLogs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/msk_cluster#broker_logs MskCluster#broker_logs}.
	BrokerLogs *MskClusterLoggingInfoBrokerLogs `field:"optional" json:"brokerLogs" yaml:"brokerLogs"`
}

