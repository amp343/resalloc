package main

import (
  "io/ioutil"
  "path/filepath"

  "gopkg.in/yaml.v2"
)

type Config struct {
  Auth Auth
  Acl AclConfig
}

type AclConfig struct {
  Host string
  Port string
}

type Auth struct {
  Username string
  Password string
}

func GetAuth() Auth {
  return GetConfig().Auth
}

func GetAclConfig() AclConfig {
  return GetConfig().Acl
}

/**
 * From some yml config file, return a Config struct
 *
 * (for now this is just hardcoded
 * to read from the relative file ./config.yml
 * which holds user credentials
 * and config values for talking to the acl)
 */
func GetConfig() Config {
  filename, _ := filepath.Abs("./config.yml")
  yamlFile, err := ioutil.ReadFile(filename)

  var config Config

  err = yaml.Unmarshal(yamlFile, &config)
  if err != nil {
      panic(err)
  }

  return config
}
