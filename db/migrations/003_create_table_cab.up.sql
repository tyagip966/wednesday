-- +goose Up
CREATE TABLE IF NOT EXISTS public.cab (
	id serial NOT NULL,
    driver int4 NOT NULL,
    rider  int4 NOT NULL,
    total_distance DOUBLE PRECISION NOT NULL,
    total_amount DOUBLE PRECISION NOT NULL,
    status text,
    payment_mode text,
    source point NOT NULL,
    destination point NOT NULL,
    date_of_travel timestamptz NOT NULL,
	CONSTRAINT pk_cab PRIMARY KEY (id),
	constraint fk_driver FOREIGN KEY (driver) REFERENCES driver (id),
	constraint fk_rider FOREIGN KEY (rider) REFERENCES rider (id)
);

-- +goose Down
DROP TABLE IF EXISTS public.cab ;