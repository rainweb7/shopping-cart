import { useState } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";

export default function LoginPage() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate();

  const handleLogin = async (e) => {
    e.preventDefault();
    try {
      const res = await axios.post("http://localhost:5000/users/login", {
        username,
        password
      });
      localStorage.setItem("token", res.data.token);
      navigate("/items");
    } catch (err) {
      console.log(err);
      window.alert("Invalid username/password");
    }
  };

  return (
    <div style={{textAlign:"center", marginTop:"100px"}}>
      <h1>Login</h1>
      <form onSubmit={handleLogin}>
        <input placeholder="Username" value={username} onChange={e=>setUsername(e.target.value)} /><br/><br/>
        <input type="password" placeholder="Password" value={password} onChange={e=>setPassword(e.target.value)} /><br/><br/>
        <button type="submit">Login</button>
      </form>
    </div>
  );
}
