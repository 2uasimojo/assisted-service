// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Bare metal inventory",
    "title": "BMInventory",
    "version": "1.0.0"
  },
  "host": "api.openshift.com",
  "basePath": "/api/bm-inventory/v1",
  "paths": {
    "/clusters": {
      "get": {
        "tags": [
          "inventory"
        ],
        "summary": "List OpenShift bare metal clusters",
        "operationId": "ListClusters",
        "responses": {
          "200": {
            "description": "Cluster list",
            "schema": {
              "$ref": "#/definitions/cluster-list"
            }
          },
          "500": {
            "description": "Internal server error"
          }
        }
      },
      "post": {
        "tags": [
          "inventory"
        ],
        "summary": "Register a new OpenShift bare metal cluster",
        "operationId": "RegisterCluster",
        "parameters": [
          {
            "description": "New cluster parameters",
            "name": "new-cluster-params",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/cluster-create-params"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Registered cluster",
            "schema": {
              "$ref": "#/definitions/cluster"
            }
          },
          "400": {
            "description": "Invalid input"
          },
          "500": {
            "description": "Internal server error"
          }
        }
      }
    },
    "/clusters/{cluster_id}": {
      "get": {
        "tags": [
          "inventory"
        ],
        "summary": "Retrieve OpenShift bare metal cluster information",
        "operationId": "GetCluster",
        "parameters": [
          {
            "type": "string",
            "description": "The ID of the cluster to retrieve",
            "name": "cluster_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Cluster information",
            "schema": {
              "$ref": "#/definitions/cluster"
            }
          },
          "404": {
            "description": "Cluster not found"
          }
        }
      },
      "delete": {
        "tags": [
          "inventory"
        ],
        "summary": "Deregister OpenShift bare metal cluster",
        "operationId": "DeregisterCluster",
        "parameters": [
          {
            "type": "string",
            "description": "The ID of the cluster to retrieve",
            "name": "cluster_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Cluster deregistered"
          },
          "404": {
            "description": "Cluster not found"
          }
        }
      }
    },
    "/images": {
      "get": {
        "tags": [
          "inventory"
        ],
        "summary": "List installation images",
        "operationId": "ListImages",
        "responses": {
          "200": {
            "description": "Image list",
            "schema": {
              "$ref": "#/definitions/image-list"
            }
          },
          "500": {
            "description": "Internal server error"
          }
        }
      },
      "post": {
        "tags": [
          "inventory"
        ],
        "summary": "Create an OpenShift bare metal cluster-assist installation image",
        "operationId": "CreateImage",
        "parameters": [
          {
            "description": "New image parameters",
            "name": "new-image-params",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/image-create-params"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created image",
            "schema": {
              "$ref": "#/definitions/image"
            }
          },
          "400": {
            "description": "Invalid input"
          },
          "500": {
            "description": "Internal server error"
          }
        }
      }
    },
    "/images/{image_id}": {
      "get": {
        "tags": [
          "inventory"
        ],
        "summary": "Retrieve installation image information",
        "operationId": "GetImage",
        "parameters": [
          {
            "type": "string",
            "description": "The ID of the image to retrieve",
            "name": "image_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Image information",
            "schema": {
              "$ref": "#/definitions/image"
            }
          },
          "404": {
            "description": "Image not found"
          }
        }
      }
    },
    "/nodes": {
      "get": {
        "tags": [
          "inventory"
        ],
        "summary": "List OpenShift bare metal nodes",
        "operationId": "ListNodes",
        "responses": {
          "200": {
            "description": "Node list",
            "schema": {
              "$ref": "#/definitions/node-list"
            }
          },
          "500": {
            "description": "Internal server error"
          }
        }
      },
      "post": {
        "tags": [
          "inventory"
        ],
        "summary": "Register a new OpenShift bare metal node",
        "operationId": "RegisterNode",
        "parameters": [
          {
            "description": "New node parameters",
            "name": "new-node-params",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/node-create-params"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Registered node",
            "schema": {
              "$ref": "#/definitions/node"
            }
          },
          "400": {
            "description": "Invalid input"
          },
          "500": {
            "description": "Internal server error"
          }
        }
      }
    },
    "/nodes/{node_id}": {
      "get": {
        "tags": [
          "inventory"
        ],
        "summary": "Retrieve OpenShift bare metal node information",
        "operationId": "GetNode",
        "parameters": [
          {
            "type": "string",
            "description": "The ID of the node to retrieve",
            "name": "node_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Node information",
            "schema": {
              "$ref": "#/definitions/node"
            }
          },
          "404": {
            "description": "Node not found"
          }
        }
      },
      "delete": {
        "tags": [
          "inventory"
        ],
        "summary": "Deregister OpenShift bare metal node",
        "operationId": "DeregisterNode",
        "parameters": [
          {
            "type": "string",
            "description": "The ID of the node to retrieve",
            "name": "node_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Node deregistered"
          },
          "400": {
            "description": "Node in use"
          },
          "404": {
            "description": "Node not found"
          }
        }
      }
    }
  },
  "definitions": {
    "base": {
      "type": "object",
      "required": [
        "kind",
        "id",
        "href"
      ],
      "properties": {
        "href": {
          "type": "string",
          "format": "uri"
        },
        "id": {
          "type": "string",
          "format": "uuid",
          "x-go-custom-tag": "gorm:\"primary_key\" query:\"filter,sort\""
        },
        "kind": {
          "type": "string",
          "enum": [
            "image",
            "node",
            "cluster"
          ]
        }
      }
    },
    "cluster": {
      "type": "object",
      "allOf": [
        {
          "$ref": "#/definitions/base"
        },
        {
          "$ref": "#/definitions/cluster-create-params"
        },
        {
          "type": "object",
          "required": [
            "status",
            "namespace"
          ],
          "properties": {
            "namespace": {
              "type": "string"
            },
            "status": {
              "type": "string",
              "enum": [
                "creating",
                "ready",
                "error"
              ]
            }
          }
        }
      ]
    },
    "cluster-create-params": {
      "type": "object",
      "required": [
        "name",
        "nodes"
      ],
      "properties": {
        "description": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "nodes": {
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "id": {
                "type": "string",
                "format": "uuid"
              },
              "role": {
                "type": "string",
                "enum": [
                  "master",
                  "worker"
                ]
              }
            }
          },
          "x-go-custom-tag": "gorm:\"type:varchar(64)[]\""
        }
      }
    },
    "cluster-list": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/cluster"
      }
    },
    "image": {
      "type": "object",
      "allOf": [
        {
          "$ref": "#/definitions/base"
        },
        {
          "$ref": "#/definitions/image-create-params"
        },
        {
          "type": "object",
          "required": [
            "status"
          ],
          "properties": {
            "download_url": {
              "type": "string",
              "format": "uri"
            },
            "status": {
              "type": "string",
              "enum": [
                "creating",
                "ready",
                "error"
              ]
            }
          }
        }
      ]
    },
    "image-create-params": {
      "type": "object",
      "required": [
        "name",
        "namespace"
      ],
      "properties": {
        "description": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "namespace": {
          "type": "string"
        },
        "proxy_ip": {
          "type": "string",
          "format": "hostname"
        },
        "proxy_port": {
          "type": "integer",
          "maximum": 65535
        }
      }
    },
    "image-list": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/image"
      }
    },
    "node": {
      "type": "object",
      "allOf": [
        {
          "$ref": "#/definitions/base"
        },
        {
          "$ref": "#/definitions/node-create-params"
        },
        {
          "type": "object",
          "required": [
            "kind",
            "status",
            "hardware_info"
          ],
          "properties": {
            "cluster_id": {
              "type": "string",
              "format": "uuid"
            },
            "status": {
              "type": "string",
              "enum": [
                "tbd"
              ]
            }
          }
        }
      ]
    },
    "node-create-params": {
      "type": "object",
      "required": [
        "namespace",
        "hardware_info",
        "serial"
      ],
      "properties": {
        "hardware_info": {
          "type": "string",
          "format": "json"
        },
        "namespace": {
          "type": "string"
        },
        "serial": {
          "type": "string"
        }
      }
    },
    "node-list": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/node"
      }
    }
  },
  "tags": [
    {
      "description": "Manage bare metal inventory",
      "name": "Bare metal inventory"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Bare metal inventory",
    "title": "BMInventory",
    "version": "1.0.0"
  },
  "host": "api.openshift.com",
  "basePath": "/api/bm-inventory/v1",
  "paths": {
    "/clusters": {
      "get": {
        "tags": [
          "inventory"
        ],
        "summary": "List OpenShift bare metal clusters",
        "operationId": "ListClusters",
        "responses": {
          "200": {
            "description": "Cluster list",
            "schema": {
              "$ref": "#/definitions/cluster-list"
            }
          },
          "500": {
            "description": "Internal server error"
          }
        }
      },
      "post": {
        "tags": [
          "inventory"
        ],
        "summary": "Register a new OpenShift bare metal cluster",
        "operationId": "RegisterCluster",
        "parameters": [
          {
            "description": "New cluster parameters",
            "name": "new-cluster-params",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/cluster-create-params"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Registered cluster",
            "schema": {
              "$ref": "#/definitions/cluster"
            }
          },
          "400": {
            "description": "Invalid input"
          },
          "500": {
            "description": "Internal server error"
          }
        }
      }
    },
    "/clusters/{cluster_id}": {
      "get": {
        "tags": [
          "inventory"
        ],
        "summary": "Retrieve OpenShift bare metal cluster information",
        "operationId": "GetCluster",
        "parameters": [
          {
            "type": "string",
            "description": "The ID of the cluster to retrieve",
            "name": "cluster_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Cluster information",
            "schema": {
              "$ref": "#/definitions/cluster"
            }
          },
          "404": {
            "description": "Cluster not found"
          }
        }
      },
      "delete": {
        "tags": [
          "inventory"
        ],
        "summary": "Deregister OpenShift bare metal cluster",
        "operationId": "DeregisterCluster",
        "parameters": [
          {
            "type": "string",
            "description": "The ID of the cluster to retrieve",
            "name": "cluster_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Cluster deregistered"
          },
          "404": {
            "description": "Cluster not found"
          }
        }
      }
    },
    "/images": {
      "get": {
        "tags": [
          "inventory"
        ],
        "summary": "List installation images",
        "operationId": "ListImages",
        "responses": {
          "200": {
            "description": "Image list",
            "schema": {
              "$ref": "#/definitions/image-list"
            }
          },
          "500": {
            "description": "Internal server error"
          }
        }
      },
      "post": {
        "tags": [
          "inventory"
        ],
        "summary": "Create an OpenShift bare metal cluster-assist installation image",
        "operationId": "CreateImage",
        "parameters": [
          {
            "description": "New image parameters",
            "name": "new-image-params",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/image-create-params"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created image",
            "schema": {
              "$ref": "#/definitions/image"
            }
          },
          "400": {
            "description": "Invalid input"
          },
          "500": {
            "description": "Internal server error"
          }
        }
      }
    },
    "/images/{image_id}": {
      "get": {
        "tags": [
          "inventory"
        ],
        "summary": "Retrieve installation image information",
        "operationId": "GetImage",
        "parameters": [
          {
            "type": "string",
            "description": "The ID of the image to retrieve",
            "name": "image_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Image information",
            "schema": {
              "$ref": "#/definitions/image"
            }
          },
          "404": {
            "description": "Image not found"
          }
        }
      }
    },
    "/nodes": {
      "get": {
        "tags": [
          "inventory"
        ],
        "summary": "List OpenShift bare metal nodes",
        "operationId": "ListNodes",
        "responses": {
          "200": {
            "description": "Node list",
            "schema": {
              "$ref": "#/definitions/node-list"
            }
          },
          "500": {
            "description": "Internal server error"
          }
        }
      },
      "post": {
        "tags": [
          "inventory"
        ],
        "summary": "Register a new OpenShift bare metal node",
        "operationId": "RegisterNode",
        "parameters": [
          {
            "description": "New node parameters",
            "name": "new-node-params",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/node-create-params"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Registered node",
            "schema": {
              "$ref": "#/definitions/node"
            }
          },
          "400": {
            "description": "Invalid input"
          },
          "500": {
            "description": "Internal server error"
          }
        }
      }
    },
    "/nodes/{node_id}": {
      "get": {
        "tags": [
          "inventory"
        ],
        "summary": "Retrieve OpenShift bare metal node information",
        "operationId": "GetNode",
        "parameters": [
          {
            "type": "string",
            "description": "The ID of the node to retrieve",
            "name": "node_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Node information",
            "schema": {
              "$ref": "#/definitions/node"
            }
          },
          "404": {
            "description": "Node not found"
          }
        }
      },
      "delete": {
        "tags": [
          "inventory"
        ],
        "summary": "Deregister OpenShift bare metal node",
        "operationId": "DeregisterNode",
        "parameters": [
          {
            "type": "string",
            "description": "The ID of the node to retrieve",
            "name": "node_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Node deregistered"
          },
          "400": {
            "description": "Node in use"
          },
          "404": {
            "description": "Node not found"
          }
        }
      }
    }
  },
  "definitions": {
    "ClusterCreateParamsNodesItems0": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "format": "uuid"
        },
        "role": {
          "type": "string",
          "enum": [
            "master",
            "worker"
          ]
        }
      }
    },
    "base": {
      "type": "object",
      "required": [
        "kind",
        "id",
        "href"
      ],
      "properties": {
        "href": {
          "type": "string",
          "format": "uri"
        },
        "id": {
          "type": "string",
          "format": "uuid",
          "x-go-custom-tag": "gorm:\"primary_key\" query:\"filter,sort\""
        },
        "kind": {
          "type": "string",
          "enum": [
            "image",
            "node",
            "cluster"
          ]
        }
      }
    },
    "cluster": {
      "type": "object",
      "allOf": [
        {
          "$ref": "#/definitions/base"
        },
        {
          "$ref": "#/definitions/cluster-create-params"
        },
        {
          "type": "object",
          "required": [
            "status",
            "namespace"
          ],
          "properties": {
            "namespace": {
              "type": "string"
            },
            "status": {
              "type": "string",
              "enum": [
                "creating",
                "ready",
                "error"
              ]
            }
          }
        }
      ]
    },
    "cluster-create-params": {
      "type": "object",
      "required": [
        "name",
        "nodes"
      ],
      "properties": {
        "description": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "nodes": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/ClusterCreateParamsNodesItems0"
          },
          "x-go-custom-tag": "gorm:\"type:varchar(64)[]\""
        }
      }
    },
    "cluster-list": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/cluster"
      }
    },
    "image": {
      "type": "object",
      "allOf": [
        {
          "$ref": "#/definitions/base"
        },
        {
          "$ref": "#/definitions/image-create-params"
        },
        {
          "type": "object",
          "required": [
            "status"
          ],
          "properties": {
            "download_url": {
              "type": "string",
              "format": "uri"
            },
            "status": {
              "type": "string",
              "enum": [
                "creating",
                "ready",
                "error"
              ]
            }
          }
        }
      ]
    },
    "image-create-params": {
      "type": "object",
      "required": [
        "name",
        "namespace"
      ],
      "properties": {
        "description": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "namespace": {
          "type": "string"
        },
        "proxy_ip": {
          "type": "string",
          "format": "hostname"
        },
        "proxy_port": {
          "type": "integer",
          "maximum": 65535,
          "minimum": 0
        }
      }
    },
    "image-list": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/image"
      }
    },
    "node": {
      "type": "object",
      "allOf": [
        {
          "$ref": "#/definitions/base"
        },
        {
          "$ref": "#/definitions/node-create-params"
        },
        {
          "type": "object",
          "required": [
            "kind",
            "status",
            "hardware_info"
          ],
          "properties": {
            "cluster_id": {
              "type": "string",
              "format": "uuid"
            },
            "status": {
              "type": "string",
              "enum": [
                "tbd"
              ]
            }
          }
        }
      ]
    },
    "node-create-params": {
      "type": "object",
      "required": [
        "namespace",
        "hardware_info",
        "serial"
      ],
      "properties": {
        "hardware_info": {
          "type": "string",
          "format": "json"
        },
        "namespace": {
          "type": "string"
        },
        "serial": {
          "type": "string"
        }
      }
    },
    "node-list": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/node"
      }
    }
  },
  "tags": [
    {
      "description": "Manage bare metal inventory",
      "name": "Bare metal inventory"
    }
  ]
}`))
}
