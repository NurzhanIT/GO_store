services_routes = {
    "healthcheckHandler": "/v1/healthcheck",
    "createItemHandler": "/v1/item",
    "showItemHandler_id": "/v1/items/:id",
    "listItemsHandler": "/v1/items",
    "updateItemHandler_id": "/v1/items/:id",
    "deleteItemHandler": "/v1/items/:id",
    "registerUserHandler": "/v1/users",
    "activateUserHandler": "v1/users/activated",
    "createAuthenticationTokenHandler": "/v1/tokens/authentication",
    "createBasketHandler": "/v1/basket",
    "showBasketHandler_id": "/v1/basket/:id",
    "updateBasketHandler_id": "/v1/basket/:id",
    "deleteBasketHandler": "/v1/basket/:id",
   "getUserInfoHandler": "/v1/get-user",
   "news-list": "/list",
       "news-detailed_id": "/news-detailed/",
       "news-add": "/add",
       "news-delete": "/delete/<int:pk>",
}

gateway_routes = ['/', '/healthcheck', '/item-create',
'/item-show-detail', '/item-list', '/items-update', '/item-delete',
'/user-register', '/user-activate', '/token-create', '/basket-create',
'/basket-show', '/basket-update', '/basket-delete', '/get-current-user'
]