import React from "react";
import { Typography, Container, Link, Box } from "@mui/material";

const Footer = () => {
  return (
    <Box
      component="footer"
      sx={{
        py: 3,
        px: 2,
        mt: 'auto',
        backgroundColor: (theme) => theme.palette.grey[200],
        textAlign: "center",
      }}
    >
      <Container maxWidth="sm">
        <Typography variant="body2" color="text.secondary">
          {"© "}
          <Link color="inherit" href="https://www.vahacreations.com/">
            Vaha Creations
          </Link>{" "}
          {new Date().getFullYear()}
          {". All rights reserved."}
        </Typography>
      </Container>
    </Box>
  );
};

export default Footer;
