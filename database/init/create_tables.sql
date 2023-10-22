CREATE TYPE user_role AS ENUM('guest','admin');

CREATE TABLE IF NOT EXISTS users(
    user_id serial PRIMARY KEY,
    name VARCHAR (255)  NOT NULL,
    password VARCHAR (255) NOT NULL,
    email VARCHAR (255) UNIQUE NOT NULL,
    photo TEXT,
    role user_role DEFAULT 'guest',
    is_active       BOOLEAN          DEFAULT 'true'                 NOT NULL,
    created_by      VARCHAR(255) DEFAULT 'SYSTEM'::CHARACTER VARYING NOT NULL,
    created_at      TIMESTAMP(0) DEFAULT NOW()                       NOT NULL,
    modified_by     VARCHAR(255) DEFAULT 'SYSTEM'::CHARACTER VARYING NOT NULL,
    modified_at     TIMESTAMP(0) DEFAULT NOW()                       NOT NULL,
    deleted_by      VARCHAR(255),
    deleted_at      TIMESTAMP
);
CREATE INDEX idx_user ON users USING btree (user_id);

INSERT INTO users (
    user_id, name, password, email, photo,
    role, is_active, created_by, created_at,
    modified_by, modified_at, deleted_by,
    deleted_at
)
VALUES
    (
        DEFAULT, 'Boilerplate Admin', '$2a$10$ri74429jiHyldWe2R5x8GOB6cQefe3JtnxHtiS37ofelfQk7OcG2q',
        'admin@boilerplate.go', 'https://lh3.googleusercontent.com/a/ACg8ocKJpoONQSu0UWewGeFmubaSFOtDYWdfoE3jYAc9moMmLhw=s96-c',
        'admin' :: user_role, DEFAULT, DEFAULT,
        DEFAULT, DEFAULT, DEFAULT, null, null
    );


CREATE TABLE categories (
    category_id     SERIAL          PRIMARY KEY,
    category_name   VARCHAR(255)    NOT NULL,
    is_active       BOOLEAN         DEFAULT 'true'                      NOT NULL,
    created_by      VARCHAR(255)    DEFAULT 'SYSTEM'::CHARACTER VARYING NOT NULL,
    created_at      TIMESTAMP(0)    DEFAULT NOW()                       NOT NULL,
    modified_by     VARCHAR(255)    DEFAULT 'SYSTEM'::CHARACTER VARYING NOT NULL,
    modified_at     TIMESTAMP(0)    DEFAULT NOW()                       NOT NULL,
    deleted_by      VARCHAR(255),
    deleted_at      TIMESTAMP
);

INSERT INTO categories (category_id, category_name, is_active, created_by, created_at, modified_by, modified_at,
                               deleted_by, deleted_at)
VALUES (DEFAULT, 'Transfer', DEFAULT, DEFAULT, DEFAULT, DEFAULT, DEFAULT, null, null);

INSERT INTO categories (category_id, category_name, is_active, created_by, created_at, modified_by, modified_at,
                               deleted_by, deleted_at)
VALUES (DEFAULT, 'Deposito', DEFAULT, DEFAULT, DEFAULT, DEFAULT, DEFAULT, null, null);

INSERT INTO categories (category_id, category_name, is_active, created_by, created_at, modified_by, modified_at,
                               deleted_by, deleted_at)
VALUES (DEFAULT, 'Reksa Dana', DEFAULT, DEFAULT, DEFAULT, DEFAULT, DEFAULT, null, null);

CREATE TYPE category_type AS ENUM('text_input','text_area', 'number_input', 'select');
CREATE TABLE extra_field_categories (
    field_id        SERIAL          PRIMARY KEY,
    category_id     INT,
    field_type      category_type   NOT NULL,
    field_label     VARCHAR(255)    NOT NULL,
    field_options   JSONB           DEFAULT '{}'::JSONB,
    is_active       BOOLEAN         DEFAULT 'true'                 NOT NULL,
    created_by      VARCHAR(255)    DEFAULT 'SYSTEM'::CHARACTER VARYING NOT NULL,
    created_at      TIMESTAMP(0)    DEFAULT NOW()                       NOT NULL,
    modified_by     VARCHAR(255)    DEFAULT 'SYSTEM'::CHARACTER VARYING NOT NULL,
    modified_at     TIMESTAMP(0)    DEFAULT NOW()                       NOT NULL,
    deleted_by      VARCHAR(255),
    deleted_at      TIMESTAMP,
    FOREIGN KEY (category_id) REFERENCES categories (category_id) ON DELETE CASCADE
);

INSERT INTO extra_field_categories (field_id, category_id, field_type, field_label, field_options, is_active,
                                           created_by, created_at, modified_by, modified_at, deleted_by, deleted_at)
VALUES (DEFAULT, 1, 'number_input'::category_type, 'Amount', '{}', DEFAULT, DEFAULT, DEFAULT, DEFAULT, DEFAULT, null, null);

INSERT INTO extra_field_categories (field_id, category_id, field_type, field_label, field_options, is_active,
                                           created_by, created_at, modified_by, modified_at, deleted_by, deleted_at)
VALUES (DEFAULT, 1, 'select'::category_type, 'Bank Name', e'[
  {
    "value": "bca",
    "name": "BCA"
  },
  {
    "value": "mandiri",
    "name": "Mandiri"
  }
]', DEFAULT, DEFAULT, DEFAULT, DEFAULT, DEFAULT, null, null);

INSERT INTO extra_field_categories (field_id, category_id, field_type, field_label, field_options, is_active,
                                           created_by, created_at, modified_by, modified_at, deleted_by, deleted_at)
VALUES (DEFAULT, 1, 'text_input'::category_type, 'Bank Account', '{}', DEFAULT, DEFAULT, DEFAULT, DEFAULT, DEFAULT,
        null, null);

INSERT INTO extra_field_categories (field_id, category_id, field_type, field_label, field_options, is_active,
                                           created_by, created_at, modified_by, modified_at, deleted_by, deleted_at)
VALUES (DEFAULT, 1, 'text_area'::category_type, 'Chronology', '{}', DEFAULT, DEFAULT, DEFAULT, DEFAULT, DEFAULT, null,
        null);



