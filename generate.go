package taskmate

// go:generate swagger generate server --exclude-main -A task-mate-server -t gen -f ./api/swagger/swagger.yaml --principal models.Principal
//go:generate swagger -q generate client -A kurirmoo -f api/swagger/swagger.yaml -c pkg/client -m gen/models --principal models.Principal
