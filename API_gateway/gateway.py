from flask import Flask, Response
from routes import gateway_routes, services_routes
import requests
from services import NEWS, SHOP
gateway = Flask(__name__)

@gateway.route('/')
def welcome():
#     return ("WELCOMW TO HOBBY SHOP!!!", [route+"\n" for route in gateway_routes])
     header_value = "WELCOME TO HOBBY SHOP!!!"
     body = [route+"\n" for route in gateway_routes]

     response = Response(response=body, status=200)
     response.headers['Custom-Header'] = header_value
     return response

@gateway.route('/healthcheck')
async def healthcheck():
    response = await requests.get(SHOP + f"{services_routes['healthcheckHandler']}")
    return "WELCOMW TO HOBBY SHOP!!!"

@gateway.route('/item-create')
def item_create():
    return "WELCOMW TO HOBBY SHOP!!!"

@gateway.route('/item-show-detail')
def item_show_detail():
    return "WELCOMW TO HOBBY SHOP!!!"

@gateway.route('/item-list')
def item_list():
    return "WELCOMW TO HOBBY SHOP!!!"

@gateway.route('/items-update')
def items_update():
    return "WELCOMW TO HOBBY SHOP!!!"

@gateway.route('/item-delete')
def item_delete():
    return "WELCOMW TO HOBBY SHOP!!!"
@gateway.route('/user-register')
def user_register():
    return "WELCOMW TO HOBBY SHOP!!!"
@gateway.route('/user-activate')
def user_activate():
    return "WELCOMW TO HOBBY SHOP!!!"

@gateway.route('/token-create')
def token_create():
    return "WELCOMW TO HOBBY SHOP!!!"

@gateway.route('/basket-create')
def basket_create():
    return "WELCOMW TO HOBBY SHOP!!!"

@gateway.route('/basket-show')
def basket_show():
    return "WELCOMW TO HOBBY SHOP!!!"

@gateway.route('/basket-update')
def basket_update():
    return "WELCOMW TO HOBBY SHOP!!!"
@gateway.route('/basket-delete')
def basket_delete():
    return "WELCOMW TO HOBBY SHOP!!!"

@gateway.route('/get-current-user')
def get_current_user():
    return "WELCOMW TO HOBBY SHOP!!!"

@gateway.route('/news-list')
def news_list():
    return "WELCOMW TO HOBBY SHOP!!!"


@gateway.route('/news-detailed/<int:pk>')
def news_detailed():
    return "WELCOMW TO HOBBY SHOP!!!"


@gateway.route('/news-add')
def news_add():
    return "WELCOMW TO HOBBY SHOP!!!"


@gateway.route('/news-delete')
def news_delete():
    return "WELCOMW TO HOBBY SHOP!!!"










if __name__ == "__main__":
    gateway.run(port=5000, host='0.0.0.0')