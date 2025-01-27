package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// Config holds all the parameters for the utility
type Config struct {
	URL     string
	Port    int
	Scheme  string
	Token   string
	CSV     bool
	File    string
	Debug   bool
	Config  string
	Version bool // Flag to indicate if the version should be printed
}

const (
	defaultOutputFilename = "channel_stats"
)

// parseConfig processes command-line flags, environment variables, and the config file
func parseConfig() *Config {
	DebugPrint("Parsing command line...")
	// Define command-line flags
	config := &Config{}
	flag.StringVar(&config.Config, "config", "./config.json", "Alternative config file (JSON)")
	flag.StringVar(&config.URL, "url", "", "Mattermost instance URL")
	flag.IntVar(&config.Port, "port", 443, "Mattermost port")
	flag.StringVar(&config.Scheme, "scheme", "https", "HTTP scheme (http/https)")
	flag.StringVar(&config.Token, "token", "", "API token")
	flag.BoolVar(&config.CSV, "csv", false, "Create CSV output")
	flag.StringVar(&config.File, "file", "", "Optional output filename")
	flag.BoolVar(&config.Debug, "debug", false, "Run in DEBUG mode")
	flag.BoolVar(&config.Version, "version", false, "Display version and exit")

	// Customise the usage/help output
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Utility Name: mm-channel-stats
Description: Collects channel statistics from a Mattermost instance and outputs them in JSON or CSV format.

Usage:
`)
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, `
Configuration Sources:
  - Command-line flags (highest priority)
  - Environment variables (e.g., MM_URL, MM_TOKEN)
  - Config file (default: ./config.json, can be overridden with -config)

Examples:
  Collect stats and output as JSON:
    ./mm-channel-stats -url https://mattermost.example.com -token YOUR_API_TOKEN

  Collect stats and output as CSV:
    ./mm-channel-stats -url https://mattermost.example.com -token YOUR_API_TOKEN -csv
`)
	}

	// Parse command-line flags
	flag.Parse()

	// Handle version flag consistently
	if config.Version {
		fmt.Printf("mm-channel-stats version: %s\n", Version)
		os.Exit(0)
	}

	// Set up Viper
	viper.SetDefault("port", 443)
	viper.SetDefault("scheme", "https")
	viper.SetDefault("csv", false)
	viper.SetDefault("debug", false)
	viper.SetDefault("config", "./config.json")

	// Bind environment variables
	viper.BindEnv("url", "MM_URL")
	viper.BindEnv("port", "MM_PORT")
	viper.BindEnv("scheme", "MM_SCHEME")
	viper.BindEnv("token", "MM_TOKEN")
	viper.BindEnv("debug", "MM_DEBUG")
	viper.BindEnv("config", "MM_CONFIG")

	// Load configuration file
	viper.SetConfigFile(config.Config)
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			confMsg := fmt.Sprintf("Config file not found, using defaults and other sources: %s\n", config.Config)
			LogMessage(infoLevel, confMsg)
		} else {
			log.Fatalf("Error reading config file: %v\n", err)
		}
	} else {
		confFileUsedMsg := fmt.Sprintf("Using config file: %s\n", viper.ConfigFileUsed())
		DebugPrint(confFileUsedMsg)
	}

	// Resolve values from environment variables and config file, only if not set via flags
	if config.URL == "" {
		config.URL = viper.GetString("url")
	}
	if config.Port == 443 {
		config.Port = viper.GetInt("port")
	}
	if config.Scheme == "https" {
		config.Scheme = viper.GetString("scheme")
	}
	if config.Token == "" {
		config.Token = viper.GetString("token")
	}
	if !config.CSV {
		config.CSV = viper.GetBool("csv")
	}
	if config.File == "" {
		config.File = viper.GetString("file")
	}
	if !config.Debug {
		config.Debug = viper.GetBool("debug")
	}

	// Ensure required parameters are provided
	if config.URL == "" || config.Token == "" {
		fmt.Println("Error: -url and -token are required parameters.")
		flag.Usage()
		os.Exit(1)
	}
	config.Scheme = strings.ToLower(config.Scheme)
	if config.Scheme != "http" && config.Scheme != "https" {
		fmt.Printf("Error: -scheme must be either 'http' or 'https'")
		flag.Usage()
		os.Exit(2)
	}

	if config.File == "" {
		if config.CSV {
			config.File = defaultOutputFilename + ".csv"
		} else {
			config.File = defaultOutputFilename + ".json"
		}
	}

	// We have a default for debugMode hard-coded into the source code, which allows us to use it
	// before it gets set during this function.  If it's 'true' in the code, we'll effectively
	// ignore anything on the command-line.
	if config.Debug {
		debugMode = true
	}

	return config
}
