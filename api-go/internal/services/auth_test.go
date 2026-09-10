package services

import (
	"strings"
	"testing"
)

func TestAuthServiceLoginWithValidCredentials(t *testing.T) {
	service := NewAuthService("admin", "admin123", "test-secret")

	token, err := service.Login("admin", "admin123")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if token == "" {
		t.Fatal("expected token, got empty string")
	}

	if !strings.Contains(token, ".") {
		t.Fatal("expected JWT token format")
	}
}

func TestAuthServiceLoginWithInvalidCredentials(t *testing.T) {
	service := NewAuthService("admin", "admin123", "test-secret")

	_, err := service.Login("admin", "wrong-password")

	if err != ErrInvalidCredentials {
		t.Fatalf("expected ErrInvalidCredentials, got %v", err)
	}
}

func TestAuthServiceValidateValidToken(t *testing.T) {
	service := NewAuthService("admin", "admin123", "test-secret")

	token, err := service.Login("admin", "admin123")
	if err != nil {
		t.Fatalf("expected no error generating token, got %v", err)
	}

	claims, err := service.ValidateToken(token)

	if err != nil {
		t.Fatalf("expected valid token, got %v", err)
	}

	if claims.Username != "admin" {
		t.Fatalf("expected username admin, got %s", claims.Username)
	}
}

func TestAuthServiceValidateInvalidToken(t *testing.T) {
	service := NewAuthService("admin", "admin123", "test-secret")

	_, err := service.ValidateToken("invalid-token")

	if err == nil {
		t.Fatal("expected error for invalid token")
	}
}

func TestAuthServiceValidateTokenWithWrongSecret(t *testing.T) {
	service := NewAuthService("admin", "admin123", "test-secret")
	otherService := NewAuthService("admin", "admin123", "different-secret")

	token, err := service.Login("admin", "admin123")
	if err != nil {
		t.Fatalf("expected no error generating token, got %v", err)
	}

	_, err = otherService.ValidateToken(token)

	if err == nil {
		t.Fatal("expected error when validating token with wrong secret")
	}
}