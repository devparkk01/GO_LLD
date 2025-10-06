The online shopping service should allow users to browse products.
Users should be able to add products to the shopping cart.
Users should be able to place orders from their cart.
The system should support multiple product categories.
The system should provide search functionality.
The system should handle inventory management and update product availability accordingly.
The system should support different payment methods.
The system should handle concurrent user requests and ensure data consistency.


ProductService → owns product metadata, categories, prices, creation of new products.

InventoryService → owns stock quantities and enforces availability.

ShopService → orchestrates user flows (Cart, Order, Checkout), delegating product lookups to ProductService and stock adjustments to InventoryService