// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateCluster creates a cluster
*/
func (a *Client) CreateCluster(params *CreateClusterParams, authInfo runtime.ClientAuthInfoWriter) (*CreateClusterCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateCluster",
		Method:             "POST",
		PathPattern:        "/api/v1/clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateClusterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateClusterCreated), nil

}

/*
GetBootstrapConfig gets bootstrap config to onboard a node
*/
func (a *Client) GetBootstrapConfig(params *GetBootstrapConfigParams, authInfo runtime.ClientAuthInfoWriter) (*GetBootstrapConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBootstrapConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetBootstrapConfig",
		Method:             "GET",
		PathPattern:        "/api/v1/clusters/{name}/bootstrap",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetBootstrapConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetBootstrapConfigOK), nil

}

/*
GetClusterCredentials gets user specific credentials to access the cluster
*/
func (a *Client) GetClusterCredentials(params *GetClusterCredentialsParams, authInfo runtime.ClientAuthInfoWriter) (*GetClusterCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterCredentialsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetClusterCredentials",
		Method:             "GET",
		PathPattern:        "/api/v1/clusters/{name}/credentials",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetClusterCredentialsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClusterCredentialsOK), nil

}

/*
GetClusterCredentialsOIDC gets user specific credentials to access the cluster with o ID c
*/
func (a *Client) GetClusterCredentialsOIDC(params *GetClusterCredentialsOIDCParams, authInfo runtime.ClientAuthInfoWriter) (*GetClusterCredentialsOIDCOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterCredentialsOIDCParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetClusterCredentialsOIDC",
		Method:             "GET",
		PathPattern:        "/api/v1/clusters/{name}/credentials/oidc",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetClusterCredentialsOIDCReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClusterCredentialsOIDCOK), nil

}

/*
GetClusterEvents gets recent events about the cluster
*/
func (a *Client) GetClusterEvents(params *GetClusterEventsParams, authInfo runtime.ClientAuthInfoWriter) (*GetClusterEventsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterEventsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetClusterEvents",
		Method:             "GET",
		PathPattern:        "/api/v1/clusters/{name}/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetClusterEventsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClusterEventsOK), nil

}

/*
GetClusterInfo gets user specific info about the cluster
*/
func (a *Client) GetClusterInfo(params *GetClusterInfoParams, authInfo runtime.ClientAuthInfoWriter) (*GetClusterInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterInfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetClusterInfo",
		Method:             "GET",
		PathPattern:        "/api/v1/clusters/{name}/info",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetClusterInfoReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClusterInfoOK), nil

}

/*
GetClusterValues gets values for cluster chart admin only
*/
func (a *Client) GetClusterValues(params *GetClusterValuesParams, authInfo runtime.ClientAuthInfoWriter) (*GetClusterValuesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterValuesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetClusterValues",
		Method:             "GET",
		PathPattern:        "/api/v1/{account}/clusters/{name}/values",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetClusterValuesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClusterValuesOK), nil

}

/*
GetOpenstackMetadata grabs bag of openstack metadata
*/
func (a *Client) GetOpenstackMetadata(params *GetOpenstackMetadataParams, authInfo runtime.ClientAuthInfoWriter) (*GetOpenstackMetadataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOpenstackMetadataParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetOpenstackMetadata",
		Method:             "GET",
		PathPattern:        "/api/v1/openstack/metadata",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOpenstackMetadataReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetOpenstackMetadataOK), nil

}

/*
Info gets info about kubernikus
*/
func (a *Client) Info(params *InfoParams) (*InfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Info",
		Method:             "GET",
		PathPattern:        "/info",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &InfoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*InfoOK), nil

}

/*
ListAPIVersions lists available api versions
*/
func (a *Client) ListAPIVersions(params *ListAPIVersionsParams) (*ListAPIVersionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAPIVersionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListAPIVersions",
		Method:             "GET",
		PathPattern:        "/api",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListAPIVersionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListAPIVersionsOK), nil

}

/*
ListClusters lists available clusters
*/
func (a *Client) ListClusters(params *ListClustersParams, authInfo runtime.ClientAuthInfoWriter) (*ListClustersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListClustersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListClusters",
		Method:             "GET",
		PathPattern:        "/api/v1/clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListClustersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListClustersOK), nil

}

/*
ShowCluster shows the specified cluser
*/
func (a *Client) ShowCluster(params *ShowClusterParams, authInfo runtime.ClientAuthInfoWriter) (*ShowClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewShowClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ShowCluster",
		Method:             "GET",
		PathPattern:        "/api/v1/clusters/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ShowClusterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ShowClusterOK), nil

}

/*
TerminateCluster terminates the specified cluster
*/
func (a *Client) TerminateCluster(params *TerminateClusterParams, authInfo runtime.ClientAuthInfoWriter) (*TerminateClusterAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTerminateClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TerminateCluster",
		Method:             "DELETE",
		PathPattern:        "/api/v1/clusters/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TerminateClusterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TerminateClusterAccepted), nil

}

/*
UpdateCluster updates the specified cluser
*/
func (a *Client) UpdateCluster(params *UpdateClusterParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateCluster",
		Method:             "PUT",
		PathPattern:        "/api/v1/clusters/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateClusterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateClusterOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
