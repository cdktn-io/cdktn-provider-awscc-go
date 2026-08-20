// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package redshiftserverlessworkgroup


type RedshiftserverlessWorkgroupWorkgroup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/redshiftserverless_workgroup#endpoint RedshiftserverlessWorkgroup#endpoint}.
	Endpoint *RedshiftserverlessWorkgroupWorkgroupEndpoint `field:"optional" json:"endpoint" yaml:"endpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/redshiftserverless_workgroup#price_performance_target RedshiftserverlessWorkgroup#price_performance_target}.
	PricePerformanceTarget *RedshiftserverlessWorkgroupWorkgroupPricePerformanceTarget `field:"optional" json:"pricePerformanceTarget" yaml:"pricePerformanceTarget"`
}

