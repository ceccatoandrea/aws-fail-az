package ecs

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/mcastellin/aws-fail-az/awsapis"
	"github.com/mcastellin/aws-fail-az/service/awsutils"
	"github.com/mcastellin/aws-fail-az/state"
)

// The resource type key to use for storing state of ECS services
const RESOURCE_TYPE string = "ecs-service"

type ECSService struct {
	Provider    *awsapis.AWSProvider
	ClusterArn  string
	ServiceName string

	stateSubnets []string
}

type ECSServiceState struct {
	ServiceName string   `json:"service"`
	ClusterArn  string   `json:"cluster"`
	Subnets     []string `json:"subnets"`
}

func (svc ECSService) Check() (bool, error) {
	isValid := true

	log.Printf("%s cluster=%s,name=%s: checking resource state before failure simulation",
		RESOURCE_TYPE, svc.ClusterArn, svc.ServiceName)

	api := awsapis.NewEcsApi(svc.Provider)

	result, err := serviceStable(api, svc.ClusterArn, svc.ServiceName)
	if err != nil {
		return false, nil
	} else {
		isValid = isValid && result
	}

	return isValid, nil
}

func (svc ECSService) Save(stateManager state.StateManager) error {
	api := awsapis.NewEcsApi(svc.Provider)

	input := &ecs.DescribeServicesInput{
		Cluster:  aws.String(svc.ClusterArn),
		Services: []string{*aws.String(svc.ServiceName)},
	}

	describeOutput, err := api.DescribeServices(context.TODO(), input)
	if err != nil {
		return err
	}

	service := describeOutput.Services[0]
	subnets := service.NetworkConfiguration.AwsvpcConfiguration.Subnets

	state := &ECSServiceState{
		ClusterArn:  svc.ClusterArn,
		ServiceName: svc.ServiceName,
		Subnets:     subnets,
	}

	data, err := json.Marshal(state)
	if err != nil {
		log.Println("Error while marshalling service state")
		return err
	}

	resourceKey := fmt.Sprintf("%s-%s", svc.ClusterArn, svc.ServiceName)
	err = stateManager.Save(RESOURCE_TYPE, resourceKey, data)
	if err != nil {
		return err
	}

	return nil
}

func (svc ECSService) Fail(azs []string) error {
	ec2Api := awsapis.NewEc2Api(svc.Provider)
	ecsApi := awsapis.NewEcsApi(svc.Provider)

	input := &ecs.DescribeServicesInput{
		Cluster:  aws.String(svc.ClusterArn),
		Services: []string{*aws.String(svc.ServiceName)},
	}

	describeOutput, err := ecsApi.DescribeServices(context.TODO(), input)
	if err != nil {
		return err
	}

	service := describeOutput.Services[0]
	subnets := service.NetworkConfiguration.AwsvpcConfiguration.Subnets

	newSubnets, err := awsutils.FilterSubnetsNotInAzs(ec2Api, subnets, azs)
	if err != nil {
		log.Printf("Error while filtering subnets by AZs: %v", err)
		return err
	}

	if len(newSubnets) == 0 {
		return fmt.Errorf("AZ failure for service %s would remove all available subnets. Service failure will now stop.", svc.ServiceName)
	}

	log.Printf("%s cluster=%s,name=%s: failing AZs %s for ecs-service",
		RESOURCE_TYPE, svc.ClusterArn, svc.ServiceName, azs)

	updatedNetworkConfig := service.NetworkConfiguration
	updatedNetworkConfig.AwsvpcConfiguration.Subnets = newSubnets

	updateServiceInput := &ecs.UpdateServiceInput{
		Cluster:              aws.String(svc.ClusterArn),
		Service:              aws.String(svc.ServiceName),
		TaskDefinition:       service.TaskDefinition,
		NetworkConfiguration: updatedNetworkConfig,
	}

	_, err = ecsApi.UpdateService(context.TODO(), updateServiceInput)
	if err != nil {
		return err
	}

	err = stopTasksInRemovedSubnets(ecsApi, svc.ClusterArn, svc.ServiceName, newSubnets)
	if err != nil {
		return err
	}

	return nil
}

func (svc ECSService) Restore() error {
	log.Printf("%s cluster=%s,name=%s: restoring AZs for ecs-service",
		RESOURCE_TYPE, svc.ClusterArn, svc.ServiceName)

	api := awsapis.NewEcsApi(svc.Provider)

	input := &ecs.DescribeServicesInput{
		Cluster:  aws.String(svc.ClusterArn),
		Services: []string{*aws.String(svc.ServiceName)},
	}

	describeOutput, err := api.DescribeServices(context.TODO(), input)
	if err != nil {
		return err
	}

	service := describeOutput.Services[0]

	updatedNetworkConfig := service.NetworkConfiguration
	updatedNetworkConfig.AwsvpcConfiguration.Subnets = svc.stateSubnets

	updateServiceInput := &ecs.UpdateServiceInput{
		Cluster:              aws.String(svc.ClusterArn),
		Service:              aws.String(svc.ServiceName),
		TaskDefinition:       service.TaskDefinition,
		NetworkConfiguration: updatedNetworkConfig,
	}

	_, err = api.UpdateService(context.TODO(), updateServiceInput)
	if err != nil {
		return err
	}
	return nil
}