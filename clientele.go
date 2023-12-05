package main

import (
	"strings"
)

type Client struct {
	Name       string
	Email      string
	Language   string
	ExternalID string
}

func CreateClientsSlice(data [][]string) []*Client {
	var clients []*Client
	for _, row := range data {
		clients = append(clients, &Client{
			Name:       strings.Trim(row[0], " "),
			Email:      strings.Trim(row[1], " "),
			Language:   strings.Trim(row[2], " "),
			ExternalID: strings.Trim(row[3], " "),
		})
	}
	return clients
}
