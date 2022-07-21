// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package evidently

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_evidently_experiment", experimentDataSourceType)
}

// experimentDataSourceType returns the Terraform awscc_evidently_experiment data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Evidently::Experiment resource type.
func experimentDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "pattern": "arn:[^:]*:[^:]*:[^:]*:[^:]*:project/[-a-zA-Z0-9._]*/experiment/[-a-zA-Z0-9._]*",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 160,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"metric_goals": {
			// Property: MetricGoals
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": true,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "DesiredChange": {
			//         "enum": [
			//           "INCREASE",
			//           "DECREASE"
			//         ],
			//         "type": "string"
			//       },
			//       "EntityIdKey": {
			//         "description": "The JSON path to reference the entity id in the event.",
			//         "type": "string"
			//       },
			//       "EventPattern": {
			//         "description": "Event patterns have the same structure as the events they match. Rules use event patterns to select events. An event pattern either matches an event or it doesn't.",
			//         "type": "string"
			//       },
			//       "MetricName": {
			//         "maxLength": 255,
			//         "minLength": 1,
			//         "pattern": "^[\\S]+$",
			//         "type": "string"
			//       },
			//       "UnitLabel": {
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "pattern": ".*",
			//         "type": "string"
			//       },
			//       "ValueKey": {
			//         "description": "The JSON path to reference the numerical metric value in the event.",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "MetricName",
			//       "EntityIdKey",
			//       "ValueKey",
			//       "EventPattern",
			//       "DesiredChange"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 3,
			//   "minItems": 1,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"desired_change": {
						// Property: DesiredChange
						Type:     types.StringType,
						Computed: true,
					},
					"entity_id_key": {
						// Property: EntityIdKey
						Description: "The JSON path to reference the entity id in the event.",
						Type:        types.StringType,
						Computed:    true,
					},
					"event_pattern": {
						// Property: EventPattern
						Description: "Event patterns have the same structure as the events they match. Rules use event patterns to select events. An event pattern either matches an event or it doesn't.",
						Type:        types.StringType,
						Computed:    true,
					},
					"metric_name": {
						// Property: MetricName
						Type:     types.StringType,
						Computed: true,
					},
					"unit_label": {
						// Property: UnitLabel
						Type:     types.StringType,
						Computed: true,
					},
					"value_key": {
						// Property: ValueKey
						Description: "The JSON path to reference the numerical metric value in the event.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 127,
			//   "minLength": 1,
			//   "pattern": "[-a-zA-Z0-9._]*",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"online_ab_config": {
			// Property: OnlineAbConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "ControlTreatmentName": {
			//       "maxLength": 127,
			//       "minLength": 1,
			//       "pattern": "[-a-zA-Z0-9._]*",
			//       "type": "string"
			//     },
			//     "TreatmentWeights": {
			//       "insertionOrder": false,
			//       "items": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "SplitWeight": {
			//             "maximum": 100000,
			//             "minimum": 0,
			//             "type": "integer"
			//           },
			//           "Treatment": {
			//             "maxLength": 127,
			//             "minLength": 1,
			//             "pattern": "[-a-zA-Z0-9._]*",
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "Treatment",
			//           "SplitWeight"
			//         ],
			//         "type": "object"
			//       },
			//       "type": "array",
			//       "uniqueItems": true
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"control_treatment_name": {
						// Property: ControlTreatmentName
						Type:     types.StringType,
						Computed: true,
					},
					"treatment_weights": {
						// Property: TreatmentWeights
						Attributes: tfsdk.SetNestedAttributes(
							map[string]tfsdk.Attribute{
								"split_weight": {
									// Property: SplitWeight
									Type:     types.Int64Type,
									Computed: true,
								},
								"treatment": {
									// Property: Treatment
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"project": {
			// Property: Project
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 2048,
			//   "minLength": 0,
			//   "pattern": "([-a-zA-Z0-9._]*)|(arn:[^:]*:[^:]*:[^:]*:[^:]*:project/[-a-zA-Z0-9._]*)",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"randomization_salt": {
			// Property: RandomizationSalt
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 127,
			//   "minLength": 0,
			//   "pattern": ".*",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"remove_segment": {
			// Property: RemoveSegment
			// CloudFormation resource type schema:
			// {
			//   "type": "boolean"
			// }
			Type:     types.BoolType,
			Computed: true,
		},
		"running_status": {
			// Property: RunningStatus
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Start Experiment. Default is False",
			//   "oneOf": [
			//     {
			//       "required": [
			//         "Status",
			//         "AnalysisCompleteTime"
			//       ]
			//     },
			//     {
			//       "required": [
			//         "Status",
			//         "Reason",
			//         "DesiredState"
			//       ]
			//     }
			//   ],
			//   "properties": {
			//     "AnalysisCompleteTime": {
			//       "description": "Provide the analysis Completion time for an experiment",
			//       "type": "string"
			//     },
			//     "DesiredState": {
			//       "description": "Provide CANCELLED or COMPLETED desired state when stopping an experiment",
			//       "pattern": "^(CANCELLED|COMPLETED)",
			//       "type": "string"
			//     },
			//     "Reason": {
			//       "description": "Reason is a required input for stopping the experiment",
			//       "type": "string"
			//     },
			//     "Status": {
			//       "description": "Provide START or STOP action to apply on an experiment",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Start Experiment. Default is False",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"analysis_complete_time": {
						// Property: AnalysisCompleteTime
						Description: "Provide the analysis Completion time for an experiment",
						Type:        types.StringType,
						Computed:    true,
					},
					"desired_state": {
						// Property: DesiredState
						Description: "Provide CANCELLED or COMPLETED desired state when stopping an experiment",
						Type:        types.StringType,
						Computed:    true,
					},
					"reason": {
						// Property: Reason
						Description: "Reason is a required input for stopping the experiment",
						Type:        types.StringType,
						Computed:    true,
					},
					"status": {
						// Property: Status
						Description: "Provide START or STOP action to apply on an experiment",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"sampling_rate": {
			// Property: SamplingRate
			// CloudFormation resource type schema:
			// {
			//   "maximum": 100000,
			//   "minimum": 0,
			//   "type": "integer"
			// }
			Type:     types.Int64Type,
			Computed: true,
		},
		"segment": {
			// Property: Segment
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 2048,
			//   "minLength": 0,
			//   "pattern": "([-a-zA-Z0-9._]*)|(arn:[^:]*:[^:]*:[^:]*:[^:]*:segment/[-a-zA-Z0-9._]*)",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"treatments": {
			// Property: Treatments
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": true,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Description": {
			//         "type": "string"
			//       },
			//       "Feature": {
			//         "pattern": "([-a-zA-Z0-9._]*)|(arn:[^:]*:[^:]*:[^:]*:[^:]*:.*)",
			//         "type": "string"
			//       },
			//       "TreatmentName": {
			//         "maxLength": 127,
			//         "minLength": 1,
			//         "pattern": "[-a-zA-Z0-9._]*",
			//         "type": "string"
			//       },
			//       "Variation": {
			//         "maxLength": 255,
			//         "minLength": 1,
			//         "pattern": "[-a-zA-Z0-9._]*",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "TreatmentName",
			//       "Feature",
			//       "Variation"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 5,
			//   "minItems": 2,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"description": {
						// Property: Description
						Type:     types.StringType,
						Computed: true,
					},
					"feature": {
						// Property: Feature
						Type:     types.StringType,
						Computed: true,
					},
					"treatment_name": {
						// Property: TreatmentName
						Type:     types.StringType,
						Computed: true,
					},
					"variation": {
						// Property: Variation
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::Evidently::Experiment",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Evidently::Experiment").WithTerraformTypeName("awscc_evidently_experiment")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"analysis_complete_time": "AnalysisCompleteTime",
		"arn":                    "Arn",
		"control_treatment_name": "ControlTreatmentName",
		"description":            "Description",
		"desired_change":         "DesiredChange",
		"desired_state":          "DesiredState",
		"entity_id_key":          "EntityIdKey",
		"event_pattern":          "EventPattern",
		"feature":                "Feature",
		"key":                    "Key",
		"metric_goals":           "MetricGoals",
		"metric_name":            "MetricName",
		"name":                   "Name",
		"online_ab_config":       "OnlineAbConfig",
		"project":                "Project",
		"randomization_salt":     "RandomizationSalt",
		"reason":                 "Reason",
		"remove_segment":         "RemoveSegment",
		"running_status":         "RunningStatus",
		"sampling_rate":          "SamplingRate",
		"segment":                "Segment",
		"split_weight":           "SplitWeight",
		"status":                 "Status",
		"tags":                   "Tags",
		"treatment":              "Treatment",
		"treatment_name":         "TreatmentName",
		"treatment_weights":      "TreatmentWeights",
		"treatments":             "Treatments",
		"unit_label":             "UnitLabel",
		"value":                  "Value",
		"value_key":              "ValueKey",
		"variation":              "Variation",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
