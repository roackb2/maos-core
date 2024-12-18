-- Drop the existing non-unique index if it exists
DROP INDEX IF EXISTS actors_name;

-- Create a new unique index on the name column
CREATE UNIQUE INDEX actors_name ON actors USING btree(name); 