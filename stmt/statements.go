package stmt

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
  deleted       int     NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_file_hash ON file(hash);
CREATE INDEX IF NOT EXISTS idx_file_ctime ON file(ctime);
CREATE INDEX IF NOT EXISTS idx_file_utime ON file(utime);

CREATE TABLE IF NOT EXISTS tag
(
  id            text    PRIMARY KEY,
  ctime         int     NOT NULL
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
  utime         int     NOT NULL
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

const GetFile = `SELECT * FROM file WHERE id=?;`
const GetFileID = `SELECT id FROM file WHERE hash=?;`
const GetFiles = `SELECT * FROM file WHERE deleted=0 ORDER BY utime;`
const InsertFile = `INSERT INTO file (
  id, name, size, type, thumb, hash, like, ctime, utime, deleted)
  VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`

const GetTag = `SELECT * FROM tag WHERE id=?;`
const GetTagCTime = `SELECT ctime FROM tag WHERE id=?;`
const InsertTag = `INSERT INTO tag (id, ctime) VALUES ( ?, ?);`
const InsertFileTag = `INSERT INTO file_tag (file_id, tag_id) VALUES (?, ?);`

const GetTagGroupID = `SELECT id FROM taggroup WHERE tags=?;`
const InsertTagGroup = `INSERT INTO taggroup (
    id, tags, protected, ctime, utime)
    VALUES (?, ?, ?, ?, ?);`
const UpdateTagGroupNow = `UPDATE taggroup SET utime=? WHERE id=?;`
const TagGroupCount = `SELECT count(*) FROM taggroup`
const LastTagGroup = `SELECT id FROM taggroup WHERE protected=0
    ORDER BY utime LIMIT 1;`
const DeleteTagGroup = `DELETE FROM taggroup WHERE id=?;`

const GetTagsByFile = `SELECT tag.id FROM file
    INNER JOIN file_tag ON file.id = file_tag.file_id
    INNER JOIN tag ON file_tag.tag_id = tag.id
    WHERE file.id=?;`

const GetFilesByTag = `SELECT file.id FROM tag
    INNER JOIN file_tag ON tag.id = file_tag.tag_id
    INNER JOIN file ON file_tag.file_id = file.id
    WHERE file.deleted=0 and tag.id=?;`