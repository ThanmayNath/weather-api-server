const express = require("express");
const cors = require("cors");
const axios = require("axios");
require("dotenv").config();

const app = express();
app.use(express.json());
app.use(cors());

const apiKey = process.env.APIKEY;
const apiUrl = process.env.APIURL;

app.post("/weather", async (req, res) => {
  const { cities } = req.body;
  console.log(cities);
  try {
    const weatherData = await getWeatherData(cities);
    res.json(weatherData);
  } catch (error) {
    console.error(error);
    res.status(500).json({ error: "An error occurred" });
  }
});

async function getWeatherData(cities) {
  const weatherPromises = cities.map(getWeather);
  return await Promise.all(weatherPromises);
}

async function getWeather(city) {
  const options = {
    method: "GET",
    url: `${apiUrl}${city}`,
    headers: {
      "X-RapidAPI-Key": apiKey,
      "X-RapidAPI-Host": "open-weather13.p.rapidapi.com",
    },
  };

  try {
    const response = await axios.request(options);
    return {
      city,
      temperature: response.data.main.temp,
    };
  } catch (error) {
    console.error(`Error fetching weather for ${city}:`, error);
    return {
      city,
      temperature: "N/A",
    };
  }
}

const PORT = process.env.PORT || 3000;
app.listen(PORT, () => {
  console.log("Backend is connected");
});
