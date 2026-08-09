// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package entityresolutionmatchingworkflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/entityresolutionmatchingworkflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference interface {
	cdktn.ComplexObject
	ApplyNormalization() interface{}
	SetApplyNormalization(val interface{})
	ApplyNormalizationInput() interface{}
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
	CustomerProfilesIntegrationConfig() EntityresolutionMatchingWorkflowOutputSourceConfigCustomerProfilesIntegrationConfigOutputReference
	CustomerProfilesIntegrationConfigInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KmsArn() *string
	SetKmsArn(val *string)
	KmsArnInput() *string
	Output() EntityresolutionMatchingWorkflowOutputSourceConfigOutputList
	OutputInput() interface{}
	OutputS3Path() *string
	SetOutputS3Path(val *string)
	OutputS3PathInput() *string
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
	PutCustomerProfilesIntegrationConfig(value *EntityresolutionMatchingWorkflowOutputSourceConfigCustomerProfilesIntegrationConfig)
	PutOutput(value interface{})
	ResetApplyNormalization()
	ResetCustomerProfilesIntegrationConfig()
	ResetKmsArn()
	ResetOutputS3Path()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference
type jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference) ApplyNormalization() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyNormalization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference) ApplyNormalizationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyNormalizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference) CustomerProfilesIntegrationConfig() EntityresolutionMatchingWorkflowOutputSourceConfigCustomerProfilesIntegrationConfigOutputReference {
	var returns EntityresolutionMatchingWorkflowOutputSourceConfigCustomerProfilesIntegrationConfigOutputReference
	_jsii_.Get(
		j,
		"customerProfilesIntegrationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference) CustomerProfilesIntegrationConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customerProfilesIntegrationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference) KmsArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference) KmsArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference) Output() EntityresolutionMatchingWorkflowOutputSourceConfigOutputList {
	var returns EntityresolutionMatchingWorkflowOutputSourceConfigOutputList
	_jsii_.Get(
		j,
		"output",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference) OutputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference) OutputS3Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputS3Path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference) OutputS3PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputS3PathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEntityresolutionMatchingWorkflowOutputSourceConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference {
	_init_.Initialize()

	if err := validateNewEntityresolutionMatchingWorkflowOutputSourceConfigOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.entityresolutionMatchingWorkflow.EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEntityresolutionMatchingWorkflowOutputSourceConfigOutputReference_Override(e EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.entityresolutionMatchingWorkflow.EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference)SetApplyNormalization(val interface{}) {
	if err := j.validateSetApplyNormalizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applyNormalization",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference)SetKmsArn(val *string) {
	if err := j.validateSetKmsArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsArn",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference)SetOutputS3Path(val *string) {
	if err := j.validateSetOutputS3PathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputS3Path",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference) PutCustomerProfilesIntegrationConfig(value *EntityresolutionMatchingWorkflowOutputSourceConfigCustomerProfilesIntegrationConfig) {
	if err := e.validatePutCustomerProfilesIntegrationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putCustomerProfilesIntegrationConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference) PutOutput(value interface{}) {
	if err := e.validatePutOutputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putOutput",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference) ResetApplyNormalization() {
	_jsii_.InvokeVoid(
		e,
		"resetApplyNormalization",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference) ResetCustomerProfilesIntegrationConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetCustomerProfilesIntegrationConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference) ResetKmsArn() {
	_jsii_.InvokeVoid(
		e,
		"resetKmsArn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference) ResetOutputS3Path() {
	_jsii_.InvokeVoid(
		e,
		"resetOutputS3Path",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := e.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflowOutputSourceConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

