// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchjobdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/batchjobdefinition/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference interface {
	cdktn.ComplexObject
	Args() *[]*string
	SetArgs(val *[]*string)
	ArgsInput() *[]*string
	Command() *[]*string
	SetCommand(val *[]*string)
	CommandInput() *[]*string
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
	Env() BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersEnvList
	EnvInput() interface{}
	// Experimental.
	Fqn() *string
	Image() *string
	SetImage(val *string)
	ImageInput() *string
	ImagePullPolicy() *string
	SetImagePullPolicy(val *string)
	ImagePullPolicyInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Resources() BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersResourcesOutputReference
	ResourcesInput() interface{}
	SecurityContext() BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersSecurityContextOutputReference
	SecurityContextInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VolumeMounts() BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersVolumeMountsList
	VolumeMountsInput() interface{}
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
	PutEnv(value interface{})
	PutResources(value *BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersResources)
	PutSecurityContext(value *BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersSecurityContext)
	PutVolumeMounts(value interface{})
	ResetArgs()
	ResetCommand()
	ResetEnv()
	ResetImage()
	ResetImagePullPolicy()
	ResetName()
	ResetResources()
	ResetSecurityContext()
	ResetVolumeMounts()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference
type jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) Args() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"args",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) ArgsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"argsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) CommandInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) Env() BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersEnvList {
	var returns BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersEnvList
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) EnvInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"envInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) ImagePullPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePullPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) ImagePullPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagePullPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) Resources() BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersResourcesOutputReference {
	var returns BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersResourcesOutputReference
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) ResourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) SecurityContext() BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersSecurityContextOutputReference {
	var returns BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersSecurityContextOutputReference
	_jsii_.Get(
		j,
		"securityContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) SecurityContextInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"securityContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) VolumeMounts() BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersVolumeMountsList {
	var returns BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersVolumeMountsList
	_jsii_.Get(
		j,
		"volumeMounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) VolumeMountsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumeMountsInput",
		&returns,
	)
	return returns
}


func NewBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference {
	_init_.Initialize()

	if err := validateNewBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.batchJobDefinition.BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference_Override(b BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.batchJobDefinition.BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference)SetArgs(val *[]*string) {
	if err := j.validateSetArgsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"args",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference)SetCommand(val *[]*string) {
	if err := j.validateSetCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"command",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference)SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference)SetImagePullPolicy(val *string) {
	if err := j.validateSetImagePullPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imagePullPolicy",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) PutEnv(value interface{}) {
	if err := b.validatePutEnvParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putEnv",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) PutResources(value *BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersResources) {
	if err := b.validatePutResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putResources",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) PutSecurityContext(value *BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersSecurityContext) {
	if err := b.validatePutSecurityContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putSecurityContext",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) PutVolumeMounts(value interface{}) {
	if err := b.validatePutVolumeMountsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putVolumeMounts",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) ResetArgs() {
	_jsii_.InvokeVoid(
		b,
		"resetArgs",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) ResetCommand() {
	_jsii_.InvokeVoid(
		b,
		"resetCommand",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) ResetEnv() {
	_jsii_.InvokeVoid(
		b,
		"resetEnv",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) ResetImage() {
	_jsii_.InvokeVoid(
		b,
		"resetImage",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) ResetImagePullPolicy() {
	_jsii_.InvokeVoid(
		b,
		"resetImagePullPolicy",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		b,
		"resetName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) ResetResources() {
	_jsii_.InvokeVoid(
		b,
		"resetResources",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) ResetSecurityContext() {
	_jsii_.InvokeVoid(
		b,
		"resetSecurityContext",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) ResetVolumeMounts() {
	_jsii_.InvokeVoid(
		b,
		"resetVolumeMounts",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesContainersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

