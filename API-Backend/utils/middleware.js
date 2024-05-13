const logger = require('../utils/logger')

const unknownEndpoint = (request, response) => {
	response.status(404).send({ error: 'unknown endpoint' })
}

const errorHandler = (error, request, response, next) => {
	logger.error(`**ERROR: ${error.name}-->${error.message}**`);
	response.status(500).json({ error: 'Internal server error' });
	next(error)
}

const requestLogger = (request, response, next) => {
	logger.log(`INFO: Method: ${request.method} || Path: ${request.path} || Body: ${request.body}`)
	next()
  }

module.exports = {
	unknownEndpoint,
	errorHandler,
	requestLogger,
}