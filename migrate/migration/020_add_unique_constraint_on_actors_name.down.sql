-- Drop the unique index
DROP INDEX IF EXISTS actors_name;

-- Recreate the original non-unique index that includes queue_id
CREATE INDEX actors_name ON actors USING btree(name, queue_id); 