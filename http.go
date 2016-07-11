package main

import (
  "io"
  "io/ioutil"
  "net/http"
)

/**
 * Build the request url based on the acl config from
 * config.yml, and the relative path
 */
func BuildRequestUrl(aclConfig AclConfig, path string) string {
  return aclConfig.Host + ":" + aclConfig.Port + "/" + path
}

/**
 * Consult the acl to get any resources leased to this user
 */
func GetLeasedResource() (*http.Response, error) {
  return RequestAcl("GET", "resources/leased")
}

/**
 * Consult the acl to get all fleet resources
 */
func GetResources() (*http.Response, error) {
  return RequestAcl("GET", "resources")
}

/**
 * Consult the acl to lease a fleet resource
 */
func LeaseResource(resourceName string) (*http.Response, error) {
  return RequestAcl("POST", "resources/" + resourceName + "/lease")
}

/**
 * Consult the acl to unlease a fleet resource
 */
func UnleaseResource(resourceName string) (*http.Response, error) {
  return RequestAcl("DELETE", "resources/" + resourceName + "/lease")
}

/**
 * Get a byte array from an http response body
 */
func ResponseBodyToBytes(body io.Reader) []byte {
  b, err := ioutil.ReadAll(body)
  if err != nil {
    panic(err)
  }

  return b
}

/**
 * Determine whether the api response constitutes an error
 */
func ResponseIsError(resp *http.Response) bool {
  if resp.StatusCode > 200 {
    return true
  } else {
    return false
  }
}

/**
 * Make an http request to the acl api and return an http response
 */
func RequestAcl(httpMethod string, path string) (*http.Response, error) {
  aclConfig := GetAclConfig()
  auth := GetAuth()

  url := BuildRequestUrl(aclConfig, path)
  req, err := http.NewRequest(httpMethod, url, nil)
  req.SetBasicAuth(auth.Username, auth.Password)

  client := &http.Client{}

  resp, err := client.Do(req)

  if err != nil {
    LogAclServiceUnavailable()
    return nil, err
  }

  return resp, err
}
