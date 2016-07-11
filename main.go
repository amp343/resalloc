package main

import (
  "os"

  "github.com/urfave/cli"
)

/**
 * Define input handlers for the cli tool, and run the app
 */
func main() {
  app := cli.NewApp()
  app.Name = "resalloc"
  app.Usage = "lease some server resources"
  app.Version = "0.1.0"
  app.Commands = []cli.Command{
    {
      Name:    "lease",
      Usage:   "lease a resource (`lease {{serverName}}`)",
      Action: Do(LeaseResourceAction),
    },
    {
      Name:    "unlease",
      Usage:   "unlease a leased resource (`unlease {{serverName}}`)",
      Action:  Do(UnleaseResourceAction),
    },
    {
      Name:    "unlease",
      Usage:   "unlease a leased resource (`unlease {{serverName}}`)",
      Action:  Do(UnleaseResourceAction),
    },
    {
      Name:    "leased",
      Usage:   "see a list of resources leased by you",
      Action:  Do(ViewLeasedResourcesAction),
    },
    {
      Name:    "list",
      Usage:   "see a list of all resources",
      Action:  Do(ViewAllResourcesAction),
    },
  }

  app.Run(os.Args)
}
