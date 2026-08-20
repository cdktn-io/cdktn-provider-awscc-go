// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdlpsetting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/quicksightdlpsetting/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference interface {
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
	Credentials() QuicksightDlpSettingProviderConfigMicrosoftPurviewCredentialsOutputReference
	CredentialsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LabelActionMappings() QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsList
	LabelActionMappingsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UnmappedAction() *string
	SetUnmappedAction(val *string)
	UnmappedActionInput() *string
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
	PutCredentials(value *QuicksightDlpSettingProviderConfigMicrosoftPurviewCredentials)
	PutLabelActionMappings(value interface{})
	ResetCredentials()
	ResetLabelActionMappings()
	ResetUnmappedAction()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference
type jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference) Credentials() QuicksightDlpSettingProviderConfigMicrosoftPurviewCredentialsOutputReference {
	var returns QuicksightDlpSettingProviderConfigMicrosoftPurviewCredentialsOutputReference
	_jsii_.Get(
		j,
		"credentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference) CredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"credentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference) LabelActionMappings() QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsList {
	var returns QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsList
	_jsii_.Get(
		j,
		"labelActionMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference) LabelActionMappingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"labelActionMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference) UnmappedAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unmappedAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference) UnmappedActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unmappedActionInput",
		&returns,
	)
	return returns
}


func NewQuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightDlpSetting.QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewQuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference_Override(q QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightDlpSetting.QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		q,
	)
}

func (j *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference)SetUnmappedAction(val *string) {
	if err := j.validateSetUnmappedActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unmappedAction",
		val,
	)
}

func (q *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := q.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := q.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := q.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		q,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := q.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := q.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := q.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		q,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := q.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		q,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := q.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := q.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		q,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := q.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference) PutCredentials(value *QuicksightDlpSettingProviderConfigMicrosoftPurviewCredentials) {
	if err := q.validatePutCredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putCredentials",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference) PutLabelActionMappings(value interface{}) {
	if err := q.validatePutLabelActionMappingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putLabelActionMappings",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference) ResetCredentials() {
	_jsii_.InvokeVoid(
		q,
		"resetCredentials",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference) ResetLabelActionMappings() {
	_jsii_.InvokeVoid(
		q,
		"resetLabelActionMappings",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference) ResetUnmappedAction() {
	_jsii_.InvokeVoid(
		q,
		"resetUnmappedAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := q.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		q,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

