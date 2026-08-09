// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kendradatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/kendradatasource/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference interface {
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
	ExcludeMimeTypes() *[]*string
	SetExcludeMimeTypes(val *[]*string)
	ExcludeMimeTypesInput() *[]*string
	ExcludeSharedDrives() *[]*string
	SetExcludeSharedDrives(val *[]*string)
	ExcludeSharedDrivesInput() *[]*string
	ExcludeUserAccounts() *[]*string
	SetExcludeUserAccounts(val *[]*string)
	ExcludeUserAccountsInput() *[]*string
	ExclusionPatterns() *[]*string
	SetExclusionPatterns(val *[]*string)
	ExclusionPatternsInput() *[]*string
	FieldMappings() KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationFieldMappingsList
	FieldMappingsInput() interface{}
	// Experimental.
	Fqn() *string
	InclusionPatterns() *[]*string
	SetInclusionPatterns(val *[]*string)
	InclusionPatternsInput() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SecretArn() *string
	SetSecretArn(val *string)
	SecretArnInput() *string
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
	PutFieldMappings(value interface{})
	ResetExcludeMimeTypes()
	ResetExcludeSharedDrives()
	ResetExcludeUserAccounts()
	ResetExclusionPatterns()
	ResetFieldMappings()
	ResetInclusionPatterns()
	ResetSecretArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference
type jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference) ExcludeMimeTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeMimeTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference) ExcludeMimeTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeMimeTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference) ExcludeSharedDrives() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeSharedDrives",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference) ExcludeSharedDrivesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeSharedDrivesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference) ExcludeUserAccounts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeUserAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference) ExcludeUserAccountsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeUserAccountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference) ExclusionPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exclusionPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference) ExclusionPatternsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exclusionPatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference) FieldMappings() KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationFieldMappingsList {
	var returns KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationFieldMappingsList
	_jsii_.Get(
		j,
		"fieldMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference) FieldMappingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fieldMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference) InclusionPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inclusionPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference) InclusionPatternsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inclusionPatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference) SecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference) SecretArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewKendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.kendraDataSource.KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference_Override(k KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.kendraDataSource.KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference)SetExcludeMimeTypes(val *[]*string) {
	if err := j.validateSetExcludeMimeTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeMimeTypes",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference)SetExcludeSharedDrives(val *[]*string) {
	if err := j.validateSetExcludeSharedDrivesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeSharedDrives",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference)SetExcludeUserAccounts(val *[]*string) {
	if err := j.validateSetExcludeUserAccountsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeUserAccounts",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference)SetExclusionPatterns(val *[]*string) {
	if err := j.validateSetExclusionPatternsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exclusionPatterns",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference)SetInclusionPatterns(val *[]*string) {
	if err := j.validateSetInclusionPatternsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inclusionPatterns",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference)SetSecretArn(val *string) {
	if err := j.validateSetSecretArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretArn",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference) PutFieldMappings(value interface{}) {
	if err := k.validatePutFieldMappingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putFieldMappings",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference) ResetExcludeMimeTypes() {
	_jsii_.InvokeVoid(
		k,
		"resetExcludeMimeTypes",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference) ResetExcludeSharedDrives() {
	_jsii_.InvokeVoid(
		k,
		"resetExcludeSharedDrives",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference) ResetExcludeUserAccounts() {
	_jsii_.InvokeVoid(
		k,
		"resetExcludeUserAccounts",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference) ResetExclusionPatterns() {
	_jsii_.InvokeVoid(
		k,
		"resetExclusionPatterns",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference) ResetFieldMappings() {
	_jsii_.InvokeVoid(
		k,
		"resetFieldMappings",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference) ResetInclusionPatterns() {
	_jsii_.InvokeVoid(
		k,
		"resetInclusionPatterns",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference) ResetSecretArn() {
	_jsii_.InvokeVoid(
		k,
		"resetSecretArn",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := k.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KendraDataSourceDataSourceConfigurationGoogleDriveConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

