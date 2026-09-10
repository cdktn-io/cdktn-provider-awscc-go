// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluesession


type GlueSessionConnections struct {
	// A list of connection names used by the session.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/glue_session#connections GlueSession#connections}
	Connections *[]*string `field:"optional" json:"connections" yaml:"connections"`
}

