package st

const CreateTables = `

CREATE TABLE IF NOT EXISTS file
(
  id            text    PRIMARY KEY,
  name          text    NOT NULL,
  size          int     NOT NULL,
  type          text    NOT NULL,
  thumb         int     NOT NULL,
  hash          text    NOT NULL UNIQUE,
  like          int     NOT NULL,
  ctime         int     NOT NULL,
  utime         int     NOT NULL,
  deleted       int     NOT NULL,
);

CREATE INDEX IF NOT EXISTS idx_file_hash ON file(hash);
CREATE INDEX IF NOT EXISTS idx_file_ctime ON file(ctime);
CREATE INDEX IF NOT EXISTS idx_file_utime ON file(utime);

CREATE TABLE IF NOT EXISTS tag
(
  id            text    PRIMARY KEY,
  name          text    NOT NULL UNIQUE,
  ctime         int     NOT NULL,
);

CREATE INDEX IF NOT EXISTS idx_tag_ctime ON tag(ctime);

CREATE TABLE IF NOT EXISTS file_tag
(
  file_id   text    REFERENCES file(id) ON DELETE CASCADE,
  tag_id    text    REFERENCES tag(id)  ON DELETE CASCADE,
  UNIQUE (file_id, tag_id)
);

CREATE TABLE IF NOT EXISTS taggroup
(
  id            text    PRIMARY KEY,
  tags          blob    NOT NULL UNIQUE,
  protected     int     NOT NULL,
  ctime         int     NOT NULL,
  utime         int     NOT NULL,
);

CREATE INDEX IF NOT EXISTS idx_taggroup_ctime ON taggroup(ctime);
CREATE INDEX IF NOT EXISTS idx_taggroup_utime ON taggroup(utime);

CREATE TABLE IF NOT EXISTS metadata
(
  name         text    NOT NULL UNIQUE,
  int_value    int     DEFAULT NULL,
  text_value   text    DEFAULT NULL
);
`

const InsertIntValue = `INSERT INTO metadata (name, int_value) VALUES (?, ?);`
const GetIntValue = `SELECT int_value FROM metadata WHERE name=?;`
const UpdateIntValue = `UPDATE metadata SET int_value=? WHERE name=?;`

const InsertTextValue = `INSERT INTO metadata (name, text_value) VALUES (?, ?);`
const GetTextValue = `SELECT text_value FROM metadata WHERE name=?;`
const UpdateTextValue = `UPDATE metadata SET text_value=? WHERE name=?;`

const InsertFile = `INSERT INTO file (
  id, name, size, type, thumb, hash, like, ctime, utime, deleted)
  VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`