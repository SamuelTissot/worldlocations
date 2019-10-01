CREATE TABLE IF NOT EXISTS "schema_migration" (
"version" TEXT NOT NULL
);
CREATE UNIQUE INDEX "schema_migration_version_idx" ON "schema_migration" (version);
CREATE TABLE country_codes
(
    alpha_2_code       VARCHAR(2) PRIMARY KEY,
    alpha_3_code       VARCHAR(3),
    numeric_code       INTEGER,
    international_name VARCHAR(255),
    is_independant     INTEGER,
    iso_status         VARCHAR(25),
    created_at         TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at         TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE subdivision_codes
(
    subdivision_code   VARCHAR(6) PRIMARY KEY,
    alpha_2_code       VARCHAR(2),
    international_name VARCHAR(255),
    category           VARCHAR(50),
    created_at         TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at         TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (alpha_2_code) REFERENCES country_codes (alpha_2_code) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE TABLE cities
(
    id                 INTEGER PRIMARY KEY,
    alpha_2_code       VARCHAR(2)   NOT NULL,
    subdivision_code   VARCHAR(6)   NOT NULL,
    locode             VARCHAR(3)   NOT NULL,
    name               VARCHAR(255) NOT NULL,
    international_name VARCHAR(255),
    iata_code          VARCHAR(3),
    latitude_longitude VARCHAR(12),
    created_at         TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at         TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (alpha_2_code) REFERENCES country_codes (alpha_2_code),
    FOREIGN KEY (subdivision_code) REFERENCES language_codes (language_alpha_2_code) ON DELETE CASCADE
);
CREATE TABLE language_codes
(
    language_alpha_2_code VARCHAR(2) PRIMARY KEY,
    language_alpha_3_code VARCHAR(3),
    created_at            TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at            TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE subdivision_names
(
    subdivision_code      VARCHAR(6),
    language_alpha_2_code VARCHAR(6)   NOT NULL,
    name                  VARCHAR(255) NOT NULL,
    local_variation       VARCHAR,
    created_at            TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at            TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (subdivision_code) REFERENCES subdivision_codes (subdivision_code) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (language_alpha_2_code) REFERENCES language_codes (language_alpha_2_code) ON DELETE CASCADE ON UPDATE CASCADE,
    PRIMARY KEY (subdivision_code, language_alpha_2_code)
);
CREATE TABLE country_languages
(
    id                    INTEGER PRIMARY KEY,
    alpha_2_code          VARCHAR(2),
    language_alpha_2_code VARCHAR(2),
    is_administrative     INTEGER NOT NULL,
    sorting_order         VARCHAR(45),
    created_at            TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at            TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (alpha_2_code) REFERENCES country_codes (alpha_2_code),
    FOREIGN KEY (language_alpha_2_code) REFERENCES language_codes (language_alpha_2_code)
);
CREATE TABLE country_names
(
    id                    INTEGER PRIMARY KEY,
    alpha_2_code          VARCHAR(2),
    language_alpha_2_code VARCHAR(2),
    name                  VARCHAR(255) NOT NULL,
    full_name             VARCHAR(255),
    created_at            TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at            TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (alpha_2_code) REFERENCES country_codes (alpha_2_code),
    FOREIGN KEY (language_alpha_2_code) REFERENCES language_codes (language_alpha_2_code)
);
