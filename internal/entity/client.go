package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewClient(name string, email string) (*Client, error) {
	client := &Client{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := client.Validate()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (client *Client) Validate() error {
	if client.Name == "" {
		return errors.New("name is required")
	}
	if client.Email == "" {
		return errors.New("email is required")
	}
	return nil
}

func (client *Client) Update(name string, email string) error {
	client.Name = name
	client.Email = email
	client.UpdatedAt = time.Now()
	err := client.Validate()
	if err != nil {
		return err
	}
	return nil
}
