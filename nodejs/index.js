const Datastore = require('@google-cloud/datastore');

const Koa = require('koa');
const Router = require('koa-router');

const projectId = 'wixgamma';

const datastore = new Datastore({ projectId });

const app = new Koa();
const router = new Router();

const KIND = 'Human';
const TEST_ENTITY_ID_1 = 'ID1';
const TEST_ENTITY_ID_2 = 'ID2';
const TEST_ENTITY_ID_3 = 'ID3';


router.get('/benchmark', async (ctx) => {
  const keys = [
    datastore.key([KIND, TEST_ENTITY_ID_1]),
    datastore.key([KIND, TEST_ENTITY_ID_2]),
    datastore.key([KIND, TEST_ENTITY_ID_3]),
  ];
  const entities = await datastore.get(keys);
  if (!entities.length) {
    throw new Error('Empty results');
  } else {
    ctx.type = 'json';
    ctx.body = entities;
  }
});

app
  .use(router.routes())
  .use(router.allowedMethods());

const port = process.env.PORT || 8080;
console.log('App started on port: ', port);
app.listen(port);
