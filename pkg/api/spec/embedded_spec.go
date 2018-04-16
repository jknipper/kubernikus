// Code generated by go-swagger; DO NOT EDIT.

package spec

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Kubernikus",
    "version": "1.0.0"
  },
  "paths": {
    "/api": {
      "get": {
        "summary": "List available api versions",
        "operationId": "ListAPIVersions",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/ApiVersions"
            }
          }
        }
      }
    },
    "/api/v1/clusters": {
      "get": {
        "summary": "List available clusters",
        "operationId": "ListClusters",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Kluster"
              }
            }
          },
          "default": {
            "$ref": "#/responses/errorResponse"
          }
        }
      },
      "post": {
        "summary": "Create a cluster",
        "operationId": "CreateCluster",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Kluster"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/Kluster"
            }
          },
          "default": {
            "$ref": "#/responses/errorResponse"
          }
        }
      }
    },
    "/api/v1/clusters/{name}": {
      "get": {
        "summary": "Show the specified cluser",
        "operationId": "ShowCluster",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/Kluster"
            }
          },
          "default": {
            "$ref": "#/responses/errorResponse"
          }
        }
      },
      "put": {
        "summary": "Update the specified cluser",
        "operationId": "UpdateCluster",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Kluster"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/Kluster"
            }
          },
          "default": {
            "$ref": "#/responses/errorResponse"
          }
        }
      },
      "delete": {
        "summary": "Terminate the specified cluster",
        "operationId": "TerminateCluster",
        "responses": {
          "202": {
            "description": "OK"
          },
          "default": {
            "$ref": "#/responses/errorResponse"
          }
        }
      },
      "parameters": [
        {
          "uniqueItems": true,
          "type": "string",
          "name": "name",
          "in": "path",
          "required": true
        }
      ]
    },
    "/api/v1/clusters/{name}/credentials": {
      "get": {
        "summary": "Get user specific credentials to access the cluster",
        "operationId": "GetClusterCredentials",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/Credentials"
            }
          },
          "default": {
            "$ref": "#/responses/errorResponse"
          }
        }
      },
      "parameters": [
        {
          "uniqueItems": true,
          "type": "string",
          "name": "name",
          "in": "path",
          "required": true
        }
      ]
    },
    "/api/v1/clusters/{name}/events": {
      "get": {
        "summary": "Get recent events about the cluster",
        "operationId": "GetClusterEvents",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Event"
              }
            }
          },
          "default": {
            "$ref": "#/responses/errorResponse"
          }
        }
      },
      "parameters": [
        {
          "uniqueItems": true,
          "type": "string",
          "name": "name",
          "in": "path",
          "required": true
        }
      ]
    },
    "/api/v1/clusters/{name}/info": {
      "get": {
        "summary": "Get user specific info about the cluster",
        "operationId": "GetClusterInfo",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/KlusterInfo"
            }
          },
          "default": {
            "$ref": "#/responses/errorResponse"
          }
        }
      },
      "parameters": [
        {
          "uniqueItems": true,
          "type": "string",
          "name": "name",
          "in": "path",
          "required": true
        }
      ]
    },
    "/api/v1/openstack/metadata": {
      "get": {
        "summary": "Grab bag of openstack metadata",
        "operationId": "GetOpenstackMetadata",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/OpenstackMetadata"
            }
          },
          "default": {
            "$ref": "#/responses/errorResponse"
          }
        }
      }
    },
    "/info": {
      "get": {
        "summary": "Get info about Kubernikus",
        "operationId": "Info",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/Info"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ApiVersions": {
      "required": [
        "versions"
      ],
      "properties": {
        "versions": {
          "description": "versions are the api versions that are available.",
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "Credentials": {
      "type": "object",
      "properties": {
        "kubeconfig": {
          "type": "string"
        }
      }
    },
    "Event": {
      "type": "object",
      "properties": {
        "count": {
          "description": "The number of times this event has occurred.",
          "type": "integer"
        },
        "firstTimestamp": {
          "description": "The time at which the event was first recorded",
          "type": "string"
        },
        "lastTimestamp": {
          "description": "The time at which the most recent occurrence of this event was recorded",
          "type": "string"
        },
        "message": {
          "description": "A human-readable description of the event",
          "type": "string"
        },
        "reason": {
          "description": "A short, machine understandable string that gives the reason for the event",
          "type": "string"
        },
        "type": {
          "description": "Type of this event",
          "type": "string",
          "enum": [
            "Normal",
            "Warning"
          ]
        }
      }
    },
    "Info": {
      "properties": {
        "version": {
          "type": "string"
        }
      }
    },
    "Kluster": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "name": {
          "description": "name of the cluster",
          "type": "string",
          "maxLength": 20,
          "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
          "x-nullable": false
        },
        "spec": {
          "$ref": "#/definitions/KlusterSpec"
        },
        "status": {
          "$ref": "#/definitions/KlusterStatus"
        }
      }
    },
    "KlusterInfo": {
      "properties": {
        "binaries": {
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "links": {
                "type": "array",
                "items": {
                  "type": "object",
                  "properties": {
                    "link": {
                      "type": "string"
                    },
                    "platform": {
                      "type": "string"
                    }
                  },
                  "x-go-name": "Link",
                  "x-nullable": false
                }
              },
              "name": {
                "type": "string"
              }
            },
            "x-go-name": "Binaries",
            "x-nullable": false
          }
        },
        "setupCommand": {
          "type": "string"
        }
      }
    },
    "KlusterPhase": {
      "type": "string",
      "enum": [
        "Pending",
        "Creating",
        "Running",
        "Terminating"
      ]
    },
    "KlusterSpec": {
      "type": "object",
      "properties": {
        "advertiseAddress": {
          "type": "string",
          "default": "1.1.1.1",
          "x-nullable": false
        },
        "clusterCIDR": {
          "description": "CIDR Range for Pods in the cluster. Can not be updated.",
          "type": "string",
          "default": "100.100.0.0/16",
          "pattern": "^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(\\/([0-9]|[1-2][0-9]|3[0-2]))$",
          "x-nullable": false
        },
        "dnsAddress": {
          "type": "string"
        },
        "dnsDomain": {
          "type": "string",
          "default": "cluster.local",
          "x-nullable": false
        },
        "name": {
          "type": "string",
          "readOnly": true
        },
        "nodePools": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/NodePool"
          }
        },
        "openstack": {
          "$ref": "#/definitions/OpenstackSpec"
        },
        "serviceCIDR": {
          "description": "CIDR Range for Services in the cluster. Can not be updated.",
          "type": "string",
          "default": "198.18.128.0/17",
          "pattern": "^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])(\\/([0-9]|[1-2][0-9]|3[0-2]))$",
          "x-nullable": false
        },
        "sshPublicKey": {
          "description": "SSH public key that is injected into spawned nodes.",
          "type": "string"
        },
        "version": {
          "type": "string",
          "pattern": "^[0-9]+\\.[0-9]+\\.[0-9]+$",
          "readOnly": true
        }
      },
      "x-nullable": false
    },
    "KlusterStatus": {
      "type": "object",
      "properties": {
        "apiserver": {
          "type": "string"
        },
        "message": {
          "type": "string"
        },
        "migrationsPending": {
          "type": "boolean"
        },
        "nodePools": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/NodePoolInfo"
          }
        },
        "phase": {
          "$ref": "#/definitions/KlusterPhase"
        },
        "specVersion": {
          "type": "integer"
        },
        "version": {
          "type": "string"
        },
        "wormhole": {
          "type": "string"
        }
      },
      "x-nullable": false,
      "readOnly": true
    },
    "NodePool": {
      "type": "object",
      "required": [
        "name",
        "flavor"
      ],
      "properties": {
        "config": {
          "$ref": "#/definitions/NodePoolConfig"
        },
        "flavor": {
          "type": "string",
          "x-nullable": false
        },
        "image": {
          "type": "string",
          "default": "coreos-stable-amd64",
          "x-nullable": false
        },
        "name": {
          "type": "string",
          "pattern": "^[a-z]([a-z0-9]*)?$",
          "x-nullable": false
        },
        "size": {
          "type": "integer",
          "default": 0,
          "maximum": 127,
          "minimum": 0,
          "x-nullable": false
        }
      },
      "x-nullable": false
    },
    "NodePoolConfig": {
      "type": "object",
      "properties": {
        "repair": {
          "type": "boolean"
        },
        "upgrade": {
          "type": "boolean"
        }
      },
      "x-nullable": false
    },
    "NodePoolInfo": {
      "type": "object",
      "properties": {
        "healthy": {
          "type": "integer"
        },
        "name": {
          "type": "string"
        },
        "running": {
          "type": "integer"
        },
        "schedulable": {
          "type": "integer"
        },
        "size": {
          "type": "integer"
        }
      },
      "x-nullable": false
    },
    "OpenstackMetadata": {
      "type": "object",
      "properties": {
        "flavors": {
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "id": {
                "type": "string"
              },
              "name": {
                "type": "string"
              },
              "ram": {
                "type": "integer"
              },
              "vcpus": {
                "type": "integer"
              }
            },
            "x-go-name": "Flavor",
            "x-nullable": false
          }
        },
        "keyPairs": {
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "name": {
                "type": "string"
              },
              "publicKey": {
                "type": "string"
              }
            },
            "x-go-name": "KeyPair"
          }
        },
        "routers": {
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "externalNetworkID": {
                "type": "string",
                "x-go-name": "ExternalNetworkID"
              },
              "id": {
                "type": "string"
              },
              "name": {
                "type": "string"
              },
              "networks": {
                "type": "array",
                "items": {
                  "type": "object",
                  "properties": {
                    "id": {
                      "type": "string"
                    },
                    "name": {
                      "type": "string"
                    },
                    "subnets": {
                      "type": "array",
                      "items": {
                        "type": "object",
                        "properties": {
                          "CIDR": {
                            "type": "string"
                          },
                          "id": {
                            "type": "string"
                          },
                          "name": {
                            "type": "string"
                          }
                        },
                        "x-go-name": "Subnet"
                      }
                    }
                  },
                  "x-go-name": "Network"
                }
              }
            },
            "x-go-name": "Router"
          }
        },
        "securityGroups": {
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "id": {
                "type": "string"
              },
              "name": {
                "type": "string"
              }
            },
            "x-go-name": "SecurityGroup"
          }
        }
      }
    },
    "OpenstackSpec": {
      "type": "object",
      "properties": {
        "lbFloatingNetworkID": {
          "type": "string",
          "x-go-name": "LBFloatingNetworkID"
        },
        "lbSubnetID": {
          "type": "string",
          "x-go-name": "LBSubnetID"
        },
        "networkID": {
          "type": "string"
        },
        "projectID": {
          "type": "string"
        },
        "routerID": {
          "type": "string"
        },
        "securityGroupName": {
          "type": "string"
        }
      },
      "x-nullable": false
    },
    "Principal": {
      "type": "object",
      "properties": {
        "account": {
          "description": "account id",
          "type": "string"
        },
        "authUrl": {
          "description": "Identity Endpoint",
          "type": "string"
        },
        "domain": {
          "description": "user's domain name",
          "type": "string"
        },
        "id": {
          "description": "userid",
          "type": "string"
        },
        "name": {
          "description": "username",
          "type": "string"
        },
        "roles": {
          "description": "list of roles the user has in the given scope",
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "error": {
      "description": "the error model is a model for all the error responses coming from Kubernikus\n",
      "type": "object",
      "required": [
        "message",
        "code"
      ],
      "properties": {
        "code": {
          "description": "The error code",
          "type": "integer"
        },
        "helpUrl": {
          "description": "link to help page explaining the error in more detail",
          "type": "string",
          "format": "uri"
        },
        "message": {
          "description": "The error message",
          "type": "string"
        }
      }
    }
  },
  "responses": {
    "errorResponse": {
      "description": "Error",
      "schema": {
        "$ref": "#/definitions/error"
      }
    }
  },
  "securityDefinitions": {
    "keystone": {
      "description": "OpenStack Keystone Authentication",
      "type": "apiKey",
      "name": "x-auth-token",
      "in": "header"
    }
  },
  "security": [
    {
      "keystone": []
    }
  ]
}`))
}
