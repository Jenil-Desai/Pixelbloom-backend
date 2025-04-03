# 🎨 Pixelbloom Backend

**Pixelbloom Backend** is the server-side codebase for the **Pixelbloom mobile application**, handling user authentication, wallpaper management, and user interactions like likes and bookmarks. Built with **Go** and **PostgreSQL**, it ensures high performance and scalability.

---

## 📑 Table of Contents

1. [Overview](#-overview)
2. [Technologies](#-technologies)
3. [Packages & Libraries Used](#-packages--libraries-used)
4. [Getting Started](#-getting-started)
5. [Setup](#-setup)
6. [Features](#-features)
7. [Demo & Screenshots](#-demo--screenshots)
8. [Acknowledgments](#-acknowledgments)
9. [License](#-license)

---

## 🌟 Overview

Pixelbloom Backend provides RESTful APIs for **user authentication, wallpaper storage, and user interactions** such as likes and bookmarks. It supports JWT-based authentication and follows best practices for API security and performance.

---

## 💻 Technologies

| Technology  | Description  |  
|------------|-------------|  
| **Go**     | Backend logic and API development |  
| **PostgreSQL** | Database management |  

---

## 📦 Packages / Libraries Used

| Package / Library | Purpose |  
|------------------|---------|  
| **Fiber**        | Web framework for Go |  
| **PGX**          | PostgreSQL driver for Go |  
| **Swagger**      | API documentation |  
| **JWT**          | JSON Web Token authentication |  
| **Crypto**       | Secure hashing and encryption |  
| **godotenv**     | Environment variable management |  

---

## 🚀 Getting Started

1. Install Go on your machine.
2. Get your PostgreSQL database set up.
3. Get an IDE or text editor of your choice.

---

## ⚙️ Setup

1. Clone the GitHub repository:
    ```shell
   git clone https://github.com/Jenil-Desai/Pixelbloom-backend.git
   ```
2. Navigate to the project directory:
    ```shell
   cd Pixelbloom-backend
   ```
3. Install dependencies:
    ```shell
   go mod tidy
   ```
4. Create a `.env` file in the root directory and set up your environment variables:
    ```env
    DATABASE_URL=""
   JWT_SECRET=""
   ```
5. Run the application:
    ```shell
   air -c .air.linux.conf
   ```
---

## 🎯 Features

✔️ **User Details Endpoints** – Manage user profiles and authentication  
✔️ **Wallpapers Endpoints** – Upload, update, and retrieve wallpapers  
✔️ **Liked Wallpapers Endpoints** – Track user interactions with wallpapers  
✔️ **Bookmarked Wallpapers Endpoints** – Allow users to save their favorite wallpapers

---

## 🔗 Demo & Screenshots

- You can test the API using [Hoppscotch](https://hoppscotch.io/) or any other API testing tool.

---

## 🙏 Acknowledgments

1. [Fiber Docs](https://docs.gofiber.io/)
2. [Medium Blog on JWT](https://medium.com/code-beyond/go-fiber-jwt-auth-eab51a7e2129)

---

## 📜 License

This project is licensed under the [MIT License](LICENSE). See the [LICENSE](LICENSE) file for details.

---

### 🎨 **Enhancing the Pixelbloom experience with a powerful backend!**  
