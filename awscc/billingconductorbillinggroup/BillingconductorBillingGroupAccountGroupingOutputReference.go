// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package billingconductorbillinggroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/billingconductorbillinggroup/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BillingconductorBillingGroupAccountGroupingOutputReference interface {
	cdktn.ComplexObject
	AutoAssociate() interface{}
	SetAutoAssociate(val interface{})
	AutoAssociateInput() interface{}
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
	LinkedAccountIds() *[]*string
	SetLinkedAccountIds(val *[]*string)
	LinkedAccountIdsInput() *[]*string
	ResponsibilityTransferArn() *string
	SetResponsibilityTransferArn(val *string)
	ResponsibilityTransferArnInput() *string
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
	ResetAutoAssociate()
	ResetLinkedAccountIds()
	ResetResponsibilityTransferArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BillingconductorBillingGroupAccountGroupingOutputReference
type jsiiProxy_BillingconductorBillingGroupAccountGroupingOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BillingconductorBillingGroupAccountGroupingOutputReference) AutoAssociate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssociate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorBillingGroupAccountGroupingOutputReference) AutoAssociateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAssociateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorBillingGroupAccountGroupingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorBillingGroupAccountGroupingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorBillingGroupAccountGroupingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorBillingGroupAccountGroupingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorBillingGroupAccountGroupingOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorBillingGroupAccountGroupingOutputReference) LinkedAccountIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"linkedAccountIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorBillingGroupAccountGroupingOutputReference) LinkedAccountIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"linkedAccountIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorBillingGroupAccountGroupingOutputReference) ResponsibilityTransferArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responsibilityTransferArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorBillingGroupAccountGroupingOutputReference) ResponsibilityTransferArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responsibilityTransferArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorBillingGroupAccountGroupingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingconductorBillingGroupAccountGroupingOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBillingconductorBillingGroupAccountGroupingOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BillingconductorBillingGroupAccountGroupingOutputReference {
	_init_.Initialize()

	if err := validateNewBillingconductorBillingGroupAccountGroupingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BillingconductorBillingGroupAccountGroupingOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.billingconductorBillingGroup.BillingconductorBillingGroupAccountGroupingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBillingconductorBillingGroupAccountGroupingOutputReference_Override(b BillingconductorBillingGroupAccountGroupingOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.billingconductorBillingGroup.BillingconductorBillingGroupAccountGroupingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BillingconductorBillingGroupAccountGroupingOutputReference)SetAutoAssociate(val interface{}) {
	if err := j.validateSetAutoAssociateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoAssociate",
		val,
	)
}

func (j *jsiiProxy_BillingconductorBillingGroupAccountGroupingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BillingconductorBillingGroupAccountGroupingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BillingconductorBillingGroupAccountGroupingOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BillingconductorBillingGroupAccountGroupingOutputReference)SetLinkedAccountIds(val *[]*string) {
	if err := j.validateSetLinkedAccountIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"linkedAccountIds",
		val,
	)
}

func (j *jsiiProxy_BillingconductorBillingGroupAccountGroupingOutputReference)SetResponsibilityTransferArn(val *string) {
	if err := j.validateSetResponsibilityTransferArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responsibilityTransferArn",
		val,
	)
}

func (j *jsiiProxy_BillingconductorBillingGroupAccountGroupingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BillingconductorBillingGroupAccountGroupingOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BillingconductorBillingGroupAccountGroupingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingconductorBillingGroupAccountGroupingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BillingconductorBillingGroupAccountGroupingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BillingconductorBillingGroupAccountGroupingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BillingconductorBillingGroupAccountGroupingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BillingconductorBillingGroupAccountGroupingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BillingconductorBillingGroupAccountGroupingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BillingconductorBillingGroupAccountGroupingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BillingconductorBillingGroupAccountGroupingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BillingconductorBillingGroupAccountGroupingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BillingconductorBillingGroupAccountGroupingOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingconductorBillingGroupAccountGroupingOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BillingconductorBillingGroupAccountGroupingOutputReference) ResetAutoAssociate() {
	_jsii_.InvokeVoid(
		b,
		"resetAutoAssociate",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BillingconductorBillingGroupAccountGroupingOutputReference) ResetLinkedAccountIds() {
	_jsii_.InvokeVoid(
		b,
		"resetLinkedAccountIds",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BillingconductorBillingGroupAccountGroupingOutputReference) ResetResponsibilityTransferArn() {
	_jsii_.InvokeVoid(
		b,
		"resetResponsibilityTransferArn",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BillingconductorBillingGroupAccountGroupingOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BillingconductorBillingGroupAccountGroupingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

