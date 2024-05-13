const config = require('./utils/config')
const express = require('express')
const app = express()
const cors = require('cors')
const redis = require('redis')


const logger = require('./utils/logger')
const boatsRouter = require('./controllers/boats')
const settingsRouter = require('./controllers/settings')
const middleware = require('./utils/middleware')

app.use(cors())
app.use(express.static('build'))
app.use(express.json())

const client = redis.createClient({ 
  host: 'localhost',
	port: 6379
})


client.on('connect', function() {
  logger.log('INFO: Connected to redis db!')
})

client.connect();

app.use((req, res, next) => {
  req.client = client
  next()
})

app.use(middleware.requestLogger)

app.use('/api/boats', boatsRouter(client))
app.use('/api/settings', settingsRouter(client))

app.use(middleware.errorHandler)
app.use(middleware.unknownEndpoint)

module.exports = app