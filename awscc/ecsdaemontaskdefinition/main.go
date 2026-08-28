// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsdaemontaskdefinition

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinition",
		reflect.TypeOf((*EcsDaemonTaskDefinition)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "containerDefinitions", GoGetter: "ContainerDefinitions"},
			_jsii_.MemberProperty{JsiiProperty: "containerDefinitionsInput", GoGetter: "ContainerDefinitionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "cpu", GoGetter: "Cpu"},
			_jsii_.MemberProperty{JsiiProperty: "cpuInput", GoGetter: "CpuInput"},
			_jsii_.MemberProperty{JsiiProperty: "daemonTaskDefinitionArn", GoGetter: "DaemonTaskDefinitionArn"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "executionRoleArn", GoGetter: "ExecutionRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "executionRoleArnInput", GoGetter: "ExecutionRoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "family", GoGetter: "Family"},
			_jsii_.MemberProperty{JsiiProperty: "familyInput", GoGetter: "FamilyInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "ipcMode", GoGetter: "IpcMode"},
			_jsii_.MemberProperty{JsiiProperty: "ipcModeInput", GoGetter: "IpcModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "markWriteOnlyAttribute", GoMethod: "MarkWriteOnlyAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "memory", GoGetter: "Memory"},
			_jsii_.MemberProperty{JsiiProperty: "memoryInput", GoGetter: "MemoryInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "pidMode", GoGetter: "PidMode"},
			_jsii_.MemberProperty{JsiiProperty: "pidModeInput", GoGetter: "PidModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putContainerDefinitions", GoMethod: "PutContainerDefinitions"},
			_jsii_.MemberMethod{JsiiMethod: "putTags", GoMethod: "PutTags"},
			_jsii_.MemberMethod{JsiiMethod: "putVolumes", GoMethod: "PutVolumes"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "registerProviderFeatureUsage", GoMethod: "RegisterProviderFeatureUsage"},
			_jsii_.MemberMethod{JsiiMethod: "resetContainerDefinitions", GoMethod: "ResetContainerDefinitions"},
			_jsii_.MemberMethod{JsiiMethod: "resetCpu", GoMethod: "ResetCpu"},
			_jsii_.MemberMethod{JsiiMethod: "resetExecutionRoleArn", GoMethod: "ResetExecutionRoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetFamily", GoMethod: "ResetFamily"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpcMode", GoMethod: "ResetIpcMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetMemory", GoMethod: "ResetMemory"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPidMode", GoMethod: "ResetPidMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTaskRoleArn", GoMethod: "ResetTaskRoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetVolumes", GoMethod: "ResetVolumes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "taskRoleArn", GoGetter: "TaskRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "taskRoleArnInput", GoGetter: "TaskRoleArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "volumes", GoGetter: "Volumes"},
			_jsii_.MemberProperty{JsiiProperty: "volumesInput", GoGetter: "VolumesInput"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_EcsDaemonTaskDefinition{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionConfig",
		reflect.TypeOf((*EcsDaemonTaskDefinitionConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitions",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsDependsOn",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsDependsOn)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsDependsOnList",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsDependsOnList)(nil)).Elem(),
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
			j := jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsDependsOnList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsDependsOnOutputReference",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsDependsOnOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "condition", GoGetter: "Condition"},
			_jsii_.MemberProperty{JsiiProperty: "conditionInput", GoGetter: "ConditionInput"},
			_jsii_.MemberProperty{JsiiProperty: "containerName", GoGetter: "ContainerName"},
			_jsii_.MemberProperty{JsiiProperty: "containerNameInput", GoGetter: "ContainerNameInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetCondition", GoMethod: "ResetCondition"},
			_jsii_.MemberMethod{JsiiMethod: "resetContainerName", GoMethod: "ResetContainerName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsDependsOnOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsEnvironment",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsEnvironment)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsEnvironmentFiles",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsEnvironmentFiles)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsEnvironmentFilesList",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsEnvironmentFilesList)(nil)).Elem(),
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
			j := jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsEnvironmentFilesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsEnvironmentFilesOutputReference",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsEnvironmentFilesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetValue", GoMethod: "ResetValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsEnvironmentFilesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsEnvironmentList",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsEnvironmentList)(nil)).Elem(),
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
			j := jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsEnvironmentList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsEnvironmentOutputReference",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsEnvironmentOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetValue", GoMethod: "ResetValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsEnvironmentOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsFirelensConfiguration",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsFirelensConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsFirelensConfigurationOutputReference",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsFirelensConfigurationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "options", GoGetter: "Options"},
			_jsii_.MemberProperty{JsiiProperty: "optionsInput", GoGetter: "OptionsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetOptions", GoMethod: "ResetOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetType", GoMethod: "ResetType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsFirelensConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsHealthCheck",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsHealthCheck)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "command", GoGetter: "Command"},
			_jsii_.MemberProperty{JsiiProperty: "commandInput", GoGetter: "CommandInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "interval", GoGetter: "Interval"},
			_jsii_.MemberProperty{JsiiProperty: "intervalInput", GoGetter: "IntervalInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCommand", GoMethod: "ResetCommand"},
			_jsii_.MemberMethod{JsiiMethod: "resetInterval", GoMethod: "ResetInterval"},
			_jsii_.MemberMethod{JsiiMethod: "resetRetries", GoMethod: "ResetRetries"},
			_jsii_.MemberMethod{JsiiMethod: "resetStartPeriod", GoMethod: "ResetStartPeriod"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeout", GoMethod: "ResetTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "retries", GoGetter: "Retries"},
			_jsii_.MemberProperty{JsiiProperty: "retriesInput", GoGetter: "RetriesInput"},
			_jsii_.MemberProperty{JsiiProperty: "startPeriod", GoGetter: "StartPeriod"},
			_jsii_.MemberProperty{JsiiProperty: "startPeriodInput", GoGetter: "StartPeriodInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "timeout", GoGetter: "Timeout"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutInput", GoGetter: "TimeoutInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsLinuxParameters",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsLinuxParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersCapabilities",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersCapabilities)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersCapabilitiesOutputReference",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersCapabilitiesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "add", GoGetter: "Add"},
			_jsii_.MemberProperty{JsiiProperty: "addInput", GoGetter: "AddInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "drop", GoGetter: "Drop"},
			_jsii_.MemberProperty{JsiiProperty: "dropInput", GoGetter: "DropInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAdd", GoMethod: "ResetAdd"},
			_jsii_.MemberMethod{JsiiMethod: "resetDrop", GoMethod: "ResetDrop"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersCapabilitiesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersDevices",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersDevices)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersDevicesList",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersDevicesList)(nil)).Elem(),
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
			j := jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersDevicesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersDevicesOutputReference",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersDevicesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "containerPath", GoGetter: "ContainerPath"},
			_jsii_.MemberProperty{JsiiProperty: "containerPathInput", GoGetter: "ContainerPathInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "hostPath", GoGetter: "HostPath"},
			_jsii_.MemberProperty{JsiiProperty: "hostPathInput", GoGetter: "HostPathInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "permissions", GoGetter: "Permissions"},
			_jsii_.MemberProperty{JsiiProperty: "permissionsInput", GoGetter: "PermissionsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetContainerPath", GoMethod: "ResetContainerPath"},
			_jsii_.MemberMethod{JsiiMethod: "resetHostPath", GoMethod: "ResetHostPath"},
			_jsii_.MemberMethod{JsiiMethod: "resetPermissions", GoMethod: "ResetPermissions"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersDevicesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "capabilities", GoGetter: "Capabilities"},
			_jsii_.MemberProperty{JsiiProperty: "capabilitiesInput", GoGetter: "CapabilitiesInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "devices", GoGetter: "Devices"},
			_jsii_.MemberProperty{JsiiProperty: "devicesInput", GoGetter: "DevicesInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "initProcessEnabled", GoGetter: "InitProcessEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "initProcessEnabledInput", GoGetter: "InitProcessEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putCapabilities", GoMethod: "PutCapabilities"},
			_jsii_.MemberMethod{JsiiMethod: "putDevices", GoMethod: "PutDevices"},
			_jsii_.MemberMethod{JsiiMethod: "putTmpfs", GoMethod: "PutTmpfs"},
			_jsii_.MemberMethod{JsiiMethod: "resetCapabilities", GoMethod: "ResetCapabilities"},
			_jsii_.MemberMethod{JsiiMethod: "resetDevices", GoMethod: "ResetDevices"},
			_jsii_.MemberMethod{JsiiMethod: "resetInitProcessEnabled", GoMethod: "ResetInitProcessEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetTmpfs", GoMethod: "ResetTmpfs"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "tmpfs", GoGetter: "Tmpfs"},
			_jsii_.MemberProperty{JsiiProperty: "tmpfsInput", GoGetter: "TmpfsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersTmpfs",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersTmpfs)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersTmpfsList",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersTmpfsList)(nil)).Elem(),
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
			j := jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersTmpfsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersTmpfsOutputReference",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersTmpfsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "containerPath", GoGetter: "ContainerPath"},
			_jsii_.MemberProperty{JsiiProperty: "containerPathInput", GoGetter: "ContainerPathInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "mountOptions", GoGetter: "MountOptions"},
			_jsii_.MemberProperty{JsiiProperty: "mountOptionsInput", GoGetter: "MountOptionsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetContainerPath", GoMethod: "ResetContainerPath"},
			_jsii_.MemberMethod{JsiiMethod: "resetMountOptions", GoMethod: "ResetMountOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetSize", GoMethod: "ResetSize"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "size", GoGetter: "Size"},
			_jsii_.MemberProperty{JsiiProperty: "sizeInput", GoGetter: "SizeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLinuxParametersTmpfsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsList",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsList)(nil)).Elem(),
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
			j := jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsLogConfiguration",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsLogConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsLogConfigurationOutputReference",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsLogConfigurationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "logDriver", GoGetter: "LogDriver"},
			_jsii_.MemberProperty{JsiiProperty: "logDriverInput", GoGetter: "LogDriverInput"},
			_jsii_.MemberProperty{JsiiProperty: "options", GoGetter: "Options"},
			_jsii_.MemberProperty{JsiiProperty: "optionsInput", GoGetter: "OptionsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putSecretOptions", GoMethod: "PutSecretOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogDriver", GoMethod: "ResetLogDriver"},
			_jsii_.MemberMethod{JsiiMethod: "resetOptions", GoMethod: "ResetOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecretOptions", GoMethod: "ResetSecretOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "secretOptions", GoGetter: "SecretOptions"},
			_jsii_.MemberProperty{JsiiProperty: "secretOptionsInput", GoGetter: "SecretOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLogConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsLogConfigurationSecretOptions",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsLogConfigurationSecretOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsLogConfigurationSecretOptionsList",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsLogConfigurationSecretOptionsList)(nil)).Elem(),
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
			j := jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLogConfigurationSecretOptionsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsLogConfigurationSecretOptionsOutputReference",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsLogConfigurationSecretOptionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetValueFrom", GoMethod: "ResetValueFrom"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "valueFrom", GoGetter: "ValueFrom"},
			_jsii_.MemberProperty{JsiiProperty: "valueFromInput", GoGetter: "ValueFromInput"},
		},
		func() interface{} {
			j := jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsLogConfigurationSecretOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsMountPoints",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsMountPoints)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsMountPointsList",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsMountPointsList)(nil)).Elem(),
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
			j := jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsMountPointsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "containerPath", GoGetter: "ContainerPath"},
			_jsii_.MemberProperty{JsiiProperty: "containerPathInput", GoGetter: "ContainerPathInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "readOnly", GoGetter: "ReadOnly"},
			_jsii_.MemberProperty{JsiiProperty: "readOnlyInput", GoGetter: "ReadOnlyInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetContainerPath", GoMethod: "ResetContainerPath"},
			_jsii_.MemberMethod{JsiiMethod: "resetReadOnly", GoMethod: "ResetReadOnly"},
			_jsii_.MemberMethod{JsiiMethod: "resetSourceVolume", GoMethod: "ResetSourceVolume"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sourceVolume", GoGetter: "SourceVolume"},
			_jsii_.MemberProperty{JsiiProperty: "sourceVolumeInput", GoGetter: "SourceVolumeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsMountPointsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsOutputReference",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "command", GoGetter: "Command"},
			_jsii_.MemberProperty{JsiiProperty: "commandInput", GoGetter: "CommandInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "cpu", GoGetter: "Cpu"},
			_jsii_.MemberProperty{JsiiProperty: "cpuInput", GoGetter: "CpuInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOnInput", GoGetter: "DependsOnInput"},
			_jsii_.MemberProperty{JsiiProperty: "entryPoint", GoGetter: "EntryPoint"},
			_jsii_.MemberProperty{JsiiProperty: "entryPointInput", GoGetter: "EntryPointInput"},
			_jsii_.MemberProperty{JsiiProperty: "environment", GoGetter: "Environment"},
			_jsii_.MemberProperty{JsiiProperty: "environmentFiles", GoGetter: "EnvironmentFiles"},
			_jsii_.MemberProperty{JsiiProperty: "environmentFilesInput", GoGetter: "EnvironmentFilesInput"},
			_jsii_.MemberProperty{JsiiProperty: "environmentInput", GoGetter: "EnvironmentInput"},
			_jsii_.MemberProperty{JsiiProperty: "essential", GoGetter: "Essential"},
			_jsii_.MemberProperty{JsiiProperty: "essentialInput", GoGetter: "EssentialInput"},
			_jsii_.MemberProperty{JsiiProperty: "firelensConfiguration", GoGetter: "FirelensConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "firelensConfigurationInput", GoGetter: "FirelensConfigurationInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "healthCheck", GoGetter: "HealthCheck"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheckInput", GoGetter: "HealthCheckInput"},
			_jsii_.MemberProperty{JsiiProperty: "image", GoGetter: "Image"},
			_jsii_.MemberProperty{JsiiProperty: "imageInput", GoGetter: "ImageInput"},
			_jsii_.MemberProperty{JsiiProperty: "interactive", GoGetter: "Interactive"},
			_jsii_.MemberProperty{JsiiProperty: "interactiveInput", GoGetter: "InteractiveInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "linuxParameters", GoGetter: "LinuxParameters"},
			_jsii_.MemberProperty{JsiiProperty: "linuxParametersInput", GoGetter: "LinuxParametersInput"},
			_jsii_.MemberProperty{JsiiProperty: "logConfiguration", GoGetter: "LogConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "logConfigurationInput", GoGetter: "LogConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "memory", GoGetter: "Memory"},
			_jsii_.MemberProperty{JsiiProperty: "memoryInput", GoGetter: "MemoryInput"},
			_jsii_.MemberProperty{JsiiProperty: "memoryReservation", GoGetter: "MemoryReservation"},
			_jsii_.MemberProperty{JsiiProperty: "memoryReservationInput", GoGetter: "MemoryReservationInput"},
			_jsii_.MemberProperty{JsiiProperty: "mountPoints", GoGetter: "MountPoints"},
			_jsii_.MemberProperty{JsiiProperty: "mountPointsInput", GoGetter: "MountPointsInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "privileged", GoGetter: "Privileged"},
			_jsii_.MemberProperty{JsiiProperty: "privilegedInput", GoGetter: "PrivilegedInput"},
			_jsii_.MemberProperty{JsiiProperty: "pseudoTerminal", GoGetter: "PseudoTerminal"},
			_jsii_.MemberProperty{JsiiProperty: "pseudoTerminalInput", GoGetter: "PseudoTerminalInput"},
			_jsii_.MemberMethod{JsiiMethod: "putDependsOn", GoMethod: "PutDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "putEnvironment", GoMethod: "PutEnvironment"},
			_jsii_.MemberMethod{JsiiMethod: "putEnvironmentFiles", GoMethod: "PutEnvironmentFiles"},
			_jsii_.MemberMethod{JsiiMethod: "putFirelensConfiguration", GoMethod: "PutFirelensConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putHealthCheck", GoMethod: "PutHealthCheck"},
			_jsii_.MemberMethod{JsiiMethod: "putLinuxParameters", GoMethod: "PutLinuxParameters"},
			_jsii_.MemberMethod{JsiiMethod: "putLogConfiguration", GoMethod: "PutLogConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "putMountPoints", GoMethod: "PutMountPoints"},
			_jsii_.MemberMethod{JsiiMethod: "putRepositoryCredentials", GoMethod: "PutRepositoryCredentials"},
			_jsii_.MemberMethod{JsiiMethod: "putRestartPolicy", GoMethod: "PutRestartPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "putSecrets", GoMethod: "PutSecrets"},
			_jsii_.MemberMethod{JsiiMethod: "putSystemControls", GoMethod: "PutSystemControls"},
			_jsii_.MemberMethod{JsiiMethod: "putUlimits", GoMethod: "PutUlimits"},
			_jsii_.MemberProperty{JsiiProperty: "readonlyRootFilesystem", GoGetter: "ReadonlyRootFilesystem"},
			_jsii_.MemberProperty{JsiiProperty: "readonlyRootFilesystemInput", GoGetter: "ReadonlyRootFilesystemInput"},
			_jsii_.MemberProperty{JsiiProperty: "repositoryCredentials", GoGetter: "RepositoryCredentials"},
			_jsii_.MemberProperty{JsiiProperty: "repositoryCredentialsInput", GoGetter: "RepositoryCredentialsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCommand", GoMethod: "ResetCommand"},
			_jsii_.MemberMethod{JsiiMethod: "resetCpu", GoMethod: "ResetCpu"},
			_jsii_.MemberMethod{JsiiMethod: "resetDependsOn", GoMethod: "ResetDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "resetEntryPoint", GoMethod: "ResetEntryPoint"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnvironment", GoMethod: "ResetEnvironment"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnvironmentFiles", GoMethod: "ResetEnvironmentFiles"},
			_jsii_.MemberMethod{JsiiMethod: "resetEssential", GoMethod: "ResetEssential"},
			_jsii_.MemberMethod{JsiiMethod: "resetFirelensConfiguration", GoMethod: "ResetFirelensConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetHealthCheck", GoMethod: "ResetHealthCheck"},
			_jsii_.MemberMethod{JsiiMethod: "resetImage", GoMethod: "ResetImage"},
			_jsii_.MemberMethod{JsiiMethod: "resetInteractive", GoMethod: "ResetInteractive"},
			_jsii_.MemberMethod{JsiiMethod: "resetLinuxParameters", GoMethod: "ResetLinuxParameters"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogConfiguration", GoMethod: "ResetLogConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetMemory", GoMethod: "ResetMemory"},
			_jsii_.MemberMethod{JsiiMethod: "resetMemoryReservation", GoMethod: "ResetMemoryReservation"},
			_jsii_.MemberMethod{JsiiMethod: "resetMountPoints", GoMethod: "ResetMountPoints"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrivileged", GoMethod: "ResetPrivileged"},
			_jsii_.MemberMethod{JsiiMethod: "resetPseudoTerminal", GoMethod: "ResetPseudoTerminal"},
			_jsii_.MemberMethod{JsiiMethod: "resetReadonlyRootFilesystem", GoMethod: "ResetReadonlyRootFilesystem"},
			_jsii_.MemberMethod{JsiiMethod: "resetRepositoryCredentials", GoMethod: "ResetRepositoryCredentials"},
			_jsii_.MemberMethod{JsiiMethod: "resetRestartPolicy", GoMethod: "ResetRestartPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecrets", GoMethod: "ResetSecrets"},
			_jsii_.MemberMethod{JsiiMethod: "resetStartTimeout", GoMethod: "ResetStartTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetStopTimeout", GoMethod: "ResetStopTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetSystemControls", GoMethod: "ResetSystemControls"},
			_jsii_.MemberMethod{JsiiMethod: "resetUlimits", GoMethod: "ResetUlimits"},
			_jsii_.MemberMethod{JsiiMethod: "resetUser", GoMethod: "ResetUser"},
			_jsii_.MemberMethod{JsiiMethod: "resetWorkingDirectory", GoMethod: "ResetWorkingDirectory"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "restartPolicy", GoGetter: "RestartPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "restartPolicyInput", GoGetter: "RestartPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "secrets", GoGetter: "Secrets"},
			_jsii_.MemberProperty{JsiiProperty: "secretsInput", GoGetter: "SecretsInput"},
			_jsii_.MemberProperty{JsiiProperty: "startTimeout", GoGetter: "StartTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "startTimeoutInput", GoGetter: "StartTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "stopTimeout", GoGetter: "StopTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "stopTimeoutInput", GoGetter: "StopTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "systemControls", GoGetter: "SystemControls"},
			_jsii_.MemberProperty{JsiiProperty: "systemControlsInput", GoGetter: "SystemControlsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "ulimits", GoGetter: "Ulimits"},
			_jsii_.MemberProperty{JsiiProperty: "ulimitsInput", GoGetter: "UlimitsInput"},
			_jsii_.MemberProperty{JsiiProperty: "user", GoGetter: "User"},
			_jsii_.MemberProperty{JsiiProperty: "userInput", GoGetter: "UserInput"},
			_jsii_.MemberProperty{JsiiProperty: "workingDirectory", GoGetter: "WorkingDirectory"},
			_jsii_.MemberProperty{JsiiProperty: "workingDirectoryInput", GoGetter: "WorkingDirectoryInput"},
		},
		func() interface{} {
			j := jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsRepositoryCredentials",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsRepositoryCredentials)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsRepositoryCredentialsOutputReference",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsRepositoryCredentialsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "credentialsParameter", GoGetter: "CredentialsParameter"},
			_jsii_.MemberProperty{JsiiProperty: "credentialsParameterInput", GoGetter: "CredentialsParameterInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetCredentialsParameter", GoMethod: "ResetCredentialsParameter"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsRepositoryCredentialsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsRestartPolicy",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsRestartPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsRestartPolicyOutputReference",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsRestartPolicyOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "ignoredExitCodes", GoGetter: "IgnoredExitCodes"},
			_jsii_.MemberProperty{JsiiProperty: "ignoredExitCodesInput", GoGetter: "IgnoredExitCodesInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetIgnoredExitCodes", GoMethod: "ResetIgnoredExitCodes"},
			_jsii_.MemberMethod{JsiiMethod: "resetRestartAttemptPeriod", GoMethod: "ResetRestartAttemptPeriod"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "restartAttemptPeriod", GoGetter: "RestartAttemptPeriod"},
			_jsii_.MemberProperty{JsiiProperty: "restartAttemptPeriodInput", GoGetter: "RestartAttemptPeriodInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsRestartPolicyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsSecrets",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsSecrets)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsSecretsList",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsSecretsList)(nil)).Elem(),
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
			j := jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsSecretsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsSecretsOutputReference",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsSecretsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetValueFrom", GoMethod: "ResetValueFrom"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "valueFrom", GoGetter: "ValueFrom"},
			_jsii_.MemberProperty{JsiiProperty: "valueFromInput", GoGetter: "ValueFromInput"},
		},
		func() interface{} {
			j := jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsSecretsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsSystemControls",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsSystemControls)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsSystemControlsList",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsSystemControlsList)(nil)).Elem(),
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
			j := jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsSystemControlsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsSystemControlsOutputReference",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsSystemControlsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "namespaceInput", GoGetter: "NamespaceInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetNamespace", GoMethod: "ResetNamespace"},
			_jsii_.MemberMethod{JsiiMethod: "resetValue", GoMethod: "ResetValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsSystemControlsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsUlimits",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsUlimits)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsUlimitsList",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsUlimitsList)(nil)).Elem(),
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
			j := jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsUlimitsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsUlimitsOutputReference",
		reflect.TypeOf((*EcsDaemonTaskDefinitionContainerDefinitionsUlimitsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "hardLimit", GoGetter: "HardLimit"},
			_jsii_.MemberProperty{JsiiProperty: "hardLimitInput", GoGetter: "HardLimitInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetHardLimit", GoMethod: "ResetHardLimit"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetSoftLimit", GoMethod: "ResetSoftLimit"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "softLimit", GoGetter: "SoftLimit"},
			_jsii_.MemberProperty{JsiiProperty: "softLimitInput", GoGetter: "SoftLimitInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsUlimitsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionTags",
		reflect.TypeOf((*EcsDaemonTaskDefinitionTags)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionTagsList",
		reflect.TypeOf((*EcsDaemonTaskDefinitionTagsList)(nil)).Elem(),
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
			j := jsiiProxy_EcsDaemonTaskDefinitionTagsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionTagsOutputReference",
		reflect.TypeOf((*EcsDaemonTaskDefinitionTagsOutputReference)(nil)).Elem(),
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
			j := jsiiProxy_EcsDaemonTaskDefinitionTagsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionVolumes",
		reflect.TypeOf((*EcsDaemonTaskDefinitionVolumes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionVolumesHost",
		reflect.TypeOf((*EcsDaemonTaskDefinitionVolumesHost)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionVolumesHostOutputReference",
		reflect.TypeOf((*EcsDaemonTaskDefinitionVolumesHostOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resetSourcePath", GoMethod: "ResetSourcePath"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sourcePath", GoGetter: "SourcePath"},
			_jsii_.MemberProperty{JsiiProperty: "sourcePathInput", GoGetter: "SourcePathInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EcsDaemonTaskDefinitionVolumesHostOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionVolumesList",
		reflect.TypeOf((*EcsDaemonTaskDefinitionVolumesList)(nil)).Elem(),
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
			j := jsiiProxy_EcsDaemonTaskDefinitionVolumesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionVolumesOutputReference",
		reflect.TypeOf((*EcsDaemonTaskDefinitionVolumesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "host", GoGetter: "Host"},
			_jsii_.MemberProperty{JsiiProperty: "hostInput", GoGetter: "HostInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberMethod{JsiiMethod: "putHost", GoMethod: "PutHost"},
			_jsii_.MemberMethod{JsiiMethod: "resetHost", GoMethod: "ResetHost"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_EcsDaemonTaskDefinitionVolumesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
}
