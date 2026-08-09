// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package groundstationconfig

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GroundstationConfigTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (g *jsiiProxy_GroundstationConfigTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GroundstationConfigTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GroundstationConfigTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GroundstationConfigTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GroundstationConfigTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GroundstationConfigTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGroundstationConfigTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

