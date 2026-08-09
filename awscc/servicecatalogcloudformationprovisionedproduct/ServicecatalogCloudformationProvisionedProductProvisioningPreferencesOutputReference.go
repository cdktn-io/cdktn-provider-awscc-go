// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package servicecatalogcloudformationprovisionedproduct

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/servicecatalogcloudformationprovisionedproduct/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference interface {
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
	StackSetAccounts() *[]*string
	SetStackSetAccounts(val *[]*string)
	StackSetAccountsInput() *[]*string
	StackSetFailureToleranceCount() *float64
	SetStackSetFailureToleranceCount(val *float64)
	StackSetFailureToleranceCountInput() *float64
	StackSetFailureTolerancePercentage() *float64
	SetStackSetFailureTolerancePercentage(val *float64)
	StackSetFailureTolerancePercentageInput() *float64
	StackSetMaxConcurrencyCount() *float64
	SetStackSetMaxConcurrencyCount(val *float64)
	StackSetMaxConcurrencyCountInput() *float64
	StackSetMaxConcurrencyPercentage() *float64
	SetStackSetMaxConcurrencyPercentage(val *float64)
	StackSetMaxConcurrencyPercentageInput() *float64
	StackSetOperationType() *string
	SetStackSetOperationType(val *string)
	StackSetOperationTypeInput() *string
	StackSetRegions() *[]*string
	SetStackSetRegions(val *[]*string)
	StackSetRegionsInput() *[]*string
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
	ResetStackSetAccounts()
	ResetStackSetFailureToleranceCount()
	ResetStackSetFailureTolerancePercentage()
	ResetStackSetMaxConcurrencyCount()
	ResetStackSetMaxConcurrencyPercentage()
	ResetStackSetOperationType()
	ResetStackSetRegions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference
type jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference) StackSetAccounts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"stackSetAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference) StackSetAccountsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"stackSetAccountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference) StackSetFailureToleranceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stackSetFailureToleranceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference) StackSetFailureToleranceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stackSetFailureToleranceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference) StackSetFailureTolerancePercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stackSetFailureTolerancePercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference) StackSetFailureTolerancePercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stackSetFailureTolerancePercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference) StackSetMaxConcurrencyCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stackSetMaxConcurrencyCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference) StackSetMaxConcurrencyCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stackSetMaxConcurrencyCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference) StackSetMaxConcurrencyPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stackSetMaxConcurrencyPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference) StackSetMaxConcurrencyPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"stackSetMaxConcurrencyPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference) StackSetOperationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackSetOperationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference) StackSetOperationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackSetOperationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference) StackSetRegions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"stackSetRegions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference) StackSetRegionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"stackSetRegionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference {
	_init_.Initialize()

	if err := validateNewServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.servicecatalogCloudformationProvisionedProduct.ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference_Override(s ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.servicecatalogCloudformationProvisionedProduct.ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference)SetStackSetAccounts(val *[]*string) {
	if err := j.validateSetStackSetAccountsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stackSetAccounts",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference)SetStackSetFailureToleranceCount(val *float64) {
	if err := j.validateSetStackSetFailureToleranceCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stackSetFailureToleranceCount",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference)SetStackSetFailureTolerancePercentage(val *float64) {
	if err := j.validateSetStackSetFailureTolerancePercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stackSetFailureTolerancePercentage",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference)SetStackSetMaxConcurrencyCount(val *float64) {
	if err := j.validateSetStackSetMaxConcurrencyCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stackSetMaxConcurrencyCount",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference)SetStackSetMaxConcurrencyPercentage(val *float64) {
	if err := j.validateSetStackSetMaxConcurrencyPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stackSetMaxConcurrencyPercentage",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference)SetStackSetOperationType(val *string) {
	if err := j.validateSetStackSetOperationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stackSetOperationType",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference)SetStackSetRegions(val *[]*string) {
	if err := j.validateSetStackSetRegionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stackSetRegions",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference) ResetStackSetAccounts() {
	_jsii_.InvokeVoid(
		s,
		"resetStackSetAccounts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference) ResetStackSetFailureToleranceCount() {
	_jsii_.InvokeVoid(
		s,
		"resetStackSetFailureToleranceCount",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference) ResetStackSetFailureTolerancePercentage() {
	_jsii_.InvokeVoid(
		s,
		"resetStackSetFailureTolerancePercentage",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference) ResetStackSetMaxConcurrencyCount() {
	_jsii_.InvokeVoid(
		s,
		"resetStackSetMaxConcurrencyCount",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference) ResetStackSetMaxConcurrencyPercentage() {
	_jsii_.InvokeVoid(
		s,
		"resetStackSetMaxConcurrencyPercentage",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference) ResetStackSetOperationType() {
	_jsii_.InvokeVoid(
		s,
		"resetStackSetOperationType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference) ResetStackSetRegions() {
	_jsii_.InvokeVoid(
		s,
		"resetStackSetRegions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServicecatalogCloudformationProvisionedProductProvisioningPreferencesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

