-- +goose Up
CREATE TABLE IF NOT EXISTS public.driver (
	id serial NOT NULL,
    name text NOT NULL ,
    mobile text NOT NULL unique,
    vehicle_no text NOT NULL unique,
    category  text NOT NULL,
    rating int4 ,
    occupied bool,
    current_location point ,
	CONSTRAINT pk_driver PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE IF EXISTS public.driver ;