// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotwirelessdeviceprofile


type IotwirelessDeviceProfileLoRaWan struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotwireless_device_profile#class_b_timeout IotwirelessDeviceProfile#class_b_timeout}.
	ClassBTimeout *float64 `field:"optional" json:"classBTimeout" yaml:"classBTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotwireless_device_profile#class_c_timeout IotwirelessDeviceProfile#class_c_timeout}.
	ClassCTimeout *float64 `field:"optional" json:"classCTimeout" yaml:"classCTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotwireless_device_profile#factory_preset_freqs_list IotwirelessDeviceProfile#factory_preset_freqs_list}.
	FactoryPresetFreqsList *[]*float64 `field:"optional" json:"factoryPresetFreqsList" yaml:"factoryPresetFreqsList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotwireless_device_profile#mac_version IotwirelessDeviceProfile#mac_version}.
	MacVersion *string `field:"optional" json:"macVersion" yaml:"macVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotwireless_device_profile#max_duty_cycle IotwirelessDeviceProfile#max_duty_cycle}.
	MaxDutyCycle *float64 `field:"optional" json:"maxDutyCycle" yaml:"maxDutyCycle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotwireless_device_profile#max_eirp IotwirelessDeviceProfile#max_eirp}.
	MaxEirp *float64 `field:"optional" json:"maxEirp" yaml:"maxEirp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotwireless_device_profile#ping_slot_dr IotwirelessDeviceProfile#ping_slot_dr}.
	PingSlotDr *float64 `field:"optional" json:"pingSlotDr" yaml:"pingSlotDr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotwireless_device_profile#ping_slot_freq IotwirelessDeviceProfile#ping_slot_freq}.
	PingSlotFreq *float64 `field:"optional" json:"pingSlotFreq" yaml:"pingSlotFreq"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotwireless_device_profile#ping_slot_period IotwirelessDeviceProfile#ping_slot_period}.
	PingSlotPeriod *float64 `field:"optional" json:"pingSlotPeriod" yaml:"pingSlotPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotwireless_device_profile#reg_params_revision IotwirelessDeviceProfile#reg_params_revision}.
	RegParamsRevision *string `field:"optional" json:"regParamsRevision" yaml:"regParamsRevision"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotwireless_device_profile#rf_region IotwirelessDeviceProfile#rf_region}.
	RfRegion *string `field:"optional" json:"rfRegion" yaml:"rfRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotwireless_device_profile#rx_data_rate_2 IotwirelessDeviceProfile#rx_data_rate_2}.
	RxDataRate2 *float64 `field:"optional" json:"rxDataRate2" yaml:"rxDataRate2"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotwireless_device_profile#rx_delay_1 IotwirelessDeviceProfile#rx_delay_1}.
	RxDelay1 *float64 `field:"optional" json:"rxDelay1" yaml:"rxDelay1"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotwireless_device_profile#rx_dr_offset_1 IotwirelessDeviceProfile#rx_dr_offset_1}.
	RxDrOffset1 *float64 `field:"optional" json:"rxDrOffset1" yaml:"rxDrOffset1"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotwireless_device_profile#rx_freq_2 IotwirelessDeviceProfile#rx_freq_2}.
	RxFreq2 *float64 `field:"optional" json:"rxFreq2" yaml:"rxFreq2"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotwireless_device_profile#supports_32_bit_f_cnt IotwirelessDeviceProfile#supports_32_bit_f_cnt}.
	Supports32BitFCnt interface{} `field:"optional" json:"supports32BitFCnt" yaml:"supports32BitFCnt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotwireless_device_profile#supports_class_b IotwirelessDeviceProfile#supports_class_b}.
	SupportsClassB interface{} `field:"optional" json:"supportsClassB" yaml:"supportsClassB"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotwireless_device_profile#supports_class_c IotwirelessDeviceProfile#supports_class_c}.
	SupportsClassC interface{} `field:"optional" json:"supportsClassC" yaml:"supportsClassC"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotwireless_device_profile#supports_join IotwirelessDeviceProfile#supports_join}.
	SupportsJoin interface{} `field:"optional" json:"supportsJoin" yaml:"supportsJoin"`
}

