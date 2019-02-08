const bunyan = require('bunyan');
const bunyanRequest = require('bunyan-request');
const express = require('express')

const app = express()
// 8080 is in use by the producer service
//const port = 8080
const port = 8081

const logger = bunyan.createLogger({name: 'dispatcher-test-service'});
const requestLogger = bunyanRequest({
  logger: logger,
  headerName: 'x-request-id'
});
app.use(requestLogger);

const bodyParser = require('body-parser')
app.use(bodyParser.json())
app.use(bodyParser.text({
    type: '*/*'
}))

app.get('/', (req, res) => res.send('{"alive":"kicking"}'))
app.post('/sub', (req, res) => {
    console.log(req.body);
    logger.info({ got: JSON.stringify(req.body), type: typeof req.body }, '/sub')
    res.json({ processed: true })
});
app.post('/reply', (req, res) => {
    logger.info({ got: JSON.stringify(req.body), type: typeof req.body }, '/reply')
    res.json({ processed: true })
});

app.listen(port, () => logger.info({ port }, 'Example app listening'))
