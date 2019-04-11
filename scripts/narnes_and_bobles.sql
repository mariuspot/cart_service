CREATE SCHEMA IF NOT EXISTS narnes_and_boble;
USE narnes_and_boble;

CREATE TABLE IF NOT EXISTS products (
   id INTEGER PRIMARY KEY AUTO_INCREMENT NOT NULL,
   title varchar(255),
   description text,
   image_url varchar(255),
   price decimal(8,2),
   created_at datetime NOT NULL,
   updated_at datetime NOT NULL
);

CREATE TABLE IF NOT EXISTS carts (
  id INTEGER PRIMARY KEY AUTO_INCREMENT NOT NULL,
  created_at datetime NOT NULL,
  updated_at datetime NOT NULL
);

CREATE TABLE IF NOT EXISTS line_items (
  id INTEGER PRIMARY KEY AUTO_INCREMENT NOT NULL,
  product_id integer,
  cart_id integer,
  created_at datetime NOT NULL,
  updated_at datetime NOT NULL,
  quantity integer DEFAULT 1,
  price decimal(8,2),
  order_id integer,
  CONSTRAINT fk_11e15d5c6b FOREIGN KEY (product_id) REFERENCES products (id),
  CONSTRAINT fk_af645e8e5f FOREIGN KEY (cart_id) REFERENCES carts (id)
);

CREATE INDEX index_line_items_on_product_id ON line_items (product_id);
CREATE INDEX index_line_items_on_cart_id ON line_items (cart_id);

CREATE TABLE IF NOT EXISTS orders (
  id INTEGER PRIMARY KEY AUTO_INCREMENT NOT NULL,
  name varchar(255),
  address text,
  email varchar(255),
  pay_type integer,
  created_at datetime NOT NULL,
  updated_at datetime NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
  id INTEGER PRIMARY KEY AUTO_INCREMENT NOT NULL,
  name varchar(255),
  password_digest varchar(255),
  created_at datetime NOT NULL,
  updated_at datetime NOT NULL
);