import axios from 'axios';

const host = "http://localhost:8080"

class HttpAPIClient {
    async get(path){
        return await axios.get(host+path);
    }

    async post(path, data){
        return await axios.post(host+path, data)
    }
}

const client = new HttpAPIClient()
export default client