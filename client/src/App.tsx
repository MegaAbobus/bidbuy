import "./App.css";
import { NavBar } from "./components/navbar/navbar";
import { Header } from "./components/header/header";
import { HomeProducts } from "./components/home-products/home-products";
import { Footer } from "./components/footer/footer";
import { Routes, Route, Link } from "react-router-dom";
import { SigIn } from "./components/register/signin/signin";
import { SigUp } from "./components/register/signup/signup";

function App() {
  return (
    <div>
      <Header />
      <div className="App">
        <NavBar />
        <Routes>
          <Route path="/" element={<HomeProducts />} />
          <Route path="/signin" element={<SigIn />} />
          <Route path="/signup" element={<SigUp />} />
        </Routes>
      </div>
      <Footer />
    </div>
  );
}

export default App;
