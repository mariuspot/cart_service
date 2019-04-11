# Cart microservice

Your mission, should you choose to accept it, is to extract a shopping cart microservice out of a monolithic application.

We ask you to push your solution to Github and send the link of the repository in reply to this email at maximum 2 business days after receiving this exercise.

## Current situation

You are given the following database schema of a monolithical order management system:

```
CREATE TABLE IF NOT EXISTS "products" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
   "title" varchar,
   "description" text,
   "image_url" varchar,
   "price" decimal(8,2),
   "created_at" datetime NOT NULL,
   "updated_at" datetime NOT NULL
);

CREATE TABLE IF NOT EXISTS "carts" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
  "created_at" datetime NOT NULL,
  "updated_at" datetime NOT NULL
);

CREATE TABLE IF NOT EXISTS "line_items" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
  "product_id" integer,
  "cart_id" integer,
  "created_at" datetime NOT NULL,
  "updated_at" datetime NOT NULL,
  "quantity" integer DEFAULT 1,
  "price" decimal,
  "order_id" integer,
  CONSTRAINT "fk_11e15d5c6b" FOREIGN KEY ("product_id") REFERENCES "products" ("id"),
  CONSTRAINT "fk_af645e8e5f" FOREIGN KEY ("cart_id") REFERENCES "carts" ("id")
);

CREATE INDEX "index_line_items_on_product_id" ON "line_items" ("product_id");
CREATE INDEX "index_line_items_on_cart_id" ON "line_items" ("cart_id");

CREATE TABLE IF NOT EXISTS "orders" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
  "name" varchar,
  "address" text,
  "email" varchar,
  "pay_type" integer,
  "created_at" datetime NOT NULL,
  "updated_at" datetime NOT NULL
);

CREATE TABLE IF NOT EXISTS "users" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
  "name" varchar,
  "password_digest" varchar,
  "created_at" datetime NOT NULL,
  "updated_at" datetime NOT NULL
);
```

## Cart operations

* Create a persisted cart
* Add items to cart
* Remove items from a cart
* Empty a cart
* Get the details of a cart

## Goals

Create a tiny microservice that models the requirements above in a structured way and exposes a simple API to the other microservices of an order management system. Think about what goes into a Cart microservice and what doesn't.

#### Instructions

* You can use a language or stack of your choice (we use **Ruby**/**Rails** or **Go**)
* You can choose to build a **RESTful API** using **JSON** or use **gRPC** using Protocol Buffers instead.
* There is no need for overengineering, we just want to see readable and tested code.
* Don't worry if you can't include everything, just make sure that the things you do include work.

#### Bonus points

* Including a way to consume your API
* Instrumenting your microservice

#### Out of scope

* Application deployment and scaling. we just want to run an instance of your service, not a Kubernetes cluster.