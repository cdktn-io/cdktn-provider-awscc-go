// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package medialivemultiplexprogram

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.medialiveMultiplexprogram.MedialiveMultiplexprogram",
		reflect.TypeOf((*MedialiveMultiplexprogram)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "channelId", GoGetter: "ChannelId"},
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
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "markWriteOnlyAttribute", GoMethod: "MarkWriteOnlyAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "multiplexId", GoGetter: "MultiplexId"},
			_jsii_.MemberProperty{JsiiProperty: "multiplexIdInput", GoGetter: "MultiplexIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "multiplexProgramSettings", GoGetter: "MultiplexProgramSettings"},
			_jsii_.MemberProperty{JsiiProperty: "multiplexProgramSettingsInput", GoGetter: "MultiplexProgramSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "packetIdentifiersMap", GoGetter: "PacketIdentifiersMap"},
			_jsii_.MemberProperty{JsiiProperty: "packetIdentifiersMapInput", GoGetter: "PacketIdentifiersMapInput"},
			_jsii_.MemberProperty{JsiiProperty: "pipelineDetails", GoGetter: "PipelineDetails"},
			_jsii_.MemberProperty{JsiiProperty: "pipelineDetailsInput", GoGetter: "PipelineDetailsInput"},
			_jsii_.MemberProperty{JsiiProperty: "preferredChannelPipeline", GoGetter: "PreferredChannelPipeline"},
			_jsii_.MemberProperty{JsiiProperty: "preferredChannelPipelineInput", GoGetter: "PreferredChannelPipelineInput"},
			_jsii_.MemberProperty{JsiiProperty: "programName", GoGetter: "ProgramName"},
			_jsii_.MemberProperty{JsiiProperty: "programNameInput", GoGetter: "ProgramNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putMultiplexProgramSettings", GoMethod: "PutMultiplexProgramSettings"},
			_jsii_.MemberMethod{JsiiMethod: "putPacketIdentifiersMap", GoMethod: "PutPacketIdentifiersMap"},
			_jsii_.MemberMethod{JsiiMethod: "putPipelineDetails", GoMethod: "PutPipelineDetails"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "registerProviderFeatureUsage", GoMethod: "RegisterProviderFeatureUsage"},
			_jsii_.MemberMethod{JsiiMethod: "resetMultiplexId", GoMethod: "ResetMultiplexId"},
			_jsii_.MemberMethod{JsiiMethod: "resetMultiplexProgramSettings", GoMethod: "ResetMultiplexProgramSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPacketIdentifiersMap", GoMethod: "ResetPacketIdentifiersMap"},
			_jsii_.MemberMethod{JsiiMethod: "resetPipelineDetails", GoMethod: "ResetPipelineDetails"},
			_jsii_.MemberMethod{JsiiMethod: "resetPreferredChannelPipeline", GoMethod: "ResetPreferredChannelPipeline"},
			_jsii_.MemberMethod{JsiiMethod: "resetProgramName", GoMethod: "ResetProgramName"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_MedialiveMultiplexprogram{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.medialiveMultiplexprogram.MedialiveMultiplexprogramConfig",
		reflect.TypeOf((*MedialiveMultiplexprogramConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.medialiveMultiplexprogram.MedialiveMultiplexprogramMultiplexProgramSettings",
		reflect.TypeOf((*MedialiveMultiplexprogramMultiplexProgramSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.medialiveMultiplexprogram.MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference",
		reflect.TypeOf((*MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "preferredChannelPipeline", GoGetter: "PreferredChannelPipeline"},
			_jsii_.MemberProperty{JsiiProperty: "preferredChannelPipelineInput", GoGetter: "PreferredChannelPipelineInput"},
			_jsii_.MemberProperty{JsiiProperty: "programNumber", GoGetter: "ProgramNumber"},
			_jsii_.MemberProperty{JsiiProperty: "programNumberInput", GoGetter: "ProgramNumberInput"},
			_jsii_.MemberMethod{JsiiMethod: "putServiceDescriptor", GoMethod: "PutServiceDescriptor"},
			_jsii_.MemberMethod{JsiiMethod: "putVideoSettings", GoMethod: "PutVideoSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetPreferredChannelPipeline", GoMethod: "ResetPreferredChannelPipeline"},
			_jsii_.MemberMethod{JsiiMethod: "resetProgramNumber", GoMethod: "ResetProgramNumber"},
			_jsii_.MemberMethod{JsiiMethod: "resetServiceDescriptor", GoMethod: "ResetServiceDescriptor"},
			_jsii_.MemberMethod{JsiiMethod: "resetVideoSettings", GoMethod: "ResetVideoSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "serviceDescriptor", GoGetter: "ServiceDescriptor"},
			_jsii_.MemberProperty{JsiiProperty: "serviceDescriptorInput", GoGetter: "ServiceDescriptorInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "videoSettings", GoGetter: "VideoSettings"},
			_jsii_.MemberProperty{JsiiProperty: "videoSettingsInput", GoGetter: "VideoSettingsInput"},
		},
		func() interface{} {
			j := jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.medialiveMultiplexprogram.MedialiveMultiplexprogramMultiplexProgramSettingsServiceDescriptor",
		reflect.TypeOf((*MedialiveMultiplexprogramMultiplexProgramSettingsServiceDescriptor)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.medialiveMultiplexprogram.MedialiveMultiplexprogramMultiplexProgramSettingsServiceDescriptorOutputReference",
		reflect.TypeOf((*MedialiveMultiplexprogramMultiplexProgramSettingsServiceDescriptorOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "providerName", GoGetter: "ProviderName"},
			_jsii_.MemberProperty{JsiiProperty: "providerNameInput", GoGetter: "ProviderNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetProviderName", GoMethod: "ResetProviderName"},
			_jsii_.MemberMethod{JsiiMethod: "resetServiceName", GoMethod: "ResetServiceName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "serviceName", GoGetter: "ServiceName"},
			_jsii_.MemberProperty{JsiiProperty: "serviceNameInput", GoGetter: "ServiceNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsServiceDescriptorOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.medialiveMultiplexprogram.MedialiveMultiplexprogramMultiplexProgramSettingsVideoSettings",
		reflect.TypeOf((*MedialiveMultiplexprogramMultiplexProgramSettingsVideoSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.medialiveMultiplexprogram.MedialiveMultiplexprogramMultiplexProgramSettingsVideoSettingsOutputReference",
		reflect.TypeOf((*MedialiveMultiplexprogramMultiplexProgramSettingsVideoSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "constantBitrate", GoGetter: "ConstantBitrate"},
			_jsii_.MemberProperty{JsiiProperty: "constantBitrateInput", GoGetter: "ConstantBitrateInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putStatmuxSettings", GoMethod: "PutStatmuxSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetConstantBitrate", GoMethod: "ResetConstantBitrate"},
			_jsii_.MemberMethod{JsiiMethod: "resetStatmuxSettings", GoMethod: "ResetStatmuxSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "statmuxSettings", GoGetter: "StatmuxSettings"},
			_jsii_.MemberProperty{JsiiProperty: "statmuxSettingsInput", GoGetter: "StatmuxSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsVideoSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.medialiveMultiplexprogram.MedialiveMultiplexprogramMultiplexProgramSettingsVideoSettingsStatmuxSettings",
		reflect.TypeOf((*MedialiveMultiplexprogramMultiplexProgramSettingsVideoSettingsStatmuxSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.medialiveMultiplexprogram.MedialiveMultiplexprogramMultiplexProgramSettingsVideoSettingsStatmuxSettingsOutputReference",
		reflect.TypeOf((*MedialiveMultiplexprogramMultiplexProgramSettingsVideoSettingsStatmuxSettingsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "maximumBitrate", GoGetter: "MaximumBitrate"},
			_jsii_.MemberProperty{JsiiProperty: "maximumBitrateInput", GoGetter: "MaximumBitrateInput"},
			_jsii_.MemberProperty{JsiiProperty: "minimumBitrate", GoGetter: "MinimumBitrate"},
			_jsii_.MemberProperty{JsiiProperty: "minimumBitrateInput", GoGetter: "MinimumBitrateInput"},
			_jsii_.MemberProperty{JsiiProperty: "priority", GoGetter: "Priority"},
			_jsii_.MemberProperty{JsiiProperty: "priorityInput", GoGetter: "PriorityInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaximumBitrate", GoMethod: "ResetMaximumBitrate"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinimumBitrate", GoMethod: "ResetMinimumBitrate"},
			_jsii_.MemberMethod{JsiiMethod: "resetPriority", GoMethod: "ResetPriority"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MedialiveMultiplexprogramMultiplexProgramSettingsVideoSettingsStatmuxSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.medialiveMultiplexprogram.MedialiveMultiplexprogramPacketIdentifiersMap",
		reflect.TypeOf((*MedialiveMultiplexprogramPacketIdentifiersMap)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.medialiveMultiplexprogram.MedialiveMultiplexprogramPacketIdentifiersMapOutputReference",
		reflect.TypeOf((*MedialiveMultiplexprogramPacketIdentifiersMapOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "audioPids", GoGetter: "AudioPids"},
			_jsii_.MemberProperty{JsiiProperty: "audioPidsInput", GoGetter: "AudioPidsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dvbSubPids", GoGetter: "DvbSubPids"},
			_jsii_.MemberProperty{JsiiProperty: "dvbSubPidsInput", GoGetter: "DvbSubPidsInput"},
			_jsii_.MemberProperty{JsiiProperty: "dvbTeletextPid", GoGetter: "DvbTeletextPid"},
			_jsii_.MemberProperty{JsiiProperty: "dvbTeletextPidInput", GoGetter: "DvbTeletextPidInput"},
			_jsii_.MemberProperty{JsiiProperty: "etvPlatformPid", GoGetter: "EtvPlatformPid"},
			_jsii_.MemberProperty{JsiiProperty: "etvPlatformPidInput", GoGetter: "EtvPlatformPidInput"},
			_jsii_.MemberProperty{JsiiProperty: "etvSignalPid", GoGetter: "EtvSignalPid"},
			_jsii_.MemberProperty{JsiiProperty: "etvSignalPidInput", GoGetter: "EtvSignalPidInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "klvDataPids", GoGetter: "KlvDataPids"},
			_jsii_.MemberProperty{JsiiProperty: "klvDataPidsInput", GoGetter: "KlvDataPidsInput"},
			_jsii_.MemberProperty{JsiiProperty: "pcrPid", GoGetter: "PcrPid"},
			_jsii_.MemberProperty{JsiiProperty: "pcrPidInput", GoGetter: "PcrPidInput"},
			_jsii_.MemberProperty{JsiiProperty: "pmtPid", GoGetter: "PmtPid"},
			_jsii_.MemberProperty{JsiiProperty: "pmtPidInput", GoGetter: "PmtPidInput"},
			_jsii_.MemberProperty{JsiiProperty: "privateMetadataPid", GoGetter: "PrivateMetadataPid"},
			_jsii_.MemberProperty{JsiiProperty: "privateMetadataPidInput", GoGetter: "PrivateMetadataPidInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAudioPids", GoMethod: "ResetAudioPids"},
			_jsii_.MemberMethod{JsiiMethod: "resetDvbSubPids", GoMethod: "ResetDvbSubPids"},
			_jsii_.MemberMethod{JsiiMethod: "resetDvbTeletextPid", GoMethod: "ResetDvbTeletextPid"},
			_jsii_.MemberMethod{JsiiMethod: "resetEtvPlatformPid", GoMethod: "ResetEtvPlatformPid"},
			_jsii_.MemberMethod{JsiiMethod: "resetEtvSignalPid", GoMethod: "ResetEtvSignalPid"},
			_jsii_.MemberMethod{JsiiMethod: "resetKlvDataPids", GoMethod: "ResetKlvDataPids"},
			_jsii_.MemberMethod{JsiiMethod: "resetPcrPid", GoMethod: "ResetPcrPid"},
			_jsii_.MemberMethod{JsiiMethod: "resetPmtPid", GoMethod: "ResetPmtPid"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrivateMetadataPid", GoMethod: "ResetPrivateMetadataPid"},
			_jsii_.MemberMethod{JsiiMethod: "resetScte27Pids", GoMethod: "ResetScte27Pids"},
			_jsii_.MemberMethod{JsiiMethod: "resetScte35Pid", GoMethod: "ResetScte35Pid"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimedMetadataPid", GoMethod: "ResetTimedMetadataPid"},
			_jsii_.MemberMethod{JsiiMethod: "resetVideoPid", GoMethod: "ResetVideoPid"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scte27Pids", GoGetter: "Scte27Pids"},
			_jsii_.MemberProperty{JsiiProperty: "scte27PidsInput", GoGetter: "Scte27PidsInput"},
			_jsii_.MemberProperty{JsiiProperty: "scte35Pid", GoGetter: "Scte35Pid"},
			_jsii_.MemberProperty{JsiiProperty: "scte35PidInput", GoGetter: "Scte35PidInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "timedMetadataPid", GoGetter: "TimedMetadataPid"},
			_jsii_.MemberProperty{JsiiProperty: "timedMetadataPidInput", GoGetter: "TimedMetadataPidInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "videoPid", GoGetter: "VideoPid"},
			_jsii_.MemberProperty{JsiiProperty: "videoPidInput", GoGetter: "VideoPidInput"},
		},
		func() interface{} {
			j := jsiiProxy_MedialiveMultiplexprogramPacketIdentifiersMapOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.medialiveMultiplexprogram.MedialiveMultiplexprogramPipelineDetails",
		reflect.TypeOf((*MedialiveMultiplexprogramPipelineDetails)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.medialiveMultiplexprogram.MedialiveMultiplexprogramPipelineDetailsList",
		reflect.TypeOf((*MedialiveMultiplexprogramPipelineDetailsList)(nil)).Elem(),
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
			j := jsiiProxy_MedialiveMultiplexprogramPipelineDetailsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.medialiveMultiplexprogram.MedialiveMultiplexprogramPipelineDetailsOutputReference",
		reflect.TypeOf((*MedialiveMultiplexprogramPipelineDetailsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "activeChannelPipeline", GoGetter: "ActiveChannelPipeline"},
			_jsii_.MemberProperty{JsiiProperty: "activeChannelPipelineInput", GoGetter: "ActiveChannelPipelineInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "pipelineId", GoGetter: "PipelineId"},
			_jsii_.MemberProperty{JsiiProperty: "pipelineIdInput", GoGetter: "PipelineIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetActiveChannelPipeline", GoMethod: "ResetActiveChannelPipeline"},
			_jsii_.MemberMethod{JsiiMethod: "resetPipelineId", GoMethod: "ResetPipelineId"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_MedialiveMultiplexprogramPipelineDetailsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
}
