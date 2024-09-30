# JWT OAuth SSO
JWT OAuth SSO is a Go-based authentication server template that provides a complete solution for handling authentication, authorization, and session management using JWT (JSON Web Tokens). It follows the OAuth2 and REST API architectural approaches, offering Single Sign-On (SSO) capabilities.

## Features:
* OAuth2-compliant access and refresh token flows
* JWT-based access tokens with a 15-minute expiration for secure API access
* Database-backed refresh tokens with IP, device, and user tracking, valid for 1 month
* SSO support for seamless user login across multiple applications
* Session management with the ability to revoke individual or all refresh tokens per device
* Customizable fields and validation for flexible authentication processes
* RESTful API for integration with both frontend and backend services
* Device-based session tracking to enhance security and user control

## Use Cases:
* Single Sign-On (SSO) systems for microservices architecture
* Secure API authentication and authorization
* User session management with custom refresh token handling

This template serves as a starting point for creating scalable and secure authentication services using Go, JWT, GORM, Sqlite3, Gin and OAuth2.