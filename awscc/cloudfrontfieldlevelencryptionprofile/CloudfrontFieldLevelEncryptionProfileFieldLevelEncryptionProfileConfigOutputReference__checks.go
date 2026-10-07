// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package cloudfrontfieldlevelencryptionprofile

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (c *jsiiProxy_CloudfrontFieldLevelEncryptionProfileFieldLevelEncryptionProfileConfigOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudfrontFieldLevelEncryptionProfileFieldLevelEncryptionProfileConfigOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudfrontFieldLevelEncryptionProfileFieldLevelEncryptionProfileConfigOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudfrontFieldLevelEncryptionProfileFieldLevelEncryptionProfileConfigOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudfrontFieldLevelEncryptionProfileFieldLevelEncryptionProfileConfigOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudfrontFieldLevelEncryptionProfileFieldLevelEncryptionProfileConfigOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudfrontFieldLevelEncryptionProfileFieldLevelEncryptionProfileConfigOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudfrontFieldLevelEncryptionProfileFieldLevelEncryptionProfileConfigOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudfrontFieldLevelEncryptionProfileFieldLevelEncryptionProfileConfigOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudfrontFieldLevelEncryptionProfileFieldLevelEncryptionProfileConfigOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CloudfrontFieldLevelEncryptionProfileFieldLevelEncryptionProfileConfigOutputReference) validatePutEncryptionEntitiesParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*CloudfrontFieldLevelEncryptionProfileFieldLevelEncryptionProfileConfigEncryptionEntities:
		value := value.(*[]*CloudfrontFieldLevelEncryptionProfileFieldLevelEncryptionProfileConfigEncryptionEntities)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*CloudfrontFieldLevelEncryptionProfileFieldLevelEncryptionProfileConfigEncryptionEntities:
		value_ := value.([]*CloudfrontFieldLevelEncryptionProfileFieldLevelEncryptionProfileConfigEncryptionEntities)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*CloudfrontFieldLevelEncryptionProfileFieldLevelEncryptionProfileConfigEncryptionEntities; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (c *jsiiProxy_CloudfrontFieldLevelEncryptionProfileFieldLevelEncryptionProfileConfigOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionProfileFieldLevelEncryptionProfileConfigOutputReference) validateSetCallerReferenceParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionProfileFieldLevelEncryptionProfileConfigOutputReference) validateSetCommentParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionProfileFieldLevelEncryptionProfileConfigOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
	switch val.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *float64:
		// ok
	case float64:
		// ok
	case *int:
		// ok
	case int:
		// ok
	case *uint:
		// ok
	case uint:
		// ok
	case *int8:
		// ok
	case int8:
		// ok
	case *int16:
		// ok
	case int16:
		// ok
	case *int32:
		// ok
	case int32:
		// ok
	case *int64:
		// ok
	case int64:
		// ok
	case *uint8:
		// ok
	case uint8:
		// ok
	case *uint16:
		// ok
	case uint16:
		// ok
	case *uint32:
		// ok
	case uint32:
		// ok
	case *uint64:
		// ok
	case uint64:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *string, *float64; received %#v (a %T)", val, val)
	}

	return nil
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionProfileFieldLevelEncryptionProfileConfigOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionProfileFieldLevelEncryptionProfileConfigOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *CloudfrontFieldLevelEncryptionProfileFieldLevelEncryptionProfileConfig:
		val := val.(*CloudfrontFieldLevelEncryptionProfileFieldLevelEncryptionProfileConfig)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CloudfrontFieldLevelEncryptionProfileFieldLevelEncryptionProfileConfig:
		val_ := val.(CloudfrontFieldLevelEncryptionProfileFieldLevelEncryptionProfileConfig)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *CloudfrontFieldLevelEncryptionProfileFieldLevelEncryptionProfileConfig; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionProfileFieldLevelEncryptionProfileConfigOutputReference) validateSetNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionProfileFieldLevelEncryptionProfileConfigOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CloudfrontFieldLevelEncryptionProfileFieldLevelEncryptionProfileConfigOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewCloudfrontFieldLevelEncryptionProfileFieldLevelEncryptionProfileConfigOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

