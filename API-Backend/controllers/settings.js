// controllers/boats.js
const express = require('express');
const settingsRouter = express.Router();
const logger = require('../utils/logger');

module.exports = function(client) {
	settingsRouter.get('/', async (request, response) => {
        try {
            const data = await client.HGET("settings", "settings")
		    const jsonData = JSON.parse(data)
            jsonData["RefreshRate"] /= 1e6
		    response.json(jsonData)
        } catch (error) {
            logger.log(`**ERROR: ${error}**`)
            response.status(500).json({ Error: "when trying to get settings  from redis db, Check log to see error description" })
        }
	})

    return settingsRouter;
}
