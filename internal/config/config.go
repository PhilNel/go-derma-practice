package config

import (
	"fmt"

	flags "github.com/jessevdk/go-flags"
)

type Config struct {
	Log      *Log      `group:"Log Options"`
	DynamoDB *DynamoDB `group:"DynamoDB Options"`
}

type Log struct {
	Level string `long:"log_level" env:"LOG_LEVEL" description:"Log level" default:"info"`
}

type DynamoDB struct {
	SpecialsTable string `long:"specials_table" env:"SPECIALS_TABLE_NAME" description:"DynamoDB table for specials"`
	Region        string `long:"aws_region_dynamodb" env:"AWS_REGION" description:"AWS region for DynamoDB" default:"eu-west-1"`
}

func Load() (*Config, error) {
	opts := &Config{}
	_, err := flags.Parse(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	return opts, nil
}
