This is a Golang package for user service (one of the kenopsia microservices)

It exposes methods which are common when adding new game to the ecosystem.

# Install

```bash
go get github.com/amirrezam75/kenopsiauser
```

# Usages

```go
package main

import "github.com/amirrezam75/kenopsiauser"

var userRepository = kenopsiauser.NewUserRepository(
	os.Getenv("KENOPSIA_USER_BASE_URL"),
)
var gameService = services.NewGameService(gameRepository, hubRepository, userRepository)
```

### Acquire User Id

Since headers cannot be set in a WebSocket connection, we use a one-time token instead. The client must provide this
token before establishing the WebSocket connection, and your game service will need a method to identify which user the
token belongs to.

https://devcenter.heroku.com/articles/websocket-security

```go
package services

type GameService struct {
	gameRepository GameRepository
	hubRepository  HubRepository
	userRepository kenopsiauser.UserRepository
}

func (game GameService) Join(gameId, ticketId string, connection *websocket.Conn) {
	userId, err := game.userRepository.AcquireUserId(ticketId)
	if err != nil {
		// Handler error
	}
}
```