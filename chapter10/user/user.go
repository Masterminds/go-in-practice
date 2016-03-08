//go:generate codecgen -o user_generated.go user.go

package user

type User struct {
	Name  string `codec:"name"`
	Email string `codec:",omitempty"`
}
