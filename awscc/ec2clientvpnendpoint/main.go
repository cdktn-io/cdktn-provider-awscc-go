// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2clientvpnendpoint

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2ClientVpnEndpoint.Ec2ClientVpnEndpoint",
		reflect.TypeOf((*Ec2ClientVpnEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "authenticationOptions", GoGetter: "AuthenticationOptions"},
			_jsii_.MemberProperty{JsiiProperty: "authenticationOptionsInput", GoGetter: "AuthenticationOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "authorizationPolicy", GoGetter: "AuthorizationPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "authorizationPolicyInput", GoGetter: "AuthorizationPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "clientCidrBlock", GoGetter: "ClientCidrBlock"},
			_jsii_.MemberProperty{JsiiProperty: "clientCidrBlockInput", GoGetter: "ClientCidrBlockInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientConnectOptions", GoGetter: "ClientConnectOptions"},
			_jsii_.MemberProperty{JsiiProperty: "clientConnectOptionsInput", GoGetter: "ClientConnectOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientLoginBannerOptions", GoGetter: "ClientLoginBannerOptions"},
			_jsii_.MemberProperty{JsiiProperty: "clientLoginBannerOptionsInput", GoGetter: "ClientLoginBannerOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientRouteEnforcementOptions", GoGetter: "ClientRouteEnforcementOptions"},
			_jsii_.MemberProperty{JsiiProperty: "clientRouteEnforcementOptionsInput", GoGetter: "ClientRouteEnforcementOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientVpnEndpointId", GoGetter: "ClientVpnEndpointId"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "connectionLogOptions", GoGetter: "ConnectionLogOptions"},
			_jsii_.MemberProperty{JsiiProperty: "connectionLogOptionsInput", GoGetter: "ConnectionLogOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "devicePostureOptions", GoGetter: "DevicePostureOptions"},
			_jsii_.MemberProperty{JsiiProperty: "devicePostureOptionsInput", GoGetter: "DevicePostureOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "disconnectOnSessionTimeout", GoGetter: "DisconnectOnSessionTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "disconnectOnSessionTimeoutInput", GoGetter: "DisconnectOnSessionTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "dnsServers", GoGetter: "DnsServers"},
			_jsii_.MemberProperty{JsiiProperty: "dnsServersInput", GoGetter: "DnsServersInput"},
			_jsii_.MemberProperty{JsiiProperty: "endpointIpAddressType", GoGetter: "EndpointIpAddressType"},
			_jsii_.MemberProperty{JsiiProperty: "endpointIpAddressTypeInput", GoGetter: "EndpointIpAddressTypeInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putAuthenticationOptions", GoMethod: "PutAuthenticationOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putAuthorizationPolicy", GoMethod: "PutAuthorizationPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "putClientConnectOptions", GoMethod: "PutClientConnectOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putClientLoginBannerOptions", GoMethod: "PutClientLoginBannerOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putClientRouteEnforcementOptions", GoMethod: "PutClientRouteEnforcementOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putConnectionLogOptions", GoMethod: "PutConnectionLogOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putDevicePostureOptions", GoMethod: "PutDevicePostureOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putTagSpecifications", GoMethod: "PutTagSpecifications"},
			_jsii_.MemberMethod{JsiiMethod: "putTransitGatewayConfiguration", GoMethod: "PutTransitGatewayConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "registerProviderFeatureUsage", GoMethod: "RegisterProviderFeatureUsage"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthorizationPolicy", GoMethod: "ResetAuthorizationPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientCidrBlock", GoMethod: "ResetClientCidrBlock"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientConnectOptions", GoMethod: "ResetClientConnectOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientLoginBannerOptions", GoMethod: "ResetClientLoginBannerOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientRouteEnforcementOptions", GoMethod: "ResetClientRouteEnforcementOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetDevicePostureOptions", GoMethod: "ResetDevicePostureOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisconnectOnSessionTimeout", GoMethod: "ResetDisconnectOnSessionTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetDnsServers", GoMethod: "ResetDnsServers"},
			_jsii_.MemberMethod{JsiiMethod: "resetEndpointIpAddressType", GoMethod: "ResetEndpointIpAddressType"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecurityGroupIds", GoMethod: "ResetSecurityGroupIds"},
			_jsii_.MemberMethod{JsiiMethod: "resetSelfServicePortal", GoMethod: "ResetSelfServicePortal"},
			_jsii_.MemberMethod{JsiiMethod: "resetSessionTimeoutHours", GoMethod: "ResetSessionTimeoutHours"},
			_jsii_.MemberMethod{JsiiMethod: "resetSplitTunnel", GoMethod: "ResetSplitTunnel"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagSpecifications", GoMethod: "ResetTagSpecifications"},
			_jsii_.MemberMethod{JsiiMethod: "resetTrafficIpAddressType", GoMethod: "ResetTrafficIpAddressType"},
			_jsii_.MemberMethod{JsiiMethod: "resetTransitGatewayConfiguration", GoMethod: "ResetTransitGatewayConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetTransportProtocol", GoMethod: "ResetTransportProtocol"},
			_jsii_.MemberMethod{JsiiMethod: "resetVpcId", GoMethod: "ResetVpcId"},
			_jsii_.MemberMethod{JsiiMethod: "resetVpnPort", GoMethod: "ResetVpnPort"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupIds", GoGetter: "SecurityGroupIds"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupIdsInput", GoGetter: "SecurityGroupIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "selfServicePortal", GoGetter: "SelfServicePortal"},
			_jsii_.MemberProperty{JsiiProperty: "selfServicePortalInput", GoGetter: "SelfServicePortalInput"},
			_jsii_.MemberProperty{JsiiProperty: "serverCertificateArn", GoGetter: "ServerCertificateArn"},
			_jsii_.MemberProperty{JsiiProperty: "serverCertificateArnInput", GoGetter: "ServerCertificateArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "sessionTimeoutHours", GoGetter: "SessionTimeoutHours"},
			_jsii_.MemberProperty{JsiiProperty: "sessionTimeoutHoursInput", GoGetter: "SessionTimeoutHoursInput"},
			_jsii_.MemberProperty{JsiiProperty: "splitTunnel", GoGetter: "SplitTunnel"},
			_jsii_.MemberProperty{JsiiProperty: "splitTunnelInput", GoGetter: "SplitTunnelInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tagSpecifications", GoGetter: "TagSpecifications"},
			_jsii_.MemberProperty{JsiiProperty: "tagSpecificationsInput", GoGetter: "TagSpecificationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "trafficIpAddressType", GoGetter: "TrafficIpAddressType"},
			_jsii_.MemberProperty{JsiiProperty: "trafficIpAddressTypeInput", GoGetter: "TrafficIpAddressTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayConfiguration", GoGetter: "TransitGatewayConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayConfigurationInput", GoGetter: "TransitGatewayConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "transportProtocol", GoGetter: "TransportProtocol"},
			_jsii_.MemberProperty{JsiiProperty: "transportProtocolInput", GoGetter: "TransportProtocolInput"},
			_jsii_.MemberProperty{JsiiProperty: "vpcId", GoGetter: "VpcId"},
			_jsii_.MemberProperty{JsiiProperty: "vpcIdInput", GoGetter: "VpcIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "vpnPort", GoGetter: "VpnPort"},
			_jsii_.MemberProperty{JsiiProperty: "vpnPortInput", GoGetter: "VpnPortInput"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2ClientVpnEndpoint{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2ClientVpnEndpoint.Ec2ClientVpnEndpointAuthenticationOptions",
		reflect.TypeOf((*Ec2ClientVpnEndpointAuthenticationOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2ClientVpnEndpoint.Ec2ClientVpnEndpointAuthenticationOptionsActiveDirectory",
		reflect.TypeOf((*Ec2ClientVpnEndpointAuthenticationOptionsActiveDirectory)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2ClientVpnEndpoint.Ec2ClientVpnEndpointAuthenticationOptionsActiveDirectoryOutputReference",
		reflect.TypeOf((*Ec2ClientVpnEndpointAuthenticationOptionsActiveDirectoryOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "directoryId", GoGetter: "DirectoryId"},
			_jsii_.MemberProperty{JsiiProperty: "directoryIdInput", GoGetter: "DirectoryIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetDirectoryId", GoMethod: "ResetDirectoryId"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsActiveDirectoryOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2ClientVpnEndpoint.Ec2ClientVpnEndpointAuthenticationOptionsFederatedAuthentication",
		reflect.TypeOf((*Ec2ClientVpnEndpointAuthenticationOptionsFederatedAuthentication)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2ClientVpnEndpoint.Ec2ClientVpnEndpointAuthenticationOptionsFederatedAuthenticationOutputReference",
		reflect.TypeOf((*Ec2ClientVpnEndpointAuthenticationOptionsFederatedAuthenticationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetSamlProviderArn", GoMethod: "ResetSamlProviderArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetSelfServiceSamlProviderArn", GoMethod: "ResetSelfServiceSamlProviderArn"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "samlProviderArn", GoGetter: "SamlProviderArn"},
			_jsii_.MemberProperty{JsiiProperty: "samlProviderArnInput", GoGetter: "SamlProviderArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "selfServiceSamlProviderArn", GoGetter: "SelfServiceSamlProviderArn"},
			_jsii_.MemberProperty{JsiiProperty: "selfServiceSamlProviderArnInput", GoGetter: "SelfServiceSamlProviderArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsFederatedAuthenticationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2ClientVpnEndpoint.Ec2ClientVpnEndpointAuthenticationOptionsList",
		reflect.TypeOf((*Ec2ClientVpnEndpointAuthenticationOptionsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2ClientVpnEndpoint.Ec2ClientVpnEndpointAuthenticationOptionsMutualAuthentication",
		reflect.TypeOf((*Ec2ClientVpnEndpointAuthenticationOptionsMutualAuthentication)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2ClientVpnEndpoint.Ec2ClientVpnEndpointAuthenticationOptionsMutualAuthenticationOutputReference",
		reflect.TypeOf((*Ec2ClientVpnEndpointAuthenticationOptionsMutualAuthenticationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "clientRootCertificateChainArn", GoGetter: "ClientRootCertificateChainArn"},
			_jsii_.MemberProperty{JsiiProperty: "clientRootCertificateChainArnInput", GoGetter: "ClientRootCertificateChainArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientRootCertificateChainArn", GoMethod: "ResetClientRootCertificateChainArn"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsMutualAuthenticationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2ClientVpnEndpoint.Ec2ClientVpnEndpointAuthenticationOptionsOutputReference",
		reflect.TypeOf((*Ec2ClientVpnEndpointAuthenticationOptionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "activeDirectory", GoGetter: "ActiveDirectory"},
			_jsii_.MemberProperty{JsiiProperty: "activeDirectoryInput", GoGetter: "ActiveDirectoryInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "federatedAuthentication", GoGetter: "FederatedAuthentication"},
			_jsii_.MemberProperty{JsiiProperty: "federatedAuthenticationInput", GoGetter: "FederatedAuthenticationInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "mutualAuthentication", GoGetter: "MutualAuthentication"},
			_jsii_.MemberProperty{JsiiProperty: "mutualAuthenticationInput", GoGetter: "MutualAuthenticationInput"},
			_jsii_.MemberMethod{JsiiMethod: "putActiveDirectory", GoMethod: "PutActiveDirectory"},
			_jsii_.MemberMethod{JsiiMethod: "putFederatedAuthentication", GoMethod: "PutFederatedAuthentication"},
			_jsii_.MemberMethod{JsiiMethod: "putMutualAuthentication", GoMethod: "PutMutualAuthentication"},
			_jsii_.MemberMethod{JsiiMethod: "resetActiveDirectory", GoMethod: "ResetActiveDirectory"},
			_jsii_.MemberMethod{JsiiMethod: "resetFederatedAuthentication", GoMethod: "ResetFederatedAuthentication"},
			_jsii_.MemberMethod{JsiiMethod: "resetMutualAuthentication", GoMethod: "ResetMutualAuthentication"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2ClientVpnEndpointAuthenticationOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2ClientVpnEndpoint.Ec2ClientVpnEndpointAuthorizationPolicy",
		reflect.TypeOf((*Ec2ClientVpnEndpointAuthorizationPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2ClientVpnEndpoint.Ec2ClientVpnEndpointAuthorizationPolicyOutputReference",
		reflect.TypeOf((*Ec2ClientVpnEndpointAuthorizationPolicyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "policyDocument", GoGetter: "PolicyDocument"},
			_jsii_.MemberProperty{JsiiProperty: "policyDocumentInput", GoGetter: "PolicyDocumentInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetPolicyDocument", GoMethod: "ResetPolicyDocument"},
			_jsii_.MemberMethod{JsiiMethod: "resetShadowMode", GoMethod: "ResetShadowMode"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "shadowMode", GoGetter: "ShadowMode"},
			_jsii_.MemberProperty{JsiiProperty: "shadowModeInput", GoGetter: "ShadowModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2ClientVpnEndpointAuthorizationPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2ClientVpnEndpoint.Ec2ClientVpnEndpointClientConnectOptions",
		reflect.TypeOf((*Ec2ClientVpnEndpointClientConnectOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2ClientVpnEndpoint.Ec2ClientVpnEndpointClientConnectOptionsOutputReference",
		reflect.TypeOf((*Ec2ClientVpnEndpointClientConnectOptionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaFunctionArn", GoGetter: "LambdaFunctionArn"},
			_jsii_.MemberProperty{JsiiProperty: "lambdaFunctionArnInput", GoGetter: "LambdaFunctionArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetLambdaFunctionArn", GoMethod: "ResetLambdaFunctionArn"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2ClientVpnEndpointClientConnectOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2ClientVpnEndpoint.Ec2ClientVpnEndpointClientLoginBannerOptions",
		reflect.TypeOf((*Ec2ClientVpnEndpointClientLoginBannerOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2ClientVpnEndpoint.Ec2ClientVpnEndpointClientLoginBannerOptionsOutputReference",
		reflect.TypeOf((*Ec2ClientVpnEndpointClientLoginBannerOptionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bannerText", GoGetter: "BannerText"},
			_jsii_.MemberProperty{JsiiProperty: "bannerTextInput", GoGetter: "BannerTextInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetBannerText", GoMethod: "ResetBannerText"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2ClientVpnEndpointClientLoginBannerOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2ClientVpnEndpoint.Ec2ClientVpnEndpointClientRouteEnforcementOptions",
		reflect.TypeOf((*Ec2ClientVpnEndpointClientRouteEnforcementOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2ClientVpnEndpoint.Ec2ClientVpnEndpointClientRouteEnforcementOptionsOutputReference",
		reflect.TypeOf((*Ec2ClientVpnEndpointClientRouteEnforcementOptionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enforced", GoGetter: "Enforced"},
			_jsii_.MemberProperty{JsiiProperty: "enforcedInput", GoGetter: "EnforcedInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnforced", GoMethod: "ResetEnforced"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2ClientVpnEndpointClientRouteEnforcementOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2ClientVpnEndpoint.Ec2ClientVpnEndpointConfig",
		reflect.TypeOf((*Ec2ClientVpnEndpointConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2ClientVpnEndpoint.Ec2ClientVpnEndpointConnectionLogOptions",
		reflect.TypeOf((*Ec2ClientVpnEndpointConnectionLogOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2ClientVpnEndpoint.Ec2ClientVpnEndpointConnectionLogOptionsOutputReference",
		reflect.TypeOf((*Ec2ClientVpnEndpointConnectionLogOptionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLogGroup", GoGetter: "CloudwatchLogGroup"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLogGroupInput", GoGetter: "CloudwatchLogGroupInput"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLogStream", GoGetter: "CloudwatchLogStream"},
			_jsii_.MemberProperty{JsiiProperty: "cloudwatchLogStreamInput", GoGetter: "CloudwatchLogStreamInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "includeAuthorizationPolicyContext", GoGetter: "IncludeAuthorizationPolicyContext"},
			_jsii_.MemberProperty{JsiiProperty: "includeAuthorizationPolicyContextInput", GoGetter: "IncludeAuthorizationPolicyContextInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloudwatchLogGroup", GoMethod: "ResetCloudwatchLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloudwatchLogStream", GoMethod: "ResetCloudwatchLogStream"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludeAuthorizationPolicyContext", GoMethod: "ResetIncludeAuthorizationPolicyContext"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2ClientVpnEndpointConnectionLogOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2ClientVpnEndpoint.Ec2ClientVpnEndpointDevicePostureOptions",
		reflect.TypeOf((*Ec2ClientVpnEndpointDevicePostureOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2ClientVpnEndpoint.Ec2ClientVpnEndpointDevicePostureOptionsOutputReference",
		reflect.TypeOf((*Ec2ClientVpnEndpointDevicePostureOptionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putTrustProviders", GoMethod: "PutTrustProviders"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetTrustProviders", GoMethod: "ResetTrustProviders"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trustProviders", GoGetter: "TrustProviders"},
			_jsii_.MemberProperty{JsiiProperty: "trustProvidersInput", GoGetter: "TrustProvidersInput"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2ClientVpnEndpointDevicePostureOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2ClientVpnEndpoint.Ec2ClientVpnEndpointDevicePostureOptionsTrustProviders",
		reflect.TypeOf((*Ec2ClientVpnEndpointDevicePostureOptionsTrustProviders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2ClientVpnEndpoint.Ec2ClientVpnEndpointDevicePostureOptionsTrustProvidersList",
		reflect.TypeOf((*Ec2ClientVpnEndpointDevicePostureOptionsTrustProvidersList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2ClientVpnEndpointDevicePostureOptionsTrustProvidersList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2ClientVpnEndpoint.Ec2ClientVpnEndpointDevicePostureOptionsTrustProvidersOutputReference",
		reflect.TypeOf((*Ec2ClientVpnEndpointDevicePostureOptionsTrustProvidersOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "publicSigningKeyUrl", GoGetter: "PublicSigningKeyUrl"},
			_jsii_.MemberProperty{JsiiProperty: "publicSigningKeyUrlInput", GoGetter: "PublicSigningKeyUrlInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetPublicSigningKeyUrl", GoMethod: "ResetPublicSigningKeyUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetTenantId", GoMethod: "ResetTenantId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTrustProviderType", GoMethod: "ResetTrustProviderType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "tenantId", GoGetter: "TenantId"},
			_jsii_.MemberProperty{JsiiProperty: "tenantIdInput", GoGetter: "TenantIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trustProviderType", GoGetter: "TrustProviderType"},
			_jsii_.MemberProperty{JsiiProperty: "trustProviderTypeInput", GoGetter: "TrustProviderTypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2ClientVpnEndpointDevicePostureOptionsTrustProvidersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2ClientVpnEndpoint.Ec2ClientVpnEndpointTagSpecifications",
		reflect.TypeOf((*Ec2ClientVpnEndpointTagSpecifications)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2ClientVpnEndpoint.Ec2ClientVpnEndpointTagSpecificationsList",
		reflect.TypeOf((*Ec2ClientVpnEndpointTagSpecificationsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2ClientVpnEndpointTagSpecificationsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2ClientVpnEndpoint.Ec2ClientVpnEndpointTagSpecificationsOutputReference",
		reflect.TypeOf((*Ec2ClientVpnEndpointTagSpecificationsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putTags", GoMethod: "PutTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetResourceType", GoMethod: "ResetResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "resourceType", GoGetter: "ResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "resourceTypeInput", GoGetter: "ResourceTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2ClientVpnEndpointTagSpecificationsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2ClientVpnEndpoint.Ec2ClientVpnEndpointTagSpecificationsTags",
		reflect.TypeOf((*Ec2ClientVpnEndpointTagSpecificationsTags)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2ClientVpnEndpoint.Ec2ClientVpnEndpointTagSpecificationsTagsList",
		reflect.TypeOf((*Ec2ClientVpnEndpointTagSpecificationsTagsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2ClientVpnEndpointTagSpecificationsTagsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2ClientVpnEndpoint.Ec2ClientVpnEndpointTagSpecificationsTagsOutputReference",
		reflect.TypeOf((*Ec2ClientVpnEndpointTagSpecificationsTagsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "key", GoGetter: "Key"},
			_jsii_.MemberProperty{JsiiProperty: "keyInput", GoGetter: "KeyInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetKey", GoMethod: "ResetKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetValue", GoMethod: "ResetValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2ClientVpnEndpointTagSpecificationsTagsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2ClientVpnEndpoint.Ec2ClientVpnEndpointTransitGatewayConfiguration",
		reflect.TypeOf((*Ec2ClientVpnEndpointTransitGatewayConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2ClientVpnEndpoint.Ec2ClientVpnEndpointTransitGatewayConfigurationOutputReference",
		reflect.TypeOf((*Ec2ClientVpnEndpointTransitGatewayConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "availabilityZoneIds", GoGetter: "AvailabilityZoneIds"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZoneIdsInput", GoGetter: "AvailabilityZoneIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZones", GoGetter: "AvailabilityZones"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZonesInput", GoGetter: "AvailabilityZonesInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetAvailabilityZoneIds", GoMethod: "ResetAvailabilityZoneIds"},
			_jsii_.MemberMethod{JsiiMethod: "resetAvailabilityZones", GoMethod: "ResetAvailabilityZones"},
			_jsii_.MemberMethod{JsiiMethod: "resetTransitGatewayId", GoMethod: "ResetTransitGatewayId"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayId", GoGetter: "TransitGatewayId"},
			_jsii_.MemberProperty{JsiiProperty: "transitGatewayIdInput", GoGetter: "TransitGatewayIdInput"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2ClientVpnEndpointTransitGatewayConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
}
