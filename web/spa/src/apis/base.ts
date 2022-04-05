import axios from "axios";
import requestFulfilled from "@/apis/interceptors/requestFulfilled";
import requestRejected from "@/apis/interceptors/requestRejected";
import responseFulfilled from "@/apis/interceptors/responseFulfilled";
import responseRejected from "@/apis/interceptors/responseRejected";

const baseAxios = axios.create({
    baseURL: 'http://localhost:8000/'
});

baseAxios.interceptors.request.use(requestFulfilled, requestRejected)
baseAxios.interceptors.response.use(responseFulfilled, responseRejected)

export default baseAxios;
