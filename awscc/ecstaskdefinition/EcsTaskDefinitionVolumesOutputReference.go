// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecstaskdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/ecstaskdefinition/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EcsTaskDefinitionVolumesOutputReference interface {
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
	ConfiguredAtLaunch() interface{}
	SetConfiguredAtLaunch(val interface{})
	ConfiguredAtLaunchInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DockerVolumeConfiguration() EcsTaskDefinitionVolumesDockerVolumeConfigurationOutputReference
	DockerVolumeConfigurationInput() interface{}
	EfsVolumeConfiguration() EcsTaskDefinitionVolumesEfsVolumeConfigurationOutputReference
	EfsVolumeConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	FsxWindowsFileServerVolumeConfiguration() EcsTaskDefinitionVolumesFsxWindowsFileServerVolumeConfigurationOutputReference
	FsxWindowsFileServerVolumeConfigurationInput() interface{}
	Host() EcsTaskDefinitionVolumesHostOutputReference
	HostInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	S3FilesVolumeConfiguration() EcsTaskDefinitionVolumesS3FilesVolumeConfigurationOutputReference
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
	PutDockerVolumeConfiguration(value *EcsTaskDefinitionVolumesDockerVolumeConfiguration)
	PutEfsVolumeConfiguration(value *EcsTaskDefinitionVolumesEfsVolumeConfiguration)
	PutFsxWindowsFileServerVolumeConfiguration(value *EcsTaskDefinitionVolumesFsxWindowsFileServerVolumeConfiguration)
	PutHost(value *EcsTaskDefinitionVolumesHost)
	PutS3FilesVolumeConfiguration(value *EcsTaskDefinitionVolumesS3FilesVolumeConfiguration)
	ResetConfiguredAtLaunch()
	ResetDockerVolumeConfiguration()
	ResetEfsVolumeConfiguration()
	ResetFsxWindowsFileServerVolumeConfiguration()
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

// The jsii proxy struct for EcsTaskDefinitionVolumesOutputReference
type jsiiProxy_EcsTaskDefinitionVolumesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) ConfiguredAtLaunch() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configuredAtLaunch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) ConfiguredAtLaunchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configuredAtLaunchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) DockerVolumeConfiguration() EcsTaskDefinitionVolumesDockerVolumeConfigurationOutputReference {
	var returns EcsTaskDefinitionVolumesDockerVolumeConfigurationOutputReference
	_jsii_.Get(
		j,
		"dockerVolumeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) DockerVolumeConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dockerVolumeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) EfsVolumeConfiguration() EcsTaskDefinitionVolumesEfsVolumeConfigurationOutputReference {
	var returns EcsTaskDefinitionVolumesEfsVolumeConfigurationOutputReference
	_jsii_.Get(
		j,
		"efsVolumeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) EfsVolumeConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"efsVolumeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) FsxWindowsFileServerVolumeConfiguration() EcsTaskDefinitionVolumesFsxWindowsFileServerVolumeConfigurationOutputReference {
	var returns EcsTaskDefinitionVolumesFsxWindowsFileServerVolumeConfigurationOutputReference
	_jsii_.Get(
		j,
		"fsxWindowsFileServerVolumeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) FsxWindowsFileServerVolumeConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fsxWindowsFileServerVolumeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) Host() EcsTaskDefinitionVolumesHostOutputReference {
	var returns EcsTaskDefinitionVolumesHostOutputReference
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) HostInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) S3FilesVolumeConfiguration() EcsTaskDefinitionVolumesS3FilesVolumeConfigurationOutputReference {
	var returns EcsTaskDefinitionVolumesS3FilesVolumeConfigurationOutputReference
	_jsii_.Get(
		j,
		"s3FilesVolumeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) S3FilesVolumeConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3FilesVolumeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEcsTaskDefinitionVolumesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EcsTaskDefinitionVolumesOutputReference {
	_init_.Initialize()

	if err := validateNewEcsTaskDefinitionVolumesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsTaskDefinitionVolumesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ecsTaskDefinition.EcsTaskDefinitionVolumesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEcsTaskDefinitionVolumesOutputReference_Override(e EcsTaskDefinitionVolumesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ecsTaskDefinition.EcsTaskDefinitionVolumesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference)SetConfiguredAtLaunch(val interface{}) {
	if err := j.validateSetConfiguredAtLaunchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configuredAtLaunch",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EcsTaskDefinitionVolumesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) PutDockerVolumeConfiguration(value *EcsTaskDefinitionVolumesDockerVolumeConfiguration) {
	if err := e.validatePutDockerVolumeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putDockerVolumeConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) PutEfsVolumeConfiguration(value *EcsTaskDefinitionVolumesEfsVolumeConfiguration) {
	if err := e.validatePutEfsVolumeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putEfsVolumeConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) PutFsxWindowsFileServerVolumeConfiguration(value *EcsTaskDefinitionVolumesFsxWindowsFileServerVolumeConfiguration) {
	if err := e.validatePutFsxWindowsFileServerVolumeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putFsxWindowsFileServerVolumeConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) PutHost(value *EcsTaskDefinitionVolumesHost) {
	if err := e.validatePutHostParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putHost",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) PutS3FilesVolumeConfiguration(value *EcsTaskDefinitionVolumesS3FilesVolumeConfiguration) {
	if err := e.validatePutS3FilesVolumeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putS3FilesVolumeConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) ResetConfiguredAtLaunch() {
	_jsii_.InvokeVoid(
		e,
		"resetConfiguredAtLaunch",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) ResetDockerVolumeConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetDockerVolumeConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) ResetEfsVolumeConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetEfsVolumeConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) ResetFsxWindowsFileServerVolumeConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetFsxWindowsFileServerVolumeConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) ResetHost() {
	_jsii_.InvokeVoid(
		e,
		"resetHost",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		e,
		"resetName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) ResetS3FilesVolumeConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetS3FilesVolumeConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := e.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsTaskDefinitionVolumesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

