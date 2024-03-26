# Undefined Bank

## Project Description

This project is a JSON API built in Golang to create a bank API. It provides a set of endpoints to perform banking operations such as creating accounts, making transactions, and retrieving account information.

## Installation

To install and set up the project, follow these steps:

1. Clone the repository: `git clone https://github.com/M-Wellerson/undefined_bank`
2. Install the required dependencies: `make run`

## API Documentation

### Login

- Endpoint: `POST /login`
- Description: Authenticate user and return JWT
- Request format: JSON with `username` and `password` fields
- Response format: JSON with `token` field

### Create Account

- Endpoint: `POST /account`
- Description: Create a new account
- Request format: JSON with `name`, `cpf` and `secret` fields
- Response format: JSON with `account_id`, `name`, `cpf` and `balance` fields

### Get Account by ID

- Endpoint: `GET /account/{id}`
- Description: Retrieve account information by ID
- Request format: None
- Response format: JSON with `account_id`, `name`, `cpf` and `balance` fields

### Transfer Funds

- Endpoint: `POST /transfer`
- Description: Transfer funds between accounts
- Request format: JSON with `account_origin_id`, `account_destination_id` and `amount` fields
- Response format: JSON with `id`, `account_origin_id`, `account_destination_id`, `amount` and `created_at` fields