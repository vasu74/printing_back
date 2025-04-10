import React from "react";
import { Link } from "react-router-dom";
import { AppBar, Toolbar, Typography, Button, Box } from "@mui/material";
import { useAuth } from "../../context/AuthContext";

const Navbar = () => {
  const { currentUser, logout } = useAuth();

  const navLinkStyle = {
    color: "white",
    textDecoration: "none",
    margin: "0 8px",
  };

  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar position="static">
        <Toolbar>
          <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
            <Link to="/" style={navLinkStyle}>
              Product Manager Genie
            </Link>
          </Typography>

          {currentUser ? (
            <>
              <Link to="/dashboard" style={navLinkStyle}>
                <Button color="inherit">Dashboard</Button>
              </Link>
              <Link to="/tenants" style={navLinkStyle}>
                <Button color="inherit">Tenants</Button>
              </Link>
              <Link to="/clients" style={navLinkStyle}>
                <Button color="inherit">Clients</Button>
              </Link>
              <Link to="/products" style={navLinkStyle}>
                <Button color="inherit">Products</Button>
              </Link>
              <Button color="inherit" onClick={logout}>
                Logout
              </Button>
            </>
          ) : (
            <Link to="/login" style={navLinkStyle}>
              <Button color="inherit">Login</Button>
            </Link>
          )}
        </Toolbar>
      </AppBar>
    </Box>
  );
};

export default Navbar;
