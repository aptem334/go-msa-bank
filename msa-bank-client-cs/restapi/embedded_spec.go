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
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is a sample server for Clients.",
    "title": "Swagger msa-bank-client-cs",
    "termsOfService": "http://swagger.io/terms/",
    "contact": {
      "email": "apiteam@swagger.io"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "host": "msa-bank-client-cs.swagger.io",
  "basePath": "/v1",
  "paths": {
    "/client": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "client-api"
        ],
        "summary": "Get all clients from the store",
        "operationId": "getClients",
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Client"
              }
            }
          },
          "405": {
            "description": "Invalid input"
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "put": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "client-api"
        ],
        "summary": "Update an existing client",
        "operationId": "updateClient",
        "parameters": [
          {
            "description": "Client object that needs to be added to the store",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Client"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Client"
            }
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Client not found"
          },
          "405": {
            "description": "Validation exception"
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "client-api"
        ],
        "summary": "Add a new client to the store",
        "operationId": "addClient",
        "parameters": [
          {
            "description": "Client object that needs to be added to the store",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Client"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "successful operation"
          },
          "405": {
            "description": "Invalid input"
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/client/{id}": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "client-api"
        ],
        "summary": "Delete client from the store",
        "operationId": "getClient",
        "parameters": [
          {
            "type": "string",
            "description": "Client object that needs to be added to the store",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Client"
            }
          },
          "405": {
            "description": "Invalid input"
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "delete": {
        "description": "Delete client",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "client-api"
        ],
        "summary": "Delete client from the store",
        "operationId": "deleteClient",
        "parameters": [
          {
            "type": "string",
            "description": "Client object that needs to be added to the store",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation"
          },
          "405": {
            "description": "Invalid input"
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Client": {
      "type": "object",
      "required": [
        "firstName",
        "lastName",
        "birthDate"
      ],
      "properties": {
        "birthDate": {
          "type": "string",
          "format": "date"
        },
        "firstName": {
          "type": "string"
        },
        "id": {
          "type": "string",
          "format": "uuid"
        },
        "lastName": {
          "type": "string"
        }
      }
    },
    "Error": {
      "type": "object",
      "title": "Error",
      "properties": {
        "errorMessage": {
          "type": "string"
        }
      }
    }
  },
  "tags": [
    {
      "description": "Client Api",
      "name": "client-api"
    }
  ],
  "externalDocs": {
    "description": "Find out more about Swagger",
    "url": "http://swagger.io"
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is a sample server for Clients.",
    "title": "Swagger msa-bank-client-cs",
    "termsOfService": "http://swagger.io/terms/",
    "contact": {
      "email": "apiteam@swagger.io"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "host": "msa-bank-client-cs.swagger.io",
  "basePath": "/v1",
  "paths": {
    "/client": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "client-api"
        ],
        "summary": "Get all clients from the store",
        "operationId": "getClients",
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Client"
              }
            }
          },
          "405": {
            "description": "Invalid input"
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "put": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "client-api"
        ],
        "summary": "Update an existing client",
        "operationId": "updateClient",
        "parameters": [
          {
            "description": "Client object that needs to be added to the store",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Client"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Client"
            }
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Client not found"
          },
          "405": {
            "description": "Validation exception"
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "client-api"
        ],
        "summary": "Add a new client to the store",
        "operationId": "addClient",
        "parameters": [
          {
            "description": "Client object that needs to be added to the store",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Client"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "successful operation"
          },
          "405": {
            "description": "Invalid input"
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/client/{id}": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "client-api"
        ],
        "summary": "Delete client from the store",
        "operationId": "getClient",
        "parameters": [
          {
            "type": "string",
            "description": "Client object that needs to be added to the store",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Client"
            }
          },
          "405": {
            "description": "Invalid input"
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "delete": {
        "description": "Delete client",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "client-api"
        ],
        "summary": "Delete client from the store",
        "operationId": "deleteClient",
        "parameters": [
          {
            "type": "string",
            "description": "Client object that needs to be added to the store",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation"
          },
          "405": {
            "description": "Invalid input"
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Client": {
      "type": "object",
      "required": [
        "firstName",
        "lastName",
        "birthDate"
      ],
      "properties": {
        "birthDate": {
          "type": "string",
          "format": "date"
        },
        "firstName": {
          "type": "string"
        },
        "id": {
          "type": "string",
          "format": "uuid"
        },
        "lastName": {
          "type": "string"
        }
      }
    },
    "Error": {
      "type": "object",
      "title": "Error",
      "properties": {
        "errorMessage": {
          "type": "string"
        }
      }
    }
  },
  "tags": [
    {
      "description": "Client Api",
      "name": "client-api"
    }
  ],
  "externalDocs": {
    "description": "Find out more about Swagger",
    "url": "http://swagger.io"
  }
}`))
}
