// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectuserhierarchystructure


type ConnectUserHierarchyStructureUserHierarchyStructure struct {
	// Information about level five.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/connect_user_hierarchy_structure#level_five ConnectUserHierarchyStructure#level_five}
	LevelFive *ConnectUserHierarchyStructureUserHierarchyStructureLevelFive `field:"optional" json:"levelFive" yaml:"levelFive"`
	// Information about level four.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/connect_user_hierarchy_structure#level_four ConnectUserHierarchyStructure#level_four}
	LevelFour *ConnectUserHierarchyStructureUserHierarchyStructureLevelFour `field:"optional" json:"levelFour" yaml:"levelFour"`
	// Information about level one.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/connect_user_hierarchy_structure#level_one ConnectUserHierarchyStructure#level_one}
	LevelOne *ConnectUserHierarchyStructureUserHierarchyStructureLevelOne `field:"optional" json:"levelOne" yaml:"levelOne"`
	// Information about level three.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/connect_user_hierarchy_structure#level_three ConnectUserHierarchyStructure#level_three}
	LevelThree *ConnectUserHierarchyStructureUserHierarchyStructureLevelThree `field:"optional" json:"levelThree" yaml:"levelThree"`
	// Information about level two.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/connect_user_hierarchy_structure#level_two ConnectUserHierarchyStructure#level_two}
	LevelTwo *ConnectUserHierarchyStructureUserHierarchyStructureLevelTwo `field:"optional" json:"levelTwo" yaml:"levelTwo"`
}

