// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package licensemanagerlicense

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/licensemanagerlicense/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LicensemanagerLicenseConsumptionConfigurationOutputReference interface {
	cdktn.ComplexObject
	BorrowConfiguration() LicensemanagerLicenseConsumptionConfigurationBorrowConfigurationOutputReference
	BorrowConfigurationInput() interface{}
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
	ProvisionalConfiguration() LicensemanagerLicenseConsumptionConfigurationProvisionalConfigurationOutputReference
	ProvisionalConfigurationInput() interface{}
	RenewType() *string
	SetRenewType(val *string)
	RenewTypeInput() *string
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
	PutBorrowConfiguration(value *LicensemanagerLicenseConsumptionConfigurationBorrowConfiguration)
	PutProvisionalConfiguration(value *LicensemanagerLicenseConsumptionConfigurationProvisionalConfiguration)
	ResetBorrowConfiguration()
	ResetProvisionalConfiguration()
	ResetRenewType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LicensemanagerLicenseConsumptionConfigurationOutputReference
type jsiiProxy_LicensemanagerLicenseConsumptionConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_LicensemanagerLicenseConsumptionConfigurationOutputReference) BorrowConfiguration() LicensemanagerLicenseConsumptionConfigurationBorrowConfigurationOutputReference {
	var returns LicensemanagerLicenseConsumptionConfigurationBorrowConfigurationOutputReference
	_jsii_.Get(
		j,
		"borrowConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConsumptionConfigurationOutputReference) BorrowConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"borrowConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConsumptionConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConsumptionConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConsumptionConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConsumptionConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConsumptionConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConsumptionConfigurationOutputReference) ProvisionalConfiguration() LicensemanagerLicenseConsumptionConfigurationProvisionalConfigurationOutputReference {
	var returns LicensemanagerLicenseConsumptionConfigurationProvisionalConfigurationOutputReference
	_jsii_.Get(
		j,
		"provisionalConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConsumptionConfigurationOutputReference) ProvisionalConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisionalConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConsumptionConfigurationOutputReference) RenewType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"renewType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConsumptionConfigurationOutputReference) RenewTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"renewTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConsumptionConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConsumptionConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLicensemanagerLicenseConsumptionConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) LicensemanagerLicenseConsumptionConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewLicensemanagerLicenseConsumptionConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LicensemanagerLicenseConsumptionConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.licensemanagerLicense.LicensemanagerLicenseConsumptionConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLicensemanagerLicenseConsumptionConfigurationOutputReference_Override(l LicensemanagerLicenseConsumptionConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.licensemanagerLicense.LicensemanagerLicenseConsumptionConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LicensemanagerLicenseConsumptionConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicenseConsumptionConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicenseConsumptionConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicenseConsumptionConfigurationOutputReference)SetRenewType(val *string) {
	if err := j.validateSetRenewTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"renewType",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicenseConsumptionConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicenseConsumptionConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LicensemanagerLicenseConsumptionConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LicensemanagerLicenseConsumptionConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LicensemanagerLicenseConsumptionConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (l *jsiiProxy_LicensemanagerLicenseConsumptionConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LicensemanagerLicenseConsumptionConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LicensemanagerLicenseConsumptionConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LicensemanagerLicenseConsumptionConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LicensemanagerLicenseConsumptionConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LicensemanagerLicenseConsumptionConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LicensemanagerLicenseConsumptionConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LicensemanagerLicenseConsumptionConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LicensemanagerLicenseConsumptionConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (l *jsiiProxy_LicensemanagerLicenseConsumptionConfigurationOutputReference) PutBorrowConfiguration(value *LicensemanagerLicenseConsumptionConfigurationBorrowConfiguration) {
	if err := l.validatePutBorrowConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putBorrowConfiguration",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LicensemanagerLicenseConsumptionConfigurationOutputReference) PutProvisionalConfiguration(value *LicensemanagerLicenseConsumptionConfigurationProvisionalConfiguration) {
	if err := l.validatePutProvisionalConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putProvisionalConfiguration",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LicensemanagerLicenseConsumptionConfigurationOutputReference) ResetBorrowConfiguration() {
	_jsii_.InvokeVoid(
		l,
		"resetBorrowConfiguration",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LicensemanagerLicenseConsumptionConfigurationOutputReference) ResetProvisionalConfiguration() {
	_jsii_.InvokeVoid(
		l,
		"resetProvisionalConfiguration",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LicensemanagerLicenseConsumptionConfigurationOutputReference) ResetRenewType() {
	_jsii_.InvokeVoid(
		l,
		"resetRenewType",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LicensemanagerLicenseConsumptionConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (l *jsiiProxy_LicensemanagerLicenseConsumptionConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

