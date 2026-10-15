// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package licensemanagerlicenseconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/licensemanagerlicenseconfiguration/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ProductInformationFilterComparator() *string
	SetProductInformationFilterComparator(val *string)
	ProductInformationFilterComparatorInput() *string
	ProductInformationFilterName() *string
	SetProductInformationFilterName(val *string)
	ProductInformationFilterNameInput() *string
	ProductInformationFilterValue() *[]*string
	SetProductInformationFilterValue(val *[]*string)
	ProductInformationFilterValueInput() *[]*string
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
	ResetProductInformationFilterComparator()
	ResetProductInformationFilterName()
	ResetProductInformationFilterValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference
type jsiiProxy_LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference) ProductInformationFilterComparator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productInformationFilterComparator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference) ProductInformationFilterComparatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productInformationFilterComparatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference) ProductInformationFilterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productInformationFilterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference) ProductInformationFilterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productInformationFilterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference) ProductInformationFilterValue() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"productInformationFilterValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference) ProductInformationFilterValueInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"productInformationFilterValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference {
	_init_.Initialize()

	if err := validateNewLicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.licensemanagerLicenseConfiguration.LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference_Override(l LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.licensemanagerLicenseConfiguration.LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference)SetProductInformationFilterComparator(val *string) {
	if err := j.validateSetProductInformationFilterComparatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"productInformationFilterComparator",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference)SetProductInformationFilterName(val *string) {
	if err := j.validateSetProductInformationFilterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"productInformationFilterName",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference)SetProductInformationFilterValue(val *[]*string) {
	if err := j.validateSetProductInformationFilterValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"productInformationFilterValue",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (l *jsiiProxy_LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (l *jsiiProxy_LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference) ResetProductInformationFilterComparator() {
	_jsii_.InvokeVoid(
		l,
		"resetProductInformationFilterComparator",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference) ResetProductInformationFilterName() {
	_jsii_.InvokeVoid(
		l,
		"resetProductInformationFilterName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference) ResetProductInformationFilterValue() {
	_jsii_.InvokeVoid(
		l,
		"resetProductInformationFilterValue",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (l *jsiiProxy_LicensemanagerLicenseConfigurationProductInformationListProductInformationFilterListStructOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

