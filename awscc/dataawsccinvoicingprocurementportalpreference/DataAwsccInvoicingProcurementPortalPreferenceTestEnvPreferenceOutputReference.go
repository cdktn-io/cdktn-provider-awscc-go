// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccinvoicingprocurementportalpreference

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccinvoicingprocurementportalpreference/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference interface {
	cdktn.ComplexObject
	BuyerDomain() *string
	BuyerIdentifier() *string
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
	InternalValue() *DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreference
	SetInternalValue(val *DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreference)
	ProcurementPortalInstanceEndpoint() *string
	ProcurementPortalSharedSecret() *string
	SupplierDomain() *string
	SupplierIdentifier() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference
type jsiiProxy_DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) BuyerDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buyerDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) BuyerIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buyerIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) InternalValue() *DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreference {
	var returns *DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreference
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) ProcurementPortalInstanceEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"procurementPortalInstanceEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) ProcurementPortalSharedSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"procurementPortalSharedSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) SupplierDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supplierDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) SupplierIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supplierIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccInvoicingProcurementPortalPreference.DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference_Override(d DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccInvoicingProcurementPortalPreference.DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference)SetInternalValue(val *DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreference) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

