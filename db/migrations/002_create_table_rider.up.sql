-- +goose Up
CREATE TABLE IF NOT EXISTS public.rider (
	id serial NOT NULL,
    name text NOT NULL ,
    mobile text NOT NULL unique,
    email text NOT NULL unique,
    current_location point ,
    start_otp int4 NOT NULL,
    end_otp int4 NOT NULL,
    rating int4 ,
	CONSTRAINT pk_rider PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE IF EXISTS public.rider ;