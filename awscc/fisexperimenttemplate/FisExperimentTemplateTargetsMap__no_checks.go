// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package fisexperimenttemplate

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FisExperimentTemplateTargetsMap) validateGetParameters(key *string) error {
	return nil
}

func (f *jsiiProxy_FisExperimentTemplateTargetsMap) validateInterpolationForAttributeParameters(property *string) error {
	return nil
}

func (f *jsiiProxy_FisExperimentTemplateTargetsMap) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FisExperimentTemplateTargetsMap) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FisExperimentTemplateTargetsMap) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FisExperimentTemplateTargetsMap) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func validateNewFisExperimentTemplateTargetsMapParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	return nil
}

