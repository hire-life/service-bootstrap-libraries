package database

import "fmt"

func MakeUrl(host, port, name, user, password string) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, name)
}
