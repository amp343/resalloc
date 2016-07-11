package main

import (
  "encoding/json"
)

/**
 * Struct for receiving errors from the acl api
 */
type AclError struct {
  Status int `json:"status"`
  Message string `json:"error"`
}

/**
 * Convert a byte array to an AclError
 */
func AclErrorFromBytes(b []byte) AclError {
  err := AclError{}
  json.Unmarshal(b, &err)

  return err
}
