import axios from 'axios'

const baseUrl = '/api/settings'

const get = async () => {
	const response = await axios.get(baseUrl)
	return await response.data
}

export default { get }