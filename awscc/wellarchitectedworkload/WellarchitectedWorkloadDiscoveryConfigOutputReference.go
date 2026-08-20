// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wellarchitectedworkload

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/wellarchitectedworkload/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WellarchitectedWorkloadDiscoveryConfigOutputReference interface {
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
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TrustedAdvisorIntegrationStatus() *string
	SetTrustedAdvisorIntegrationStatus(val *string)
	TrustedAdvisorIntegrationStatusInput() *string
	WorkloadResourceDefinition() *[]*string
	SetWorkloadResourceDefinition(val *[]*string)
	WorkloadResourceDefinitionInput() *[]*string
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
	ResetTrustedAdvisorIntegrationStatus()
	ResetWorkloadResourceDefinition()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WellarchitectedWorkloadDiscoveryConfigOutputReference
type jsiiProxy_WellarchitectedWorkloadDiscoveryConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_WellarchitectedWorkloadDiscoveryConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WellarchitectedWorkloadDiscoveryConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WellarchitectedWorkloadDiscoveryConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WellarchitectedWorkloadDiscoveryConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WellarchitectedWorkloadDiscoveryConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WellarchitectedWorkloadDiscoveryConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WellarchitectedWorkloadDiscoveryConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WellarchitectedWorkloadDiscoveryConfigOutputReference) TrustedAdvisorIntegrationStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustedAdvisorIntegrationStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WellarchitectedWorkloadDiscoveryConfigOutputReference) TrustedAdvisorIntegrationStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustedAdvisorIntegrationStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WellarchitectedWorkloadDiscoveryConfigOutputReference) WorkloadResourceDefinition() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"workloadResourceDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WellarchitectedWorkloadDiscoveryConfigOutputReference) WorkloadResourceDefinitionInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"workloadResourceDefinitionInput",
		&returns,
	)
	return returns
}


func NewWellarchitectedWorkloadDiscoveryConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) WellarchitectedWorkloadDiscoveryConfigOutputReference {
	_init_.Initialize()

	if err := validateNewWellarchitectedWorkloadDiscoveryConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WellarchitectedWorkloadDiscoveryConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.wellarchitectedWorkload.WellarchitectedWorkloadDiscoveryConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWellarchitectedWorkloadDiscoveryConfigOutputReference_Override(w WellarchitectedWorkloadDiscoveryConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.wellarchitectedWorkload.WellarchitectedWorkloadDiscoveryConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WellarchitectedWorkloadDiscoveryConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WellarchitectedWorkloadDiscoveryConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WellarchitectedWorkloadDiscoveryConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WellarchitectedWorkloadDiscoveryConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WellarchitectedWorkloadDiscoveryConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WellarchitectedWorkloadDiscoveryConfigOutputReference)SetTrustedAdvisorIntegrationStatus(val *string) {
	if err := j.validateSetTrustedAdvisorIntegrationStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustedAdvisorIntegrationStatus",
		val,
	)
}

func (j *jsiiProxy_WellarchitectedWorkloadDiscoveryConfigOutputReference)SetWorkloadResourceDefinition(val *[]*string) {
	if err := j.validateSetWorkloadResourceDefinitionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workloadResourceDefinition",
		val,
	)
}

func (w *jsiiProxy_WellarchitectedWorkloadDiscoveryConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WellarchitectedWorkloadDiscoveryConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WellarchitectedWorkloadDiscoveryConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WellarchitectedWorkloadDiscoveryConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WellarchitectedWorkloadDiscoveryConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WellarchitectedWorkloadDiscoveryConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WellarchitectedWorkloadDiscoveryConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WellarchitectedWorkloadDiscoveryConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WellarchitectedWorkloadDiscoveryConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WellarchitectedWorkloadDiscoveryConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WellarchitectedWorkloadDiscoveryConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WellarchitectedWorkloadDiscoveryConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WellarchitectedWorkloadDiscoveryConfigOutputReference) ResetTrustedAdvisorIntegrationStatus() {
	_jsii_.InvokeVoid(
		w,
		"resetTrustedAdvisorIntegrationStatus",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WellarchitectedWorkloadDiscoveryConfigOutputReference) ResetWorkloadResourceDefinition() {
	_jsii_.InvokeVoid(
		w,
		"resetWorkloadResourceDefinition",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WellarchitectedWorkloadDiscoveryConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := w.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WellarchitectedWorkloadDiscoveryConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

