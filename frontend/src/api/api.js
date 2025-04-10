import axios from "axios";

const API_URL = "http://localhost:8080";

// Create axios instance
const api = axios.create({
  baseURL: API_URL,
  headers: {
    "Content-Type": "application/json",
  },
});

// Add request interceptor to add auth token
api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem("token");
    if (token) {
      config.headers["Authorization"] = `Bearer ${token}`;
    }
    console.log("API Request:", config.method.toUpperCase(), config.url);
    return config;
  },
  (error) => {
    console.error("API Request Error:", error);
    return Promise.reject(error);
  }
);

// Add response interceptor for debugging
api.interceptors.response.use(
  (response) => {
    console.log("API Response:", response.status, response.config.url);
    return response;
  },
  (error) => {
    console.error(
      "API Response Error:",
      error.response?.status,
      error.response?.config?.url,
      error.response?.data
    );
    return Promise.reject(error);
  }
);

// Auth APIs
export const login = (credentials) => {
  console.log("Login API call with:", JSON.stringify(credentials));
  return api.post("/login", credentials);
};

// Tenant APIs
export const getAllTenants = () => api.get("/tenants");
export const createTenant = (tenantData) => api.post("/addtenants", tenantData);

// Client APIs
export const getAllClients = () => api.get("/client");
export const createClient = (clientData) => api.post("/addclient", clientData);

// Product APIs
export const createProduct = (productData) =>
  api.post("/addproduct", productData);
export const calculatePrice = (priceData) =>
  api.post("/calculateprice", priceData);

// Role APIs
export const createRole = (roleData) => api.post("/roles", roleData);
export const getRoleById = (id) => api.get(`/roles/${id}`);
export const getRolePermissions = (id) => api.get(`/rolespermission/${id}`);
export const assignPermissionToRole = (roleId, permissionId) =>
  api.post(`/roles/${roleId}/permissions`, { permissionId });

// Organization APIs
export const createOrganization = (orgData) =>
  api.post("/createOrganization", orgData);

export default api;
