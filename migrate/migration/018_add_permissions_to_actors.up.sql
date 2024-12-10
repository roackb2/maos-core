-- Add "deploying" status to deployment_status enum
ALTER TABLE actors
  ADD COLUMN IF NOT EXISTS permissions varchar(255)[] NOT NULL DEFAULT '{read:invocation}' ::varchar(255)[];
