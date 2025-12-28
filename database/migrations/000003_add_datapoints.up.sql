PRAGMA foreign_keys = ON;
BEGIN TRANSACTION;
-- A user can have many datapoints
CREATE TABLE IF NOT EXISTS datapoints
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id    BLOB     NOT NULL,
    name       TEXT     NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id)
);
-- A datapoint can have many dataentries
CREATE TABLE IF NOT EXISTS dataentries
(
    id           INTEGER PRIMARY KEY AUTOINCREMENT,
    datapoint_id INTEGER             NOT NULL,
    type         TEXT                NOT NULL,
    text_value   TEXT,
    int_value    INTEGER,
                 created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (datapoint_id) REFERENCES datapoints (id)
);
COMMIT;
