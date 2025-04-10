import React, { createContext, useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { login as apiLogin } from "../api/api";

export const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
  const [currentUser, setCurrentUser] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const navigate = useNavigate();

  useEffect(() => {
    const token = localStorage.getItem("token");
    const user = localStorage.getItem("user");

    if (token && user) {
      try {
        setCurrentUser(JSON.parse(user));
      } catch (e) {
        localStorage.removeItem("token");
        localStorage.removeItem("user");
      }
    }

    setLoading(false);
  }, []);

  const login = async (credentials) => {
    try {
      setError(null);
      console.log("Attempting login with credentials:", credentials);
      const response = await apiLogin(credentials);

      console.log("Login response:", response);

      if (response.data && response.data.token) {
        localStorage.setItem("token", response.data.token);

        // Extract user info
        const user = {
          name: credentials.name,
          id: response.data.user?.id,
          // Add other user details from response if available
        };

        localStorage.setItem("user", JSON.stringify(user));
        setCurrentUser(user);

        return true;
      } else {
        // This handles when we get a response but no token
        console.error("No token in response:", response);
        setError("Login successful but no authentication token received");
        return false;
      }
    } catch (err) {
      console.error("Login error:", err);
      // Extract the error message from the response if available
      const errorMessage =
        err.response?.data?.message ||
        err.response?.data?.error ||
        "Login failed";
      console.error("Error details:", err.response?.data);
      setError(errorMessage);
      return false;
    }
  };

  const logout = () => {
    localStorage.removeItem("token");
    localStorage.removeItem("user");
    setCurrentUser(null);
    navigate("/login");
  };

  const value = {
    currentUser,
    loading,
    error,
    login,
    logout,
  };

  return (
    <AuthContext.Provider value={value}>
      {!loading && children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => {
  const context = React.useContext(AuthContext);
  if (context === undefined) {
    throw new Error("useAuth must be used within an AuthProvider");
  }
  return context;
};
