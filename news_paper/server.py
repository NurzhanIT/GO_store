from flask import Flask, request
from data import news_list
import json

news_paper = Flask(__name__)


#
@news_paper.route('/list/', methods=['GET'])
def get_news_list():
    if request.method == 'GET':
        return json.dumps(news_list)


@news_paper.route('/list/<int:pk>', methods=['GET'])
def get_news_list_by_id(**kwargs):
    if request.method == 'GET':
        if kwargs.get('pk'):
            pk = kwargs.get('pk')
            return json.dumps(news_list[pk])


@news_paper.route('/add', methods=['POST'])
def add_news(**kwargs):
    data = request.get_json()
    print(data)
    news_list.append(data)
    return f"Success added {data}"


@news_paper.route('/delete/<int:pk>', methods=['GET'])
def delete_news(**kwargs):
    if request.method == 'GET':
        if kwargs.get('pk'):
            pk = kwargs.get('pk')
            popped = json.dumps(news_list.pop(pk))
            return f"Deleted news = {popped}"


if __name__ == '__main__':
    news_paper.run(host='0.0.0.0', port=5000)
