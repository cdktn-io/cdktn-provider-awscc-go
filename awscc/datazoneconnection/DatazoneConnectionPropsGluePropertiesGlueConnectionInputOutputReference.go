// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazoneconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/datazoneconnection/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference interface {
	cdktn.ComplexObject
	AthenaProperties() *map[string]*string
	SetAthenaProperties(val *map[string]*string)
	AthenaPropertiesInput() *map[string]*string
	AuthenticationConfiguration() DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference
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
	ConnectionProperties() *map[string]*string
	SetConnectionProperties(val *map[string]*string)
	ConnectionPropertiesInput() *map[string]*string
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
	MatchCriteria() *string
	SetMatchCriteria(val *string)
	MatchCriteriaInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	PhysicalConnectionRequirements() DatazoneConnectionPropsGluePropertiesGlueConnectionInputPhysicalConnectionRequirementsOutputReference
	PhysicalConnectionRequirementsInput() interface{}
	PythonProperties() *map[string]*string
	SetPythonProperties(val *map[string]*string)
	PythonPropertiesInput() *map[string]*string
	SparkProperties() *map[string]*string
	SetSparkProperties(val *map[string]*string)
	SparkPropertiesInput() *map[string]*string
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
	PutAuthenticationConfiguration(value *DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfiguration)
	PutPhysicalConnectionRequirements(value *DatazoneConnectionPropsGluePropertiesGlueConnectionInputPhysicalConnectionRequirements)
	ResetAthenaProperties()
	ResetAuthenticationConfiguration()
	ResetConnectionProperties()
	ResetConnectionType()
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

// The jsii proxy struct for DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference
type jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) AthenaProperties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"athenaProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) AthenaPropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"athenaPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) AuthenticationConfiguration() DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference {
	var returns DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfigurationOutputReference
	_jsii_.Get(
		j,
		"authenticationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) AuthenticationConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authenticationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) ConnectionProperties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"connectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) ConnectionPropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"connectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) ConnectionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) ConnectionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) MatchCriteria() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) MatchCriteriaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) PhysicalConnectionRequirements() DatazoneConnectionPropsGluePropertiesGlueConnectionInputPhysicalConnectionRequirementsOutputReference {
	var returns DatazoneConnectionPropsGluePropertiesGlueConnectionInputPhysicalConnectionRequirementsOutputReference
	_jsii_.Get(
		j,
		"physicalConnectionRequirements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) PhysicalConnectionRequirementsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"physicalConnectionRequirementsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) PythonProperties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"pythonProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) PythonPropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"pythonPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) SparkProperties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sparkProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) SparkPropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sparkPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) ValidateCredentials() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validateCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) ValidateCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validateCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) ValidateForComputeEnvironments() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"validateForComputeEnvironments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) ValidateForComputeEnvironmentsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"validateForComputeEnvironmentsInput",
		&returns,
	)
	return returns
}


func NewDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference {
	_init_.Initialize()

	if err := validateNewDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.datazoneConnection.DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference_Override(d DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.datazoneConnection.DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference)SetAthenaProperties(val *map[string]*string) {
	if err := j.validateSetAthenaPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"athenaProperties",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference)SetConnectionProperties(val *map[string]*string) {
	if err := j.validateSetConnectionPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionProperties",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference)SetConnectionType(val *string) {
	if err := j.validateSetConnectionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionType",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference)SetMatchCriteria(val *string) {
	if err := j.validateSetMatchCriteriaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchCriteria",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference)SetPythonProperties(val *map[string]*string) {
	if err := j.validateSetPythonPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pythonProperties",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference)SetSparkProperties(val *map[string]*string) {
	if err := j.validateSetSparkPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparkProperties",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference)SetValidateCredentials(val interface{}) {
	if err := j.validateSetValidateCredentialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validateCredentials",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference)SetValidateForComputeEnvironments(val *[]*string) {
	if err := j.validateSetValidateForComputeEnvironmentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validateForComputeEnvironments",
		val,
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) PutAuthenticationConfiguration(value *DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfiguration) {
	if err := d.validatePutAuthenticationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAuthenticationConfiguration",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) PutPhysicalConnectionRequirements(value *DatazoneConnectionPropsGluePropertiesGlueConnectionInputPhysicalConnectionRequirements) {
	if err := d.validatePutPhysicalConnectionRequirementsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPhysicalConnectionRequirements",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) ResetAthenaProperties() {
	_jsii_.InvokeVoid(
		d,
		"resetAthenaProperties",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) ResetAuthenticationConfiguration() {
	_jsii_.InvokeVoid(
		d,
		"resetAuthenticationConfiguration",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) ResetConnectionProperties() {
	_jsii_.InvokeVoid(
		d,
		"resetConnectionProperties",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) ResetConnectionType() {
	_jsii_.InvokeVoid(
		d,
		"resetConnectionType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) ResetMatchCriteria() {
	_jsii_.InvokeVoid(
		d,
		"resetMatchCriteria",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) ResetPhysicalConnectionRequirements() {
	_jsii_.InvokeVoid(
		d,
		"resetPhysicalConnectionRequirements",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) ResetPythonProperties() {
	_jsii_.InvokeVoid(
		d,
		"resetPythonProperties",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) ResetSparkProperties() {
	_jsii_.InvokeVoid(
		d,
		"resetSparkProperties",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) ResetValidateCredentials() {
	_jsii_.InvokeVoid(
		d,
		"resetValidateCredentials",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) ResetValidateForComputeEnvironments() {
	_jsii_.InvokeVoid(
		d,
		"resetValidateForComputeEnvironments",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DatazoneConnectionPropsGluePropertiesGlueConnectionInputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

