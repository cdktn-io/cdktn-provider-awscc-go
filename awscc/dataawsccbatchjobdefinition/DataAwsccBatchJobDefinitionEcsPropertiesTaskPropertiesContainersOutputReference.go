// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccbatchjobdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccbatchjobdefinition/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference interface {
	cdktn.ComplexObject
	Command() *[]*string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DependsOn() DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersDependsOnList
	Environment() DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersEnvironmentList
	Essential() cdktn.IResolvable
	FirelensConfiguration() DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersFirelensConfigurationOutputReference
	// Experimental.
	Fqn() *string
	Image() *string
	InternalValue() *DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainers
	SetInternalValue(val *DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainers)
	LinuxParameters() DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersLinuxParametersOutputReference
	LogConfiguration() DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersLogConfigurationOutputReference
	MountPoints() DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersMountPointsList
	Name() *string
	Privileged() cdktn.IResolvable
	ReadonlyRootFilesystem() cdktn.IResolvable
	RepositoryCredentials() DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersRepositoryCredentialsOutputReference
	ResourceRequirements() DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersResourceRequirementsList
	Secrets() DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersSecretsList
	StartTimeout() *float64
	StopTimeout() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Ulimits() DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersUlimitsList
	User() *string
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

// The jsii proxy struct for DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference
type jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) DependsOn() DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersDependsOnList {
	var returns DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersDependsOnList
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) Environment() DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersEnvironmentList {
	var returns DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersEnvironmentList
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) Essential() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"essential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) FirelensConfiguration() DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersFirelensConfigurationOutputReference {
	var returns DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersFirelensConfigurationOutputReference
	_jsii_.Get(
		j,
		"firelensConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) InternalValue() *DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainers {
	var returns *DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainers
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) LinuxParameters() DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersLinuxParametersOutputReference {
	var returns DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersLinuxParametersOutputReference
	_jsii_.Get(
		j,
		"linuxParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) LogConfiguration() DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersLogConfigurationOutputReference {
	var returns DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersLogConfigurationOutputReference
	_jsii_.Get(
		j,
		"logConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) MountPoints() DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersMountPointsList {
	var returns DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersMountPointsList
	_jsii_.Get(
		j,
		"mountPoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) Privileged() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"privileged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) ReadonlyRootFilesystem() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"readonlyRootFilesystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) RepositoryCredentials() DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersRepositoryCredentialsOutputReference {
	var returns DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersRepositoryCredentialsOutputReference
	_jsii_.Get(
		j,
		"repositoryCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) ResourceRequirements() DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersResourceRequirementsList {
	var returns DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersResourceRequirementsList
	_jsii_.Get(
		j,
		"resourceRequirements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) Secrets() DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersSecretsList {
	var returns DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersSecretsList
	_jsii_.Get(
		j,
		"secrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) StartTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) StopTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stopTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) Ulimits() DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersUlimitsList {
	var returns DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersUlimitsList
	_jsii_.Get(
		j,
		"ulimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}


func NewDataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccBatchJobDefinition.DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference_Override(d DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccBatchJobDefinition.DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference)SetInternalValue(val *DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainers) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccBatchJobDefinitionEcsPropertiesTaskPropertiesContainersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

