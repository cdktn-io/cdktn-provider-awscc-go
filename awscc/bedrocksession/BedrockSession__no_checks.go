// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package bedrocksession

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BedrockSession) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (b *jsiiProxy_BedrockSession) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (b *jsiiProxy_BedrockSession) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_BedrockSession) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_BedrockSession) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_BedrockSession) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_BedrockSession) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_BedrockSession) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_BedrockSession) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_BedrockSession) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_BedrockSession) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_BedrockSession) validateImportFromParameters(id *string) error {
	return nil
}

func (b *jsiiProxy_BedrockSession) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (b *jsiiProxy_BedrockSession) validateMarkWriteOnlyAttributeParameters(value interface{}) error {
	return nil
}

func (b *jsiiProxy_BedrockSession) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (b *jsiiProxy_BedrockSession) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (b *jsiiProxy_BedrockSession) validateMoveToIdParameters(id *string) error {
	return nil
}

func (b *jsiiProxy_BedrockSession) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (b *jsiiProxy_BedrockSession) validatePutTagsParameters(value interface{}) error {
	return nil
}

func (b *jsiiProxy_BedrockSession) validateRegisterProviderFeatureUsageParameters(feature cdktn.ProviderFeature) error {
	return nil
}

func validateBedrockSession_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateBedrockSession_IsConstructParameters(x interface{}) error {
	return nil
}

func validateBedrockSession_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateBedrockSession_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_BedrockSession) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BedrockSession) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BedrockSession) validateSetEncryptionKeyArnParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BedrockSession) validateSetLifecycleParameters(val *cdktn.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_BedrockSession) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_BedrockSession) validateSetSessionMetadataParameters(val *map[string]*string) error {
	return nil
}

func validateNewBedrockSessionParameters(scope constructs.Construct, id *string, config *BedrockSessionConfig) error {
	return nil
}

