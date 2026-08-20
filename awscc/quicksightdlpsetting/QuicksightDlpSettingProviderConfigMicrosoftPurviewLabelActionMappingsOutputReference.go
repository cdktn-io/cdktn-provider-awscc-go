// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdlpsetting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/quicksightdlpsetting/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference interface {
	cdktn.ComplexObject
	Action() *string
	SetAction(val *string)
	ActionInput() *string
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LabelId() *string
	SetLabelId(val *string)
	LabelIdInput() *string
	LabelName() *string
	SetLabelName(val *string)
	LabelNameInput() *string
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
	ResetAction()
	ResetLabelId()
	ResetLabelName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference
type jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference) LabelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference) LabelIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference) LabelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference) LabelNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewQuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightDlpSetting.QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewQuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference_Override(q QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightDlpSetting.QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		q,
	)
}

func (j *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference)SetAction(val *string) {
	if err := j.validateSetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference)SetLabelId(val *string) {
	if err := j.validateSetLabelIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labelId",
		val,
	)
}

func (j *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference)SetLabelName(val *string) {
	if err := j.validateSetLabelNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labelName",
		val,
	)
}

func (j *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (q *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (q *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (q *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (q *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (q *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (q *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (q *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (q *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (q *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference) ResetAction() {
	_jsii_.InvokeVoid(
		q,
		"resetAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference) ResetLabelId() {
	_jsii_.InvokeVoid(
		q,
		"resetLabelId",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference) ResetLabelName() {
	_jsii_.InvokeVoid(
		q,
		"resetLabelName",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (q *jsiiProxy_QuicksightDlpSettingProviderConfigMicrosoftPurviewLabelActionMappingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

