// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/glueconnection/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueConnectionConnectionInputOutputReference interface {
	cdktn.ComplexObject
	AthenaProperties() *string
	SetAthenaProperties(val *string)
	AthenaPropertiesInput() *string
	AuthenticationConfiguration() GlueConnectionConnectionInputAuthenticationConfigurationOutputReference
	AuthenticationConfigurationInput() interface{}
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
	ConnectionProperties() *string
	SetConnectionProperties(val *string)
	ConnectionPropertiesInput() *string
	ConnectionType() *string
	SetConnectionType(val *string)
	ConnectionTypeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MatchCriteria() *[]*string
	SetMatchCriteria(val *[]*string)
	MatchCriteriaInput() *[]*string
	Name() *string
	SetName(val *string)
	NameInput() *string
	PhysicalConnectionRequirements() GlueConnectionConnectionInputPhysicalConnectionRequirementsOutputReference
	PhysicalConnectionRequirementsInput() interface{}
	PythonProperties() *string
	SetPythonProperties(val *string)
	PythonPropertiesInput() *string
	SparkProperties() *string
	SetSparkProperties(val *string)
	SparkPropertiesInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	ValidateCredentials() interface{}
	SetValidateCredentials(val interface{})
	ValidateCredentialsInput() interface{}
	ValidateForComputeEnvironments() *[]*string
	SetValidateForComputeEnvironments(val *[]*string)
	ValidateForComputeEnvironmentsInput() *[]*string
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
	PutAuthenticationConfiguration(value *GlueConnectionConnectionInputAuthenticationConfiguration)
	PutPhysicalConnectionRequirements(value *GlueConnectionConnectionInputPhysicalConnectionRequirements)
	ResetAthenaProperties()
	ResetAuthenticationConfiguration()
	ResetConnectionProperties()
	ResetDescription()
	ResetMatchCriteria()
	ResetName()
	ResetPhysicalConnectionRequirements()
	ResetPythonProperties()
	ResetSparkProperties()
	ResetValidateCredentials()
	ResetValidateForComputeEnvironments()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GlueConnectionConnectionInputOutputReference
type jsiiProxy_GlueConnectionConnectionInputOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference) AthenaProperties() *string {
	var returns *string
	_jsii_.Get(
		j,
		"athenaProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference) AthenaPropertiesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"athenaPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference) AuthenticationConfiguration() GlueConnectionConnectionInputAuthenticationConfigurationOutputReference {
	var returns GlueConnectionConnectionInputAuthenticationConfigurationOutputReference
	_jsii_.Get(
		j,
		"authenticationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference) AuthenticationConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authenticationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference) ConnectionProperties() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference) ConnectionPropertiesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference) ConnectionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference) ConnectionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference) MatchCriteria() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference) MatchCriteriaInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference) PhysicalConnectionRequirements() GlueConnectionConnectionInputPhysicalConnectionRequirementsOutputReference {
	var returns GlueConnectionConnectionInputPhysicalConnectionRequirementsOutputReference
	_jsii_.Get(
		j,
		"physicalConnectionRequirements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference) PhysicalConnectionRequirementsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"physicalConnectionRequirementsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference) PythonProperties() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pythonProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference) PythonPropertiesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pythonPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference) SparkProperties() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparkProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference) SparkPropertiesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparkPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference) ValidateCredentials() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validateCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference) ValidateCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validateCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference) ValidateForComputeEnvironments() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"validateForComputeEnvironments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference) ValidateForComputeEnvironmentsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"validateForComputeEnvironmentsInput",
		&returns,
	)
	return returns
}


func NewGlueConnectionConnectionInputOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GlueConnectionConnectionInputOutputReference {
	_init_.Initialize()

	if err := validateNewGlueConnectionConnectionInputOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueConnectionConnectionInputOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.glueConnection.GlueConnectionConnectionInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGlueConnectionConnectionInputOutputReference_Override(g GlueConnectionConnectionInputOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.glueConnection.GlueConnectionConnectionInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference)SetAthenaProperties(val *string) {
	if err := j.validateSetAthenaPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"athenaProperties",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference)SetConnectionProperties(val *string) {
	if err := j.validateSetConnectionPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionProperties",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference)SetConnectionType(val *string) {
	if err := j.validateSetConnectionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionType",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference)SetMatchCriteria(val *[]*string) {
	if err := j.validateSetMatchCriteriaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchCriteria",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference)SetPythonProperties(val *string) {
	if err := j.validateSetPythonPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pythonProperties",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference)SetSparkProperties(val *string) {
	if err := j.validateSetSparkPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparkProperties",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference)SetValidateCredentials(val interface{}) {
	if err := j.validateSetValidateCredentialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validateCredentials",
		val,
	)
}

func (j *jsiiProxy_GlueConnectionConnectionInputOutputReference)SetValidateForComputeEnvironments(val *[]*string) {
	if err := j.validateSetValidateForComputeEnvironmentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validateForComputeEnvironments",
		val,
	)
}

func (g *jsiiProxy_GlueConnectionConnectionInputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionConnectionInputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GlueConnectionConnectionInputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueConnectionConnectionInputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GlueConnectionConnectionInputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GlueConnectionConnectionInputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GlueConnectionConnectionInputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GlueConnectionConnectionInputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GlueConnectionConnectionInputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GlueConnectionConnectionInputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GlueConnectionConnectionInputOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueConnectionConnectionInputOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GlueConnectionConnectionInputOutputReference) PutAuthenticationConfiguration(value *GlueConnectionConnectionInputAuthenticationConfiguration) {
	if err := g.validatePutAuthenticationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAuthenticationConfiguration",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueConnectionConnectionInputOutputReference) PutPhysicalConnectionRequirements(value *GlueConnectionConnectionInputPhysicalConnectionRequirements) {
	if err := g.validatePutPhysicalConnectionRequirementsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPhysicalConnectionRequirements",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueConnectionConnectionInputOutputReference) ResetAthenaProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetAthenaProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionConnectionInputOutputReference) ResetAuthenticationConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthenticationConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionConnectionInputOutputReference) ResetConnectionProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetConnectionProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionConnectionInputOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionConnectionInputOutputReference) ResetMatchCriteria() {
	_jsii_.InvokeVoid(
		g,
		"resetMatchCriteria",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionConnectionInputOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		g,
		"resetName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionConnectionInputOutputReference) ResetPhysicalConnectionRequirements() {
	_jsii_.InvokeVoid(
		g,
		"resetPhysicalConnectionRequirements",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionConnectionInputOutputReference) ResetPythonProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetPythonProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionConnectionInputOutputReference) ResetSparkProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetSparkProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionConnectionInputOutputReference) ResetValidateCredentials() {
	_jsii_.InvokeVoid(
		g,
		"resetValidateCredentials",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionConnectionInputOutputReference) ResetValidateForComputeEnvironments() {
	_jsii_.InvokeVoid(
		g,
		"resetValidateForComputeEnvironments",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueConnectionConnectionInputOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GlueConnectionConnectionInputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

