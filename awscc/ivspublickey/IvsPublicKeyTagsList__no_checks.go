// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ivspublickey

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IvsPublicKeyTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_IvsPublicKeyTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IvsPublicKeyTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IvsPublicKeyTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IvsPublicKeyTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IvsPublicKeyTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IvsPublicKeyTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIvsPublicKeyTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

