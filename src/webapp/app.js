import React, { Component } from "react";
import "./app.css";
import ReactImage from "./react.png";

export default function App() {
  const [username, setUsername] = React.useState(null);
  const [error, setError] = React.useState("");
  const [success, setSuccess] = React.useState(false);
  React.useEffect(() => {
    fetch("/api/username")
      .then((res) => res.json())
      .then((user) => setUsername(user.username));
  }, []);

  return (
    <div>
      <div>
        <button type="button" onClick={() => testMethod(null)}>
          Test
        </button>
        <div>
          {error ? <p>error: {error}</p> : null}
          {success ? <p>success: {success}</p> : null}
        </div>
      </div>
      {username ? (
        <h1>{`Hello ${username}`}</h1>
      ) : (
        <h1>Loading.. please wait!</h1>
      )}
      <img src={ReactImage} alt="react" />
    </div>
  );

  function testMethod(test) {
    const response = fetch("/api/test", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "X-CSRF-Token": getCookie("_gorilla_csrf"),
      },
      body: JSON.stringify({ test }),
    });

    if (!response.ok) {
      setError(response.statusText);
    }
  }
}

// Function to get a specific cookie by name
function getCookie(cookieName) {
  const name = `${cookieName}=`;
  const decodedCookie = decodeURIComponent(document.cookie);
  const cookieArray = decodedCookie.split(";");
  for (let i = 0; i < cookieArray.length; i++) {
    let cookie = cookieArray[i];
    while (cookie.charAt(0) === " ") {
      cookie = cookie.substring(1);
    }
    if (cookie.indexOf(name) === 0) {
      return cookie.substring(name.length, cookie.length);
    }
  }
  return null; // Cookie not found
}
