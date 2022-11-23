// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_apigateway_documentation_part", documentationPartDataSource)
}

// documentationPartDataSource returns the Terraform awscc_apigateway_documentation_part data source.
// This Terraform data source corresponds to the CloudFormation AWS::ApiGateway::DocumentationPart resource.
func documentationPartDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"documentation_part_id": {
			// Property: DocumentationPartId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The identifier of the documentation Part.",
			//	  "type": "string"
			//	}
			Description: "The identifier of the documentation Part.",
			Type:        types.StringType,
			Computed:    true,
		},
		"location": {
			// Property: Location
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "The location of the API entity that the documentation applies to.",
			//	  "properties": {
			//	    "Method": {
			//	      "description": "The HTTP verb of a method.",
			//	      "type": "string"
			//	    },
			//	    "Name": {
			//	      "description": "The name of the targeted API entity.",
			//	      "type": "string"
			//	    },
			//	    "Path": {
			//	      "description": "The URL path of the target.",
			//	      "type": "string"
			//	    },
			//	    "StatusCode": {
			//	      "description": "The HTTP status code of a response.",
			//	      "type": "string"
			//	    },
			//	    "Type": {
			//	      "description": "The type of API entity that the documentation content applies to.",
			//	      "enum": [
			//	        "API",
			//	        "AUTHORIZER",
			//	        "MODEL",
			//	        "RESOURCE",
			//	        "METHOD",
			//	        "PATH_PARAMETER",
			//	        "QUERY_PARAMETER",
			//	        "REQUEST_HEADER",
			//	        "REQUEST_BODY",
			//	        "RESPONSE",
			//	        "RESPONSE_HEADER",
			//	        "RESPONSE_BODY"
			//	      ],
			//	      "type": "string"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Description: "The location of the API entity that the documentation applies to.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"method": {
						// Property: Method
						Description: "The HTTP verb of a method.",
						Type:        types.StringType,
						Computed:    true,
					},
					"name": {
						// Property: Name
						Description: "The name of the targeted API entity.",
						Type:        types.StringType,
						Computed:    true,
					},
					"path": {
						// Property: Path
						Description: "The URL path of the target.",
						Type:        types.StringType,
						Computed:    true,
					},
					"status_code": {
						// Property: StatusCode
						Description: "The HTTP status code of a response.",
						Type:        types.StringType,
						Computed:    true,
					},
					"type": {
						// Property: Type
						Description: "The type of API entity that the documentation content applies to.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
			),
			Computed: true,
		},
		"properties": {
			// Property: Properties
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The documentation content map of the targeted API entity.",
			//	  "type": "string"
			//	}
			Description: "The documentation content map of the targeted API entity.",
			Type:        types.StringType,
			Computed:    true,
		},
		"rest_api_id": {
			// Property: RestApiId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Identifier of the targeted API entity",
			//	  "type": "string"
			//	}
			Description: "Identifier of the targeted API entity",
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
		Description: "Data Source schema for AWS::ApiGateway::DocumentationPart",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ApiGateway::DocumentationPart").WithTerraformTypeName("awscc_apigateway_documentation_part")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"documentation_part_id": "DocumentationPartId",
		"location":              "Location",
		"method":                "Method",
		"name":                  "Name",
		"path":                  "Path",
		"properties":            "Properties",
		"rest_api_id":           "RestApiId",
		"status_code":           "StatusCode",
		"type":                  "Type",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
