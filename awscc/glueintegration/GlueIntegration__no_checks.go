// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package glueintegration

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GlueIntegration) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (g *jsiiProxy_GlueIntegration) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (g *jsiiProxy_GlueIntegration) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GlueIntegration) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GlueIntegration) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GlueIntegration) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GlueIntegration) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GlueIntegration) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GlueIntegration) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GlueIntegration) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GlueIntegration) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GlueIntegration) validateImportFromParameters(id *string) error {
	return nil
}

func (g *jsiiProxy_GlueIntegration) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (g *jsiiProxy_GlueIntegration) validateMarkWriteOnlyAttributeParameters(value interface{}) error {
	return nil
}

func (g *jsiiProxy_GlueIntegration) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (g *jsiiProxy_GlueIntegration) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (g *jsiiProxy_GlueIntegration) validateMoveToIdParameters(id *string) error {
	return nil
}

func (g *jsiiProxy_GlueIntegration) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (g *jsiiProxy_GlueIntegration) validatePutIntegrationConfigParameters(value *GlueIntegrationIntegrationConfig) error {
	return nil
}

func (g *jsiiProxy_GlueIntegration) validatePutTagsParameters(value interface{}) error {
	return nil
}

func (g *jsiiProxy_GlueIntegration) validateRegisterProviderFeatureUsageParameters(feature cdktn.ProviderFeature) error {
	return nil
}

func validateGlueIntegration_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateGlueIntegration_IsConstructParameters(x interface{}) error {
	return nil
}

func validateGlueIntegration_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateGlueIntegration_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_GlueIntegration) validateSetAdditionalEncryptionContextParameters(val *map[string]*string) error {
	return nil
}

func (j *jsiiProxy_GlueIntegration) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GlueIntegration) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GlueIntegration) validateSetDataFilterParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GlueIntegration) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GlueIntegration) validateSetIntegrationNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GlueIntegration) validateSetKmsKeyIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GlueIntegration) validateSetLifecycleParameters(val *cdktn.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_GlueIntegration) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_GlueIntegration) validateSetSourceArnParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GlueIntegration) validateSetTargetArnParameters(val *string) error {
	return nil
}

func validateNewGlueIntegrationParameters(scope constructs.Construct, id *string, config *GlueIntegrationConfig) error {
	return nil
}

