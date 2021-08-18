// Code generated by generators/resource/main.go; DO NOT EDIT.

package logs

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	providertypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddResourceTypeFactory("awscc_logs_query_definition", queryDefinitionResourceType)
}

// queryDefinitionResourceType returns the Terraform awscc_logs_query_definition resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Logs::QueryDefinition resource type.
func queryDefinitionResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"log_group_names": {
			// Property: LogGroupNames
			// CloudFormation resource type schema:
			// {
			//   "description": "Optionally define specific log groups as part of your query definition",
			//   "insertionOrder": false,
			//   "items": {
			//     "maxLength": 512,
			//     "minLength": 1,
			//     "pattern": "",
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "Optionally define specific log groups as part of your query definition",
			Type:        providertypes.MultisetType{ListType: types.ListType{ElemType: types.StringType}},
			Optional:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "A name for the saved query definition",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "A name for the saved query definition",
			Type:        types.StringType,
			Required:    true,
		},
		"query_definition_id": {
			// Property: QueryDefinitionId
			// CloudFormation resource type schema:
			// {
			//   "description": "Unique identifier of a query definition",
			//   "maxLength": 256,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Description: "Unique identifier of a query definition",
			Type:        types.StringType,
			Computed:    true,
		},
		"query_string": {
			// Property: QueryString
			// CloudFormation resource type schema:
			// {
			//   "description": "The query string to use for this definition",
			//   "maxLength": 10000,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "The query string to use for this definition",
			Type:        types.StringType,
			Required:    true,
		},
	}

	// Required for acceptance testing.
	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "The resource schema for AWSLogs QueryDefinition",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Logs::QueryDefinition").WithTerraformTypeName("awscc_logs_query_definition").WithTerraformSchema(schema)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_logs_query_definition", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
