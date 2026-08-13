// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ivschatroom


type IvschatRoomMessageReviewHandler struct {
	// Specifies the fallback behavior if the handler does not return a valid response, encounters an error, or times out.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ivschat_room#fallback_result IvschatRoom#fallback_result}
	FallbackResult *string `field:"optional" json:"fallbackResult" yaml:"fallbackResult"`
	// Identifier of the message review handler.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ivschat_room#uri IvschatRoom#uri}
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

