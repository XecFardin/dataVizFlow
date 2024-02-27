import axios from 'axios';

const API_URL = 'http://localhost:8000/datasets';

const ApiService = {
  async getData() {
    
      const response = await axios.get(API_URL);
      console.log(response)
      return response;
   
  }
}

export default ApiService;