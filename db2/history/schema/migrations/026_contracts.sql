-- +migrate Up

CREATE TABLE history_contracts
(
  contract_id     BIGINT      NOT NULL CHECK (contract_id >= 0),
  contractor      VARCHAR(56) NOT NULL,
  customer        VARCHAR(56) NOT NULL,
  escrow          VARCHAR(56) NOT NULL,
  disputer        VARCHAR(56) DEFAULT NULL,
  start_time      TIMESTAMP without time zone NOT NULL,
  end_time        TIMESTAMP without time zone NOT NULL,
  details         jsonb       NOT NULL,
  invoices        BIGINT[]    DEFAULT NULL,
  dispute_reason  jsonb       DEFAULT NULL,
  state           INT         NOT NULL,
  PRIMARY KEY (contract_id)
);

-- +migrate Down

DROP TABLE IF EXISTS history_contracts;