CREATE TABLE customers (
	id uuid NOT NULL,
	created_at timestamp WITH TIME ZONE NOT NULL DEFAULT NOW(),
	updated_at timestamp WITH TIME ZONE NOT NULL DEFAULT NOW()
);

 CREATE TABLE transportation_ordering_officers (
	id uuid NOT NULL,
	created_at timestamp WITH TIME ZONE NOT NULL DEFAULT NOW(),
	updated_at timestamp WITH TIME ZONE NOT NULL DEFAULT NOW()
 );

 CREATE TABLE transportation_invoicing_officers (
	id uuid NOT NULL,
	created_at timestamp WITH TIME ZONE NOT NULL DEFAULT NOW(),
	updated_at timestamp WITH TIME ZONE NOT NULL DEFAULT NOW()
 );

 CREATE TABLE contracting_officers (
	id uuid NOT NULL,
	created_at timestamp WITH TIME ZONE NOT NULL DEFAULT NOW(),
	updated_at timestamp WITH TIME ZONE NOT NULL DEFAULT NOW()
 );

 CREATE TABLE ppm_office_users (
	id uuid NOT NULL,
	created_at timestamp WITH TIME ZONE NOT NULL DEFAULT NOW(),
	updated_at timestamp WITH TIME ZONE NOT NULL DEFAULT NOW()
 );