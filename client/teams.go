package client

import (
	"reflect"
	
	"github.com/pkg/errors"
	"github.com/google/jsonapi"
	rootlygo "github.com/rootlyhq/terraform-provider-rootly/schema"
)

type Team struct {
	ID string `jsonapi:"primary,groups"`
	Name string `jsonapi:"attr,name,omitempty"`
  Slug string `jsonapi:"attr,slug,omitempty"`
  Description string `jsonapi:"attr,description,omitempty"`
  NotifyEmails []interface{} `jsonapi:"attr,notify_emails,omitempty"`
  Color string `jsonapi:"attr,color,omitempty"`
  Position int `jsonapi:"attr,position,omitempty"`
  PagerdutyId string `jsonapi:"attr,pagerduty_id,omitempty"`
  OpsgenieId string `jsonapi:"attr,opsgenie_id,omitempty"`
  VictorOpsId string `jsonapi:"attr,victor_ops_id,omitempty"`
  PagertreeId string `jsonapi:"attr,pagertree_id,omitempty"`
  SlackChannels []interface{} `jsonapi:"attr,slack_channels,omitempty"`
  SlackAliases []interface{} `jsonapi:"attr,slack_aliases,omitempty"`
}

func (c *Client) ListTeams(params *rootlygo.ListTeamsParams) ([]interface{}, error) {
	req, err := rootlygo.NewListTeamsRequest(c.Rootly.Server, params)
	if err != nil {
		return nil, errors.Errorf("Error building request: %s", err.Error())
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, errors.Errorf("Failed to make request: %s", err.Error())
	}

	teams, err := jsonapi.UnmarshalManyPayload(resp.Body, reflect.TypeOf(new(Team)))
	if err != nil {
		return nil, errors.Errorf("Error unmarshaling: %s", err.Error())
	}

	return teams, nil
}

func (c *Client) CreateTeam(d *Team) (*Team, error) {
	buffer, err := MarshalData(d)
	if err != nil {
		return nil, errors.Errorf("Error marshaling team: %s", err.Error())
	}

	req, err := rootlygo.NewCreateTeamRequestWithBody(c.Rootly.Server, c.ContentType, buffer)
	if err != nil {
		return nil, errors.Errorf("Error building request: %s", err.Error())
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, errors.Errorf("Failed to perform request to create team: %s", err.Error())
	}

	data, err := UnmarshalData(resp.Body, new(Team))
	if err != nil {
		return nil, errors.Errorf("Error unmarshaling team: %s", err.Error())
	}

	return data.(*Team), nil
}

func (c *Client) GetTeam(id string) (*Team, error) {
	req, err := rootlygo.NewGetTeamRequest(c.Rootly.Server, id)
	if err != nil {
		return nil, errors.Errorf("Error building request: %s", err.Error())
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, errors.Errorf("Failed to make request to get team: %s", err.Error())
	}

	data, err := UnmarshalData(resp.Body, new(Team))
	if err != nil {
		return nil, errors.Errorf("Error unmarshaling team: %s", err.Error())
	}

	return data.(*Team), nil
}

func (c *Client) UpdateTeam(id string, team *Team) (*Team, error) {
	buffer, err := MarshalData(team)
	if err != nil {
		return nil, errors.Errorf("Error marshaling team: %s", err.Error())
	}

	req, err := rootlygo.NewUpdateTeamRequestWithBody(c.Rootly.Server, id, c.ContentType, buffer)
	if err != nil {
		return nil, errors.Errorf("Error building request: %s", err.Error())
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, errors.Errorf("Failed to make request to update team: %s", err.Error())
	}

	data, err := UnmarshalData(resp.Body, new(Team))
	if err != nil {
		return nil, errors.Errorf("Error unmarshaling team: %s", err.Error())
	}

	return data.(*Team), nil
}

func (c *Client) DeleteTeam(id string) error {
	req, err := rootlygo.NewDeleteTeamRequest(c.Rootly.Server, id)
	if err != nil {
		return errors.Errorf("Error building request: %s", err.Error())
	}

	_, err = c.Do(req)
	if err != nil {
		return errors.Errorf("Failed to make request to delete team: %s", err.Error())
	}

	return nil
}
