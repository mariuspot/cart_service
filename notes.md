# Design

## Calls

* create cart 
* add item
* remote item
* remove all / empty
* get items



## Sume possibly wrong assumtions

Since the line_items has a refference to products I assumed it should be product aware (not just a generic cart service that stores id,quantities and prices)
Since there are no refferences between users and carts I assumed that the calling side will handle linking users to carts