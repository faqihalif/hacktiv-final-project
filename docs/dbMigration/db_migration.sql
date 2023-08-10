create table patients(
    patient_id character varying(20) COLLATE pg_catalog."default" NOT NULL,
    first_name character varying(100) COLLATE pg_catalog."default",
    last_name character varying(100) COLLATE pg_catalog."default",
    date_of_birth timestamp without time zone,
    address text COLLATE pg_catalog."default",
    email character varying(100) COLLATE pg_catalog."default",
    phone_number character varying(50) COLLATE pg_catalog."default",
    is_active boolean,
    create_at timestamp without time zone,
    user_create character varying(20) COLLATE pg_catalog."default",
    CONSTRAINT patients_pkey PRIMARY KEY (patient_id)
);