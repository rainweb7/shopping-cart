import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import LoginPage from "./pages/LoginPage";
import ListItems from "./pages/ListItems";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<LoginPage />} />
        <Route path="/items" element={<ListItems />} />
      </Routes>
    </Router>
  );
}

export default App;
