package issuer

import (
	"context"

	"github.com/desmos-labs/juno/client"

	"github.com/forbole/bdjuno/database"
	"github.com/forbole/bdjuno/modules/utils"

	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"

	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
)

// RegisterPeriodicOps returns the AdditionalOperation that periodically runs fetches from
// the LCD to make sure that constantly changing data are synced properly.
func RegisterPeriodicOps(scheduler *gocron.Scheduler, minClient minttypes.QueryClient, db *database.Db) error {
	log.Debug().Str("module", "issuer").Msg("setting up periodic tasks")

	// Setup a cron job to run every 6 hours
	if _, err := scheduler.Every(6).Hours().At("00:00").StartImmediately().Do(func() {
		utils.WatchMethod(func() error { return updateInflation(minClient, db) })
	}); err != nil {
		return err
	}

	return nil
}

// updateInflation fetches from the REST APIs the latest value for the
// inflation, and saves it inside the database.
func updateInflation(mintClient minttypes.QueryClient, db *database.Db) error {
	log.Debug().
		Str("module", "issuer").
		Str("operation", "inflation").
		Msg("getting inflation data")

	height, err := db.GetLastBlockHeight()
	if err != nil {
		return err
	}

	// Get the inflation
	res, err := mintClient.Inflation(
		context.Background(),
		&minttypes.QueryInflationRequest{},
		client.GetHeightRequestHeader(height),
	)
	if err != nil {
		return err
	}

	return db.SaveInflation(res.Inflation, height)
}
