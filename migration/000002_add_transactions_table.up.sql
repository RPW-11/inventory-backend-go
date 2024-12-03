CREATE TABLE "Transaction" (
  "transaction_id" serial PRIMARY KEY,
  "product_id" varchar(100) NOT NULL,
  "warehouse_id" varchar(100) NOT NULL,
  "quantity" integer NOT NULL CHECK ("quantity" > 0),
  "total_price" integer NOT NULL CHECK ("total_price" >= 0),
  "description" text,
  "employee_in_charge" varchar(100) NOT NULL,
  "transaction_type" VARCHAR NOT NULL,
  "payment_status" varchar(100) DEFAULT 'unpaid',
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  FOREIGN KEY ("product_id") REFERENCES "Product" ("id") ON DELETE CASCADE,
  FOREIGN KEY ("warehouse_id") REFERENCES "Warehouse" ("id") ON DELETE CASCADE,
  FOREIGN KEY ("employee_in_charge") REFERENCES "User" ("id") ON DELETE CASCADE
);
