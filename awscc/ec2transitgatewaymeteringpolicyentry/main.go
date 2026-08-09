// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2transitgatewaymeteringpolicyentry

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2TransitGatewayMeteringPolicyEntry.Ec2TransitGatewayMeteringPolicyEntry",
		reflect.TypeOf((*Ec2TransitGatewayMeteringPolicyEntry)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "destinationCidrBlock", GoGetter: "DestinationCidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "destinationCidrBlockInput", GoGetter: "DestinationCidrBlockInput"},
			_jsii_.MemberProperty{JsiiProperty: "destinationPortRange", GoGetter: "DestinationPortRange"},
			_jsii_.MemberProperty{JsiiProperty: "destinationPortRangeInput", GoGetter: "DestinationPortRangeInput"},
			_jsii_.MemberProperty{JsiiProperty: "destinationTransitGatewayAttachmentId", GoGetter: "DestinationTransitGatewayAttachmentId"},
			_jsii_.MemberProperty{JsiiProperty: "destinationTransitGatewayAttachmentIdInput", GoGetter: "DestinationTransitGatewayAttachmentIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "destinationTransitGatewayAttachmentType", GoGetter: "DestinationTransitGatewayAttachmentType"},
			_jsii_.MemberProperty{JsiiProperty: "destinationTransitGatewayAttachmentTypeInput", GoGetter: "DestinationTransitGatewayAttachmentTypeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "markWriteOnlyAttribute", GoMethod: "MarkWriteOnlyAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "meteredAccount", GoGetter: "MeteredAccount"},
			_jsii_.MemberProperty{JsiiProperty: "meteredAccountInput", GoGetter: "MeteredAccountInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "policyRuleNumber", GoGetter: "PolicyRuleNumber"},
			_jsii_.MemberProperty{JsiiProperty: "policyRuleNumberInput", GoGetter: "PolicyRuleNumberInput"},
			_jsii_.MemberProperty{JsiiProperty: "protocol", GoGetter: "Protocol"},
			_jsii_.MemberProperty{JsiiProperty: "protocolInput", GoGetter: "ProtocolInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "registerProviderFeatureUsage", GoMethod: "RegisterProviderFeatureUsage"},
			_jsii_.MemberMethod{JsiiMethod: "resetDestinationCidrBlock", GoMethod: "ResetDestinationCidrBlock"},
			_jsii_.MemberMethod{JsiiMethod: "resetDestinationPortRange", GoMethod: "ResetDestinationPortRange"},
			_jsii_.MemberMethod{JsiiMethod: "resetDestinationTransitGatewayAttachmentId", GoMethod: "ResetDestinationTransitGatewayAttachmentId"},
			_jsii_.MemberMethod{JsiiMethod: "resetDestinationTransitGatewayAttachmentType", GoMethod: "ResetDestinationTransitGatewayAttachmentType"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetProtocol", GoMethod: "ResetProtocol"},
			_jsii_.MemberMethod{JsiiMethod: "resetSourceCidrBlock", GoMethod: "ResetSourceCidrBlock"},
			_jsii_.MemberMethod{JsiiMethod: "resetSourcePortRange", GoMethod: "ResetSourcePortRange"},
			_jsii_.MemberMethod{JsiiMethod: "resetSourceTransitGatewayAttachmentId", GoMethod: "ResetSourceTransitGatewayAttachmentId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSourceTransitGatewayAttachmentType", GoMethod: "ResetSourceTransitGatewayAttachmentType"},
			_jsii_.MemberProperty{JsiiProperty: "sourceCidrBlock", GoGetter: "SourceCidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "sourceCidrBlockInput", GoGetter: "SourceCidrBlockInput"},
			_jsii_.MemberProperty{JsiiProperty: "sourcePortRange", GoGetter: "SourcePortRange"},
			_jsii_.MemberProperty{JsiiProperty: "sourcePortRangeInput", GoGetter: "SourcePortRangeInput"},
			_jsii_.MemberProperty{JsiiProperty: "sourceTransitGatewayAttachmentId", GoGetter: "SourceTransitGatewayAttachmentId"},
			_jsii_.MemberProperty{JsiiProperty: "sourceTransitGatewayAttachmentIdInput", GoGetter: "SourceTransitGatewayAttachmentIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "sourceTransitGatewayAttachmentType", GoGetter: "SourceTransitGatewayAttachmentType"},
			_jsii_.MemberProperty{JsiiProperty: "sourceTransitGatewayAttachmentTypeInput", GoGetter: "SourceTransitGatewayAttachmentTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "state", GoGetter: "State"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayMeteringPolicyId", GoGetter: "TransitGatewayMeteringPolicyId"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayMeteringPolicyIdInput", GoGetter: "TransitGatewayMeteringPolicyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "updateEffectiveAt", GoGetter: "UpdateEffectiveAt"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2TransitGatewayMeteringPolicyEntry{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2TransitGatewayMeteringPolicyEntry.Ec2TransitGatewayMeteringPolicyEntryConfig",
		reflect.TypeOf((*Ec2TransitGatewayMeteringPolicyEntryConfig)(nil)).Elem(),
	)
}
