// controllers/boats.js
const express = require('express');
const boatsRouter = express.Router();
const logger = require('../utils/logger');

module.exports = function(client) {
    boatsRouter.get('/', async (request, response, next) => {
        try {
            const data = await client.HGETALL("boats")
            logger.log("INFO: Receving DATA from redis db")
            const boatdata = Object.keys(data).map(key => {
                const boat = JSON.parse(data[key]);
                boat.id = key; // Ajouter l'ID du bateau à l'objet
                return boat;
            })
            response.json(boatdata)
        }catch (error) {
            logger.log(`**ERROR: ${error}**`)
            response.status(500).json({ Error: "when trying to get all data from redis db, Check log to see error description" })
        }
    })
    
	boatsRouter.get('/:id', async (request, response) => {
        const id = request.params.id
        
        try {
            const data = await client.HGET("boats", id)
		    const jsonData = JSON.parse(data)
		    response.json(jsonData)
        } catch (error) {
            logger.log(`**ERROR: ${error}**`)
            response.status(500).json({ Error: "when trying to get data from 1 boat  from redis db, Check log to see error description" })
        }
	})

    return boatsRouter;
}
