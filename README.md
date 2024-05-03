# MyMandap Task

## How to get it running?

1. Clone the repository
2. Navigate to the root directory of the project
3. Run the following command to build the executable
```bash
make setup
```
4. Run the executable
```bash
./main
```

## Code Structure

```
📦MyMandap
 ┣ 📂helpers
 ┃ ┣ 📜fact.go
 ┃ ┣ 📜log.go
 ┃ ┗ 📜message.go
 ┣ 📂models
 ┃ ┗ 📜user.go
 ┣ 📂utils
 ┃ ┗ 📜input.go
 ┣ 📜.gitignore
 ┣ 📜Makefile
 ┣ 📜README.md
 ┣ 📜go.mod
 ┣ 📜main
 ┗ 📜main.go
```

Functionalities:
1. Send a message to an user
![Send a message to an user](static/user.png)
1.1 without a message
![without a message](static/fact.png)
2. Send a message to all users
![Send a message to all users](static/broadcast.png)
3. View message log of an user
![View message log of an user](static/view.png)