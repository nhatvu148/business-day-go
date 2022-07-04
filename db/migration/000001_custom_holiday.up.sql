CREATE TYPE CATEGORY AS ENUM ('Holiday', 'Business day');
CREATE TABLE custom_holiday (
    id BIGSERIAL PRIMARY KEY,
    date DATE UNIQUE NOT NULL,
    category CATEGORY NOT NULL
);