ALTER TABLE actors
  ADD COLUMN IF NOT EXISTS mcp_enabled boolean NOT NULL DEFAULT false; 