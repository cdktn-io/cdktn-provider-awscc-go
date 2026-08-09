// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccgameliftcontainergroupdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccgameliftcontainergroupdefinition/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference interface {
	cdktn.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	ContainerName() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DependsOn() DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsDependsOnList
	EnvironmentOverride() DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsEnvironmentOverrideList
	Essential() cdktn.IResolvable
	// Experimental.
	Fqn() *string
	HealthCheck() DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsHealthCheckOutputReference
	ImageUri() *string
	InternalValue() *DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitions
	SetInternalValue(val *DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitions)
	LinuxCapabilities() DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsLinuxCapabilitiesOutputReference
	MemoryHardLimitMebibytes() *float64
	MountPoints() DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsMountPointsList
	PortConfiguration() DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsPortConfigurationOutputReference
	ResolvedImageDigest() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Vcpu() *float64
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference
type jsiiProxy_DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) ContainerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) DependsOn() DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsDependsOnList {
	var returns DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsDependsOnList
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) EnvironmentOverride() DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsEnvironmentOverrideList {
	var returns DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsEnvironmentOverrideList
	_jsii_.Get(
		j,
		"environmentOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) Essential() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"essential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) HealthCheck() DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsHealthCheckOutputReference {
	var returns DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsHealthCheckOutputReference
	_jsii_.Get(
		j,
		"healthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) ImageUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) InternalValue() *DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitions {
	var returns *DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) LinuxCapabilities() DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsLinuxCapabilitiesOutputReference {
	var returns DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsLinuxCapabilitiesOutputReference
	_jsii_.Get(
		j,
		"linuxCapabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) MemoryHardLimitMebibytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryHardLimitMebibytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) MountPoints() DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsMountPointsList {
	var returns DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsMountPointsList
	_jsii_.Get(
		j,
		"mountPoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) PortConfiguration() DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsPortConfigurationOutputReference {
	var returns DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsPortConfigurationOutputReference
	_jsii_.Get(
		j,
		"portConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) ResolvedImageDigest() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resolvedImageDigest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) Vcpu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vcpu",
		&returns,
	)
	return returns
}


func NewDataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccGameliftContainerGroupDefinition.DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference_Override(d DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccGameliftContainerGroupDefinition.DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference)SetInternalValue(val *DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccGameliftContainerGroupDefinitionSupportContainerDefinitionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

