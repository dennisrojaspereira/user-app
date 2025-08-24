package domain

type User struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

type CreateUserInput struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}
