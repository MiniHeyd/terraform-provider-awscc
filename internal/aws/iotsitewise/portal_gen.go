// Code generated by generators/resource/main.go; DO NOT EDIT.

package iotsitewise

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("awscc_iotsitewise_portal", portalResourceType)
}

// portalResourceType returns the Terraform awscc_iotsitewise_portal resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::IoTSiteWise::Portal resource type.
func portalResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"alarms": {
			// Property: Alarms
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Contains the configuration information of an alarm created in an AWS IoT SiteWise Monitor portal. You can use the alarm to monitor an asset property and get notified when the asset property value is outside a specified range.",
			//   "properties": {
			//     "AlarmRoleArn": {
			//       "description": "The ARN of the IAM role that allows the alarm to perform actions and access AWS resources and services, such as AWS IoT Events.",
			//       "type": "string"
			//     },
			//     "NotificationLambdaArn": {
			//       "description": "The ARN of the AWS Lambda function that manages alarm notifications. For more information, see Managing alarm notifications in the AWS IoT Events Developer Guide.",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Contains the configuration information of an alarm created in an AWS IoT SiteWise Monitor portal. You can use the alarm to monitor an asset property and get notified when the asset property value is outside a specified range.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"alarm_role_arn": {
						// Property: AlarmRoleArn
						Description: "The ARN of the IAM role that allows the alarm to perform actions and access AWS resources and services, such as AWS IoT Events.",
						Type:        types.StringType,
						Optional:    true,
					},
					"notification_lambda_arn": {
						// Property: NotificationLambdaArn
						Description: "The ARN of the AWS Lambda function that manages alarm notifications. For more information, see Managing alarm notifications in the AWS IoT Events Developer Guide.",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Optional: true,
		},
		"notification_sender_email": {
			// Property: NotificationSenderEmail
			// CloudFormation resource type schema:
			// {
			//   "description": "The email address that sends alarm notifications.",
			//   "type": "string"
			// }
			Description: "The email address that sends alarm notifications.",
			Type:        types.StringType,
			Optional:    true,
		},
		"portal_arn": {
			// Property: PortalArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the portal, which has the following format.",
			//   "type": "string"
			// }
			Description: "The ARN of the portal, which has the following format.",
			Type:        types.StringType,
			Computed:    true,
		},
		"portal_auth_mode": {
			// Property: PortalAuthMode
			// CloudFormation resource type schema:
			// {
			//   "description": "The service to use to authenticate users to the portal. Choose from SSO or IAM. You can't change this value after you create a portal.",
			//   "type": "string"
			// }
			Description: "The service to use to authenticate users to the portal. Choose from SSO or IAM. You can't change this value after you create a portal.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			// PortalAuthMode is a force-new attribute.
		},
		"portal_client_id": {
			// Property: PortalClientId
			// CloudFormation resource type schema:
			// {
			//   "description": "The AWS SSO application generated client ID (used with AWS SSO APIs).",
			//   "type": "string"
			// }
			Description: "The AWS SSO application generated client ID (used with AWS SSO APIs).",
			Type:        types.StringType,
			Computed:    true,
		},
		"portal_contact_email": {
			// Property: PortalContactEmail
			// CloudFormation resource type schema:
			// {
			//   "description": "The AWS administrator's contact email address.",
			//   "type": "string"
			// }
			Description: "The AWS administrator's contact email address.",
			Type:        types.StringType,
			Required:    true,
		},
		"portal_description": {
			// Property: PortalDescription
			// CloudFormation resource type schema:
			// {
			//   "description": "A description for the portal.",
			//   "type": "string"
			// }
			Description: "A description for the portal.",
			Type:        types.StringType,
			Optional:    true,
		},
		"portal_id": {
			// Property: PortalId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the portal.",
			//   "type": "string"
			// }
			Description: "The ID of the portal.",
			Type:        types.StringType,
			Computed:    true,
		},
		"portal_name": {
			// Property: PortalName
			// CloudFormation resource type schema:
			// {
			//   "description": "A friendly name for the portal.",
			//   "type": "string"
			// }
			Description: "A friendly name for the portal.",
			Type:        types.StringType,
			Required:    true,
		},
		"portal_start_url": {
			// Property: PortalStartUrl
			// CloudFormation resource type schema:
			// {
			//   "description": "The public root URL for the AWS IoT AWS IoT SiteWise Monitor application portal.",
			//   "type": "string"
			// }
			Description: "The public root URL for the AWS IoT AWS IoT SiteWise Monitor application portal.",
			Type:        types.StringType,
			Computed:    true,
		},
		"role_arn": {
			// Property: RoleArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of a service role that allows the portal's users to access your AWS IoT SiteWise resources on your behalf.",
			//   "type": "string"
			// }
			Description: "The ARN of a service role that allows the portal's users to access your AWS IoT SiteWise resources on your behalf.",
			Type:        types.StringType,
			Required:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of key-value pairs that contain metadata for the portal.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "To add or update tag, provide both key and value. To delete tag, provide only tag key to be deleted.",
			//     "properties": {
			//       "Key": {
			//         "type": "string"
			//       },
			//       "Value": {
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
			//   "uniqueItems": false
			// }
			Description: "A list of key-value pairs that contain metadata for the portal.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
			// Tags is a write-only attribute.
		},
	}

	// Required for acceptance testing.
	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
	}

	schema := tfsdk.Schema{
		Description: "Resource schema for AWS::IoTSiteWise::Portal",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTSiteWise::Portal").WithTerraformTypeName("awscc_iotsitewise_portal").WithTerraformSchema(schema)

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Tags",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_iotsitewise_portal", "schema", hclog.Fmt("%v", schema))

	return resourceType, nil
}
