// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package licensemanagerlicenseassetruleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/licensemanagerlicenseassetruleset/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference interface {
	cdktn.ComplexObject
	AndRuleStatement() LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementAndRuleStatementOutputReference
	AndRuleStatementInput() interface{}
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
	MatchingRuleStatement() LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementMatchingRuleStatementOutputReference
	MatchingRuleStatementInput() interface{}
	OrRuleStatement() LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOrRuleStatementOutputReference
	OrRuleStatementInput() interface{}
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
	PutAndRuleStatement(value *LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementAndRuleStatement)
	PutMatchingRuleStatement(value *LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementMatchingRuleStatement)
	PutOrRuleStatement(value *LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOrRuleStatement)
	ResetAndRuleStatement()
	ResetMatchingRuleStatement()
	ResetOrRuleStatement()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference
type jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference) AndRuleStatement() LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementAndRuleStatementOutputReference {
	var returns LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementAndRuleStatementOutputReference
	_jsii_.Get(
		j,
		"andRuleStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference) AndRuleStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"andRuleStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference) MatchingRuleStatement() LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementMatchingRuleStatementOutputReference {
	var returns LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementMatchingRuleStatementOutputReference
	_jsii_.Get(
		j,
		"matchingRuleStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference) MatchingRuleStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"matchingRuleStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference) OrRuleStatement() LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOrRuleStatementOutputReference {
	var returns LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOrRuleStatementOutputReference
	_jsii_.Get(
		j,
		"orRuleStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference) OrRuleStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"orRuleStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference {
	_init_.Initialize()

	if err := validateNewLicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.licensemanagerLicenseAssetRuleSet.LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference_Override(l LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.licensemanagerLicenseAssetRuleSet.LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference) PutAndRuleStatement(value *LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementAndRuleStatement) {
	if err := l.validatePutAndRuleStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putAndRuleStatement",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference) PutMatchingRuleStatement(value *LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementMatchingRuleStatement) {
	if err := l.validatePutMatchingRuleStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putMatchingRuleStatement",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference) PutOrRuleStatement(value *LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOrRuleStatement) {
	if err := l.validatePutOrRuleStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putOrRuleStatement",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference) ResetAndRuleStatement() {
	_jsii_.InvokeVoid(
		l,
		"resetAndRuleStatement",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference) ResetMatchingRuleStatement() {
	_jsii_.InvokeVoid(
		l,
		"resetMatchingRuleStatement",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference) ResetOrRuleStatement() {
	_jsii_.InvokeVoid(
		l,
		"resetOrRuleStatement",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := l.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

