import React, { useState, useEffect } from "react";
import {
  Container,
  Typography,
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Button,
  Box,
  CircularProgress,
} from "@mui/material";
import { useNavigate } from "react-router-dom";
import { getAllTenants } from "../../api/api";

const TenantsList = () => {
  const [tenants, setTenants] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const navigate = useNavigate();

  useEffect(() => {
    const fetchTenants = async () => {
      try {
        setLoading(true);
        const response = await getAllTenants();
        setTenants(response.data || []);
        setError(null);
      } catch (err) {
        console.error("Error fetching tenants:", err);
        setError("Failed to load tenants. Please try again later.");
      } finally {
        setLoading(false);
      }
    };

    fetchTenants();
  }, []);

  return (
    <Container maxWidth="lg" sx={{ pt: 4, pb: 4 }}>
      <Box
        sx={{
          display: "flex",
          justifyContent: "space-between",
          alignItems: "center",
          mb: 3,
        }}
      >
        <Typography variant="h4">Tenants Management</Typography>
        <Button
          variant="contained"
          color="primary"
          onClick={() => navigate("/tenants/create")}
        >
          Add New Tenant
        </Button>
      </Box>

      {error && (
        <Box mt={2} mb={2}>
          <Typography color="error" align="center">
            {error}
          </Typography>
        </Box>
      )}

      {loading ? (
        <Box sx={{ display: "flex", justifyContent: "center", p: 4 }}>
          <CircularProgress />
        </Box>
      ) : (
        <TableContainer component={Paper} sx={{ mt: 3 }}>
          <Table>
            <TableHead>
              <TableRow>
                <TableCell>ID</TableCell>
                <TableCell>Name</TableCell>
                <TableCell>Email</TableCell>
                <TableCell>Phone</TableCell>
                <TableCell>GST No</TableCell>
                <TableCell>Role</TableCell>
                <TableCell>Address</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {tenants.length > 0 ? (
                tenants.map((tenant) => (
                  <TableRow key={tenant.id}>
                    <TableCell>{tenant.id}</TableCell>
                    <TableCell>{tenant.name}</TableCell>
                    <TableCell>{tenant.email}</TableCell>
                    <TableCell>{tenant.phoneNo}</TableCell>
                    <TableCell>{tenant.gstNo}</TableCell>
                    <TableCell>{tenant.roleName}</TableCell>
                    <TableCell>{tenant.address}</TableCell>
                  </TableRow>
                ))
              ) : (
                <TableRow>
                  <TableCell colSpan={7} align="center">
                    No tenants found
                  </TableCell>
                </TableRow>
              )}
            </TableBody>
          </Table>
        </TableContainer>
      )}
    </Container>
  );
};

export default TenantsList;
