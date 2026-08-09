// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package greengrassv2deployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/greengrassv2deployment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Greengrassv2DeploymentIotJobConfigurationOutputReference interface {
	cdktn.ComplexObject
	AbortConfig() Greengrassv2DeploymentIotJobConfigurationAbortConfigOutputReference
	AbortConfigInput() interface{}
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
	JobExecutionsRolloutConfig() Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigOutputReference
	JobExecutionsRolloutConfigInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeoutConfig() Greengrassv2DeploymentIotJobConfigurationTimeoutConfigOutputReference
	TimeoutConfigInput() interface{}
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
	PutAbortConfig(value *Greengrassv2DeploymentIotJobConfigurationAbortConfig)
	PutJobExecutionsRolloutConfig(value *Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfig)
	PutTimeoutConfig(value *Greengrassv2DeploymentIotJobConfigurationTimeoutConfig)
	ResetAbortConfig()
	ResetJobExecutionsRolloutConfig()
	ResetTimeoutConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Greengrassv2DeploymentIotJobConfigurationOutputReference
type jsiiProxy_Greengrassv2DeploymentIotJobConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationOutputReference) AbortConfig() Greengrassv2DeploymentIotJobConfigurationAbortConfigOutputReference {
	var returns Greengrassv2DeploymentIotJobConfigurationAbortConfigOutputReference
	_jsii_.Get(
		j,
		"abortConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationOutputReference) AbortConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"abortConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationOutputReference) JobExecutionsRolloutConfig() Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigOutputReference {
	var returns Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfigOutputReference
	_jsii_.Get(
		j,
		"jobExecutionsRolloutConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationOutputReference) JobExecutionsRolloutConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jobExecutionsRolloutConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationOutputReference) TimeoutConfig() Greengrassv2DeploymentIotJobConfigurationTimeoutConfigOutputReference {
	var returns Greengrassv2DeploymentIotJobConfigurationTimeoutConfigOutputReference
	_jsii_.Get(
		j,
		"timeoutConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationOutputReference) TimeoutConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutConfigInput",
		&returns,
	)
	return returns
}


func NewGreengrassv2DeploymentIotJobConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) Greengrassv2DeploymentIotJobConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewGreengrassv2DeploymentIotJobConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Greengrassv2DeploymentIotJobConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.greengrassv2Deployment.Greengrassv2DeploymentIotJobConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGreengrassv2DeploymentIotJobConfigurationOutputReference_Override(g Greengrassv2DeploymentIotJobConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.greengrassv2Deployment.Greengrassv2DeploymentIotJobConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationOutputReference) PutAbortConfig(value *Greengrassv2DeploymentIotJobConfigurationAbortConfig) {
	if err := g.validatePutAbortConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAbortConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationOutputReference) PutJobExecutionsRolloutConfig(value *Greengrassv2DeploymentIotJobConfigurationJobExecutionsRolloutConfig) {
	if err := g.validatePutJobExecutionsRolloutConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putJobExecutionsRolloutConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationOutputReference) PutTimeoutConfig(value *Greengrassv2DeploymentIotJobConfigurationTimeoutConfig) {
	if err := g.validatePutTimeoutConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeoutConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationOutputReference) ResetAbortConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetAbortConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationOutputReference) ResetJobExecutionsRolloutConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetJobExecutionsRolloutConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationOutputReference) ResetTimeoutConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeoutConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_Greengrassv2DeploymentIotJobConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

