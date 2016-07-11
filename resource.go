package main

import (
  "strconv"
  "encoding/json"
)

/**
 * A ResourceArray is a collection of Resource-s
 */
type ResourceArray struct {
  Collection []Resource
}

/**
 * Define a struct for resource records
 * received from the acl api
 */
type Resource struct {
  Id int `json:"id"`
  LeasedToUser string `json:"leased_to_user"`
  Name string `json:"name"`
  Os string `json:"os"`
  IsAvailable bool `json:"is_available"`
  LeasedAt string `json:"leased_at"`
  LeasedUntil string `json:"leased_until"`
  LeaseTimeRemaining string `json:"lease_time_remaining"`
}

/**
 * Given a resource struct, return a table row array
 */
func (r Resource) ToTableRow() []string {
  return []string{r.Name, r.Os, strconv.FormatBool(r.IsAvailable), r.LeasedToUser, r.LeaseTimeRemaining}
}

/**
 * Given a byte array, return a Resources struct
 */
func ResourceArrayFromBytes(b []byte) []Resource {
  resources := make([]Resource, 0)
  json.Unmarshal(b, &resources)

  return resources
}

/**
 * Given a byte array, return a Resource struct
 */
func ResourceFromBytes(b []byte) Resource {
  resource := Resource{}
  json.Unmarshal(b, &resource)

  return resource
}
