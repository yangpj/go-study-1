package main

// 注意原先的redis go 客户端godis 好像变更了目录结构了！

import (
	"github.com/kavu/go-resque"             // Import this package
	godis "github.com/simonz05/godis/redis" // For now go-resque support only godis package
)

func main() {
	client := godis.New("tcp:127.0.0.1:6379", 0, "") // Create new Redis client to use for enqueuing

	// Enqueue the job into the "go" queue with appropriate client
	// resque.Enqueue(client, "go", "Demo::Job")
	resque.Enqueue(client, "hello", "Hello")

	// Enqueue into the "default" queue with passing one parameter to the Demo::Job.perform
	// resque.Enqueue(client, "default", "Demo::Job", 1)

	// Enqueue into the "default" queue with passing multiple
	// parameters to the Demo::Job.perform so it will fail
	// resque.Enqueue(client, "default", "Demo::Job", 1, 2, "woot")

}
