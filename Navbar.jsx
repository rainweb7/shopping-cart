export default function Navbar({ onCheckout, onCart, onOrders }) {
  return (
    <div style={{display:"flex", justifyContent:"space-around", padding:"10px", borderBottom:"1px solid gray"}}>
      <button onClick={onCheckout}>Checkout</button>
      <button onClick={onCart}>Cart</button>
      <button onClick={onOrders}>Order History</button>
    </div>
  );
}
