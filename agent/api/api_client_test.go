package api

import (
	"strings"
	"testing"

	"github.com/aws/amazon-ecs-agent/agent/auth"
	"github.com/aws/amazon-ecs-agent/agent/config"
	svc "github.com/aws/amazon-ecs-agent/agent/ecs_client/ecs_autogenerated_client"
)

type MockAmazonEC2ContainerServiceV20141113Client struct {
	svc.AmazonEC2ContainerServiceV20141113Client

	lastRequest interface{}
}

func (mock *MockAmazonEC2ContainerServiceV20141113Client) SubmitContainerStateChange(req svc.SubmitContainerStateChangeRequest) (svc.SubmitContainerStateChangeResponse, error) {
	mock.lastRequest = req

	return svc.NewSubmitContainerStateChangeResponse(), nil
}

func NewMockClient() (ECSClient, *MockAmazonEC2ContainerServiceV20141113Client) {
	client := NewECSClient(auth.TestCredentialProvider{}, &config.Config{Cluster: "default"}, false)
	mockSvcClient := &MockAmazonEC2ContainerServiceV20141113Client{}
	client.(*ApiECSClient).serviceClientFn = func() (svc.AmazonEC2ContainerServiceV20141113, error) {
		return mockSvcClient, nil
	}
	return client, mockSvcClient
}

func TestSubmitContainerStateChange(t *testing.T) {
	client, mockSvcClient := NewMockClient()
	err := client.SubmitContainerStateChange(ContainerStateChange{
		TaskArn:       "arn",
		ContainerName: "cont",
		Status:        ContainerRunning,
	})
	if err != nil {
		t.Error("Unable to submit trivial container state change: %v", err)
	}
	req := mockSvcClient.lastRequest.(svc.SubmitContainerStateChangeRequest)
	if req == nil {
		t.Error("Expected request to be present")
	}
	if *req.Cluster() != "default" {
		t.Error("Submitted wrong cluster")
	}
	if *req.Task() != "arn" {
		t.Error("Submitted wrong arn")
	}
	if *req.ContainerName() != "cont" {
		t.Error("Submitted wrong container name")
	}
	if *req.Status() != "RUNNING" {
		t.Error("Submitted wrong status")
	}
}

func TestSubmitContainerStateChange2(t *testing.T) {
	client, mockSvcClient := NewMockClient()
	exitCode := 20
	err := client.SubmitContainerStateChange(ContainerStateChange{
		TaskArn:       "arn",
		ContainerName: "cont",
		Status:        ContainerStopped,
		ExitCode:      &exitCode,
		Reason:        "I exited",
	})
	if err != nil {
		t.Error("Unable to submit container state change: %v", err)
	}
	req := mockSvcClient.lastRequest.(svc.SubmitContainerStateChangeRequest)
	if req == nil {
		t.Error("Expected request to be present")
	}
	if *req.Reason() != "I exited" {
		t.Error("Submitted wrong reason")
	}
	if *req.ExitCode() != 20 {
		t.Error("Submitted wrong exit code")
	}
}

func TestSubmitCOntainerStateChangeReason(t *testing.T) {
	client, mockSvcClient := NewMockClient()

	err := client.SubmitContainerStateChange(ContainerStateChange{
		TaskArn:       "arn",
		ContainerName: "cont",
		Status:        ContainerStopped,
		Reason:        strings.Repeat("a", EcsMaxReasonLength),
	})
	if err != nil {
		t.Error("Unable to submit container state change: %v", err)
	}
	req := mockSvcClient.lastRequest.(svc.SubmitContainerStateChangeRequest)
	if *req.Reason() != strings.Repeat("a", EcsMaxReasonLength) {
		t.Error("Submitted wrong reason")
	}
}

func TestSubmitContainerStateChangeLongReason(t *testing.T) {
	client, mockSvcClient := NewMockClient()
	// Test reason gets trimmed
	err := client.SubmitContainerStateChange(ContainerStateChange{
		TaskArn:       "arn",
		ContainerName: "cont",
		Status:        ContainerStopped,
		Reason:        strings.Repeat("a", EcsMaxReasonLength+1),
	})
	if err != nil {
		t.Error("Unable to submit container state change: %v", err)
	}
	req := mockSvcClient.lastRequest.(svc.SubmitContainerStateChangeRequest)
	if *req.Reason() != strings.Repeat("a", EcsMaxReasonLength) {
		t.Error("Submitted wrong reason")
	}
}
