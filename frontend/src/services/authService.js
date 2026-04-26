import api from "../api/http"
import { API_ENDPOINTS } from "../constants/api"

export const loginService = (payload) => {
  return api.post(API_ENDPOINTS.AUTH.LOGIN, payload)
}