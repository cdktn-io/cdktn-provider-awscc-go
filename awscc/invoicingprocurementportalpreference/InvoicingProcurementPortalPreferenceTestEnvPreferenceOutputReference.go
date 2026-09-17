// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package invoicingprocurementportalpreference

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/invoicingprocurementportalpreference/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference interface {
	cdktn.ComplexObject
	BuyerDomain() *string
	SetBuyerDomain(val *string)
	BuyerDomainInput() *string
	BuyerIdentifier() *string
	SetBuyerIdentifier(val *string)
	BuyerIdentifierInput() *string
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
	ProcurementPortalInstanceEndpoint() *string
	SetProcurementPortalInstanceEndpoint(val *string)
	ProcurementPortalInstanceEndpointInput() *string
	ProcurementPortalSharedSecret() *string
	SetProcurementPortalSharedSecret(val *string)
	ProcurementPortalSharedSecretInput() *string
	SupplierDomain() *string
	SetSupplierDomain(val *string)
	SupplierDomainInput() *string
	SupplierIdentifier() *string
	SetSupplierIdentifier(val *string)
	SupplierIdentifierInput() *string
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
	ResetBuyerDomain()
	ResetBuyerIdentifier()
	ResetProcurementPortalInstanceEndpoint()
	ResetProcurementPortalSharedSecret()
	ResetSupplierDomain()
	ResetSupplierIdentifier()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference
type jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) BuyerDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buyerDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) BuyerDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buyerDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) BuyerIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buyerIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) BuyerIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buyerIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) ProcurementPortalInstanceEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"procurementPortalInstanceEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) ProcurementPortalInstanceEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"procurementPortalInstanceEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) ProcurementPortalSharedSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"procurementPortalSharedSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) ProcurementPortalSharedSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"procurementPortalSharedSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) SupplierDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supplierDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) SupplierDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supplierDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) SupplierIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supplierIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) SupplierIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supplierIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference {
	_init_.Initialize()

	if err := validateNewInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.invoicingProcurementPortalPreference.InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewInvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference_Override(i InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.invoicingProcurementPortalPreference.InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference)SetBuyerDomain(val *string) {
	if err := j.validateSetBuyerDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buyerDomain",
		val,
	)
}

func (j *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference)SetBuyerIdentifier(val *string) {
	if err := j.validateSetBuyerIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buyerIdentifier",
		val,
	)
}

func (j *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference)SetProcurementPortalInstanceEndpoint(val *string) {
	if err := j.validateSetProcurementPortalInstanceEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"procurementPortalInstanceEndpoint",
		val,
	)
}

func (j *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference)SetProcurementPortalSharedSecret(val *string) {
	if err := j.validateSetProcurementPortalSharedSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"procurementPortalSharedSecret",
		val,
	)
}

func (j *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference)SetSupplierDomain(val *string) {
	if err := j.validateSetSupplierDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supplierDomain",
		val,
	)
}

func (j *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference)SetSupplierIdentifier(val *string) {
	if err := j.validateSetSupplierIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supplierIdentifier",
		val,
	)
}

func (j *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) ResetBuyerDomain() {
	_jsii_.InvokeVoid(
		i,
		"resetBuyerDomain",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) ResetBuyerIdentifier() {
	_jsii_.InvokeVoid(
		i,
		"resetBuyerIdentifier",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) ResetProcurementPortalInstanceEndpoint() {
	_jsii_.InvokeVoid(
		i,
		"resetProcurementPortalInstanceEndpoint",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) ResetProcurementPortalSharedSecret() {
	_jsii_.InvokeVoid(
		i,
		"resetProcurementPortalSharedSecret",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) ResetSupplierDomain() {
	_jsii_.InvokeVoid(
		i,
		"resetSupplierDomain",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) ResetSupplierIdentifier() {
	_jsii_.InvokeVoid(
		i,
		"resetSupplierIdentifier",
		nil, // no parameters
	)
}

func (i *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := i.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InvoicingProcurementPortalPreferenceTestEnvPreferenceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

