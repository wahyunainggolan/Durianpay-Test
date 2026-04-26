import api from "../api/http"
import { API_ENDPOINTS } from "../constants/api"

export const getPayments = (params) => {
  return api.get(API_ENDPOINTS.PAYMENT.LIST, { params })
}