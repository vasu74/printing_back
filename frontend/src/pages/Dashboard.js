import React, { useState, useEffect } from "react";
import {
  Container,
  Grid,
  Paper,
  Typography,
  Card,
  CardContent,
  CardHeader,
  Button,
  Box,
} from "@mui/material";
import { useNavigate } from "react-router-dom";
import { getAllTenants, getAllClients } from "../api/api";
import { useAuth } from "../context/AuthContext";

const Dashboard = () => {
  const [stats, setStats] = useState({
    tenants: 0,
    clients: 0,
  });
  const navigate = useNavigate();
  const { currentUser } = useAuth();

  useEffect(() => {
    const fetchStats = async () => {
      try {
        // Fetch tenants count
        const tenantsResponse = await getAllTenants();

        // Fetch clients count
        const clientsResponse = await getAllClients();

        setStats({
          tenants: tenantsResponse.data.length || 0,
          clients: clientsResponse.data.length || 0,
        });
      } catch (error) {
        console.error("Error fetching dashboard data:", error);
      }
    };

    fetchStats();
  }, []);

  return (
    <Container maxWidth="lg" sx={{ pt: 4, pb: 4 }}>
      <Typography variant="h4" gutterBottom>
        Dashboard
      </Typography>
      <Typography variant="subtitle1" gutterBottom>
        Welcome back, {currentUser?.name || "User"}
      </Typography>

      <Grid container spacing={3} sx={{ mt: 2 }}>
        {/* Tenants Card */}
        <Grid item xs={12} md={4}>
          <Card
            sx={{ height: "100%", display: "flex", flexDirection: "column" }}
          >
            <CardHeader
              title="Tenants"
              sx={{
                bgcolor: "primary.light",
                color: "primary.contrastText",
              }}
            />
            <CardContent sx={{ flexGrow: 1 }}>
              <Typography
                sx={{
                  fontSize: "2rem",
                  fontWeight: "bold",
                  textAlign: "center",
                  mt: 2,
                }}
              >
                {stats.tenants}
              </Typography>
              <Typography variant="body2" color="text.secondary" align="center">
                Total registered tenants
              </Typography>
              <Button
                variant="contained"
                color="primary"
                fullWidth
                sx={{ mt: 2 }}
                onClick={() => navigate("/tenants")}
              >
                Manage Tenants
              </Button>
            </CardContent>
          </Card>
        </Grid>

        {/* Clients Card */}
        <Grid item xs={12} md={4}>
          <Card
            sx={{ height: "100%", display: "flex", flexDirection: "column" }}
          >
            <CardHeader
              title="Clients"
              sx={{
                bgcolor: "primary.light",
                color: "primary.contrastText",
              }}
            />
            <CardContent sx={{ flexGrow: 1 }}>
              <Typography
                sx={{
                  fontSize: "2rem",
                  fontWeight: "bold",
                  textAlign: "center",
                  mt: 2,
                }}
              >
                {stats.clients}
              </Typography>
              <Typography variant="body2" color="text.secondary" align="center">
                Total registered clients
              </Typography>
              <Button
                variant="contained"
                color="primary"
                fullWidth
                sx={{ mt: 2 }}
                onClick={() => navigate("/clients")}
              >
                Manage Clients
              </Button>
            </CardContent>
          </Card>
        </Grid>

        {/* Products Card */}
        <Grid item xs={12} md={4}>
          <Card
            sx={{ height: "100%", display: "flex", flexDirection: "column" }}
          >
            <CardHeader
              title="Products"
              sx={{
                bgcolor: "primary.light",
                color: "primary.contrastText",
              }}
            />
            <CardContent sx={{ flexGrow: 1 }}>
              <Typography variant="body2" color="text.secondary" align="center">
                Manage your product catalog
              </Typography>
              <Button
                variant="contained"
                color="primary"
                fullWidth
                sx={{ mt: 2 }}
                onClick={() => navigate("/products")}
              >
                Manage Products
              </Button>
            </CardContent>
          </Card>
        </Grid>
      </Grid>
    </Container>
  );
};

export default Dashboard;
