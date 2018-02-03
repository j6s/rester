package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

type BackupOptions struct {
	// TODO Support all `restic backup` command line options
	Excludes []string
	Hostname string

	// Additional properties
	Sources    []string
	Repository string
	Password   string

	KeepLast    int
	KeepHourly  int
	KeepDaily   int
	KeepWeekly  int
	KeepMonthly int
	KeepYearly  int
}

func main() {
	options := readBackupOptionsFromFile(os.Args[1])

	// Ensure that the required properties are set
	// TODO surely there are better ways to do this
	if options.Repository == "" {
		panic("Key `repository` must be set!")
	}
	if options.Password == "" {
		panic("Key `password` must be set!")
	}
	if len(options.Sources) == 0 {
		panic("Key `sources` must be set!")
	}

	runBackupCommand(options)
	runForgetCommand(options)
}

// Runs the `restic forget` command in order to prune old
// backups according to the configured policy.
func runForgetCommand(options BackupOptions) {
	arguments := []string{"forget"}

	// TODO There must be a better way of mapping these arguments
	if options.KeepLast != 0 {
		arguments = append(arguments, "--keep-last", strconv.Itoa(options.KeepLast))
	}
	if options.KeepHourly != 0 {
		arguments = append(arguments, "--keep-hourly", strconv.Itoa(options.KeepHourly))
	}
	if options.KeepDaily != 0 {
		arguments = append(arguments, "--keep-daily", strconv.Itoa(options.KeepDaily))
	}
	if options.KeepWeekly != 0 {
		arguments = append(arguments, "--keep-weekly", strconv.Itoa(options.KeepWeekly))
	}
	if options.KeepMonthly != 0 {
		arguments = append(arguments, "--keep-monthly", strconv.Itoa(options.KeepMonthly))
	}
	if options.KeepYearly != 0 {
		arguments = append(arguments, "--keep-yearly", strconv.Itoa(options.KeepYearly))
	}

	executeResticCommand(arguments, options)
}

//
// Runs the `restic backup` command for the given options
//
func runBackupCommand(options BackupOptions) {
	arguments := []string{"backup", "--exclude-caches"}

	// --exclude
	for _, exclude := range options.Excludes {
		arguments = append(arguments, "--exclude", exclude)
	}

	// --hostname
	if options.Hostname != "" {
		arguments = append(arguments, "--hostname", options.Hostname)
	}

	// Add sources
	arguments = append(arguments, options.Sources...)

	// Execute
	executeResticCommand(arguments, options)
}

//
// Executes restic with the given list of arguments.
// Stdout and Stderr will be connected to Stderr and Stdout of
// the current application and the password will be passed via
// environment variable.
//
func executeResticCommand(arguments []string, options BackupOptions) {
	arguments = append(arguments, "--repo", options.Repository)
	fmt.Println("restic", arguments)

	command := exec.Command("restic", arguments...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Env = append(os.Environ(), "RESTIC_PASSWORD="+options.Password)
	err := command.Run()
	if err != nil {
		panic(err)
	}
}

//
// Reads the backup options from a json file and returns a struct
// containing the properties from that file.
//
func readBackupOptionsFromFile(jsonPath string) BackupOptions {
	jsonString, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	//var dat map[string]interface{}
	var options BackupOptions
	if err := json.Unmarshal(jsonString, &options); err != nil {
		panic(err)
	}
	return options
}
