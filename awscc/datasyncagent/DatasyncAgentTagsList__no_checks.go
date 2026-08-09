// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datasyncagent

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatasyncAgentTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DatasyncAgentTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DatasyncAgentTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DatasyncAgentTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DatasyncAgentTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DatasyncAgentTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DatasyncAgentTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDatasyncAgentTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

