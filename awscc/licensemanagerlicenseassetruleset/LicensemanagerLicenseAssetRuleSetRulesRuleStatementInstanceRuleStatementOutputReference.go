// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package licensemanagerlicenseassetruleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/licensemanagerlicenseassetruleset/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference interface {
	cdktn.ComplexObject
	AndRuleStatement() LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementAndRuleStatementOutputReference
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
	MatchingRuleStatement() LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementMatchingRuleStatementOutputReference
	MatchingRuleStatementInput() interface{}
	OrRuleStatement() LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOrRuleStatementOutputReference
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
	PutAndRuleStatement(value *LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementAndRuleStatement)
	PutMatchingRuleStatement(value *LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementMatchingRuleStatement)
	PutOrRuleStatement(value *LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOrRuleStatement)
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

// The jsii proxy struct for LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference
type jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference) AndRuleStatement() LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementAndRuleStatementOutputReference {
	var returns LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementAndRuleStatementOutputReference
	_jsii_.Get(
		j,
		"andRuleStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference) AndRuleStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"andRuleStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference) MatchingRuleStatement() LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementMatchingRuleStatementOutputReference {
	var returns LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementMatchingRuleStatementOutputReference
	_jsii_.Get(
		j,
		"matchingRuleStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference) MatchingRuleStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"matchingRuleStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference) OrRuleStatement() LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOrRuleStatementOutputReference {
	var returns LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOrRuleStatementOutputReference
	_jsii_.Get(
		j,
		"orRuleStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference) OrRuleStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"orRuleStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference {
	_init_.Initialize()

	if err := validateNewLicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.licensemanagerLicenseAssetRuleSet.LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference_Override(l LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.licensemanagerLicenseAssetRuleSet.LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference) PutAndRuleStatement(value *LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementAndRuleStatement) {
	if err := l.validatePutAndRuleStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putAndRuleStatement",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference) PutMatchingRuleStatement(value *LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementMatchingRuleStatement) {
	if err := l.validatePutMatchingRuleStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putMatchingRuleStatement",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference) PutOrRuleStatement(value *LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOrRuleStatement) {
	if err := l.validatePutOrRuleStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putOrRuleStatement",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference) ResetAndRuleStatement() {
	_jsii_.InvokeVoid(
		l,
		"resetAndRuleStatement",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference) ResetMatchingRuleStatement() {
	_jsii_.InvokeVoid(
		l,
		"resetMatchingRuleStatement",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference) ResetOrRuleStatement() {
	_jsii_.InvokeVoid(
		l,
		"resetOrRuleStatement",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

