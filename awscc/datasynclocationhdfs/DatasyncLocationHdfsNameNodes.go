// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasynclocationhdfs


type DatasyncLocationHdfsNameNodes struct {
	// The DNS name or IP address of the Name Node in the customer's on premises HDFS cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/datasync_location_hdfs#hostname DatasyncLocationHdfs#hostname}
	Hostname *string `field:"required" json:"hostname" yaml:"hostname"`
	// The port on which the Name Node is listening on for client requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/datasync_location_hdfs#port DatasyncLocationHdfs#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

