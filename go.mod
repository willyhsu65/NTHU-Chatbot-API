module nthu-chatbot-api

go 1.13

require (
	github.com/DataDog/zstd v1.4.4 // indirect
	github.com/gin-contrib/sessions v0.0.3
	github.com/gin-gonic/gin v1.5.0
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/google/go-cmp v0.3.1 // indirect
	github.com/pkg/errors v0.8.1 // indirect
	github.com/robfig/cron v1.2.0
	github.com/tidwall/pretty v1.0.0 // indirect
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c // indirect
	github.com/xdg/stringprep v1.0.0 // indirect
	go.mongodb.org/mongo-driver v1.2.0
	golang.org/x/crypto v0.0.0-20191206172530-e9b2fee46413 // indirect
	golang.org/x/net v0.0.0-20190503192946-f4e77d36d62c
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e // indirect
	google.golang.org/api v0.17.0
)

replace (
	nthu-chatbot-api/database => /Users/jensonsu/MyProject/nthu-chatbot-api/database
	nthu-chatbot-api/models => /Users/jensonsu/MyProject/nthu-chatbot-api/models
	nthu-chatbot-api/pkg => /Users/jensonsu/MyProject/nthu-chatbot-api/pkg
	nthu-chatbot-api/router => /Users/jensonsu/MyProject/nthu-chatbot-api/router
	nthu-chatbot-api/utils => /Users/jensonsu/MyProject/nthu-chatbot-api/utils
	nthu-chatbot-api/vendors => /Users/jensonsu/MyProject/nthu-chatbot-api/vendors
)
