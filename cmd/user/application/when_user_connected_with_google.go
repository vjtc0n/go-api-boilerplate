package application

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"

	"github.com/vardius/go-api-boilerplate/cmd/user/domain/user"
	"github.com/vardius/go-api-boilerplate/cmd/user/infrastructure/persistence"
	"github.com/vardius/go-api-boilerplate/pkg/domain"
	"github.com/vardius/go-api-boilerplate/pkg/eventbus"
)

// WhenUserConnectedWithGoogle handles event
func WhenUserConnectedWithGoogle(db *sql.DB, repository persistence.UserRepository) eventbus.EventHandler {
	fn := func(ctx context.Context, event domain.Event) {
		// this goroutine runs independently to request's goroutine,
		// there for recover middlewears will not recover
		// recover from panic to prevent crash
		defer func() {
			if r := recover(); r != nil {
				log.Printf("[EventHandler] Recovered in %v", r)
			}
		}()

		log.Printf("[EventHandler] %s", event.Payload)

		e := &user.ConnectedWithGoogle{}

		err := json.Unmarshal(event.Payload, e)
		if err != nil {
			log.Printf("[EventHandler] Error: %v", err)
			return
		}

		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			log.Printf("[EventHandler] Error: %v", err)
			return
		}
		defer tx.Rollback()

		err = repository.UpdateGoogleID(ctx, e.ID.String(), e.GoogleID)
		if err != nil {
			log.Printf("[EventHandler] Error: %v", err)
			return
		}
		tx.Commit()
	}

	return eventbus.EventHandler(fn)
}
