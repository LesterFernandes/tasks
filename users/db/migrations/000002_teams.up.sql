CREATE TABLE teams (
	team_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	title VARCHAR(255) NOT NULL,
	descr TEXT,
	leader_id UUID,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (leader_id) REFERENCES users(user_id) ON DELETE SET NULL
);

CREATE TYPE user_role AS ENUM ('owner', 'member');

CREATE TABLE team_members (
  user_id UUID,
  team_id UUID,
  role user_role NOT NULL DEFAULT 'member',
  CONSTRAINT user_team_unique UNIQUE(user_id, team_id),
  FOREIGN KEY (team_id) REFERENCES teams(team_id) ON DELETE CASCADE,
  FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

