// Code generated by generators/resource/main.go; DO NOT EDIT.

package autoscaling_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSAutoScalingScheduledAction_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::AutoScaling::ScheduledAction", "awscc_autoscaling_scheduled_action", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}
