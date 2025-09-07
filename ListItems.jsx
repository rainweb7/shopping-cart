import { useEffect, useState } from "react";
import axios from "axios";
import Navbar from "../components/Navbar";

export default function ListItems() {
  const [items, setItems] = useState([]);
  const [cart, setCart] = useState([]);
  const [orders, setOrders] = useState([]);

  const token = localStorage.getItem("token");

  // Fetch items from backend
  useEffect(() => {
    axios.get("http://localhost:5000/items", {
      headers: { Authorization: `Bearer ${token}` }
    })
    .then(res => setItems(res.data))
    .catch(err => console.log(err));
  }, [token]);

  const addToCart = (item) => {
    setCart([...cart, item]);
    window.alert(`${item.name} added to cart`);
  };

  const handleCheckout = async () => {
    if(cart.length === 0) return window.alert("Cart is empty");
    try {
      await axios.post("http://localhost:5000/orders", { items: cart }, {
        headers: { Authorization: `Bearer ${token}` }
      });
      setCart([]);
      window.alert("Order successful");
      fetchOrders();
    } catch(err) {
      console.log(err);
      window.alert("Order failed");
    }
  };

  const showCart = () => {
    if(cart.length === 0) return window.alert("Cart is empty");
    const cartList = cart.map((i, index) => `cart_id:${index+1}, item_id:${i.id}`).join("\n");
    window.alert(cartList);
  };

  const [fetchedOrders, setFetchedOrders] = useState([]);
  const fetchOrders = async () => {
    try {
      const res = await axios.get("http://localhost:5000/orders", {
        headers: { Authorization: `Bearer ${token}` }
      });
      setFetchedOrders(res.data);
    } catch(err) {
      console.log(err);
    }
  };

  const showOrders = () => {
    if(fetchedOrders.length === 0) return window.alert("No orders yet");
    const orderList = fetchedOrders.map(o => `Order ID: ${o.id}`).join("\n");
    window.alert(orderList);
  };

  return (
    <div>
      <Navbar onCheckout={handleCheckout} onCart={showCart} onOrders={showOrders} />
      <h2 style={{textAlign:"center"}}>Items List</h2>
      <div style={{display:"flex", justifyContent:"center", gap:"20px"}}>
        {items.map(item => (
          <div key={item.id} style={{border:"1px solid black", padding:"10px", cursor:"pointer"}} onClick={()=>addToCart(item)}>
            <p>{item.name}</p>
            <p>₹{item.price}</p>
          </div>
        ))}
      </div>
    </div>
  );
}
