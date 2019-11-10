import axios from 'axios';

const host = "http://localhost:8080"



class HttpAPIClient {
    async get(path, token){
        return await axios.get(host+path, { headers: {'Token': token}});
    }

    async post(path, data, token){
        return await axios.post(host+path, data, { headers: {'Token': token}});
    }
}

const client = new HttpAPIClient()
export default client