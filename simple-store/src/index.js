import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import {
  createBrowserRouter,
  createRoutesFromElements,
  Route,
  RouterProvider, 
  Routes
} from 'react-router-dom'
import ProductPage from './pages/ProductPage/ProductPage';
import ProductsPage from './pages/ProductsPage/ProductsPage';
import ProductsByCategoriesPage from './pages/ProductsPage/ProductsByCategoryPage';
import CallRequestForm from './components/OrderRequest/OrderRequest';

const router = createBrowserRouter(
  createRoutesFromElements(
      <Route>
        <Route index={true} element={<CallRequestForm />} />
        <Route path="products" element={<ProductsPage/>} />
        <Route path="product/:name" element={<><ProductPage /><CallRequestForm /></>} />
        <Route path="category/:name" element={<ProductsByCategoriesPage />} />
        <Route path="contact" element={<CallRequestForm />} />
      </Route>
  )
);

ReactDOM.createRoot(document.getElementById("root")).render(
  <React.StrictMode>
    <App/>
    <RouterProvider router={router} />
  </React.StrictMode>
);


// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
