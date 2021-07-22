[![Build Status](https://app.travis-ci.com/michalq/generate.svg?branch=master)](https://app.travis-ci.com/michalq/generate)

# generate

Generates Go (golang) Structs and Validation code from JSON schema.

# Differences between original lib

## Rendering enums

This
```json
{
  "type": "object",
  "required": [
    "currencyCode"
  ],
  "properties": {
    "currencyCode": {
      "type": "string",
      "enum": [
        "PLN",
        "USD",
        "EUR"
      ]
    }
  },
  "$schema": "http://json-schema.org/draft-07/schema#"
}
```
Will generate code
```go
// EnumCurrencyCode
type EnumCurrencyCode string
const (
    EnumCurrencyCodeEUR EnumCurrencyCode = "EUR"
    EnumCurrencyCodePLN EnumCurrencyCode = "PLN"
    EnumCurrencyCodeUSD EnumCurrencyCode = "USD"
)

// Root 
type Root struct {
    CurrencyCode EnumCurrencyCode `json:"currencyCode" valid:"required"`
}
```
## Required and not required field rendering

All fields that are not required are treated as nullable, since its value can be undefined.

If given field is an array item or map value then its value is always required (unless it is nullable). Since value in map or array cannot be undefined.

## Objects rendering
 
Objects are rendered as value unless they are not required or nullable.

## OneOf rendering

If oneOf consists of two values and one of them is null then type is rendered with pointer.

If oneOf consists of more than two values then it is rendered as interface{}, but all types if not primitive will be created.

# Requirements

* Go 1.13+

# Usage

Install

```console
$ go get -u github.com/michalq/generate/...
```

or

Build

```console
$ make
```

Run

```console
$ schema-generate exampleschema.json
```

# Example

This schema

```json
{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "title": "Example",
  "id": "http://example.com/exampleschema.json",
  "type": "object",
  "description": "An example JSON Schema",
  "properties": {
    "name": {
      "type": "string"
    },
    "address": {
      "$ref": "#/definitions/address"
    },
    "status": {
      "$ref": "#/definitions/status"
    }
  },
  "definitions": {
    "address": {
      "id": "address",
      "type": "object",
      "description": "Address",
      "properties": {
        "street": {
          "type": "string",
          "description": "Address 1",
          "maxLength": 40
        },
        "houseNumber": {
          "type": "integer",
          "description": "House Number"
        }
      }
    },
    "status": {
      "type": "object",
      "properties": {
        "favouritecat": {
          "enum": [
            "A",
            "B",
            "C"
          ],
          "type": "string",
          "description": "The favourite cat.",
          "maxLength": 1
        }
      }
    }
  }
}
```

generates

```go
package main

type Address struct {
  HouseNumber int `json:"houseNumber,omitempty"`
  Street string `json:"street,omitempty"`
}

type Example struct {
  Address *Address `json:"address,omitempty"`
  Name string `json:"name,omitempty"`
  Status *Status `json:"status,omitempty"`
}

type Status struct {
  Favouritecat string `json:"favouritecat,omitempty"`
}
```

See the [test/](./test/) directory for more examples.
