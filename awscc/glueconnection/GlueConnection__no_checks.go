// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package glueconnection

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GlueConnection) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (g *jsiiProxy_GlueConnection) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (g *jsiiProxy_GlueConnection) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GlueConnection) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GlueConnection) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GlueConnection) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GlueConnection) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GlueConnection) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GlueConnection) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GlueConnection) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GlueConnection) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GlueConnection) validateImportFromParameters(id *string) error {
	return nil
}

func (g *jsiiProxy_GlueConnection) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GlueConnection) validateMarkWriteOnlyAttributeParameters(value interface{}) error {
	return nil
}

func (g *jsiiProxy_GlueConnection) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (g *jsiiProxy_GlueConnection) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (g *jsiiProxy_GlueConnection) validateMoveToIdParameters(id *string) error {
	return nil
}

func (g *jsiiProxy_GlueConnection) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (g *jsiiProxy_GlueConnection) validatePutConnectionInputParameters(value *GlueConnectionConnectionInput) error {
	return nil
}

func (g *jsiiProxy_GlueConnection) validateRegisterProviderFeatureUsageParameters(feature cdktn.ProviderFeature) error {
	return nil
}

func validateGlueConnection_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateGlueConnection_IsConstructParameters(x interface{}) error {
	return nil
}

func validateGlueConnection_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateGlueConnection_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_GlueConnection) validateSetCatalogIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GlueConnection) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GlueConnection) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GlueConnection) validateSetLifecycleParameters(val *cdktn.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_GlueConnection) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_GlueConnection) validateSetTagsParameters(val *string) error {
	return nil
}

func validateNewGlueConnectionParameters(scope constructs.Construct, id *string, config *GlueConnectionConfig) error {
	return nil
}

