const fs = require('fs');
const path = require('path'); // For path manipulation
const moment = require('moment-timezone')

const info = (...params) => console.log(...params)

const error = (...params) => console.error(...params)

async function log(message) {
	const logFile = path.join(__dirname, '../server.log')
	const helsinkiTime = moment().tz('Europe/Helsinki')
	const timestamp = helsinkiTime.format('DD/MM/YY HH:mm:ss')
	const logEntry = `${timestamp}: ${message}\n`

    fs.writeFile(logFile, logEntry, {
		encoded: 'utf8',
		flag: "a"},
		(err) => {
			if (err){
				console.log(err)
				console.error('Error creating log file:', err);
				return
			} else {
				console.log(logEntry);
			}
	})
}

module.exports = {
	info,
	error,
	log,
}