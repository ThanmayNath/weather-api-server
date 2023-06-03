import "./App.css";
import React, { useState } from "react";
import axios from "axios";

function App() {
  const [cityInput, setCityInput] = useState("");
  const [cityList, setCityList] = useState([]);
  const [apiResponse, setAPIResponse] = useState(null);

  const handleInputChange = (event) => {
    setCityInput(event.target.value);
  };

  const handleAddCity = () => {
    if (cityInput.trim() !== "") {
      setCityList([...cityList, cityInput.trim()]);
      setCityInput("");
    }
  };

  const handleRemoveCity = (city) => {
    const updatedList = cityList.filter((c) => c !== city);
    setCityList(updatedList);
  };

  const makeAPICall = () => {
    const requestBody = {
      cities: cityList,
    };

    axios
      .post("http://localhost:8800/weather", requestBody)
      .then((response) => {
        console.log(response.data);
        setAPIResponse(response.data);
      })
      .catch((error) => {
        console.error(error);
      });
  };

  return (
    <div className="App">
      <h1>City Input</h1>
      <div className="city-input-container">
        <input
          type="text"
          placeholder="Enter a city"
          value={cityInput}
          onChange={handleInputChange}
        />
        <button onClick={handleAddCity}>Add City</button>
      </div>
      <div className="city-list">
        {cityList.map((city, index) => (
          <div key={index} className="city-list-item">
            <span>{city}</span>
            <button onClick={() => handleRemoveCity(city)}>Remove</button>
          </div>
        ))}
      </div>
      <button onClick={makeAPICall} disabled={cityList.length === 0}>
        Make API Call
      </button>

      <div className="api-response">
        <h2>API Response</h2>
        {apiResponse &&
          apiResponse.map((data, index) => (
            <p key={index}>
              {data.city} - {data.temperature} °F
            </p>
          ))}
      </div>
    </div>
  );
}

export default App;
