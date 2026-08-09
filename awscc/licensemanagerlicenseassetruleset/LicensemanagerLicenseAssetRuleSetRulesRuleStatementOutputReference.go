// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package licensemanagerlicenseassetruleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/licensemanagerlicenseassetruleset/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InstanceRuleStatement() LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference
	InstanceRuleStatementInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LicenseConfigurationRuleStatement() LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseConfigurationRuleStatementOutputReference
	LicenseConfigurationRuleStatementInput() interface{}
	LicenseRuleStatement() LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference
	LicenseRuleStatementInput() interface{}
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
	PutInstanceRuleStatement(value *LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatement)
	PutLicenseConfigurationRuleStatement(value *LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseConfigurationRuleStatement)
	PutLicenseRuleStatement(value *LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatement)
	ResetInstanceRuleStatement()
	ResetLicenseConfigurationRuleStatement()
	ResetLicenseRuleStatement()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference
type jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference) InstanceRuleStatement() LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference {
	var returns LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatementOutputReference
	_jsii_.Get(
		j,
		"instanceRuleStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference) InstanceRuleStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceRuleStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference) LicenseConfigurationRuleStatement() LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseConfigurationRuleStatementOutputReference {
	var returns LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseConfigurationRuleStatementOutputReference
	_jsii_.Get(
		j,
		"licenseConfigurationRuleStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference) LicenseConfigurationRuleStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"licenseConfigurationRuleStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference) LicenseRuleStatement() LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference {
	var returns LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatementOutputReference
	_jsii_.Get(
		j,
		"licenseRuleStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference) LicenseRuleStatementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"licenseRuleStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference {
	_init_.Initialize()

	if err := validateNewLicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.licensemanagerLicenseAssetRuleSet.LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference_Override(l LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.licensemanagerLicenseAssetRuleSet.LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference) PutInstanceRuleStatement(value *LicensemanagerLicenseAssetRuleSetRulesRuleStatementInstanceRuleStatement) {
	if err := l.validatePutInstanceRuleStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putInstanceRuleStatement",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference) PutLicenseConfigurationRuleStatement(value *LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseConfigurationRuleStatement) {
	if err := l.validatePutLicenseConfigurationRuleStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putLicenseConfigurationRuleStatement",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference) PutLicenseRuleStatement(value *LicensemanagerLicenseAssetRuleSetRulesRuleStatementLicenseRuleStatement) {
	if err := l.validatePutLicenseRuleStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putLicenseRuleStatement",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference) ResetInstanceRuleStatement() {
	_jsii_.InvokeVoid(
		l,
		"resetInstanceRuleStatement",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference) ResetLicenseConfigurationRuleStatement() {
	_jsii_.InvokeVoid(
		l,
		"resetLicenseConfigurationRuleStatement",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference) ResetLicenseRuleStatement() {
	_jsii_.InvokeVoid(
		l,
		"resetLicenseRuleStatement",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (l *jsiiProxy_LicensemanagerLicenseAssetRuleSetRulesRuleStatementOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

