import json

from flask import Flask
from flask.views import MethodView

from google.appengine.ext import ndb


app = Flask(__name__)

TEST_ENTITY_ID_1 = 'ID1'
TEST_ENTITY_ID_2 = 'ID2'
TEST_ENTITY_ID_3 = 'ID3'


class HumanModel(ndb.Model):
    FirstName = ndb.TextProperty(indexed=True, required=True)
    LastName = ndb.TextProperty(indexed=True, required=True)
    Company = ndb.TextProperty(indexed=True, required=True)

    @classmethod
    def _get_kind(cls):
        return 'Human'


class BenchmarkView(MethodView):
    def get(self):
        keys = [
            ndb.Key(HumanModel, TEST_ENTITY_ID_1),
            ndb.Key(HumanModel, TEST_ENTITY_ID_2),
            ndb.Key(HumanModel, TEST_ENTITY_ID_3),
        ]

        entities = ndb.get_multi(keys)

        if not entities:
            raise Exception('Empty results')
        else:
            payload = [{
                'FirstName': entity.FirstName,
                'LastName': entity.LastName,
                'Company': entity.Company
            } for entity in entities]
            return json.dumps(payload), 200, {'Content-Type': 'application/json'}


app.add_url_rule('/benchmark', view_func=BenchmarkView.as_view('benchmark'))
