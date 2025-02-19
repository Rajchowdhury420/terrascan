/*
    Copyright (C) 2022 Tenable, Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
*/

package config

import (
	"github.com/awslabs/goformation/v6/cloudformation/route53"
	"github.com/tenable/terrascan/pkg/mapper/iac-providers/cft/functions"
)

// Route53RecordConfig holds config for aws_route53_record
type Route53RecordConfig struct {
	Config
	ResourceRecords []string `json:"records"`
}

// GetRoute53RecordConfig returns config for aws_route53_record
func GetRoute53RecordConfig(r *route53.RecordSet) []AWSResourceConfig {
	cf := Route53RecordConfig{
		Config: Config{
			Name: r.Name,
		},
		ResourceRecords: functions.GetVal(r.ResourceRecords),
	}
	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: r.AWSCloudFormationMetadata,
	}}
}
