// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package evsenvironment


type EvsEnvironmentVcfHostnames struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/evs_environment#cloud_builder EvsEnvironment#cloud_builder}.
	CloudBuilder *string `field:"optional" json:"cloudBuilder" yaml:"cloudBuilder"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/evs_environment#nsx EvsEnvironment#nsx}.
	Nsx *string `field:"optional" json:"nsx" yaml:"nsx"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/evs_environment#nsx_edge_1 EvsEnvironment#nsx_edge_1}.
	NsxEdge1 *string `field:"optional" json:"nsxEdge1" yaml:"nsxEdge1"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/evs_environment#nsx_edge_2 EvsEnvironment#nsx_edge_2}.
	NsxEdge2 *string `field:"optional" json:"nsxEdge2" yaml:"nsxEdge2"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/evs_environment#nsx_manager_1 EvsEnvironment#nsx_manager_1}.
	NsxManager1 *string `field:"optional" json:"nsxManager1" yaml:"nsxManager1"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/evs_environment#nsx_manager_2 EvsEnvironment#nsx_manager_2}.
	NsxManager2 *string `field:"optional" json:"nsxManager2" yaml:"nsxManager2"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/evs_environment#nsx_manager_3 EvsEnvironment#nsx_manager_3}.
	NsxManager3 *string `field:"optional" json:"nsxManager3" yaml:"nsxManager3"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/evs_environment#sddc_manager EvsEnvironment#sddc_manager}.
	SddcManager *string `field:"optional" json:"sddcManager" yaml:"sddcManager"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/evs_environment#v_center EvsEnvironment#v_center}.
	VCenter *string `field:"optional" json:"vCenter" yaml:"vCenter"`
}

