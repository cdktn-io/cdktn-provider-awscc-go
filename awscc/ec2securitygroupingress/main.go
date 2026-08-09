// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2securitygroupingress

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2SecurityGroupIngress.Ec2SecurityGroupIngress",
		reflect.TypeOf((*Ec2SecurityGroupIngress)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "cidrIp", GoGetter: "CidrIp"},
			_jsii_.MemberProperty{JsiiProperty: "cidrIpInput", GoGetter: "CidrIpInput"},
			_jsii_.MemberProperty{JsiiProperty: "cidrIpv6", GoGetter: "CidrIpv6"},
			_jsii_.MemberProperty{JsiiProperty: "cidrIpv6Input", GoGetter: "CidrIpv6Input"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "fromPort", GoGetter: "FromPort"},
			_jsii_.MemberProperty{JsiiProperty: "fromPortInput", GoGetter: "FromPortInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "groupId", GoGetter: "GroupId"},
			_jsii_.MemberProperty{JsiiProperty: "groupIdInput", GoGetter: "GroupIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "groupName", GoGetter: "GroupName"},
			_jsii_.MemberProperty{JsiiProperty: "groupNameInput", GoGetter: "GroupNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipProtocol", GoGetter: "IpProtocol"},
			_jsii_.MemberProperty{JsiiProperty: "ipProtocolInput", GoGetter: "IpProtocolInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "markWriteOnlyAttribute", GoMethod: "MarkWriteOnlyAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "registerProviderFeatureUsage", GoMethod: "RegisterProviderFeatureUsage"},
			_jsii_.MemberMethod{JsiiMethod: "resetCidrIp", GoMethod: "ResetCidrIp"},
			_jsii_.MemberMethod{JsiiMethod: "resetCidrIpv6", GoMethod: "ResetCidrIpv6"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetFromPort", GoMethod: "ResetFromPort"},
			_jsii_.MemberMethod{JsiiMethod: "resetGroupId", GoMethod: "ResetGroupId"},
			_jsii_.MemberMethod{JsiiMethod: "resetGroupName", GoMethod: "ResetGroupName"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSourcePrefixListId", GoMethod: "ResetSourcePrefixListId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSourceSecurityGroupId", GoMethod: "ResetSourceSecurityGroupId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSourceSecurityGroupName", GoMethod: "ResetSourceSecurityGroupName"},
			_jsii_.MemberMethod{JsiiMethod: "resetSourceSecurityGroupOwnerId", GoMethod: "ResetSourceSecurityGroupOwnerId"},
			_jsii_.MemberMethod{JsiiMethod: "resetToPort", GoMethod: "ResetToPort"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupIngressId", GoGetter: "SecurityGroupIngressId"},
			_jsii_.MemberProperty{JsiiProperty: "sourcePrefixListId", GoGetter: "SourcePrefixListId"},
			_jsii_.MemberProperty{JsiiProperty: "sourcePrefixListIdInput", GoGetter: "SourcePrefixListIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "sourceSecurityGroupId", GoGetter: "SourceSecurityGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "sourceSecurityGroupIdInput", GoGetter: "SourceSecurityGroupIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "sourceSecurityGroupName", GoGetter: "SourceSecurityGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "sourceSecurityGroupNameInput", GoGetter: "SourceSecurityGroupNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "sourceSecurityGroupOwnerId", GoGetter: "SourceSecurityGroupOwnerId"},
			_jsii_.MemberProperty{JsiiProperty: "sourceSecurityGroupOwnerIdInput", GoGetter: "SourceSecurityGroupOwnerIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "toPort", GoGetter: "ToPort"},
			_jsii_.MemberProperty{JsiiProperty: "toPortInput", GoGetter: "ToPortInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2SecurityGroupIngress{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2SecurityGroupIngress.Ec2SecurityGroupIngressConfig",
		reflect.TypeOf((*Ec2SecurityGroupIngressConfig)(nil)).Elem(),
	)
}
