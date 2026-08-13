// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package redshiftserverlessworkgroup


type RedshiftserverlessWorkgroupPricePerformanceTarget struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/redshiftserverless_workgroup#level RedshiftserverlessWorkgroup#level}.
	Level *float64 `field:"optional" json:"level" yaml:"level"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/redshiftserverless_workgroup#status RedshiftserverlessWorkgroup#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

