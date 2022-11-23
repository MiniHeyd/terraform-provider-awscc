// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package elasticbeanstalk

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_elasticbeanstalk_environment", environmentDataSource)
}

// environmentDataSource returns the Terraform awscc_elasticbeanstalk_environment data source.
// This Terraform data source corresponds to the CloudFormation AWS::ElasticBeanstalk::Environment resource.
func environmentDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"application_name": {
			// Property: ApplicationName
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The name of the application that is associated with this environment.",
			//	  "type": "string"
			//	}
			Description: "The name of the application that is associated with this environment.",
			Type:        types.StringType,
			Computed:    true,
		},
		"cname_prefix": {
			// Property: CNAMEPrefix
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "If specified, the environment attempts to use this value as the prefix for the CNAME in your Elastic Beanstalk environment URL. If not specified, the CNAME is generated automatically by appending a random alphanumeric string to the environment name.",
			//	  "type": "string"
			//	}
			Description: "If specified, the environment attempts to use this value as the prefix for the CNAME in your Elastic Beanstalk environment URL. If not specified, the CNAME is generated automatically by appending a random alphanumeric string to the environment name.",
			Type:        types.StringType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Your description for this environment.",
			//	  "type": "string"
			//	}
			Description: "Your description for this environment.",
			Type:        types.StringType,
			Computed:    true,
		},
		"endpoint_url": {
			// Property: EndpointURL
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"environment_name": {
			// Property: EnvironmentName
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "A unique name for the environment.",
			//	  "type": "string"
			//	}
			Description: "A unique name for the environment.",
			Type:        types.StringType,
			Computed:    true,
		},
		"operations_role": {
			// Property: OperationsRole
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The Amazon Resource Name (ARN) of an existing IAM role to be used as the environment's operations role.",
			//	  "type": "string"
			//	}
			Description: "The Amazon Resource Name (ARN) of an existing IAM role to be used as the environment's operations role.",
			Type:        types.StringType,
			Computed:    true,
		},
		"option_settings": {
			// Property: OptionSettings
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Key-value pairs defining configuration options for this environment, such as the instance type.",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "properties": {
			//	      "Namespace": {
			//	        "description": "A unique namespace that identifies the option's associated AWS resource.",
			//	        "type": "string"
			//	      },
			//	      "OptionName": {
			//	        "description": "The name of the configuration option.",
			//	        "type": "string"
			//	      },
			//	      "ResourceName": {
			//	        "description": "A unique resource name for the option setting. Use it for a time–based scaling configuration option.",
			//	        "type": "string"
			//	      },
			//	      "Value": {
			//	        "description": "The current value for the configuration option.",
			//	        "type": "string"
			//	      }
			//	    },
			//	    "required": [
			//	      "Namespace",
			//	      "OptionName"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "type": "array",
			//	  "uniqueItems": false
			//	}
			Description: "Key-value pairs defining configuration options for this environment, such as the instance type.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"namespace": {
						// Property: Namespace
						Description: "A unique namespace that identifies the option's associated AWS resource.",
						Type:        types.StringType,
						Computed:    true,
					},
					"option_name": {
						// Property: OptionName
						Description: "The name of the configuration option.",
						Type:        types.StringType,
						Computed:    true,
					},
					"resource_name": {
						// Property: ResourceName
						Description: "A unique resource name for the option setting. Use it for a time–based scaling configuration option.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The current value for the configuration option.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"platform_arn": {
			// Property: PlatformArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The Amazon Resource Name (ARN) of the custom platform to use with the environment.",
			//	  "type": "string"
			//	}
			Description: "The Amazon Resource Name (ARN) of the custom platform to use with the environment.",
			Type:        types.StringType,
			Computed:    true,
		},
		"solution_stack_name": {
			// Property: SolutionStackName
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The name of an Elastic Beanstalk solution stack (platform version) to use with the environment.",
			//	  "type": "string"
			//	}
			Description: "The name of an Elastic Beanstalk solution stack (platform version) to use with the environment.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Specifies the tags applied to resources in the environment.",
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "properties": {
			//	      "Key": {
			//	        "description": "The key name of the tag.",
			//	        "type": "string"
			//	      },
			//	      "Value": {
			//	        "description": "The value for the tag.",
			//	        "type": "string"
			//	      }
			//	    },
			//	    "required": [
			//	      "Value",
			//	      "Key"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "type": "array",
			//	  "uniqueItems": false
			//	}
			Description: "Specifies the tags applied to resources in the environment.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"template_name": {
			// Property: TemplateName
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The name of the Elastic Beanstalk configuration template to use with the environment.",
			//	  "type": "string"
			//	}
			Description: "The name of the Elastic Beanstalk configuration template to use with the environment.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tier": {
			// Property: Tier
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "Specifies the tier to use in creating this environment. The environment tier that you choose determines whether Elastic Beanstalk provisions resources to support a web application that handles HTTP(S) requests or a web application that handles background-processing tasks.",
			//	  "properties": {
			//	    "Name": {
			//	      "description": "The name of this environment tier.",
			//	      "type": "string"
			//	    },
			//	    "Type": {
			//	      "description": "The type of this environment tier.",
			//	      "type": "string"
			//	    },
			//	    "Version": {
			//	      "description": "The version of this environment tier. When you don't set a value to it, Elastic Beanstalk uses the latest compatible worker tier version.",
			//	      "type": "string"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Description: "Specifies the tier to use in creating this environment. The environment tier that you choose determines whether Elastic Beanstalk provisions resources to support a web application that handles HTTP(S) requests or a web application that handles background-processing tasks.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"name": {
						// Property: Name
						Description: "The name of this environment tier.",
						Type:        types.StringType,
						Computed:    true,
					},
					"type": {
						// Property: Type
						Description: "The type of this environment tier.",
						Type:        types.StringType,
						Computed:    true,
					},
					"version": {
						// Property: Version
						Description: "The version of this environment tier. When you don't set a value to it, Elastic Beanstalk uses the latest compatible worker tier version.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"version_label": {
			// Property: VersionLabel
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The name of the application version to deploy.",
			//	  "type": "string"
			//	}
			Description: "The name of the application version to deploy.",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::ElasticBeanstalk::Environment",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ElasticBeanstalk::Environment").WithTerraformTypeName("awscc_elasticbeanstalk_environment")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"application_name":    "ApplicationName",
		"cname_prefix":        "CNAMEPrefix",
		"description":         "Description",
		"endpoint_url":        "EndpointURL",
		"environment_name":    "EnvironmentName",
		"key":                 "Key",
		"name":                "Name",
		"namespace":           "Namespace",
		"operations_role":     "OperationsRole",
		"option_name":         "OptionName",
		"option_settings":     "OptionSettings",
		"platform_arn":        "PlatformArn",
		"resource_name":       "ResourceName",
		"solution_stack_name": "SolutionStackName",
		"tags":                "Tags",
		"template_name":       "TemplateName",
		"tier":                "Tier",
		"type":                "Type",
		"value":               "Value",
		"version":             "Version",
		"version_label":       "VersionLabel",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
