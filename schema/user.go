package schema

type UserCreateRequest struct {
    Name string `json:"name"`
}

type UserUpdateRequest struct {
    Name string `json:"name"`
}
