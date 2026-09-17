// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3expressaccesspoint


type S3ExpressAccessPointScope struct {
	// You can include one or more API operations as permissions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3express_access_point#permissions S3ExpressAccessPoint#permissions}
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
	// You can specify any amount of pre?xes, but the total length of characters of all pre?xes must be less than 256 bytes in size.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/s3express_access_point#prefixes S3ExpressAccessPoint#prefixes}
	Prefixes *[]*string `field:"optional" json:"prefixes" yaml:"prefixes"`
}

