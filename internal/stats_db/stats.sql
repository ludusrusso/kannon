-- name: InsertPrepared :exec
INSERT INTO prepared (email, message_id, timestamp, first_timestamp, domain) VALUES (@email, @message_id, @timestamp, @timestamp, @domain)
	ON CONFLICT (email, message_id, domain) DO UPDATE
	SET timestamp = @timestamp;

-- name: InsertAccepted :exec
INSERT INTO accepted (email, message_id, timestamp, domain) VALUES (@email, @message_id, @timestamp, @domain);

-- name: InsertHardBounced :exec
INSERT INTO hard_bounced (email, message_id, timestamp, domain, err_code, err_msg) VALUES  (@email, @message_id, @timestamp, @domain, @err_code, @err_msg);

-- name: InsertOpen :exec
INSERT INTO open (email, message_id, timestamp, domain, ip, user_agent) VALUES  (@email, @message_id, @timestamp, @domain, @ip, @user_agent);

-- name: InsertClick :exec
INSERT INTO click (email, message_id, timestamp, domain, ip, user_agent, url) VALUES  (@email, @message_id, @timestamp, @domain, @ip, @user_agent, @url);