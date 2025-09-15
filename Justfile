set quiet

run:
  go run main.go

build url="https://webhook.site/85cae907-8b39-456f-99be-5947b10fcad8":
  go build -ldflags "-X main.URL={{url}}"