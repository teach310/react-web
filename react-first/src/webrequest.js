import axios from 'axios';

const host = "http://localhost:8080"



class HttpAPIClient {
    async get(path, authContext){
        let result
        try {
            result = await axios.get(host+path, { headers: {'Token': authContext.accessToken}});
        } catch (error){
            if (error.response.status === 401){ // Unauthorized
                const newToken = await authContext.refreshAccessToken()
                result = await axios.get(host+path, { headers: {'Token': newToken}});
            }
        }
        
        return result
    }

    async post(path, data, authContext){
        let result
        try {
             result = await axios.post(host+path, data, { headers: {'Token': authContext.accessToken}});
        }catch (error){
            if (error.response.status === 401){ // Unauthorized
                const newToken = await authContext.refreshAccessToken()
                result = await axios.post(host+path, data, { headers: {'Token': newToken}});
            }
        }
        return result
    }
}

const client = new HttpAPIClient()
export default client