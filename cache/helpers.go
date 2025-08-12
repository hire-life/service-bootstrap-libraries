package cache

import "fmt"

func MakeUrl(host, port string) string {
	return fmt.Sprintf("redis://%s:%s", host, port)
}
