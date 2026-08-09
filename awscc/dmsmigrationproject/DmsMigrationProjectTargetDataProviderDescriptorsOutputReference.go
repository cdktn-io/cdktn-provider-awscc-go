// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dmsmigrationproject

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dmsmigrationproject/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DmsMigrationProjectTargetDataProviderDescriptorsOutputReference interface {
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
	DataProviderArn() *string
	SetDataProviderArn(val *string)
	DataProviderArnInput() *string
	DataProviderIdentifier() *string
	SetDataProviderIdentifier(val *string)
	DataProviderIdentifierInput() *string
	DataProviderName() *string
	SetDataProviderName(val *string)
	DataProviderNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SecretsManagerAccessRoleArn() *string
	SetSecretsManagerAccessRoleArn(val *string)
	SecretsManagerAccessRoleArnInput() *string
	SecretsManagerSecretId() *string
	SetSecretsManagerSecretId(val *string)
	SecretsManagerSecretIdInput() *string
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
	ResetDataProviderArn()
	ResetDataProviderIdentifier()
	ResetDataProviderName()
	ResetSecretsManagerAccessRoleArn()
	ResetSecretsManagerSecretId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DmsMigrationProjectTargetDataProviderDescriptorsOutputReference
type jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference) DataProviderArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataProviderArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference) DataProviderArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataProviderArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference) DataProviderIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataProviderIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference) DataProviderIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataProviderIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference) DataProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataProviderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference) DataProviderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataProviderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference) SecretsManagerAccessRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerAccessRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference) SecretsManagerAccessRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerAccessRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference) SecretsManagerSecretId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerSecretId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference) SecretsManagerSecretIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretsManagerSecretIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDmsMigrationProjectTargetDataProviderDescriptorsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DmsMigrationProjectTargetDataProviderDescriptorsOutputReference {
	_init_.Initialize()

	if err := validateNewDmsMigrationProjectTargetDataProviderDescriptorsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dmsMigrationProject.DmsMigrationProjectTargetDataProviderDescriptorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDmsMigrationProjectTargetDataProviderDescriptorsOutputReference_Override(d DmsMigrationProjectTargetDataProviderDescriptorsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dmsMigrationProject.DmsMigrationProjectTargetDataProviderDescriptorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference)SetDataProviderArn(val *string) {
	if err := j.validateSetDataProviderArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataProviderArn",
		val,
	)
}

func (j *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference)SetDataProviderIdentifier(val *string) {
	if err := j.validateSetDataProviderIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataProviderIdentifier",
		val,
	)
}

func (j *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference)SetDataProviderName(val *string) {
	if err := j.validateSetDataProviderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataProviderName",
		val,
	)
}

func (j *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference)SetSecretsManagerAccessRoleArn(val *string) {
	if err := j.validateSetSecretsManagerAccessRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretsManagerAccessRoleArn",
		val,
	)
}

func (j *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference)SetSecretsManagerSecretId(val *string) {
	if err := j.validateSetSecretsManagerSecretIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretsManagerSecretId",
		val,
	)
}

func (j *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference) ResetDataProviderArn() {
	_jsii_.InvokeVoid(
		d,
		"resetDataProviderArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference) ResetDataProviderIdentifier() {
	_jsii_.InvokeVoid(
		d,
		"resetDataProviderIdentifier",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference) ResetDataProviderName() {
	_jsii_.InvokeVoid(
		d,
		"resetDataProviderName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference) ResetSecretsManagerAccessRoleArn() {
	_jsii_.InvokeVoid(
		d,
		"resetSecretsManagerAccessRoleArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference) ResetSecretsManagerSecretId() {
	_jsii_.InvokeVoid(
		d,
		"resetSecretsManagerSecretId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DmsMigrationProjectTargetDataProviderDescriptorsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

