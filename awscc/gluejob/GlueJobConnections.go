// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluejob


type GlueJobConnections struct {
	// A list of connections used by the job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/glue_job#connections GlueJob#connections}
	Connections *[]*string `field:"optional" json:"connections" yaml:"connections"`
}

