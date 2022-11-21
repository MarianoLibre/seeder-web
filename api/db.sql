CREATE TABLE
  IF NOT EXISTS products (
    `id` int NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `description` text NOT NULL,
    expiration_rate float NOT NULL,
    freezing_rate float NOT NULL,
    height float NOT NULL,
    lenght float NOT NULL,
    netweight float NOT NULL,
    product_code text NOT NULL,
    recommended_freezing_temperature float NOT NULL,
    width float NOT NULL,
    id_product_type int NOT NULL,
    id_seller int NOT NULL
  );

CREATE TABLE
  IF NOT EXISTS employees (
    `id` int NOT NULL PRIMARY KEY AUTO_INCREMENT,
    card_number_id text NOT NULL,
    first_name text NOT NULL,
    last_name text NOT NULL,
    warehouse_id int NOT NULL
  );

CREATE TABLE
  IF NOT EXISTS warehouses (
    `id` int NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `address` text NULL,
    telephone text NULL,
    warehouse_code varchar(255) NULL UNIQUE,
    minimum_capacity int NULL,
    minimum_temperature int NULL
  );

CREATE TABLE
  IF NOT EXISTS sections (
    `id` int NOT NULL PRIMARY KEY AUTO_INCREMENT,
    section_number int NOT NULL,
    current_temperature int NOT NULL,
    minimum_temperature int NOT NULL,
    current_capacity int NOT NULL,
    minimum_capacity int NOT NULL,
    maximum_capacity int NOT NULL,
    warehouse_id int NOT NULL,
    id_product_type int NOT NULL
  );

CREATE TABLE
  IF NOT EXISTS sellers (
    `id` int NOT NULL PRIMARY KEY AUTO_INCREMENT,
    cid int NOT NULL,
    company_name text NOT NULL,
    `address` text NOT NULL,
    telephone varchar(15) NOT NULL,
    locality_id int NOT NULL
  );

CREATE TABLE
  IF NOT EXISTS buyers (
    `id` int NOT NULL PRIMARY KEY AUTO_INCREMENT,
    card_number_id text NOT NULL,
    first_name text NOT NULL,
    last_name text NOT NULL,
    purchase_orders_count int NOT NULL
  );

CREATE TABLE
  IF NOT EXISTS product_records (
    `id` int NOT NULL PRIMARY KEY AUTO_INCREMENT,
    last_updated_date date NOT NULL,
    purchase_price float NOT NULL,
    sale_price float NOT NULL,
    product_id int NOT NULL
  );

CREATE TABLE
  IF NOT EXISTS carries (
    `id` int NOT NULL PRIMARY KEY AUTO_INCREMENT,
    cid int NOT NULL UNIQUE,
    company_name varchar(50) NOT NULL,
    `address` text NOT NULL,
    telephone varchar(15) NOT NULL,
    locality_id int NOT NULL
  );

CREATE TABLE
  IF NOT EXISTS product_batches (
    `id` int NOT NULL PRIMARY KEY AUTO_INCREMENT,
    batch_number int NOT NULL,
    current_quantity int NOT NULL,
    current_temperature int NOT NULL,
    due_date date NOT NULL,
    initial_quantity int NOT NULL,
    manufacturing_date date NOT NULL,
    manufacturing_hour time NOT NULL,
    minimum_temperature int NOT NULL,
    product_id int NOT NULL,
    section_id int NOT NULL
  );

CREATE TABLE
  IF NOT EXISTS purchase_orders (
    `id` int NOT NULL PRIMARY KEY AUTO_INCREMENT,
    order_number int NOT NULL,
    order_date datetime NOT NULL,
    tracking_code varchar(15) NOT NULL,
    buyer_id int NOT NULL,
    product_record_id int NOT NULL,
    order_status_id int NOT NULL
  );

CREATE TABLE
  IF NOT EXISTS inbound_orders (
    `id` int NOT NULL PRIMARY KEY AUTO_INCREMENT,
    order_date datetime NOT NULL,
    order_number varchar(15) NOT NULL,
    employee_id int NOT NULL,
    product_batch_id int NOT NULL,
    warehouse_id int NOT NULL
  );

CREATE TABLE
  IF NOT EXISTS localities (
    `id` int NOT NULL PRIMARY KEY AUTO_INCREMENT,
    locality_name varchar(30) NOT NULL,
    province_name varchar(30) NOT NULL,
    country_name varchar(30) NOT NULL,
    zip_code varchar(15) NOT NULL UNIQUE
  );

ALTER TABLE
  sellers ADD CONSTRAINT FK_SELLERS_LOCALITIES FOREIGN KEY (locality_id) REFERENCES localities (id);

ALTER TABLE
  employees ADD CONSTRAINT FK_EMPLOYEES_WAREHOUSES FOREIGN KEY (warehouse_id) REFERENCES warehouses (id);

ALTER TABLE
  inbound_orders ADD CONSTRAINT FK_INBOUNDORDERS_EMPLOYEES FOREIGN KEY (employee_id) REFERENCES employees (id);

ALTER TABLE
  inbound_orders ADD CONSTRAINT FK_INBOUNDORDERS_PRODUCTBATCHES FOREIGN KEY (product_batch_id) REFERENCES product_batches (id);

ALTER TABLE
  inbound_orders ADD CONSTRAINT FK_INBOUNDORDERS_WAREHOUSES FOREIGN KEY (warehouse_id) REFERENCES warehouses (id);

ALTER TABLE
  inbound_orders ADD UNIQUE INDEX `order_number_UNIQUE` (`order_number` ASC) VISIBLE;

ALTER TABLE
  carries ADD CONSTRAINT FK_CARRIES_LOCALITIES FOREIGN KEY (locality_id) REFERENCES localities (id);

ALTER TABLE
  product_batches ADD CONSTRAINT FK_PRODUCTBATCHES_PRODUCTS FOREIGN KEY (product_id) REFERENCES products (id);

ALTER TABLE
  product_batches ADD CONSTRAINT FK_PRODUCTBATCHES_SECTIONS FOREIGN KEY (section_id) REFERENCES sections (id);

ALTER TABLE
  product_records ADD CONSTRAINT FK_PRODUCTRECORDS_PRODUCTS FOREIGN KEY (product_id) REFERENCES products (id);

ALTER TABLE
  purchase_orders ADD CONSTRAINT FK_PURCHASEORDERS_BUYERS FOREIGN KEY (buyer_id) REFERENCES buyers (id);

ALTER TABLE
  purchase_orders ADD CONSTRAINT FK_PURCHASEORDERS_PRODUCTRECORDS FOREIGN KEY (product_record_id) REFERENCES product_records (id);
