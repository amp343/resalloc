package main

import (
  "strconv"
  "github.com/urfave/cli"
)

type Action func(c *cli.Context) (error)

/**
 * wrap actions in lines of padding for
 * consistent output
 */
func Do(a Action) Action {
  return func (c *cli.Context) error {
    LogN()
    a(c)
    LogN()
    return nil
  }
}

/**
 * invoked by the cli to unlease a resource
 * (resalloc unlease {{ serverName }})
 */
func UnleaseResourceAction(c *cli.Context) error {

  // get cli argument
  requestedResource := c.Args().First()
  if requestedResource == "" {
    LogError("No resource specified; use: `resalloc unlease {{ serverName }}``")
    return nil
  }

  // request acl to unlease resource
  Log("Attempting to unlease resource: " + requestedResource + " ...")
  resp, err := UnleaseResource(requestedResource)

  // handle acl service unavailable
  if err != nil {
    return err
  }

  b := ResponseBodyToBytes(resp.Body)
  LogStatusCode(resp.StatusCode)

  // error case
  if ResponseIsError(resp) {
    err := AclErrorFromBytes(b)
    LogError(err.Message)
    Log("Try `resalloc list` to see leasable servers or `resalloc leased` to see servers leased to you")
    return nil
  }

  // success case
  resource := ResourceFromBytes(b)
  resources := []Resource{resource}

  // render table
  LogSuccess("Successfully unleased resource:")
  RenderResourceTable(resources)

  return nil
}

/**
 * invoked by the cli to lease a resource
 * (resalloc lease {{ serverName }})
 */
func LeaseResourceAction(c *cli.Context) error {

  // get cli argument
  requestedResource := c.Args().First()
  if requestedResource == "" {
    LogError("No resource specified; use: `resalloc lease {{ serverName }}``")
    return nil
  }

  // request acl to lease resource
  Log("Attempting to lease resource: " + requestedResource + " ...")
  resp, err := LeaseResource(requestedResource)

  // handle acl service unavailable
  if err != nil {
    return err
  }

  b := ResponseBodyToBytes(resp.Body)
  LogStatusCode(resp.StatusCode)

  // error case
  if ResponseIsError(resp) {
    err := AclErrorFromBytes(b)
    LogError(err.Message)
    Log("Try `resalloc list` to see leasable servers or `resalloc leased` to see servers leased to you")
    return nil
  }

  // success case
  resource := ResourceFromBytes(b)
  resources := []Resource{resource}

  // render table
  LogSuccess("Successfully leased resource:")
  RenderResourceTable(resources)

  return nil
}

/**
 * invoked by the cli to view resources leased to this user
 * (resalloc leased)
 */
func ViewLeasedResourcesAction(c *cli.Context) error {

  // request acl to view leased resources
  Log("Attempting to view resources leased to you...")
  resp, err := GetLeasedResource()

  // handle acl service unavailable
  if err != nil {
    return err
  }

  b := ResponseBodyToBytes(resp.Body)
  LogStatusCode(resp.StatusCode)

  // error case
  if ResponseIsError(resp) {
    err := AclErrorFromBytes(b)
    LogError(err.Message)
    return nil
  }

  // success case
  resources := ResourceArrayFromBytes(b)

  // render table
  numServersStr := ""
  if len(resources) == 0 {
    numServersStr = "are 0 servers"
  } else {
    // oversimplified, but we know the acl will only allow 1 server lease per user
    numServersStr = "is 1 server"
  }
  Log("There " + numServersStr + " leased to you:")
  RenderResourceTable(resources)

  return nil
}

/**
 * Invoked by the cli to view all fleet resources
 * (resalloc list)
 */
func ViewAllResourcesAction(c *cli.Context) error {

  // request acl to view leased resources
  Log("Attempting to view all resources...")
  resp, err := GetResources()

  // handle acl service unavailable
  if err != nil {
    return err
  }

  b := ResponseBodyToBytes(resp.Body)
  LogStatusCode(resp.StatusCode)

  // error case
  if ResponseIsError(resp) {
    err := AclErrorFromBytes(b)
    LogError(err.Message)
    return nil
  }

  // success case
  resources := ResourceArrayFromBytes(b)

  // render table
  Log("There are " +  strconv.Itoa(len(resources)) + " server resources in the fleet:")
  RenderResourceTable(resources)

  return nil
}
