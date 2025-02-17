// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: pool.sql

package sqlc

import (
	"context"
)

const prepareForSend = `-- name: PrepareForSend :many
UPDATE sending_pool_emails AS sp
    SET status = 'sending'
    FROM (
            SELECT id FROM sending_pool_emails
            WHERE scheduled_time <= NOW() AND status = 'scheduled'
            LIMIT $1
        ) AS t
    WHERE sp.id = t.id
    RETURNING sp.id, sp.status, sp.scheduled_time, sp.original_scheduled_time, sp.send_attempts_cnt, sp.email, sp.message_id, sp.error_msg, sp.error_code
`

func (q *Queries) PrepareForSend(ctx context.Context, limit int32) ([]SendingPoolEmail, error) {
	rows, err := q.query(ctx, q.prepareForSendStmt, prepareForSend, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SendingPoolEmail
	for rows.Next() {
		var i SendingPoolEmail
		if err := rows.Scan(
			&i.ID,
			&i.Status,
			&i.ScheduledTime,
			&i.OriginalScheduledTime,
			&i.SendAttemptsCnt,
			&i.Email,
			&i.MessageID,
			&i.ErrorMsg,
			&i.ErrorCode,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const setSendingPoolDelivered = `-- name: SetSendingPoolDelivered :exec
UPDATE sending_pool_emails 
	SET status = 'sent' WHERE email = $1 AND message_id = $2
`

type SetSendingPoolDeliveredParams struct {
	Email     string
	MessageID string
}

func (q *Queries) SetSendingPoolDelivered(ctx context.Context, arg SetSendingPoolDeliveredParams) error {
	_, err := q.exec(ctx, q.setSendingPoolDeliveredStmt, setSendingPoolDelivered, arg.Email, arg.MessageID)
	return err
}

const setSendingPoolError = `-- name: SetSendingPoolError :exec
UPDATE sending_pool_emails 
	SET status = 'error', error_msg = $1 
    WHERE email = $2 AND message_id = $3
`

type SetSendingPoolErrorParams struct {
	ErrorMsg  string
	Email     string
	MessageID string
}

func (q *Queries) SetSendingPoolError(ctx context.Context, arg SetSendingPoolErrorParams) error {
	_, err := q.exec(ctx, q.setSendingPoolErrorStmt, setSendingPoolError, arg.ErrorMsg, arg.Email, arg.MessageID)
	return err
}
