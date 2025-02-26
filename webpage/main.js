const express = require("express");
const path = require("path");

const app = express();
const PORT = 5500;

// Serve static files (your HTML, CSS, JS)
app.use(express.static(path.join(__dirname, "public")));

// Handle short ID redirection
app.get("/:shortId", async (req, res) => {
  const shortId = req.params.shortId;

  try {
    const response = await fetch(`http://api:8000/api/v1/${shortId}`);
    const data = await response.json();

    if (data.data) {
      console.log(`Redirecting to: ${data.data}`);
      res.redirect(data.data); // Redirect to the original URL
    } else {
      res.status(404).send("Short URL not found");
    }
  } catch (error) {
    console.error("Error fetching short URL:", error);
    res.status(500).send("Internal Server Error");
  }
});

// Start the server
app.listen(PORT, () => {
  console.log(`Server running at http://localhost:${PORT}`);
});
