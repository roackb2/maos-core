-- Add "deploying" status to deployment_status enum
ALTER TABLE actors DROP COLUMN IF EXISTS permissions;
