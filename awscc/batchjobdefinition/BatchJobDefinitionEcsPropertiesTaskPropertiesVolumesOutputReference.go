// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchjobdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/batchjobdefinition/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EfsVolumeConfiguration() BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationOutputReference
	EfsVolumeConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	Host() BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesHostOutputReference
	HostInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	S3FilesVolumeConfiguration() BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesS3FilesVolumeConfigurationOutputReference
	S3FilesVolumeConfigurationInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	PutEfsVolumeConfiguration(value *BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfiguration)
	PutHost(value *BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesHost)
	PutS3FilesVolumeConfiguration(value *BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesS3FilesVolumeConfiguration)
	ResetEfsVolumeConfiguration()
	ResetHost()
	ResetName()
	ResetS3FilesVolumeConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference
type jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference) EfsVolumeConfiguration() BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationOutputReference {
	var returns BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfigurationOutputReference
	_jsii_.Get(
		j,
		"efsVolumeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference) EfsVolumeConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"efsVolumeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference) Host() BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesHostOutputReference {
	var returns BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesHostOutputReference
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference) HostInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference) S3FilesVolumeConfiguration() BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesS3FilesVolumeConfigurationOutputReference {
	var returns BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesS3FilesVolumeConfigurationOutputReference
	_jsii_.Get(
		j,
		"s3FilesVolumeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference) S3FilesVolumeConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3FilesVolumeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference {
	_init_.Initialize()

	if err := validateNewBatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.batchJobDefinition.BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference_Override(b BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.batchJobDefinition.BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference) PutEfsVolumeConfiguration(value *BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfiguration) {
	if err := b.validatePutEfsVolumeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putEfsVolumeConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference) PutHost(value *BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesHost) {
	if err := b.validatePutHostParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putHost",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference) PutS3FilesVolumeConfiguration(value *BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesS3FilesVolumeConfiguration) {
	if err := b.validatePutS3FilesVolumeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putS3FilesVolumeConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference) ResetEfsVolumeConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetEfsVolumeConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference) ResetHost() {
	_jsii_.InvokeVoid(
		b,
		"resetHost",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		b,
		"resetName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference) ResetS3FilesVolumeConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetS3FilesVolumeConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := b.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

