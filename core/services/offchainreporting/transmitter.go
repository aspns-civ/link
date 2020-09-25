package offchainreporting

import (
	"context"
	"database/sql"

	gethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

type transmitter struct {
	db          *sql.DB
	ctx         context.Context
	fromAddress gethCommon.Address
	gasLimit    uint64
}

// NewTransmitter creates a new eth transmitter
func NewTransmitter(ctx context.Context, sqldb *sql.DB, fromAddress gethCommon.Address, gasLimit uint64) Transmitter {
	return &transmitter{
		db:          sqldb,
		ctx:         ctx,
		fromAddress: fromAddress,
		gasLimit:    gasLimit,
	}
}

func (t *transmitter) CreateEthTransaction(toAddress gethCommon.Address, payload []byte) error {
	_, err := t.db.ExecContext(t.ctx, `
INSERT INTO eth_txes (from_address, to_address, encoded_payload, value, gas_limit, state, created_at)
VALUES ($1,$2,$3,$4,$5,'unstarted',NOW())
`, t.fromAddress, toAddress, payload, 0, t.gasLimit)

	return errors.Wrap(err, "failed to create eth_tx")
}
