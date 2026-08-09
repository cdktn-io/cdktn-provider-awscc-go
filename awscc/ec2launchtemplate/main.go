// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2launchtemplate

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplate",
		reflect.TypeOf((*Ec2LaunchTemplate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "defaultVersionNumber", GoGetter: "DefaultVersionNumber"},
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
			_jsii_.MemberProperty{JsiiProperty: "latestVersionNumber", GoGetter: "LatestVersionNumber"},
			_jsii_.MemberProperty{JsiiProperty: "launchTemplateData", GoGetter: "LaunchTemplateData"},
			_jsii_.MemberProperty{JsiiProperty: "launchTemplateDataInput", GoGetter: "LaunchTemplateDataInput"},
			_jsii_.MemberProperty{JsiiProperty: "launchTemplateId", GoGetter: "LaunchTemplateId"},
			_jsii_.MemberProperty{JsiiProperty: "launchTemplateName", GoGetter: "LaunchTemplateName"},
			_jsii_.MemberProperty{JsiiProperty: "launchTemplateNameInput", GoGetter: "LaunchTemplateNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "markWriteOnlyAttribute", GoMethod: "MarkWriteOnlyAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putLaunchTemplateData", GoMethod: "PutLaunchTemplateData"},
			_jsii_.MemberMethod{JsiiMethod: "putTagSpecifications", GoMethod: "PutTagSpecifications"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "registerProviderFeatureUsage", GoMethod: "RegisterProviderFeatureUsage"},
			_jsii_.MemberMethod{JsiiMethod: "resetLaunchTemplateName", GoMethod: "ResetLaunchTemplateName"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagSpecifications", GoMethod: "ResetTagSpecifications"},
			_jsii_.MemberMethod{JsiiMethod: "resetVersionDescription", GoMethod: "ResetVersionDescription"},
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
			_jsii_.MemberProperty{JsiiProperty: "versionDescription", GoGetter: "VersionDescription"},
			_jsii_.MemberProperty{JsiiProperty: "versionDescriptionInput", GoGetter: "VersionDescriptionInput"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2LaunchTemplate{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateConfig",
		reflect.TypeOf((*Ec2LaunchTemplateConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateData",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateData)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappings",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsEbs",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsEbs)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsEbsOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsEbsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "deleteOnTermination", GoGetter: "DeleteOnTermination"},
			_jsii_.MemberProperty{JsiiProperty: "deleteOnTerminationInput", GoGetter: "DeleteOnTerminationInput"},
			_jsii_.MemberProperty{JsiiProperty: "encrypted", GoGetter: "Encrypted"},
			_jsii_.MemberProperty{JsiiProperty: "encryptedInput", GoGetter: "EncryptedInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "iops", GoGetter: "Iops"},
			_jsii_.MemberProperty{JsiiProperty: "iopsInput", GoGetter: "IopsInput"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyId", GoGetter: "KmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyIdInput", GoGetter: "KmsKeyIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeleteOnTermination", GoMethod: "ResetDeleteOnTermination"},
			_jsii_.MemberMethod{JsiiMethod: "resetEncrypted", GoMethod: "ResetEncrypted"},
			_jsii_.MemberMethod{JsiiMethod: "resetIops", GoMethod: "ResetIops"},
			_jsii_.MemberMethod{JsiiMethod: "resetKmsKeyId", GoMethod: "ResetKmsKeyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSnapshotId", GoMethod: "ResetSnapshotId"},
			_jsii_.MemberMethod{JsiiMethod: "resetThroughput", GoMethod: "ResetThroughput"},
			_jsii_.MemberMethod{JsiiMethod: "resetVolumeSize", GoMethod: "ResetVolumeSize"},
			_jsii_.MemberMethod{JsiiMethod: "resetVolumeType", GoMethod: "ResetVolumeType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "snapshotId", GoGetter: "SnapshotId"},
			_jsii_.MemberProperty{JsiiProperty: "snapshotIdInput", GoGetter: "SnapshotIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "throughput", GoGetter: "Throughput"},
			_jsii_.MemberProperty{JsiiProperty: "throughputInput", GoGetter: "ThroughputInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "volumeSize", GoGetter: "VolumeSize"},
			_jsii_.MemberProperty{JsiiProperty: "volumeSizeInput", GoGetter: "VolumeSizeInput"},
			_jsii_.MemberProperty{JsiiProperty: "volumeType", GoGetter: "VolumeType"},
			_jsii_.MemberProperty{JsiiProperty: "volumeTypeInput", GoGetter: "VolumeTypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsEbsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsList",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsList)(nil)).Elem(),
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
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "deviceName", GoGetter: "DeviceName"},
			_jsii_.MemberProperty{JsiiProperty: "deviceNameInput", GoGetter: "DeviceNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "ebs", GoGetter: "Ebs"},
			_jsii_.MemberProperty{JsiiProperty: "ebsInput", GoGetter: "EbsInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "noDevice", GoGetter: "NoDevice"},
			_jsii_.MemberProperty{JsiiProperty: "noDeviceInput", GoGetter: "NoDeviceInput"},
			_jsii_.MemberMethod{JsiiMethod: "putEbs", GoMethod: "PutEbs"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeviceName", GoMethod: "ResetDeviceName"},
			_jsii_.MemberMethod{JsiiMethod: "resetEbs", GoMethod: "ResetEbs"},
			_jsii_.MemberMethod{JsiiMethod: "resetNoDevice", GoMethod: "ResetNoDevice"},
			_jsii_.MemberMethod{JsiiMethod: "resetVirtualName", GoMethod: "ResetVirtualName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "virtualName", GoGetter: "VirtualName"},
			_jsii_.MemberProperty{JsiiProperty: "virtualNameInput", GoGetter: "VirtualNameInput"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataBlockDeviceMappingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataCapacityReservationSpecification",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataCapacityReservationSpecification)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataCapacityReservationSpecificationCapacityReservationTarget",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataCapacityReservationSpecificationCapacityReservationTarget)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataCapacityReservationSpecificationCapacityReservationTargetOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataCapacityReservationSpecificationCapacityReservationTargetOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "capacityReservationId", GoGetter: "CapacityReservationId"},
			_jsii_.MemberProperty{JsiiProperty: "capacityReservationIdInput", GoGetter: "CapacityReservationIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "capacityReservationResourceGroupArn", GoGetter: "CapacityReservationResourceGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "capacityReservationResourceGroupArnInput", GoGetter: "CapacityReservationResourceGroupArnInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetCapacityReservationId", GoMethod: "ResetCapacityReservationId"},
			_jsii_.MemberMethod{JsiiMethod: "resetCapacityReservationResourceGroupArn", GoMethod: "ResetCapacityReservationResourceGroupArn"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataCapacityReservationSpecificationCapacityReservationTargetOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataCapacityReservationSpecificationOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataCapacityReservationSpecificationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "capacityReservationPreference", GoGetter: "CapacityReservationPreference"},
			_jsii_.MemberProperty{JsiiProperty: "capacityReservationPreferenceInput", GoGetter: "CapacityReservationPreferenceInput"},
			_jsii_.MemberProperty{JsiiProperty: "capacityReservationTarget", GoGetter: "CapacityReservationTarget"},
			_jsii_.MemberProperty{JsiiProperty: "capacityReservationTargetInput", GoGetter: "CapacityReservationTargetInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putCapacityReservationTarget", GoMethod: "PutCapacityReservationTarget"},
			_jsii_.MemberMethod{JsiiMethod: "resetCapacityReservationPreference", GoMethod: "ResetCapacityReservationPreference"},
			_jsii_.MemberMethod{JsiiMethod: "resetCapacityReservationTarget", GoMethod: "ResetCapacityReservationTarget"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataCapacityReservationSpecificationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataCpuOptions",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataCpuOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataCpuOptionsOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataCpuOptionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "amdSevSnp", GoGetter: "AmdSevSnp"},
			_jsii_.MemberProperty{JsiiProperty: "amdSevSnpInput", GoGetter: "AmdSevSnpInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "coreCount", GoGetter: "CoreCount"},
			_jsii_.MemberProperty{JsiiProperty: "coreCountInput", GoGetter: "CoreCountInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAmdSevSnp", GoMethod: "ResetAmdSevSnp"},
			_jsii_.MemberMethod{JsiiMethod: "resetCoreCount", GoMethod: "ResetCoreCount"},
			_jsii_.MemberMethod{JsiiMethod: "resetThreadsPerCore", GoMethod: "ResetThreadsPerCore"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "threadsPerCore", GoGetter: "ThreadsPerCore"},
			_jsii_.MemberProperty{JsiiProperty: "threadsPerCoreInput", GoGetter: "ThreadsPerCoreInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataCpuOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataCreditSpecification",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataCreditSpecification)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataCreditSpecificationOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataCreditSpecificationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "cpuCredits", GoGetter: "CpuCredits"},
			_jsii_.MemberProperty{JsiiProperty: "cpuCreditsInput", GoGetter: "CpuCreditsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetCpuCredits", GoMethod: "ResetCpuCredits"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataCreditSpecificationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataElasticGpuSpecifications",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataElasticGpuSpecifications)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataElasticGpuSpecificationsList",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataElasticGpuSpecificationsList)(nil)).Elem(),
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
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataElasticGpuSpecificationsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataElasticGpuSpecificationsOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataElasticGpuSpecificationsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetType", GoMethod: "ResetType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataElasticGpuSpecificationsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataElasticInferenceAccelerators",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataElasticInferenceAccelerators)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataElasticInferenceAcceleratorsList",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataElasticInferenceAcceleratorsList)(nil)).Elem(),
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
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataElasticInferenceAcceleratorsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataElasticInferenceAcceleratorsOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataElasticInferenceAcceleratorsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "countInput", GoGetter: "CountInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetCount", GoMethod: "ResetCount"},
			_jsii_.MemberMethod{JsiiMethod: "resetType", GoMethod: "ResetType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataElasticInferenceAcceleratorsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataEnclaveOptions",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataEnclaveOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataEnclaveOptionsOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataEnclaveOptionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataEnclaveOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataHibernationOptions",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataHibernationOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataHibernationOptionsOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataHibernationOptionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "configured", GoGetter: "Configured"},
			_jsii_.MemberProperty{JsiiProperty: "configuredInput", GoGetter: "ConfiguredInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetConfigured", GoMethod: "ResetConfigured"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataHibernationOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataIamInstanceProfile",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataIamInstanceProfile)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataIamInstanceProfileOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataIamInstanceProfileOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "arn", GoGetter: "Arn"},
			_jsii_.MemberProperty{JsiiProperty: "arnInput", GoGetter: "ArnInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetArn", GoMethod: "ResetArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataIamInstanceProfileOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptions",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "marketType", GoGetter: "MarketType"},
			_jsii_.MemberProperty{JsiiProperty: "marketTypeInput", GoGetter: "MarketTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "putSpotOptions", GoMethod: "PutSpotOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetMarketType", GoMethod: "ResetMarketType"},
			_jsii_.MemberMethod{JsiiMethod: "resetSpotOptions", GoMethod: "ResetSpotOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "spotOptions", GoGetter: "SpotOptions"},
			_jsii_.MemberProperty{JsiiProperty: "spotOptionsInput", GoGetter: "SpotOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptions",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "blockDurationMinutes", GoGetter: "BlockDurationMinutes"},
			_jsii_.MemberProperty{JsiiProperty: "blockDurationMinutesInput", GoGetter: "BlockDurationMinutesInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "instanceInterruptionBehavior", GoGetter: "InstanceInterruptionBehavior"},
			_jsii_.MemberProperty{JsiiProperty: "instanceInterruptionBehaviorInput", GoGetter: "InstanceInterruptionBehaviorInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "maxPrice", GoGetter: "MaxPrice"},
			_jsii_.MemberProperty{JsiiProperty: "maxPriceInput", GoGetter: "MaxPriceInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetBlockDurationMinutes", GoMethod: "ResetBlockDurationMinutes"},
			_jsii_.MemberMethod{JsiiMethod: "resetInstanceInterruptionBehavior", GoMethod: "ResetInstanceInterruptionBehavior"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxPrice", GoMethod: "ResetMaxPrice"},
			_jsii_.MemberMethod{JsiiMethod: "resetSpotInstanceType", GoMethod: "ResetSpotInstanceType"},
			_jsii_.MemberMethod{JsiiMethod: "resetValidUntil", GoMethod: "ResetValidUntil"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "spotInstanceType", GoGetter: "SpotInstanceType"},
			_jsii_.MemberProperty{JsiiProperty: "spotInstanceTypeInput", GoGetter: "SpotInstanceTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "validUntil", GoGetter: "ValidUntil"},
			_jsii_.MemberProperty{JsiiProperty: "validUntilInput", GoGetter: "ValidUntilInput"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataInstanceRequirements",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataInstanceRequirements)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsAcceleratorCount",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsAcceleratorCount)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsAcceleratorCountOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsAcceleratorCountOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "max", GoGetter: "Max"},
			_jsii_.MemberProperty{JsiiProperty: "maxInput", GoGetter: "MaxInput"},
			_jsii_.MemberProperty{JsiiProperty: "min", GoGetter: "Min"},
			_jsii_.MemberProperty{JsiiProperty: "minInput", GoGetter: "MinInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMax", GoMethod: "ResetMax"},
			_jsii_.MemberMethod{JsiiMethod: "resetMin", GoMethod: "ResetMin"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsAcceleratorCountOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsAcceleratorTotalMemoryMiB",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsAcceleratorTotalMemoryMiB)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsAcceleratorTotalMemoryMiBOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsAcceleratorTotalMemoryMiBOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "max", GoGetter: "Max"},
			_jsii_.MemberProperty{JsiiProperty: "maxInput", GoGetter: "MaxInput"},
			_jsii_.MemberProperty{JsiiProperty: "min", GoGetter: "Min"},
			_jsii_.MemberProperty{JsiiProperty: "minInput", GoGetter: "MinInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMax", GoMethod: "ResetMax"},
			_jsii_.MemberMethod{JsiiMethod: "resetMin", GoMethod: "ResetMin"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsAcceleratorTotalMemoryMiBOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsBaselineEbsBandwidthMbps",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsBaselineEbsBandwidthMbps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsBaselineEbsBandwidthMbpsOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsBaselineEbsBandwidthMbpsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "max", GoGetter: "Max"},
			_jsii_.MemberProperty{JsiiProperty: "maxInput", GoGetter: "MaxInput"},
			_jsii_.MemberProperty{JsiiProperty: "min", GoGetter: "Min"},
			_jsii_.MemberProperty{JsiiProperty: "minInput", GoGetter: "MinInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMax", GoMethod: "ResetMax"},
			_jsii_.MemberMethod{JsiiMethod: "resetMin", GoMethod: "ResetMin"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsBaselineEbsBandwidthMbpsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsBaselinePerformanceFactors",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsBaselinePerformanceFactors)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsBaselinePerformanceFactorsCpu",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsBaselinePerformanceFactorsCpu)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "putReferences", GoMethod: "PutReferences"},
			_jsii_.MemberProperty{JsiiProperty: "references", GoGetter: "References"},
			_jsii_.MemberProperty{JsiiProperty: "referencesInput", GoGetter: "ReferencesInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetReferences", GoMethod: "ResetReferences"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsBaselinePerformanceFactorsCpuOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsBaselinePerformanceFactorsCpuReferences",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsBaselinePerformanceFactorsCpuReferences)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsBaselinePerformanceFactorsCpuReferencesList",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsBaselinePerformanceFactorsCpuReferencesList)(nil)).Elem(),
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
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsBaselinePerformanceFactorsCpuReferencesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsBaselinePerformanceFactorsCpuReferencesOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsBaselinePerformanceFactorsCpuReferencesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "instanceFamily", GoGetter: "InstanceFamily"},
			_jsii_.MemberProperty{JsiiProperty: "instanceFamilyInput", GoGetter: "InstanceFamilyInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetInstanceFamily", GoMethod: "ResetInstanceFamily"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsBaselinePerformanceFactorsCpuReferencesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsBaselinePerformanceFactorsOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsBaselinePerformanceFactorsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "cpu", GoGetter: "Cpu"},
			_jsii_.MemberProperty{JsiiProperty: "cpuInput", GoGetter: "CpuInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putCpu", GoMethod: "PutCpu"},
			_jsii_.MemberMethod{JsiiMethod: "resetCpu", GoMethod: "ResetCpu"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsBaselinePerformanceFactorsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsMemoryGiBPerVCpu",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsMemoryGiBPerVCpu)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsMemoryGiBPerVCpuOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsMemoryGiBPerVCpuOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "max", GoGetter: "Max"},
			_jsii_.MemberProperty{JsiiProperty: "maxInput", GoGetter: "MaxInput"},
			_jsii_.MemberProperty{JsiiProperty: "min", GoGetter: "Min"},
			_jsii_.MemberProperty{JsiiProperty: "minInput", GoGetter: "MinInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMax", GoMethod: "ResetMax"},
			_jsii_.MemberMethod{JsiiMethod: "resetMin", GoMethod: "ResetMin"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsMemoryGiBPerVCpuOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsMemoryMiB",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsMemoryMiB)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsMemoryMiBOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsMemoryMiBOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "max", GoGetter: "Max"},
			_jsii_.MemberProperty{JsiiProperty: "maxInput", GoGetter: "MaxInput"},
			_jsii_.MemberProperty{JsiiProperty: "min", GoGetter: "Min"},
			_jsii_.MemberProperty{JsiiProperty: "minInput", GoGetter: "MinInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMax", GoMethod: "ResetMax"},
			_jsii_.MemberMethod{JsiiMethod: "resetMin", GoMethod: "ResetMin"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsMemoryMiBOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsNetworkBandwidthGbps",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsNetworkBandwidthGbps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsNetworkBandwidthGbpsOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsNetworkBandwidthGbpsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "max", GoGetter: "Max"},
			_jsii_.MemberProperty{JsiiProperty: "maxInput", GoGetter: "MaxInput"},
			_jsii_.MemberProperty{JsiiProperty: "min", GoGetter: "Min"},
			_jsii_.MemberProperty{JsiiProperty: "minInput", GoGetter: "MinInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMax", GoMethod: "ResetMax"},
			_jsii_.MemberMethod{JsiiMethod: "resetMin", GoMethod: "ResetMin"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsNetworkBandwidthGbpsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsNetworkInterfaceCount",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsNetworkInterfaceCount)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsNetworkInterfaceCountOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsNetworkInterfaceCountOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "max", GoGetter: "Max"},
			_jsii_.MemberProperty{JsiiProperty: "maxInput", GoGetter: "MaxInput"},
			_jsii_.MemberProperty{JsiiProperty: "min", GoGetter: "Min"},
			_jsii_.MemberProperty{JsiiProperty: "minInput", GoGetter: "MinInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMax", GoMethod: "ResetMax"},
			_jsii_.MemberMethod{JsiiMethod: "resetMin", GoMethod: "ResetMin"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsNetworkInterfaceCountOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "acceleratorCount", GoGetter: "AcceleratorCount"},
			_jsii_.MemberProperty{JsiiProperty: "acceleratorCountInput", GoGetter: "AcceleratorCountInput"},
			_jsii_.MemberProperty{JsiiProperty: "acceleratorManufacturers", GoGetter: "AcceleratorManufacturers"},
			_jsii_.MemberProperty{JsiiProperty: "acceleratorManufacturersInput", GoGetter: "AcceleratorManufacturersInput"},
			_jsii_.MemberProperty{JsiiProperty: "acceleratorNames", GoGetter: "AcceleratorNames"},
			_jsii_.MemberProperty{JsiiProperty: "acceleratorNamesInput", GoGetter: "AcceleratorNamesInput"},
			_jsii_.MemberProperty{JsiiProperty: "acceleratorTotalMemoryMiB", GoGetter: "AcceleratorTotalMemoryMiB"},
			_jsii_.MemberProperty{JsiiProperty: "acceleratorTotalMemoryMiBInput", GoGetter: "AcceleratorTotalMemoryMiBInput"},
			_jsii_.MemberProperty{JsiiProperty: "acceleratorTypes", GoGetter: "AcceleratorTypes"},
			_jsii_.MemberProperty{JsiiProperty: "acceleratorTypesInput", GoGetter: "AcceleratorTypesInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowedInstanceTypes", GoGetter: "AllowedInstanceTypes"},
			_jsii_.MemberProperty{JsiiProperty: "allowedInstanceTypesInput", GoGetter: "AllowedInstanceTypesInput"},
			_jsii_.MemberProperty{JsiiProperty: "bareMetal", GoGetter: "BareMetal"},
			_jsii_.MemberProperty{JsiiProperty: "bareMetalInput", GoGetter: "BareMetalInput"},
			_jsii_.MemberProperty{JsiiProperty: "baselineEbsBandwidthMbps", GoGetter: "BaselineEbsBandwidthMbps"},
			_jsii_.MemberProperty{JsiiProperty: "baselineEbsBandwidthMbpsInput", GoGetter: "BaselineEbsBandwidthMbpsInput"},
			_jsii_.MemberProperty{JsiiProperty: "baselinePerformanceFactors", GoGetter: "BaselinePerformanceFactors"},
			_jsii_.MemberProperty{JsiiProperty: "baselinePerformanceFactorsInput", GoGetter: "BaselinePerformanceFactorsInput"},
			_jsii_.MemberProperty{JsiiProperty: "burstablePerformance", GoGetter: "BurstablePerformance"},
			_jsii_.MemberProperty{JsiiProperty: "burstablePerformanceInput", GoGetter: "BurstablePerformanceInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "cpuManufacturers", GoGetter: "CpuManufacturers"},
			_jsii_.MemberProperty{JsiiProperty: "cpuManufacturersInput", GoGetter: "CpuManufacturersInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "excludedInstanceTypes", GoGetter: "ExcludedInstanceTypes"},
			_jsii_.MemberProperty{JsiiProperty: "excludedInstanceTypesInput", GoGetter: "ExcludedInstanceTypesInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "instanceGenerations", GoGetter: "InstanceGenerations"},
			_jsii_.MemberProperty{JsiiProperty: "instanceGenerationsInput", GoGetter: "InstanceGenerationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "localStorage", GoGetter: "LocalStorage"},
			_jsii_.MemberProperty{JsiiProperty: "localStorageInput", GoGetter: "LocalStorageInput"},
			_jsii_.MemberProperty{JsiiProperty: "localStorageTypes", GoGetter: "LocalStorageTypes"},
			_jsii_.MemberProperty{JsiiProperty: "localStorageTypesInput", GoGetter: "LocalStorageTypesInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxSpotPriceAsPercentageOfOptimalOnDemandPrice", GoGetter: "MaxSpotPriceAsPercentageOfOptimalOnDemandPrice"},
			_jsii_.MemberProperty{JsiiProperty: "maxSpotPriceAsPercentageOfOptimalOnDemandPriceInput", GoGetter: "MaxSpotPriceAsPercentageOfOptimalOnDemandPriceInput"},
			_jsii_.MemberProperty{JsiiProperty: "memoryGiBPerVCpu", GoGetter: "MemoryGiBPerVCpu"},
			_jsii_.MemberProperty{JsiiProperty: "memoryGiBPerVCpuInput", GoGetter: "MemoryGiBPerVCpuInput"},
			_jsii_.MemberProperty{JsiiProperty: "memoryMiB", GoGetter: "MemoryMiB"},
			_jsii_.MemberProperty{JsiiProperty: "memoryMiBInput", GoGetter: "MemoryMiBInput"},
			_jsii_.MemberProperty{JsiiProperty: "networkBandwidthGbps", GoGetter: "NetworkBandwidthGbps"},
			_jsii_.MemberProperty{JsiiProperty: "networkBandwidthGbpsInput", GoGetter: "NetworkBandwidthGbpsInput"},
			_jsii_.MemberProperty{JsiiProperty: "networkInterfaceCount", GoGetter: "NetworkInterfaceCount"},
			_jsii_.MemberProperty{JsiiProperty: "networkInterfaceCountInput", GoGetter: "NetworkInterfaceCountInput"},
			_jsii_.MemberProperty{JsiiProperty: "onDemandMaxPricePercentageOverLowestPrice", GoGetter: "OnDemandMaxPricePercentageOverLowestPrice"},
			_jsii_.MemberProperty{JsiiProperty: "onDemandMaxPricePercentageOverLowestPriceInput", GoGetter: "OnDemandMaxPricePercentageOverLowestPriceInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAcceleratorCount", GoMethod: "PutAcceleratorCount"},
			_jsii_.MemberMethod{JsiiMethod: "putAcceleratorTotalMemoryMiB", GoMethod: "PutAcceleratorTotalMemoryMiB"},
			_jsii_.MemberMethod{JsiiMethod: "putBaselineEbsBandwidthMbps", GoMethod: "PutBaselineEbsBandwidthMbps"},
			_jsii_.MemberMethod{JsiiMethod: "putBaselinePerformanceFactors", GoMethod: "PutBaselinePerformanceFactors"},
			_jsii_.MemberMethod{JsiiMethod: "putMemoryGiBPerVCpu", GoMethod: "PutMemoryGiBPerVCpu"},
			_jsii_.MemberMethod{JsiiMethod: "putMemoryMiB", GoMethod: "PutMemoryMiB"},
			_jsii_.MemberMethod{JsiiMethod: "putNetworkBandwidthGbps", GoMethod: "PutNetworkBandwidthGbps"},
			_jsii_.MemberMethod{JsiiMethod: "putNetworkInterfaceCount", GoMethod: "PutNetworkInterfaceCount"},
			_jsii_.MemberMethod{JsiiMethod: "putTotalLocalStorageGb", GoMethod: "PutTotalLocalStorageGb"},
			_jsii_.MemberMethod{JsiiMethod: "putVCpuCount", GoMethod: "PutVCpuCount"},
			_jsii_.MemberProperty{JsiiProperty: "requireHibernateSupport", GoGetter: "RequireHibernateSupport"},
			_jsii_.MemberProperty{JsiiProperty: "requireHibernateSupportInput", GoGetter: "RequireHibernateSupportInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAcceleratorCount", GoMethod: "ResetAcceleratorCount"},
			_jsii_.MemberMethod{JsiiMethod: "resetAcceleratorManufacturers", GoMethod: "ResetAcceleratorManufacturers"},
			_jsii_.MemberMethod{JsiiMethod: "resetAcceleratorNames", GoMethod: "ResetAcceleratorNames"},
			_jsii_.MemberMethod{JsiiMethod: "resetAcceleratorTotalMemoryMiB", GoMethod: "ResetAcceleratorTotalMemoryMiB"},
			_jsii_.MemberMethod{JsiiMethod: "resetAcceleratorTypes", GoMethod: "ResetAcceleratorTypes"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedInstanceTypes", GoMethod: "ResetAllowedInstanceTypes"},
			_jsii_.MemberMethod{JsiiMethod: "resetBareMetal", GoMethod: "ResetBareMetal"},
			_jsii_.MemberMethod{JsiiMethod: "resetBaselineEbsBandwidthMbps", GoMethod: "ResetBaselineEbsBandwidthMbps"},
			_jsii_.MemberMethod{JsiiMethod: "resetBaselinePerformanceFactors", GoMethod: "ResetBaselinePerformanceFactors"},
			_jsii_.MemberMethod{JsiiMethod: "resetBurstablePerformance", GoMethod: "ResetBurstablePerformance"},
			_jsii_.MemberMethod{JsiiMethod: "resetCpuManufacturers", GoMethod: "ResetCpuManufacturers"},
			_jsii_.MemberMethod{JsiiMethod: "resetExcludedInstanceTypes", GoMethod: "ResetExcludedInstanceTypes"},
			_jsii_.MemberMethod{JsiiMethod: "resetInstanceGenerations", GoMethod: "ResetInstanceGenerations"},
			_jsii_.MemberMethod{JsiiMethod: "resetLocalStorage", GoMethod: "ResetLocalStorage"},
			_jsii_.MemberMethod{JsiiMethod: "resetLocalStorageTypes", GoMethod: "ResetLocalStorageTypes"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxSpotPriceAsPercentageOfOptimalOnDemandPrice", GoMethod: "ResetMaxSpotPriceAsPercentageOfOptimalOnDemandPrice"},
			_jsii_.MemberMethod{JsiiMethod: "resetMemoryGiBPerVCpu", GoMethod: "ResetMemoryGiBPerVCpu"},
			_jsii_.MemberMethod{JsiiMethod: "resetMemoryMiB", GoMethod: "ResetMemoryMiB"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetworkBandwidthGbps", GoMethod: "ResetNetworkBandwidthGbps"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetworkInterfaceCount", GoMethod: "ResetNetworkInterfaceCount"},
			_jsii_.MemberMethod{JsiiMethod: "resetOnDemandMaxPricePercentageOverLowestPrice", GoMethod: "ResetOnDemandMaxPricePercentageOverLowestPrice"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequireHibernateSupport", GoMethod: "ResetRequireHibernateSupport"},
			_jsii_.MemberMethod{JsiiMethod: "resetSpotMaxPricePercentageOverLowestPrice", GoMethod: "ResetSpotMaxPricePercentageOverLowestPrice"},
			_jsii_.MemberMethod{JsiiMethod: "resetTotalLocalStorageGb", GoMethod: "ResetTotalLocalStorageGb"},
			_jsii_.MemberMethod{JsiiMethod: "resetVCpuCount", GoMethod: "ResetVCpuCount"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "spotMaxPricePercentageOverLowestPrice", GoGetter: "SpotMaxPricePercentageOverLowestPrice"},
			_jsii_.MemberProperty{JsiiProperty: "spotMaxPricePercentageOverLowestPriceInput", GoGetter: "SpotMaxPricePercentageOverLowestPriceInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "totalLocalStorageGb", GoGetter: "TotalLocalStorageGb"},
			_jsii_.MemberProperty{JsiiProperty: "totalLocalStorageGbInput", GoGetter: "TotalLocalStorageGbInput"},
			_jsii_.MemberProperty{JsiiProperty: "vCpuCount", GoGetter: "VCpuCount"},
			_jsii_.MemberProperty{JsiiProperty: "vCpuCountInput", GoGetter: "VCpuCountInput"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsTotalLocalStorageGb",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsTotalLocalStorageGb)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsTotalLocalStorageGbOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsTotalLocalStorageGbOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "max", GoGetter: "Max"},
			_jsii_.MemberProperty{JsiiProperty: "maxInput", GoGetter: "MaxInput"},
			_jsii_.MemberProperty{JsiiProperty: "min", GoGetter: "Min"},
			_jsii_.MemberProperty{JsiiProperty: "minInput", GoGetter: "MinInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMax", GoMethod: "ResetMax"},
			_jsii_.MemberMethod{JsiiMethod: "resetMin", GoMethod: "ResetMin"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsTotalLocalStorageGbOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsVCpuCount",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsVCpuCount)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsVCpuCountOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsVCpuCountOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "max", GoGetter: "Max"},
			_jsii_.MemberProperty{JsiiProperty: "maxInput", GoGetter: "MaxInput"},
			_jsii_.MemberProperty{JsiiProperty: "min", GoGetter: "Min"},
			_jsii_.MemberProperty{JsiiProperty: "minInput", GoGetter: "MinInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMax", GoMethod: "ResetMax"},
			_jsii_.MemberMethod{JsiiMethod: "resetMin", GoMethod: "ResetMin"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceRequirementsVCpuCountOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataLicenseSpecifications",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataLicenseSpecifications)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataLicenseSpecificationsList",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataLicenseSpecificationsList)(nil)).Elem(),
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
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataLicenseSpecificationsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataLicenseSpecificationsOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataLicenseSpecificationsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "licenseConfigurationArn", GoGetter: "LicenseConfigurationArn"},
			_jsii_.MemberProperty{JsiiProperty: "licenseConfigurationArnInput", GoGetter: "LicenseConfigurationArnInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetLicenseConfigurationArn", GoMethod: "ResetLicenseConfigurationArn"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataLicenseSpecificationsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataMaintenanceOptions",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataMaintenanceOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataMaintenanceOptionsOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataMaintenanceOptionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "autoRecovery", GoGetter: "AutoRecovery"},
			_jsii_.MemberProperty{JsiiProperty: "autoRecoveryInput", GoGetter: "AutoRecoveryInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAutoRecovery", GoMethod: "ResetAutoRecovery"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataMaintenanceOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataMetadataOptions",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataMetadataOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataMetadataOptionsOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataMetadataOptionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "httpEndpoint", GoGetter: "HttpEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "httpEndpointInput", GoGetter: "HttpEndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "httpProtocolIpv6", GoGetter: "HttpProtocolIpv6"},
			_jsii_.MemberProperty{JsiiProperty: "httpProtocolIpv6Input", GoGetter: "HttpProtocolIpv6Input"},
			_jsii_.MemberProperty{JsiiProperty: "httpPutResponseHopLimit", GoGetter: "HttpPutResponseHopLimit"},
			_jsii_.MemberProperty{JsiiProperty: "httpPutResponseHopLimitInput", GoGetter: "HttpPutResponseHopLimitInput"},
			_jsii_.MemberProperty{JsiiProperty: "httpTokens", GoGetter: "HttpTokens"},
			_jsii_.MemberProperty{JsiiProperty: "httpTokensInput", GoGetter: "HttpTokensInput"},
			_jsii_.MemberProperty{JsiiProperty: "instanceMetadataTags", GoGetter: "InstanceMetadataTags"},
			_jsii_.MemberProperty{JsiiProperty: "instanceMetadataTagsInput", GoGetter: "InstanceMetadataTagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpEndpoint", GoMethod: "ResetHttpEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpProtocolIpv6", GoMethod: "ResetHttpProtocolIpv6"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpPutResponseHopLimit", GoMethod: "ResetHttpPutResponseHopLimit"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpTokens", GoMethod: "ResetHttpTokens"},
			_jsii_.MemberMethod{JsiiMethod: "resetInstanceMetadataTags", GoMethod: "ResetInstanceMetadataTags"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataMetadataOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataMonitoring",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataMonitoring)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataMonitoringOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataMonitoringOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataMonitoringOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataNetworkInterfaces",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataNetworkInterfaces)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesConnectionTrackingSpecification",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesConnectionTrackingSpecification)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesConnectionTrackingSpecificationOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesConnectionTrackingSpecificationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetTcpEstablishedTimeout", GoMethod: "ResetTcpEstablishedTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetUdpStreamTimeout", GoMethod: "ResetUdpStreamTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetUdpTimeout", GoMethod: "ResetUdpTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "tcpEstablishedTimeout", GoGetter: "TcpEstablishedTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "tcpEstablishedTimeoutInput", GoGetter: "TcpEstablishedTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "udpStreamTimeout", GoGetter: "UdpStreamTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "udpStreamTimeoutInput", GoGetter: "UdpStreamTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "udpTimeout", GoGetter: "UdpTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "udpTimeoutInput", GoGetter: "UdpTimeoutInput"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesConnectionTrackingSpecificationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesEnaSrdSpecification",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesEnaSrdSpecification)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesEnaSrdSpecificationEnaSrdUdpSpecification",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesEnaSrdSpecificationEnaSrdUdpSpecification)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesEnaSrdSpecificationEnaSrdUdpSpecificationOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesEnaSrdSpecificationEnaSrdUdpSpecificationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enaSrdUdpEnabled", GoGetter: "EnaSrdUdpEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "enaSrdUdpEnabledInput", GoGetter: "EnaSrdUdpEnabledInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetEnaSrdUdpEnabled", GoMethod: "ResetEnaSrdUdpEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesEnaSrdSpecificationEnaSrdUdpSpecificationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesEnaSrdSpecificationOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesEnaSrdSpecificationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enaSrdEnabled", GoGetter: "EnaSrdEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "enaSrdEnabledInput", GoGetter: "EnaSrdEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "enaSrdUdpSpecification", GoGetter: "EnaSrdUdpSpecification"},
			_jsii_.MemberProperty{JsiiProperty: "enaSrdUdpSpecificationInput", GoGetter: "EnaSrdUdpSpecificationInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putEnaSrdUdpSpecification", GoMethod: "PutEnaSrdUdpSpecification"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnaSrdEnabled", GoMethod: "ResetEnaSrdEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnaSrdUdpSpecification", GoMethod: "ResetEnaSrdUdpSpecification"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesEnaSrdSpecificationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv4Prefixes",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv4Prefixes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv4PrefixesList",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv4PrefixesList)(nil)).Elem(),
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
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv4PrefixesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv4PrefixesOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv4PrefixesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "ipv4Prefix", GoGetter: "Ipv4Prefix"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4PrefixInput", GoGetter: "Ipv4PrefixInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpv4Prefix", GoMethod: "ResetIpv4Prefix"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv4PrefixesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv6Addresses",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv6Addresses)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv6AddressesList",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv6AddressesList)(nil)).Elem(),
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
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv6AddressesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv6AddressesOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv6AddressesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "ipv6Address", GoGetter: "Ipv6Address"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6AddressInput", GoGetter: "Ipv6AddressInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpv6Address", GoMethod: "ResetIpv6Address"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv6AddressesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv6Prefixes",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv6Prefixes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv6PrefixesList",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv6PrefixesList)(nil)).Elem(),
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
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv6PrefixesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv6PrefixesOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv6PrefixesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "ipv6Prefix", GoGetter: "Ipv6Prefix"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6PrefixInput", GoGetter: "Ipv6PrefixInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpv6Prefix", GoMethod: "ResetIpv6Prefix"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv6PrefixesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesList",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesList)(nil)).Elem(),
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
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "associateCarrierIpAddress", GoGetter: "AssociateCarrierIpAddress"},
			_jsii_.MemberProperty{JsiiProperty: "associateCarrierIpAddressInput", GoGetter: "AssociateCarrierIpAddressInput"},
			_jsii_.MemberProperty{JsiiProperty: "associatePublicIpAddress", GoGetter: "AssociatePublicIpAddress"},
			_jsii_.MemberProperty{JsiiProperty: "associatePublicIpAddressInput", GoGetter: "AssociatePublicIpAddressInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "connectionTrackingSpecification", GoGetter: "ConnectionTrackingSpecification"},
			_jsii_.MemberProperty{JsiiProperty: "connectionTrackingSpecificationInput", GoGetter: "ConnectionTrackingSpecificationInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "deleteOnTermination", GoGetter: "DeleteOnTermination"},
			_jsii_.MemberProperty{JsiiProperty: "deleteOnTerminationInput", GoGetter: "DeleteOnTerminationInput"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "deviceIndex", GoGetter: "DeviceIndex"},
			_jsii_.MemberProperty{JsiiProperty: "deviceIndexInput", GoGetter: "DeviceIndexInput"},
			_jsii_.MemberProperty{JsiiProperty: "enaSrdSpecification", GoGetter: "EnaSrdSpecification"},
			_jsii_.MemberProperty{JsiiProperty: "enaSrdSpecificationInput", GoGetter: "EnaSrdSpecificationInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "groups", GoGetter: "Groups"},
			_jsii_.MemberProperty{JsiiProperty: "groupsInput", GoGetter: "GroupsInput"},
			_jsii_.MemberProperty{JsiiProperty: "interfaceType", GoGetter: "InterfaceType"},
			_jsii_.MemberProperty{JsiiProperty: "interfaceTypeInput", GoGetter: "InterfaceTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4PrefixCount", GoGetter: "Ipv4PrefixCount"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4PrefixCountInput", GoGetter: "Ipv4PrefixCountInput"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4Prefixes", GoGetter: "Ipv4Prefixes"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4PrefixesInput", GoGetter: "Ipv4PrefixesInput"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6AddressCount", GoGetter: "Ipv6AddressCount"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6AddressCountInput", GoGetter: "Ipv6AddressCountInput"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6Addresses", GoGetter: "Ipv6Addresses"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6AddressesInput", GoGetter: "Ipv6AddressesInput"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6PrefixCount", GoGetter: "Ipv6PrefixCount"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6PrefixCountInput", GoGetter: "Ipv6PrefixCountInput"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6Prefixes", GoGetter: "Ipv6Prefixes"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6PrefixesInput", GoGetter: "Ipv6PrefixesInput"},
			_jsii_.MemberProperty{JsiiProperty: "networkCardIndex", GoGetter: "NetworkCardIndex"},
			_jsii_.MemberProperty{JsiiProperty: "networkCardIndexInput", GoGetter: "NetworkCardIndexInput"},
			_jsii_.MemberProperty{JsiiProperty: "networkInterfaceId", GoGetter: "NetworkInterfaceId"},
			_jsii_.MemberProperty{JsiiProperty: "networkInterfaceIdInput", GoGetter: "NetworkInterfaceIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "primaryIpv6", GoGetter: "PrimaryIpv6"},
			_jsii_.MemberProperty{JsiiProperty: "primaryIpv6Input", GoGetter: "PrimaryIpv6Input"},
			_jsii_.MemberProperty{JsiiProperty: "privateIpAddress", GoGetter: "PrivateIpAddress"},
			_jsii_.MemberProperty{JsiiProperty: "privateIpAddresses", GoGetter: "PrivateIpAddresses"},
			_jsii_.MemberProperty{JsiiProperty: "privateIpAddressesInput", GoGetter: "PrivateIpAddressesInput"},
			_jsii_.MemberProperty{JsiiProperty: "privateIpAddressInput", GoGetter: "PrivateIpAddressInput"},
			_jsii_.MemberMethod{JsiiMethod: "putConnectionTrackingSpecification", GoMethod: "PutConnectionTrackingSpecification"},
			_jsii_.MemberMethod{JsiiMethod: "putEnaSrdSpecification", GoMethod: "PutEnaSrdSpecification"},
			_jsii_.MemberMethod{JsiiMethod: "putIpv4Prefixes", GoMethod: "PutIpv4Prefixes"},
			_jsii_.MemberMethod{JsiiMethod: "putIpv6Addresses", GoMethod: "PutIpv6Addresses"},
			_jsii_.MemberMethod{JsiiMethod: "putIpv6Prefixes", GoMethod: "PutIpv6Prefixes"},
			_jsii_.MemberMethod{JsiiMethod: "putPrivateIpAddresses", GoMethod: "PutPrivateIpAddresses"},
			_jsii_.MemberMethod{JsiiMethod: "resetAssociateCarrierIpAddress", GoMethod: "ResetAssociateCarrierIpAddress"},
			_jsii_.MemberMethod{JsiiMethod: "resetAssociatePublicIpAddress", GoMethod: "ResetAssociatePublicIpAddress"},
			_jsii_.MemberMethod{JsiiMethod: "resetConnectionTrackingSpecification", GoMethod: "ResetConnectionTrackingSpecification"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeleteOnTermination", GoMethod: "ResetDeleteOnTermination"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeviceIndex", GoMethod: "ResetDeviceIndex"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnaSrdSpecification", GoMethod: "ResetEnaSrdSpecification"},
			_jsii_.MemberMethod{JsiiMethod: "resetGroups", GoMethod: "ResetGroups"},
			_jsii_.MemberMethod{JsiiMethod: "resetInterfaceType", GoMethod: "ResetInterfaceType"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpv4PrefixCount", GoMethod: "ResetIpv4PrefixCount"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpv4Prefixes", GoMethod: "ResetIpv4Prefixes"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpv6AddressCount", GoMethod: "ResetIpv6AddressCount"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpv6Addresses", GoMethod: "ResetIpv6Addresses"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpv6PrefixCount", GoMethod: "ResetIpv6PrefixCount"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpv6Prefixes", GoMethod: "ResetIpv6Prefixes"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetworkCardIndex", GoMethod: "ResetNetworkCardIndex"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetworkInterfaceId", GoMethod: "ResetNetworkInterfaceId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrimaryIpv6", GoMethod: "ResetPrimaryIpv6"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrivateIpAddress", GoMethod: "ResetPrivateIpAddress"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrivateIpAddresses", GoMethod: "ResetPrivateIpAddresses"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecondaryPrivateIpAddressCount", GoMethod: "ResetSecondaryPrivateIpAddressCount"},
			_jsii_.MemberMethod{JsiiMethod: "resetSubnetId", GoMethod: "ResetSubnetId"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryPrivateIpAddressCount", GoGetter: "SecondaryPrivateIpAddressCount"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryPrivateIpAddressCountInput", GoGetter: "SecondaryPrivateIpAddressCountInput"},
			_jsii_.MemberProperty{JsiiProperty: "subnetId", GoGetter: "SubnetId"},
			_jsii_.MemberProperty{JsiiProperty: "subnetIdInput", GoGetter: "SubnetIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesPrivateIpAddresses",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesPrivateIpAddresses)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesPrivateIpAddressesList",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesPrivateIpAddressesList)(nil)).Elem(),
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
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesPrivateIpAddressesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesPrivateIpAddressesOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesPrivateIpAddressesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "primary", GoGetter: "Primary"},
			_jsii_.MemberProperty{JsiiProperty: "primaryInput", GoGetter: "PrimaryInput"},
			_jsii_.MemberProperty{JsiiProperty: "privateIpAddress", GoGetter: "PrivateIpAddress"},
			_jsii_.MemberProperty{JsiiProperty: "privateIpAddressInput", GoGetter: "PrivateIpAddressInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrimary", GoMethod: "ResetPrimary"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrivateIpAddress", GoMethod: "ResetPrivateIpAddress"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataNetworkInterfacesPrivateIpAddressesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "blockDeviceMappings", GoGetter: "BlockDeviceMappings"},
			_jsii_.MemberProperty{JsiiProperty: "blockDeviceMappingsInput", GoGetter: "BlockDeviceMappingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "capacityReservationSpecification", GoGetter: "CapacityReservationSpecification"},
			_jsii_.MemberProperty{JsiiProperty: "capacityReservationSpecificationInput", GoGetter: "CapacityReservationSpecificationInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "cpuOptions", GoGetter: "CpuOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cpuOptionsInput", GoGetter: "CpuOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "creditSpecification", GoGetter: "CreditSpecification"},
			_jsii_.MemberProperty{JsiiProperty: "creditSpecificationInput", GoGetter: "CreditSpecificationInput"},
			_jsii_.MemberProperty{JsiiProperty: "disableApiStop", GoGetter: "DisableApiStop"},
			_jsii_.MemberProperty{JsiiProperty: "disableApiStopInput", GoGetter: "DisableApiStopInput"},
			_jsii_.MemberProperty{JsiiProperty: "disableApiTermination", GoGetter: "DisableApiTermination"},
			_jsii_.MemberProperty{JsiiProperty: "disableApiTerminationInput", GoGetter: "DisableApiTerminationInput"},
			_jsii_.MemberProperty{JsiiProperty: "ebsOptimized", GoGetter: "EbsOptimized"},
			_jsii_.MemberProperty{JsiiProperty: "ebsOptimizedInput", GoGetter: "EbsOptimizedInput"},
			_jsii_.MemberProperty{JsiiProperty: "elasticGpuSpecifications", GoGetter: "ElasticGpuSpecifications"},
			_jsii_.MemberProperty{JsiiProperty: "elasticGpuSpecificationsInput", GoGetter: "ElasticGpuSpecificationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "elasticInferenceAccelerators", GoGetter: "ElasticInferenceAccelerators"},
			_jsii_.MemberProperty{JsiiProperty: "elasticInferenceAcceleratorsInput", GoGetter: "ElasticInferenceAcceleratorsInput"},
			_jsii_.MemberProperty{JsiiProperty: "enclaveOptions", GoGetter: "EnclaveOptions"},
			_jsii_.MemberProperty{JsiiProperty: "enclaveOptionsInput", GoGetter: "EnclaveOptionsInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "hibernationOptions", GoGetter: "HibernationOptions"},
			_jsii_.MemberProperty{JsiiProperty: "hibernationOptionsInput", GoGetter: "HibernationOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "iamInstanceProfile", GoGetter: "IamInstanceProfile"},
			_jsii_.MemberProperty{JsiiProperty: "iamInstanceProfileInput", GoGetter: "IamInstanceProfileInput"},
			_jsii_.MemberProperty{JsiiProperty: "imageId", GoGetter: "ImageId"},
			_jsii_.MemberProperty{JsiiProperty: "imageIdInput", GoGetter: "ImageIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "instanceInitiatedShutdownBehavior", GoGetter: "InstanceInitiatedShutdownBehavior"},
			_jsii_.MemberProperty{JsiiProperty: "instanceInitiatedShutdownBehaviorInput", GoGetter: "InstanceInitiatedShutdownBehaviorInput"},
			_jsii_.MemberProperty{JsiiProperty: "instanceMarketOptions", GoGetter: "InstanceMarketOptions"},
			_jsii_.MemberProperty{JsiiProperty: "instanceMarketOptionsInput", GoGetter: "InstanceMarketOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "instanceRequirements", GoGetter: "InstanceRequirements"},
			_jsii_.MemberProperty{JsiiProperty: "instanceRequirementsInput", GoGetter: "InstanceRequirementsInput"},
			_jsii_.MemberProperty{JsiiProperty: "instanceType", GoGetter: "InstanceType"},
			_jsii_.MemberProperty{JsiiProperty: "instanceTypeInput", GoGetter: "InstanceTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "kernelId", GoGetter: "KernelId"},
			_jsii_.MemberProperty{JsiiProperty: "kernelIdInput", GoGetter: "KernelIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "keyName", GoGetter: "KeyName"},
			_jsii_.MemberProperty{JsiiProperty: "keyNameInput", GoGetter: "KeyNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "licenseSpecifications", GoGetter: "LicenseSpecifications"},
			_jsii_.MemberProperty{JsiiProperty: "licenseSpecificationsInput", GoGetter: "LicenseSpecificationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "maintenanceOptions", GoGetter: "MaintenanceOptions"},
			_jsii_.MemberProperty{JsiiProperty: "maintenanceOptionsInput", GoGetter: "MaintenanceOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "metadataOptions", GoGetter: "MetadataOptions"},
			_jsii_.MemberProperty{JsiiProperty: "metadataOptionsInput", GoGetter: "MetadataOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "monitoring", GoGetter: "Monitoring"},
			_jsii_.MemberProperty{JsiiProperty: "monitoringInput", GoGetter: "MonitoringInput"},
			_jsii_.MemberProperty{JsiiProperty: "networkInterfaces", GoGetter: "NetworkInterfaces"},
			_jsii_.MemberProperty{JsiiProperty: "networkInterfacesInput", GoGetter: "NetworkInterfacesInput"},
			_jsii_.MemberProperty{JsiiProperty: "placement", GoGetter: "Placement"},
			_jsii_.MemberProperty{JsiiProperty: "placementInput", GoGetter: "PlacementInput"},
			_jsii_.MemberProperty{JsiiProperty: "privateDnsNameOptions", GoGetter: "PrivateDnsNameOptions"},
			_jsii_.MemberProperty{JsiiProperty: "privateDnsNameOptionsInput", GoGetter: "PrivateDnsNameOptionsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putBlockDeviceMappings", GoMethod: "PutBlockDeviceMappings"},
			_jsii_.MemberMethod{JsiiMethod: "putCapacityReservationSpecification", GoMethod: "PutCapacityReservationSpecification"},
			_jsii_.MemberMethod{JsiiMethod: "putCpuOptions", GoMethod: "PutCpuOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putCreditSpecification", GoMethod: "PutCreditSpecification"},
			_jsii_.MemberMethod{JsiiMethod: "putElasticGpuSpecifications", GoMethod: "PutElasticGpuSpecifications"},
			_jsii_.MemberMethod{JsiiMethod: "putElasticInferenceAccelerators", GoMethod: "PutElasticInferenceAccelerators"},
			_jsii_.MemberMethod{JsiiMethod: "putEnclaveOptions", GoMethod: "PutEnclaveOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putHibernationOptions", GoMethod: "PutHibernationOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putIamInstanceProfile", GoMethod: "PutIamInstanceProfile"},
			_jsii_.MemberMethod{JsiiMethod: "putInstanceMarketOptions", GoMethod: "PutInstanceMarketOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putInstanceRequirements", GoMethod: "PutInstanceRequirements"},
			_jsii_.MemberMethod{JsiiMethod: "putLicenseSpecifications", GoMethod: "PutLicenseSpecifications"},
			_jsii_.MemberMethod{JsiiMethod: "putMaintenanceOptions", GoMethod: "PutMaintenanceOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putMetadataOptions", GoMethod: "PutMetadataOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putMonitoring", GoMethod: "PutMonitoring"},
			_jsii_.MemberMethod{JsiiMethod: "putNetworkInterfaces", GoMethod: "PutNetworkInterfaces"},
			_jsii_.MemberMethod{JsiiMethod: "putPlacement", GoMethod: "PutPlacement"},
			_jsii_.MemberMethod{JsiiMethod: "putPrivateDnsNameOptions", GoMethod: "PutPrivateDnsNameOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putTagSpecifications", GoMethod: "PutTagSpecifications"},
			_jsii_.MemberProperty{JsiiProperty: "ramDiskId", GoGetter: "RamDiskId"},
			_jsii_.MemberProperty{JsiiProperty: "ramDiskIdInput", GoGetter: "RamDiskIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetBlockDeviceMappings", GoMethod: "ResetBlockDeviceMappings"},
			_jsii_.MemberMethod{JsiiMethod: "resetCapacityReservationSpecification", GoMethod: "ResetCapacityReservationSpecification"},
			_jsii_.MemberMethod{JsiiMethod: "resetCpuOptions", GoMethod: "ResetCpuOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreditSpecification", GoMethod: "ResetCreditSpecification"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisableApiStop", GoMethod: "ResetDisableApiStop"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisableApiTermination", GoMethod: "ResetDisableApiTermination"},
			_jsii_.MemberMethod{JsiiMethod: "resetEbsOptimized", GoMethod: "ResetEbsOptimized"},
			_jsii_.MemberMethod{JsiiMethod: "resetElasticGpuSpecifications", GoMethod: "ResetElasticGpuSpecifications"},
			_jsii_.MemberMethod{JsiiMethod: "resetElasticInferenceAccelerators", GoMethod: "ResetElasticInferenceAccelerators"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnclaveOptions", GoMethod: "ResetEnclaveOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetHibernationOptions", GoMethod: "ResetHibernationOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetIamInstanceProfile", GoMethod: "ResetIamInstanceProfile"},
			_jsii_.MemberMethod{JsiiMethod: "resetImageId", GoMethod: "ResetImageId"},
			_jsii_.MemberMethod{JsiiMethod: "resetInstanceInitiatedShutdownBehavior", GoMethod: "ResetInstanceInitiatedShutdownBehavior"},
			_jsii_.MemberMethod{JsiiMethod: "resetInstanceMarketOptions", GoMethod: "ResetInstanceMarketOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetInstanceRequirements", GoMethod: "ResetInstanceRequirements"},
			_jsii_.MemberMethod{JsiiMethod: "resetInstanceType", GoMethod: "ResetInstanceType"},
			_jsii_.MemberMethod{JsiiMethod: "resetKernelId", GoMethod: "ResetKernelId"},
			_jsii_.MemberMethod{JsiiMethod: "resetKeyName", GoMethod: "ResetKeyName"},
			_jsii_.MemberMethod{JsiiMethod: "resetLicenseSpecifications", GoMethod: "ResetLicenseSpecifications"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaintenanceOptions", GoMethod: "ResetMaintenanceOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetMetadataOptions", GoMethod: "ResetMetadataOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetMonitoring", GoMethod: "ResetMonitoring"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetworkInterfaces", GoMethod: "ResetNetworkInterfaces"},
			_jsii_.MemberMethod{JsiiMethod: "resetPlacement", GoMethod: "ResetPlacement"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrivateDnsNameOptions", GoMethod: "ResetPrivateDnsNameOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetRamDiskId", GoMethod: "ResetRamDiskId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecurityGroupIds", GoMethod: "ResetSecurityGroupIds"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecurityGroups", GoMethod: "ResetSecurityGroups"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagSpecifications", GoMethod: "ResetTagSpecifications"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserData", GoMethod: "ResetUserData"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupIds", GoGetter: "SecurityGroupIds"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupIdsInput", GoGetter: "SecurityGroupIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroups", GoGetter: "SecurityGroups"},
			_jsii_.MemberProperty{JsiiProperty: "securityGroupsInput", GoGetter: "SecurityGroupsInput"},
			_jsii_.MemberProperty{JsiiProperty: "tagSpecifications", GoGetter: "TagSpecifications"},
			_jsii_.MemberProperty{JsiiProperty: "tagSpecificationsInput", GoGetter: "TagSpecificationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "userData", GoGetter: "UserData"},
			_jsii_.MemberProperty{JsiiProperty: "userDataInput", GoGetter: "UserDataInput"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataPlacement",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataPlacement)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataPlacementOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataPlacementOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "affinity", GoGetter: "Affinity"},
			_jsii_.MemberProperty{JsiiProperty: "affinityInput", GoGetter: "AffinityInput"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZone", GoGetter: "AvailabilityZone"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZoneInput", GoGetter: "AvailabilityZoneInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "groupId", GoGetter: "GroupId"},
			_jsii_.MemberProperty{JsiiProperty: "groupIdInput", GoGetter: "GroupIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "groupName", GoGetter: "GroupName"},
			_jsii_.MemberProperty{JsiiProperty: "groupNameInput", GoGetter: "GroupNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "hostId", GoGetter: "HostId"},
			_jsii_.MemberProperty{JsiiProperty: "hostIdInput", GoGetter: "HostIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "hostResourceGroupArn", GoGetter: "HostResourceGroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "hostResourceGroupArnInput", GoGetter: "HostResourceGroupArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "partitionNumber", GoGetter: "PartitionNumber"},
			_jsii_.MemberProperty{JsiiProperty: "partitionNumberInput", GoGetter: "PartitionNumberInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAffinity", GoMethod: "ResetAffinity"},
			_jsii_.MemberMethod{JsiiMethod: "resetAvailabilityZone", GoMethod: "ResetAvailabilityZone"},
			_jsii_.MemberMethod{JsiiMethod: "resetGroupId", GoMethod: "ResetGroupId"},
			_jsii_.MemberMethod{JsiiMethod: "resetGroupName", GoMethod: "ResetGroupName"},
			_jsii_.MemberMethod{JsiiMethod: "resetHostId", GoMethod: "ResetHostId"},
			_jsii_.MemberMethod{JsiiMethod: "resetHostResourceGroupArn", GoMethod: "ResetHostResourceGroupArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetPartitionNumber", GoMethod: "ResetPartitionNumber"},
			_jsii_.MemberMethod{JsiiMethod: "resetSpreadDomain", GoMethod: "ResetSpreadDomain"},
			_jsii_.MemberMethod{JsiiMethod: "resetTenancy", GoMethod: "ResetTenancy"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "spreadDomain", GoGetter: "SpreadDomain"},
			_jsii_.MemberProperty{JsiiProperty: "spreadDomainInput", GoGetter: "SpreadDomainInput"},
			_jsii_.MemberProperty{JsiiProperty: "tenancy", GoGetter: "Tenancy"},
			_jsii_.MemberProperty{JsiiProperty: "tenancyInput", GoGetter: "TenancyInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataPlacementOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataPrivateDnsNameOptions",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataPrivateDnsNameOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataPrivateDnsNameOptionsOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataPrivateDnsNameOptionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enableResourceNameDnsAaaaRecord", GoGetter: "EnableResourceNameDnsAaaaRecord"},
			_jsii_.MemberProperty{JsiiProperty: "enableResourceNameDnsAaaaRecordInput", GoGetter: "EnableResourceNameDnsAaaaRecordInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableResourceNameDnsARecord", GoGetter: "EnableResourceNameDnsARecord"},
			_jsii_.MemberProperty{JsiiProperty: "enableResourceNameDnsARecordInput", GoGetter: "EnableResourceNameDnsARecordInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "hostnameType", GoGetter: "HostnameType"},
			_jsii_.MemberProperty{JsiiProperty: "hostnameTypeInput", GoGetter: "HostnameTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableResourceNameDnsAaaaRecord", GoMethod: "ResetEnableResourceNameDnsAaaaRecord"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableResourceNameDnsARecord", GoMethod: "ResetEnableResourceNameDnsARecord"},
			_jsii_.MemberMethod{JsiiMethod: "resetHostnameType", GoMethod: "ResetHostnameType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataPrivateDnsNameOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataTagSpecifications",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataTagSpecifications)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataTagSpecificationsList",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataTagSpecificationsList)(nil)).Elem(),
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
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataTagSpecificationsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataTagSpecificationsOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataTagSpecificationsOutputReference)(nil)).Elem(),
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
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataTagSpecificationsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataTagSpecificationsTags",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataTagSpecificationsTags)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataTagSpecificationsTagsList",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataTagSpecificationsTagsList)(nil)).Elem(),
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
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataTagSpecificationsTagsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataTagSpecificationsTagsOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateLaunchTemplateDataTagSpecificationsTagsOutputReference)(nil)).Elem(),
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
			j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataTagSpecificationsTagsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateTagSpecifications",
		reflect.TypeOf((*Ec2LaunchTemplateTagSpecifications)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateTagSpecificationsList",
		reflect.TypeOf((*Ec2LaunchTemplateTagSpecificationsList)(nil)).Elem(),
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
			j := jsiiProxy_Ec2LaunchTemplateTagSpecificationsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateTagSpecificationsOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateTagSpecificationsOutputReference)(nil)).Elem(),
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
			j := jsiiProxy_Ec2LaunchTemplateTagSpecificationsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateTagSpecificationsTags",
		reflect.TypeOf((*Ec2LaunchTemplateTagSpecificationsTags)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateTagSpecificationsTagsList",
		reflect.TypeOf((*Ec2LaunchTemplateTagSpecificationsTagsList)(nil)).Elem(),
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
			j := jsiiProxy_Ec2LaunchTemplateTagSpecificationsTagsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ec2LaunchTemplate.Ec2LaunchTemplateTagSpecificationsTagsOutputReference",
		reflect.TypeOf((*Ec2LaunchTemplateTagSpecificationsTagsOutputReference)(nil)).Elem(),
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
			j := jsiiProxy_Ec2LaunchTemplateTagSpecificationsTagsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
}
