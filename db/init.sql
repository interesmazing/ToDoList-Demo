-- Create the 'todos' table if it doesn't already exist.
CREATE TABLE IF NOT EXISTS todos (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    completed BOOLEAN DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Optional: Insert some initial data for testing
-- INSERT INTO todos (title, completed) VALUES ('Learn Go', true);
-- INSERT INTO todos (title) VALUES ('Learn Vue.js');
-- INSERT INTO todos (title) VALUES ('Build a Full-Stack App');
