// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2vpccidrblock

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2VpcCidrBlock.Ec2VpcCidrBlock",
		reflect.TypeOf((*Ec2VpcCidrBlock)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "amazonProvidedIpv6CidrBlock", GoGetter: "AmazonProvidedIpv6CidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "amazonProvidedIpv6CidrBlockInput", GoGetter: "AmazonProvidedIpv6CidrBlockInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "cidrBlock", GoGetter: "CidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "cidrBlockInput", GoGetter: "CidrBlockInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipSource", GoGetter: "IpSource"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4IpamPoolId", GoGetter: "Ipv4IpamPoolId"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4IpamPoolIdInput", GoGetter: "Ipv4IpamPoolIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4NetmaskLength", GoGetter: "Ipv4NetmaskLength"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4NetmaskLengthInput", GoGetter: "Ipv4NetmaskLengthInput"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6AddressAttribute", GoGetter: "Ipv6AddressAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6CidrBlock", GoGetter: "Ipv6CidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6CidrBlockInput", GoGetter: "Ipv6CidrBlockInput"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6CidrBlockNetworkBorderGroup", GoGetter: "Ipv6CidrBlockNetworkBorderGroup"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6CidrBlockNetworkBorderGroupInput", GoGetter: "Ipv6CidrBlockNetworkBorderGroupInput"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6IpamPoolId", GoGetter: "Ipv6IpamPoolId"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6IpamPoolIdInput", GoGetter: "Ipv6IpamPoolIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6NetmaskLength", GoGetter: "Ipv6NetmaskLength"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6NetmaskLengthInput", GoGetter: "Ipv6NetmaskLengthInput"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6Pool", GoGetter: "Ipv6Pool"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6PoolInput", GoGetter: "Ipv6PoolInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAmazonProvidedIpv6CidrBlock", GoMethod: "ResetAmazonProvidedIpv6CidrBlock"},
			_jsii_.MemberMethod{JsiiMethod: "resetCidrBlock", GoMethod: "ResetCidrBlock"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpv4IpamPoolId", GoMethod: "ResetIpv4IpamPoolId"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpv4NetmaskLength", GoMethod: "ResetIpv4NetmaskLength"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpv6CidrBlock", GoMethod: "ResetIpv6CidrBlock"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpv6CidrBlockNetworkBorderGroup", GoMethod: "ResetIpv6CidrBlockNetworkBorderGroup"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpv6IpamPoolId", GoMethod: "ResetIpv6IpamPoolId"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpv6NetmaskLength", GoMethod: "ResetIpv6NetmaskLength"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpv6Pool", GoMethod: "ResetIpv6Pool"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "vpcCidrBlockId", GoGetter: "VpcCidrBlockId"},
			_jsii_.MemberProperty{JsiiProperty: "vpcId", GoGetter: "VpcId"},
			_jsii_.MemberProperty{JsiiProperty: "vpcIdInput", GoGetter: "VpcIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2VpcCidrBlock{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2VpcCidrBlock.Ec2VpcCidrBlockConfig",
		reflect.TypeOf((*Ec2VpcCidrBlockConfig)(nil)).Elem(),
	)
}
