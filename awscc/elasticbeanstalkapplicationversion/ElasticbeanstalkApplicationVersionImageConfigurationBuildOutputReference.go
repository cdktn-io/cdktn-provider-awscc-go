// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticbeanstalkapplicationversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/elasticbeanstalkapplicationversion/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference interface {
	cdktn.ComplexObject
	Architecture() *string
	SetArchitecture(val *string)
	ArchitectureInput() *string
	Buildpack() *string
	SetBuildpack(val *string)
	BuildpackInput() *string
	CodeBuildServiceRole() *string
	SetCodeBuildServiceRole(val *string)
	CodeBuildServiceRoleInput() *string
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
	ComputeType() *string
	SetComputeType(val *string)
	ComputeTypeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DockerfileLocation() *string
	SetDockerfileLocation(val *string)
	DockerfileLocationInput() *string
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
	TimeoutInMinutes() *float64
	SetTimeoutInMinutes(val *float64)
	TimeoutInMinutesInput() *float64
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetArchitecture()
	ResetBuildpack()
	ResetCodeBuildServiceRole()
	ResetComputeType()
	ResetDockerfileLocation()
	ResetTimeoutInMinutes()
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference
type jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference) Architecture() *string {
	var returns *string
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference) ArchitectureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"architectureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference) Buildpack() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildpack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference) BuildpackInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildpackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference) CodeBuildServiceRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeBuildServiceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference) CodeBuildServiceRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeBuildServiceRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference) ComputeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference) ComputeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference) DockerfileLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerfileLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference) DockerfileLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dockerfileLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference) TimeoutInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference) TimeoutInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference {
	_init_.Initialize()

	if err := validateNewElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.elasticbeanstalkApplicationVersion.ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference_Override(e ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.elasticbeanstalkApplicationVersion.ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference)SetArchitecture(val *string) {
	if err := j.validateSetArchitectureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"architecture",
		val,
	)
}

func (j *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference)SetBuildpack(val *string) {
	if err := j.validateSetBuildpackParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buildpack",
		val,
	)
}

func (j *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference)SetCodeBuildServiceRole(val *string) {
	if err := j.validateSetCodeBuildServiceRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"codeBuildServiceRole",
		val,
	)
}

func (j *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference)SetComputeType(val *string) {
	if err := j.validateSetComputeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeType",
		val,
	)
}

func (j *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference)SetDockerfileLocation(val *string) {
	if err := j.validateSetDockerfileLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dockerfileLocation",
		val,
	)
}

func (j *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference)SetTimeoutInMinutes(val *float64) {
	if err := j.validateSetTimeoutInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutInMinutes",
		val,
	)
}

func (j *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (e *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference) ResetArchitecture() {
	_jsii_.InvokeVoid(
		e,
		"resetArchitecture",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference) ResetBuildpack() {
	_jsii_.InvokeVoid(
		e,
		"resetBuildpack",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference) ResetCodeBuildServiceRole() {
	_jsii_.InvokeVoid(
		e,
		"resetCodeBuildServiceRole",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference) ResetComputeType() {
	_jsii_.InvokeVoid(
		e,
		"resetComputeType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference) ResetDockerfileLocation() {
	_jsii_.InvokeVoid(
		e,
		"resetDockerfileLocation",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference) ResetTimeoutInMinutes() {
	_jsii_.InvokeVoid(
		e,
		"resetTimeoutInMinutes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		e,
		"resetType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (e *jsiiProxy_ElasticbeanstalkApplicationVersionImageConfigurationBuildOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

