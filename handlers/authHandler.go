package handlers

import (
	"Pixelbloom-Backend/models"
	"Pixelbloom-Backend/utils"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// @summary		Sign Up
// @description	Sign up a new user
// @tags		auth
// @accept		json
// @produce		json
// @param		body	body	SignUpRequest	true	"User Sign Up"
// @success		200	{object}	map[string]interface{}	"User created successfully"
// @failure		400	{object}	map[string]interface{}	"Invalid request or Email already exists"
// @failure		500	{object}	map[string]interface{}	"Database query error, Failed to hash password, or Failed to create user"
// @router		/auth/signup [post]
func SignUpHandler(c *fiber.Ctx) error {
	ctx := context.Background()
	body := new(SignUpRequest)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	db := utils.Database()
	defer func(db *pgx.Conn, ctx context.Context) {
		err := db.Close(ctx)
		if err != nil {
			fmt.Printf("Failed to close database connection %v", err)
		}
	}(db, ctx)

	row, err := db.Query(ctx, `SELECT * FROM "Users" WHERE email = $1 LIMIT 1`, body.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database query error",
		})
	}
	defer row.Close()

	if row.Next() {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email already exists",
		})
	}

	hashedPassword, err := utils.HashPassword(body.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	args := pgx.NamedArgs{
		"email":    body.Email,
		"password": hashedPassword,
	}

	if _, err := db.Exec(ctx, `INSERT INTO "Users"(email,password) VALUES (@{email},@{password})`, args); err != nil {
		fmt.Println("Error creating user:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User created successfully"})
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// @summary		Sign In
// @description	Sign in an existing user
// @tags		auth
// @accept		json
// @produce		json
// @param		body	body	SignInRequest	true	"User Sign In"
// @success		200	{object}	map[string]interface{}	"User signed in successfully"
// @failure		400	{object}	map[string]interface{}	"Invalid request"
// @failure		401	{object}	map[string]interface{}	"Invalid User or Password"
// @failure		500	{object}	map[string]interface{}	"Failed to generate token"
// @router		/auth/signin [post]
func SignInHandler(c *fiber.Ctx) error {
	ctx := context.Background()
	body := new(SignInRequest)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	db := utils.Database()
	defer func(db *pgx.Conn, ctx context.Context) {
		err := db.Close(ctx)
		if err != nil {
			fmt.Printf("Failed to close database connection %v", err)
		}
	}(db, ctx)

	rows, err := db.Query(ctx, `SELECT * FROM "Users" WHERE email = $1 LIMIT 1`, body.Email)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User not found",
		})
	}
	defer rows.Close()

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.User])
	if err != nil || len(users) == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	if res := utils.CheckPasswordHash(body.Password, users[0].Password); !res {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid Password",
		})
	}

	token, err := utils.GenerateToken(users[0].Id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}
