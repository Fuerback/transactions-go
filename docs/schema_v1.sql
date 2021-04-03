CREATE TABLE account (
    id              INTEGER PRIMARY KEY AUTOINCREMENT,
    document_number TEXT    NOT NULL
)
WITHOUT ROWID;

CREATE TABLE [transaction] (
    id             INTEGER  PRIMARY KEY AUTOINCREMENT
                            NOT NULL,
    account_id     INT      NOT NULL
                            REFERENCES account (id),
    amount         DECIMAL  NOT NULL,
    event_date     DATETIME NOT NULL,
    operation_type INT      NOT NULL
);