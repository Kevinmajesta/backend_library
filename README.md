# DEPUBLIC

## Library Transaction 📚

---

Library Transaction API adalah RESTful backend service untuk mengelola sistem peminjaman perpustakaan, dibangun menggunakan Golang (Echo Framework) dan PostgreSQL. API ini mendukung proses peminjaman dan pengembalian buku dengan transactional logic yang bersifat atomic, memastikan validasi stok buku dan kuota peminjaman member berjalan aman dalam satu database transaction untuk mencegah race condition. Seluruh aplikasi dijalankan menggunakan Docker dan docker-compose, sehingga mudah dijalankan secara konsisten di berbagai environment.

---

## Installation 👨🏻‍💻

Depublic build by [Go](https://go.dev/dl/) Go 1.13+ to run.

1. Clone Repository
   By use terminal/cmd

```sh
git clone https://github.com/Kevinmajesta/backend_library.git
```

2. Open Repository
   By use terminal/cmd

```sh
cd backendlibrary
```

2. Check the .env file and configure with your device

3. Enable the PostgreSQL database
   Option you can use :

   - [pgAdmin](https://www.pgadmin.org/)
   - [NaviCat](https://www.navicat.com/en/download/navicat-premium?gad_source=1&gclid=CjwKCAjwmYCzBhA6EiwAxFwfgFWv6YNc_nwrdL5BByjvaEmUNbzD0vvg-tHgv7x6rFyIx-zSdWYQWhoCRP0QAvD_BwE)
   - Or anything else you usualy use

4. Run the command to create the database and migrate it.
   Make sure you have install migrate cli.
   If you dont, install first by

**If you MAC user** 🍏

- First if you dont have [Home Brew](https://brew.sh/)
  Open terminal and copy code below :

```sh
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

Then install migrate cli

```sh
brew install golang-migrate
```

**If you windows user**🪟

- Open PowerShell. You can do this by searching for PowerShell in the start menu.
- Inside PowerShell, type the code below

```sh
iwr -useb get.scoop.sh | iex
```

Then Install use scoop

```sh
scoop install migrate
```

After all, migrate it by

```sh
migrate -path database/migration/ -database "migrate -path migrations -database "postgres://library_c:library_c@localhost:5432/library_c?sslmode=disable" up" -verbose up


```

5. Configure Docker
   **First Install Docker**
   - Windows User[Docker](https://docs.docker.com/desktop/install/windows-install/)
   - Mac User [Docker](https://docs.docker.com/desktop/install/mac-install/)
   - Then compose it by

```sh
docker-compose up -d
```

6. Run the application

```sh
go mod tidy
go run cmd/app/main.go
```

---

## ✨ Features

- **User Management**
  - Member data management
  - Borrowing quota enforcement (max active borrow limit)

- **Book Management**
  - Create and retrieve books
  - Stock management with concurrency safety

- **Borrowing Transaction**
  - Borrow book with atomic database transaction
  - Stock validation with locking
  - Member quota validation based on active borrow records
  - Return book and restore stock

---

## Tech 💻

Book Management uses a number of open source projects to work properly :

1. **Golang** - High-performance language for scalable apps.
2. **Echo Framework** - Web framework for Go.
3. **PostgreSQL** - Reliable open-source database.
5. **Docker** - Container platform for consistent deployment.

## 🛠️ Tech Stack

- **Language**: Go (Golang)
- **Framework**: Echo
- **Database**: PostgreSQL
- **Communication**: RESTful API
- **Containerization**: Docker


## Development

This project app develope by 1 people
| Name | Github |
| ------ | ------ |
| Kevin | https://github.com/Kevinmajesta |


By using github for development for staging and production.


## API Documentation 🔗

Documentation for API can be get by :

```sh
link web : 
https://www.postman.com/lunar-resonance-148572/workspace/kevin-work/collection/33423852-6f3c1929-6965-46e1-b080-8fca327d092a?action=share&creator=33423852
for the the api spec :)



