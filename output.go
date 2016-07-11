package main

import (
  "fmt"
  "strconv"
  "os"

  "github.com/fatih/color"
  "github.com/olekukonko/tablewriter"
)

/**
 * Log a newline
 */
func LogN() {
  fmt.Printf("\n")
}

/**
 * Log a string with the Resalloc: prefix
 */
func Log(str string) {
  fmt.Printf(color.CyanString("Resalloc: ") + str + "\n")
}

/**
 * Log an error string
 */
func LogError(str string) {
  Log(color.RedString("Error: " + str))
}

/**
 * Log a success string
 */
func LogSuccess(str string) {
  Log(color.GreenString(str))
}

/**
 * Log a consistent message for when the acl server is unreachable
 */
func LogAclServiceUnavailable() {
  LogError("ACL service is offline or otherwise unavailable")
}

/**
 * Log an http status code, with appropriate color formatting
 */
func LogStatusCode(code int) {
  if code < 300 {
    Log(color.GreenString("Success: " + strconv.Itoa(code)))
  } else {
    Log(color.RedString("Error: " + strconv.Itoa(code)))
  }
}

/**
 * Render a display table for an array of resources
 */
func RenderResourceTable(resources []Resource) {
  table := tablewriter.NewWriter(os.Stdout)
  table.SetHeader([]string{"Server Name", "OS", "Available", "Leased By", "Time Remaining"})
  for _, r := range resources {
    table.Append(r.ToTableRow())
  }
  table.Render()
}
