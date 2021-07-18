if (!process.env.PRODUCTION) {
    require('dotenv').config();
}

const { MongoClient } = require('mongodb');
const {createConnection} = require('net');

const uri = process.env.MONGO_URI;
const client = new MongoClient(uri, { useNewUrlParser: true, useUnifiedTopology: true });
client.connect(async (err) => {
  if (err) {
      console.error(err);
  }
  const socket = createConnection({port:8080, host: 'localhost'});
  socket.on('error', (err) => {
    console.error(err);
  })
  socket.on('connect', async () => {
    console.log("connected!");
    console.time();
    client.db()
          .collection("values")
          .find({})
          .forEach(async val => {
              delete val.class, val._id, val.__v; // remove unnecessary fields
              const  sendValue = () => {
                const res = socket.write(JSON.stringify(val))
                setImmediate(sendValue);
              }
              sendValue();
              // await sleep(1000);
          })
          .then(() => {
            console.timeEnd();
            client.close();
          })
          .catch(err => console.error(err));
  })  
});



function sleep(ms) {
  return new Promise((resolve) => {
    setTimeout(resolve, ms);
  });
}   
