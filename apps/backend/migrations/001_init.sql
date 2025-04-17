-- 001_init.sql

-- Create the `countries` table
CREATE TABLE countries (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    code CHAR(2) NOT NULL UNIQUE,
    flag_url TEXT
);

-- Create the `tariff_type` enum
CREATE TYPE tariff_type AS ENUM ('standard', 'embargo', 'legislation');

-- Create the `tariffs` table
CREATE TABLE tariffs (
    id SERIAL PRIMARY KEY,
    country_id INTEGER NOT NULL REFERENCES countries(id) ON DELETE CASCADE,
    target_country INTEGER NOT NULL REFERENCES countries(id) ON DELETE CASCADE,
    product TEXT NOT NULL,
    type tariff_type NOT NULL,
    tariff NUMERIC(5,2) NOT NULL,
    last_updated TIMESTAMP WITHOUT TIME ZONE DEFAULT now()
);

-- Create indexes for better query performance
CREATE INDEX idx_tariffs_country_id ON tariffs(country_id);
CREATE INDEX idx_tariffs_target_country ON tariffs(target_country);
CREATE INDEX idx_tariffs_product ON tariffs(product);

-- Insert sample data into `countries`
INSERT INTO countries (name, code, flag_url) VALUES
('United States', 'US', 'https://flagcdn.com/us.svg'),
('China', 'CN', 'https://flagcdn.com/cn.svg'),
('India', 'IN', 'https://flagcdn.com/in.svg'),
('Germany', 'DE', 'https://flagcdn.com/de.svg'),
('Canada', 'CA', 'https://flagcdn.com/ca.svg');

-- Insert sample data into `tariffs`
INSERT INTO tariffs (country_id, target_country, product, type, tariff, last_updated) VALUES
(1, 2, 'Steel', 'standard', 25.00, '2023-01-01'),
(2, 1, 'Electronics', 'embargo', 40.00, '2023-02-01'),
(3, 1, 'Pharmaceuticals', 'standard', 30.00, '2023-03-01'),
(1, 3, 'Textiles', 'legislation', 15.00, '2023-04-01'),
(4, 1, 'Automobiles', 'standard', 20.00, '2023-05-01'),
(5, 1, 'Lumber', 'standard', 10.00, '2023-06-01');
