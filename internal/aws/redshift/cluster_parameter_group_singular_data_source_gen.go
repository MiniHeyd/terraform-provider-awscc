// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_redshift_cluster_parameter_group", clusterParameterGroupDataSource)
}

// clusterParameterGroupDataSource returns the Terraform awscc_redshift_cluster_parameter_group data source.
// This Terraform data source corresponds to the CloudFormation AWS::Redshift::ClusterParameterGroup resource.
func clusterParameterGroupDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "A description of the parameter group.",
			//	  "type": "string"
			//	}
			Description: "A description of the parameter group.",
			Type:        types.StringType,
			Computed:    true,
		},
		"parameter_group_family": {
			// Property: ParameterGroupFamily
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The Amazon Redshift engine version to which the cluster parameter group applies. The cluster engine version determines the set of parameters.",
			//	  "type": "string"
			//	}
			Description: "The Amazon Redshift engine version to which the cluster parameter group applies. The cluster engine version determines the set of parameters.",
			Type:        types.StringType,
			Computed:    true,
		},
		"parameter_group_name": {
			// Property: ParameterGroupName
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The name of the cluster parameter group.",
			//	  "maxLength": 255,
			//	  "type": "string"
			//	}
			Description: "The name of the cluster parameter group.",
			Type:        types.StringType,
			Computed:    true,
		},
		"parameters": {
			// Property: Parameters
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "An array of parameters to be modified. A maximum of 20 parameters can be modified in a single request.",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "properties": {
			//	      "ParameterName": {
			//	        "description": "The name of the parameter.",
			//	        "type": "string"
			//	      },
			//	      "ParameterValue": {
			//	        "description": "The value of the parameter. If `ParameterName` is `wlm_json_configuration`, then the maximum size of `ParameterValue` is 8000 characters.",
			//	        "type": "string"
			//	      }
			//	    },
			//	    "required": [
			//	      "ParameterValue",
			//	      "ParameterName"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "type": "array"
			//	}
			Description: "An array of parameters to be modified. A maximum of 20 parameters can be modified in a single request.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"parameter_name": {
						// Property: ParameterName
						Description: "The name of the parameter.",
						Type:        types.StringType,
						Computed:    true,
					},
					"parameter_value": {
						// Property: ParameterValue
						Description: "The value of the parameter. If `ParameterName` is `wlm_json_configuration`, then the maximum size of `ParameterValue` is 8000 characters.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "An array of key-value pairs to apply to this resource.",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "description": "A key-value pair to associate with a resource.",
			//	    "properties": {
			//	      "Key": {
			//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//	        "maxLength": 128,
			//	        "minLength": 1,
			//	        "type": "string"
			//	      },
			//	      "Value": {
			//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//	        "maxLength": 256,
			//	        "minLength": 0,
			//	        "type": "string"
			//	      }
			//	    },
			//	    "required": [
			//	      "Key",
			//	      "Value"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "type": "array"
			//	}
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: tfsdk.ListNestedAttributes(
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
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::Redshift::ClusterParameterGroup",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Redshift::ClusterParameterGroup").WithTerraformTypeName("awscc_redshift_cluster_parameter_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"description":            "Description",
		"key":                    "Key",
		"parameter_group_family": "ParameterGroupFamily",
		"parameter_group_name":   "ParameterGroupName",
		"parameter_name":         "ParameterName",
		"parameter_value":        "ParameterValue",
		"parameters":             "Parameters",
		"tags":                   "Tags",
		"value":                  "Value",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
