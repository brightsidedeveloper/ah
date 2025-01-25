# Usage

```js
get('/api/users', undefined, (b) => Users.decode(b))

post(
  '/api/user',
  User.create({
    name: 'test2',
  }),
  (u) => User.encode(u).finish(),
  (b) => User.decode(b)
)

// etc...
```

```go

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {

	qUsers, err := h.query.ListUsers(context.Background())
	if err != nil {
		fmt.Println(err)
		h.bin.WriteError(w, http.StatusInternalServerError, "Failed to get users")
		return
	}

	users := make([]*buf.User, 0, len(qUsers))
	for _, u := range qUsers {
		users = append(users, &buf.User{
			Id:   u.ID,
			Name: u.Name,
		})
	}

	h.bin.ProtoWrite(w, http.StatusOK, &buf.Users{
		Users: users,
	})
}

func (h *Handler) PostUser(w http.ResponseWriter, r *http.Request) {
	var user buf.User
	if err := h.bin.UnmarshalBody(r.Body, &user); err != nil {
		h.bin.WriteError(w, http.StatusBadRequest, "Failed to decode body")
		return
	}

	if user.Name == "" {
		h.bin.WriteError(w, http.StatusBadRequest, "Name is required")
		return
	}

	qUser, err := h.query.CreateUser(r.Context(), user.Name)
	if err != nil {
		h.bin.WriteError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	h.bin.ProtoWrite(w, http.StatusCreated, &buf.User{
		Id:   qUser.ID,
		Name: qUser.Name,
	})
}
```
