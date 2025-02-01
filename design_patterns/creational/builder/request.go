package main

import (
	"encoding/json"
	"fmt"
)

// Create a struct representing the product to be built
type Response struct {
	Status  string            `json:"status"`
	Message string            `json:"message"`
	Data    map[string]string `json:"data"`
}

// Create a Builder that holds an instance of the product class
type ResponseBuilder struct {
	response Response
}

// Implement constructor for the builder
func NewResponseBuilder() *ResponseBuilder {
	return &ResponseBuilder{response: Response{
		Data: make(map[string]string),
	}}
}

// Add methods to set each field
func (b *ResponseBuilder) SetStatus(status string) *ResponseBuilder {
	b.response.Status = status
	return b
}

func (b *ResponseBuilder) SetMessage(message string) *ResponseBuilder {
	b.response.Message = message
	return b
}
func (b *ResponseBuilder) AddData(key, value string) *ResponseBuilder {
	b.response.Data[key] = value
	return b
}

func (b *ResponseBuilder) Build() string {
	jsonData, _ := json.Marshal(b.response)
	return string(jsonData)
}

func main() {
	response := NewResponseBuilder().
		SetStatus("success").
		SetMessage("User created").
		AddData("userID", "123").
		Build()

	fmt.Println(response)
}
