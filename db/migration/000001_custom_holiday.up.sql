CREATE TYPE CATEGORY AS ENUM ('Holiday', 'Business day');
CREATE TABLE custom_holiday (
    date DATE primary key,
    category CATEGORY not null
);