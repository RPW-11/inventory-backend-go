CREATE TABLE "Product" (
  "id" varchar(100) PRIMARY KEY,
  "name" varchar(100) NOT NULL,
  "description" varchar(2000) NOT NULL,
  "price" decimal(10,2) NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "ProductImage" (
  "id" serial PRIMARY KEY,
  "product_id" varchar(100) REFERENCES "Product"("id") ON DELETE CASCADE,
  "image_url" varchar(255) NOT NULL,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "User" (
  "id" varchar(100) PRIMARY KEY,
  "full_name" varchar(100) NOT NULL,
  "position" varchar(100) NOT NULL,
  "email" varchar(100) NOT NULL,
  "password" varchar(100) NOT NULL,
  "phone_number" varchar(20) NOT NULL,
  "image_url" varchar(255),
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "Inventory" (
  "id" serial PRIMARY KEY,
  "quantity" integer NOT NULL,
  "product_id" varchar(100) NOT NULL,
  "warehouse_id" varchar(100) NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "Warehouse" (
  "id" varchar(100) PRIMARY KEY,
  "name" varchar(100) NOT NULL,
  "address" varchar(300) NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

ALTER TABLE "Inventory" ADD FOREIGN KEY ("product_id") REFERENCES "Product" ("id") ON DELETE CASCADE;

ALTER TABLE "Inventory" ADD FOREIGN KEY ("warehouse_id") REFERENCES "Warehouse" ("id") ON DELETE CASCADE;